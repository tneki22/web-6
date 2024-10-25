package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int = 0

func main() {
	http.HandleFunc("/count", counthHandler)
	http.ListenAndServe(":3333", nil)
}

func counthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "%d", counter)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		countstr := r.FormValue("count")
		count, err := strconv.Atoi(countstr)
		if err != nil {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		counter += count
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
