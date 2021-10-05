package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: weight <command>\nAvailable commands:\n\tadd\n\tview")
		return
	}
	command := args[0]
	if command == "add" {
		if len(args) < 2 {
			fmt.Println("Usage: weight add <weight>")
			return
		}
		db, err := sql.Open("sqlite3", "weight_data.db")
		check(err)

		weightValue := args[1]
		date := strconv.FormatInt(time.Now().Unix(), 10)
		query := "INSERT INTO weight (weight, dateAdded) VALUES (?,?);"
		_, execErr := db.Exec(query, weightValue, date)
		check(execErr)
		fmt.Println("Weight added!")
	}
}
