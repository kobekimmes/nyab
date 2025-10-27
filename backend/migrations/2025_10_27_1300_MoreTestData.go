

// backend/migrations/2025_10_25_1325_UploadTestData.go

package migrations

import (
	"fmt"
	"database/sql"
	"github.com/kobekimmes/nyab/backend/models"
	"github.com/kobekimmes/nyab/backend/db"
)

func init() {

	exampleProducts := []models.Product {
		{
			Name: "Ex. Product 2", 
			Description: "This product contains a description with a lot of text because we have to see what a product with a large text body looks like, we are also gonna test" +
			" how the photo arrays render, this will contain 4 images (all the same) and we will see how they render, maybe it would be cool if when the thumbnail is hovered from the" +
			" gallery it would slide through but thats too much for now maybe.",
			Price: 100.0, 
			Discount: 0, 
			Images: []string{"https://i.imgur.com/zQFqfv4.jpeg", "https://i.imgur.com/zQFqfv4.jpeg","https://i.imgur.com/zQFqfv4.jpeg", "https://i.imgur.com/zQFqfv4.jpeg"}, 
			Sold: false,
		},
		{
			Name: "Ex. Product 3", 
			Description: "Lets get more products cooking yippeee", 
			Price: 100.0, 
			Discount: 10.0, 
			Images: []string{"https://i.imgur.com/zQFqfv4.jpeg"}, 
			Sold: false,
		},
	}

	Register(
		models.Migration {
			Name: "2025_10_27_1300_UploadTestData.go",
			Up: func(_db *sql.DB) error {
				tx, err := _db.Begin()
				if err != nil {
					return err
				}
				defer tx.Rollback()

				for _, p := range exampleProducts {
					if _, err := db.CreateProduct(p, tx); err != nil {
						return fmt.Errorf("failed to create product %s: %w", p.Name, err)
					}
				}

				return tx.Commit()
			},
			Down: func(_db *sql.DB) error {
				tx, err := _db.Begin()
				if err != nil {
					return err
				}
				defer tx.Rollback()

				for _, p := range exampleProducts {
					if err := db.DeleteProduct(p, tx); err != nil {
						return fmt.Errorf("failed to delete product %s: %w", p.Name, err)
					}
				}

				return tx.Commit()
			},
		},
	)
}