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

func main() {
    // Load .env
    _ = godotenv.Load()

    // Handle CLI migration commands
    args := os.Args[1:]
    if len(args) > 0 {
        switch args[0] {
        case "migrate":
            log.Println("----- Running migrations UP -----")
            migrations.RunMigrationsUp()
            return
        case "rollback":
            log.Println("----- Running migrations DOWN -----")
            migrations.RunMigrationsDown()
            return
        }
    }

    // Database
    db.DbInit()
    defer db.Db.Close()

    mux := http.NewServeMux()

    // API routes
    mux.Handle("/api/checkout", middleware.Limit(middleware.CORS(http.HandlerFunc(api.HandleCheckout))))
    mux.Handle("/api/products", middleware.Limit(middleware.CORS(http.HandlerFunc(api.HandleProducts))))

    // Check if we want Go to serve frontend
    if os.Getenv("SERVE_FRONTEND") == "true" {
        feBuildDir := os.Getenv("FE_BUILD_DIR")
        if feBuildDir == "" {
            feBuildDir = "./frontend/dist"
        }
        feStaticHtml := os.Getenv("FE_STATIC_HTML")
        if feStaticHtml == "" {
            feStaticHtml = fmt.Sprintf("%s/index.html", feBuildDir)
        }

        // Static assets
        fs := http.FileServer(http.Dir(feBuildDir))
        mux.Handle("/static/", http.StripPrefix("/static/", fs))

        // Catch-all for frontend routes
        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            if strings.HasPrefix(r.URL.Path, "/api/") {
                http.NotFound(w, r)
                return
            }
            http.ServeFile(w, r, feStaticHtml)
        })

        log.Println("Frontend serving enabled")
    }

    // Start server
    bePort := os.Getenv("BE_PORT")
    if bePort == "" {
        bePort = "8080"
    }
    log.Printf("Server running on port %s\n", bePort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", bePort), middleware.Logger(mux)))
}
