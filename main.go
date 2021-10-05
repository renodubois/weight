package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeWeight(db sql.DB, weight string) {
	// First, look to see if we've got a date for today
	date := time.Now().Format("2006/01/02")
	selectQuery := "SELECT id FROM weight WHERE dateAdded = ?"
	result := db.QueryRow(selectQuery, date)
	var id int
	scanErr := result.Scan(&id)
	if scanErr == nil {
		// If so, update
		query := "UPDATE weight SET weight=? WHERE id=?"
		_, err := db.Exec(query, weight, id)
		check(err)
		fmt.Println("Today's weight updated!")
		return
	} else {
		// Otherwise, add new entry
		query := "INSERT INTO weight (weight, dateAdded) VALUES (?,?)"
		_, err := db.Exec(query, weight, date)
		check(err)
		fmt.Println("Today's weight added!")
		return
	}

}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: weight <command>\nAvailable commands:\n\tadd\n\tview")
		return
	}
	command := args[0]
	db, err := sql.Open("sqlite3", "weight_data.db")
	check(err)
	if command == "add" {
		if len(args) < 2 {
			fmt.Println("Usage: weight add <weight>")
			return
		}
		weightValue := args[1]
		writeWeight(*db, weightValue)
	} else if command == "view" {
		// View today's weight entry.
		// TODO(reno): allow for past dates to be queried
		// now := strconv.FormatInt(time.Now().Unix(), 10)
		return
	}
}
