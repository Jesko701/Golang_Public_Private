package main

import (
	"fmt"
	"net/http"
	"simple_web_server/controllers"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "How is it going!")
}

func main() {
	// * Anonymous Function (Closure)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is using default route")
	})
	http.HandleFunc("/index", index)
	http.HandleFunc("/users", controllers.Users)
	http.HandleFunc("/user", controllers.User)

	fmt.Println("Starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
