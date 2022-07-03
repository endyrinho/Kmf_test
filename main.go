package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	request := fmt.Sprintf("https://nationalbank.kz/rss/get_rates.cfm?fdate=%v", id)
	req, err := http.Get(request)
	if err == nil {
		fmt.Println(err)
	}
	response := fmt.Sprintf("Product %s", id)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	fmt.Fprint(w, response)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", productsHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
