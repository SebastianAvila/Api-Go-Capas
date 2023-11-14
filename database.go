// database.go
package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // import the MySQL driver
)

func withDB(fn func(db *sql.DB) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := getDBConnection()
		if err != nil {
			http.Error(w, "Error de conexión a la base de datos", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		err = fn(db)
		if err != nil {
			http.Error(w, "Error en la operación de base de datos", http.StatusInternalServerError)
			return
		}
	}
}
