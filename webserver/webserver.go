package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type dash struct {
	Title string
	Body  string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
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
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "accessToken", Value: username, Expires: expiration, Path: "/", MaxAge: 90000, Secure: true, HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.ServeFile(w, r, "./static/success.html")
		fmt.Println("Cookie has been set.")
	} else {
		http.ServeFile(w, r, "./static/fail.html")
	}
}

func checkIfGET(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}

	cookie, err := r.Cookie("accessToken")

	if err != nil {
		fmt.Fprintf(w, "cookie error: %v", err)
		return
	}

	if cookie.Value == "admin" {
		userDash := dash{Title: cookie.Value + "'s Dashoard", Body: "This is some sample body text"}
		t := template.Must(template.New("dashboards").ParseFiles("static/dashboard.html"))
		err = t.ExecuteTemplate(w, "dashboard.html", userDash)
		if err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "You do not have access to this page.", http.StatusForbidden)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/success.html", checkIfGET)
	http.HandleFunc("/dashboard", dashboardHandler)
	rh := http.RedirectHandler("/dashboard", http.StatusMisdirectedRequest)
	http.Handle("/dashboard.html", rh)

	fmt.Println("Starting Webserver on port 8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
