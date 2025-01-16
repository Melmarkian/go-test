package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tabelle erstellen, wenn nicht vorhanden
	createTable()

	r := mux.NewRouter()
	r.HandleFunc("/", renderHomePage).Methods("GET")
	r.HandleFunc("/validate-email", validateEmail).Methods("POST")
	r.HandleFunc("/register", registerUser).Methods("POST")

	fmt.Println("Server läuft auf http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

// Tabelle erstellen
func createTable() {
	query := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT UNIQUE
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Fehler beim Erstellen der Tabelle:", err)
	}
}

// Startseite mit HTMX-Formular
func renderHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// E-Mail validieren (Live-Check)
func validateEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		http.Error(w, "Datenbankfehler", http.StatusInternalServerError)
		return
	}

	if exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "<p style='color: red;'>E-Mail bereits vergeben!</p>")
	} else {
		fmt.Fprintf(w, "<p style='color: green;'>E-Mail verfügbar!</p>")
	}
}

// Benutzer registrieren
func registerUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")

	_, err := db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", name, email)
	if err != nil {
		http.Error(w, "E-Mail bereits vergeben oder Datenbankfehler", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "<p style='color: green;'>Benutzer erfolgreich registriert!</p>")
}
