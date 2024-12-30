package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect establishes a connection to the MySQL database
// @Description Establishes a connection to the MySQL database using the provided DSN (Data Source Name).
// Logs a fatal error and stops the application if the connection fails.
// @Param dsn string The Data Source Name containing the database connection details (e.g., username, password, host, port, database name).
// @Return *sql.DB A pointer to the SQL database connection.
func Connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
