package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tpl *template.Template

func main() {

	// get ENV variable
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load env file")
		os.Exit(1)
	}
	var (
		mongoURI = os.Getenv("SECURE_URL")
	)

	//fmt.Printf("SECURE_URL: %s\n", os.Getenv("SECURE_URL"))
	// DB Connection
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to the Mongo DB")
	}
	defer client.Disconnect(ctx)

	// Connect/create to DB instance
	goDemoDB := client.Database("godemo")
	err = goDemoDB.CreateCollection(ctx, "cats")
	if err != nil {
		log.Fatal((err))
	}
	catsCollection := goDemoDB.Collection("cats")
	defer catsCollection.Drop(ctx)
	result, err := catsCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Mocha"},
		{Key: "breed", Value: "Turkish Van"},
	})
	if err != nil {
		log.Fatal((err))
	}
	fmt.Println("Result:", result)

	// Routes
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
