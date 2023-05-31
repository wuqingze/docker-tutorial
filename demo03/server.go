package main

import (
	"fmt"
    "net/http"
)

func main() {
    fmt.Printf("hello world -++-- %d\n", 199);
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
