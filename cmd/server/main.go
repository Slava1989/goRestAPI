package main

import (
	"fmt"

	"github.com/Slava1989/goRestAPI/internal/db"
)

// Run - is going to be responsible for
// the instantiation and startup of
// go app
func Run() error {
	fmt.Println("starting up")

	db, err := db.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println((err))
	}
}