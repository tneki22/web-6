package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/api/user", userHandler)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello,%s!", name)
}
