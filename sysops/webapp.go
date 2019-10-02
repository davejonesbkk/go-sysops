package main 

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The homepage - Go Home")

}

func about_handler(w, http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the about page")
}

func main() {
		http.Handlefunc("/", index_handler)
		http.Handlefunc("/about/", about_handler)
		http.ListenAndServe(":8080", nil)

}