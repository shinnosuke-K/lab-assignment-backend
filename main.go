package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/shinnosuke-K/lab-assignment-backend/config/db"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")
	var form struct {
		StudentNum string `json:"student_num"`
		Password   string `json:"password"`
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

func SaveHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Save")

	body, _ := io.ReadAll(r.Body)
	type Request struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	log.Println(string(body))

	w.Header().Set("Content-Type", "application/json")

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	log.Println(req)

	type Response struct {
		Msg string `json:"msg"`
	}

	res := Response{Msg: "thank you"}

	b, _ := json.Marshal(res)
	fmt.Fprintf(w, string(b))
	return
}

var DB = db.Open()

func InsertLabsHandler(w http.ResponseWriter, r *http.Request) {

	if err := db.InsertLabs(DB); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

func InsertUsersHandler(w http.ResponseWriter, r *http.Request) {

	if err := db.InsertUsers(DB); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {

	http.HandleFunc("/auth/login", LoginHandler)
	http.HandleFunc("/auth/logout", LogoutHandler)
	http.HandleFunc("/lab/save", SaveHandler)

	http.HandleFunc("/lab", InsertLabsHandler)
	http.HandleFunc("/student", InsertStudentsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
