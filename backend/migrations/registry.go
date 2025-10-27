
package migrations

import (
	"github.com/kobekimmes/nyab/backend/models"
)

var All = []models.Migration{}

func Register(m models.Migration) {
	All = append(All, m)
}