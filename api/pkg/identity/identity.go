package identity

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-validation/api/pkg/persistence"
	"k8s.io/apimachinery/pkg/util/rand"
)

func Generate(n int) error {
	db := persistence.MustGetDBSession()

	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	defer tx.Rollback()

	for i := 0; i < n; i++ {
		id := strings.ToLower(rand.String(32))

		query := `INSERT INTO identity (id) VALUES (?)`
		_, err = tx.Exec(query, id)
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				i--
			} else {
				return errors.Wrap(err, "failed to add an id")
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failerd to commit transaction")
	}

	return nil
}

func List() ([]string, error) {
	db := persistence.MustGetDBSession()

	rows, err := db.Query(`SELECT id FROM identity ORDER BY id ASC`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query clusters")
	}
	defer rows.Close()

	ids := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}

		ids = append(ids, id)
	}

	return ids, nil
}
