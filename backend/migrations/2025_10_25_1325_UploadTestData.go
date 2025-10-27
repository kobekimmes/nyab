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
			Name: "Example Product 1", 
			Description: "Test", 
			Price: 100.0, 
			Discount: 10.0, 
			Images: []string{"https://imgur.com/a/erzMKbC"}, 
			Sold: false,
		},
	}

	Register(
		models.Migration {
			Name: "2025_10_25_1325_UploadTestData.go",
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