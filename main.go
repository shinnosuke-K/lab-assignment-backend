package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	var form struct {
		UserId   string `json:"user_id"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ng")
		return
	}

	res := struct {
		User struct {
			UserID  string `json:"user_id"`
			Entered bool   `json:"entered"`
		} `json:"user"`
	}{}

	w.Header().Set("Content-Type", "application/json")

	switch form.Password {
	case "a":
		res.User.UserID = "1316190124"
		res.User.Entered = false
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(res)
		return

	case "b":
		res.User.UserID = "1316190123"
		res.User.Entered = true
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(res)
		return

	default:
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "ng")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout")
	log.Println(w.Header())
}

func main() {
	http.HandleFunc("/auth/login", LoginHandler)
	http.HandleFunc("/auth/logout", LogoutHandler)
	http.HandleFunc("/lab/save", SaveHandler)

	http.ListenAndServe(":8080", nil)
}
