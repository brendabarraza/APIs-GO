package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pagina 1")
}
func index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pagina 2")
}
func index3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pagina 3")
}
func main() {

	Router := mux.NewRouter()
	Router.HandleFunc("/", index1)
	Router.HandleFunc("/index2", index2)
	Router.HandleFunc("/index3", index3)
	http.Handle("/", Router)

	Server := http.ListenAndServe(":8080", nil)
	log.Fatal(Server)

}
