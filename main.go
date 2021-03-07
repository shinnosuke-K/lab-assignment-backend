package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {

		var form struct {
			UserId   string `json:"user_id"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ng")
			return
		}

		log.Println(form)

		if form.Password != "a" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ng")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	})

	http.ListenAndServe(":8080", nil)
}
