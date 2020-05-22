package controllers

import (
	"database/sql"
	"log"
	"net/http"
)

// Protected handles the logic underneath the protected endpoint
func (c Controller) Protected(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SUCCESS")
	}
}
