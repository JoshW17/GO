package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	fmt.Println("Hello There! - Obi Wan")

	http.ServeFile(w, r, "./static/hello.html")

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Println("POST request successful!")

	username := r.FormValue("username")
	password := r.FormValue("password")

	if (username == "admin") && (password == "admin") {
		http.ServeFile(w, r, "./static/success.html")
	} else {
		http.ServeFile(w, r, "./static/fail.html")
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Starting Webserver on port 8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
