package main

import (
	"fmt"
	"net/http"
	"readcontainer/Read" // Import the Read package
)

func main() {

	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/read", Read.ReadTasks) // Use the exported function

	httpMux.HandleFunc("/readConnTest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("connected to read container")
	})

	port := ":8002"
	fmt.Printf("Starting to listen on Port: %s\n", port)
	http.ListenAndServe(port, httpMux)
}
