package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	//tpl, _ = template.ParseFiles("index.html")
	tpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login-submit", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Starting the server on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "register.html", nil)
}

// Ninja Login Handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": //todo: input boxes
		login(w, r)
	case "/login-submit": // handle the login
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Sup Ninjas")
	}
}

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Fprint(w, "hello world")
// 	tpl.ExecuteTemplate(w, "login.html", nil)
// }

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}

}

// Very Super Dumb DB
var userDB = map[string]string{
	"root": "example",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You're in. Welcome!")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not in the DB")
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dashboard.html", nil)
}
