package main

import (
	"fmt"
	"net/http"
	"online-compiler/internal/handler"
)

func main() {

	http.HandleFunc("/api/v1/run", handler.HandleRun)

	// Run the server on port 8080
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

	//content := "public class Solution {\n" +
	//	"\n" +
	//	"    public static void main(String[] args) {\n" +
	//	"        for(int i=0; i< 10000; i++)" +
	//	"			for(int j=0; j< 10000; j++)" +
	//	"				System.out.println(\"Hello World\");\n" +
	//	"    }\n" +
	//	"\n" +
	//	"}"
	//
	//input := "Amit"
	//compiler := "jdk8"

	//runner.Run(content, input, docker.GetContainerName(compiler))
}
