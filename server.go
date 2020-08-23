package main

import (
	"github.com/YtaloWill/1sti.challenge.backend/database"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	database.ConnectDb()
	defer database.Db.Close()

	// Load Schemas
	schema, err := GetSchema()
	if err != nil {
		log.Fatal("error compiling schema: ", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   false,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	log.Printf("Server running at http://localhost:%s/graphql", port)

	http.ListenAndServe(":"+port, nil)
}
