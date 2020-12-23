package main

import (
	"fmt"
	"os"
)

func addr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}
