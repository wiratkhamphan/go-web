package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "web_lek"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}

func QueryData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM emp")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age string // Add variables for other columns

		err := rows.Scan(&id, &name, &age) // Scan data into variables
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		fmt.Printf("ID: %d, Name: %s, age: %s\n", id, name, age) // Print other columns as needed
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
	}
}

// Example usage
func BD() {
	db, err := DBConnection()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	QueryData(db)
}
