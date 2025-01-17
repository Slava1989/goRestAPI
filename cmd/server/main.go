package main

import (
	"fmt"

	"github.com/Slava1989/goRestAPI/internal/comment"
	"github.com/Slava1989/goRestAPI/internal/db"
	transportHttp "github.com/Slava1989/goRestAPI/internal/transport/http"
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

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println((err))
	}
}