package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/kobekimmes/nyab/backend/db"
    "github.com/kobekimmes/nyab/backend/middleware"
    "github.com/kobekimmes/nyab/backend/api"
    "github.com/kobekimmes/nyab/backend/migrations"

    "github.com/joho/godotenv"
)

// Read environment variables
var bePort string 
var beDomain string 
var feBuildDir string 
var feStaticHtml string 


func main() {

    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, failed to initialize database connection")
    }

    args := os.Args[1:]

    // Run migration CLI if performing manual migrations
    if len(args) > 0 {
        if args[0] == "migrate" {
            fmt.Println("----- Running migrations UP -----")
            migrations.RunMigrationsUp()
            return
        }
        if args[0] == "rollback" {
            fmt.Println("----- Running migrations DOWN -----")
            migrations.RunMigrationsDown()
            return
        }
        return
    }
    
    // Open connection to database
    db.DbInit()
    defer db.Db.Close()

    bePort = os.Getenv("BE_PORT")
    beDomain = os.Getenv("BE_DOMAIN")
    feBuildDir = os.Getenv("FE_BUILD_DIR")
    feStaticHtml = os.Getenv("FE_STATIC_HTML")

    mux := http.NewServeMux()

    // API route handling
    mux.Handle("/api/checkout", middleware.Limit(middleware.CORS(http.HandlerFunc(api.HandleCheckout))))
    mux.Handle("/api/products", middleware.Limit(middleware.CORS(http.HandlerFunc(api.HandleProducts))))

    // Static asset serving
    fs := http.FileServer(http.Dir(feBuildDir))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    
    // Catch-all for frontend routing
    mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        if strings.HasPrefix(req.URL.Path, "/api/") {
            http.NotFound(res, req)
            return
        }
        http.ServeFile(res, req, feStaticHtml)
    })

    log.Printf("Server running on %s:%s\n", beDomain, bePort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", bePort), middleware.Logger(mux)))
}