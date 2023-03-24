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
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Starting the server on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	tpl.ExecuteTemplate(w, "login.html", nil)
}
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "hello world")
	tpl.ExecuteTemplate(w, "dashboard.html", nil)
}
