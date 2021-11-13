package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float32

func (d dollar) String() string {
	return fmt.Sprintf("%.2f円", d)
}

type database map[string]dollar

func (d database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d database) price(w http.ResponseWriter, req *http.Request) {
	queryItem := req.URL.Query().Get("item")
	price, ok := d[queryItem]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s not found in database", queryItem)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", queryItem, price)
}

func (d database) index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "welcome to starbuck")
}

func main() {
	db := database{"ice-cream": 199, "coffee": 299}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
