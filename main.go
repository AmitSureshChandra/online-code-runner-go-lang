package main

import (
	"fmt"
	"net/http"
	"online-compiler/internal/handler"
)

func main() {
	http.HandleFunc("/api/v1/run/compilers", handler.HandleGetCompilers)
	http.HandleFunc("/api/v1/run", handler.HandleRun)

	fmt.Println("server starting on 8010")

	// Run the server on port 8010
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
