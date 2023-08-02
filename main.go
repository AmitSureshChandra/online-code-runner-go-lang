package main

import (
	"fmt"
	"net/http"
	"online-compiler/internal/handler"
)

func main() {
	http.HandleFunc("/api/v1/run", handler.HandleRun)

	// Run the server on port 8010
	err := http.ListenAndServe(":8010", nil)
	fmt.Println("server started on 8010")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
