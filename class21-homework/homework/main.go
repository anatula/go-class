package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// run a webserver and query it
// db is map of string to prices
// money is represented by floating point (WRONG)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// it does not return anything, in a handler we return by writing on the ResponseWriter so it goes back to original client
// method value, closed over the db, still have 2 other params
// ok to have value receiver, db is a map descriptor (all method will refer same hash table)
// copying the map header (which contains the pointer to the data), not the actual map data itself.
// list is a method with 2 regular param and receiver is db
// by creating a method value I turn that into a function that has 2 params
// this is a function that can be passed to HandleFunc
// HandleFunc wants a handler takes 2 params (responsewriter and request)
// cast it to a certain type that then provides the method http
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s \n", item, price)
	}

}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	// variables in the url
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	fmt.Println(item, price)

	// check if item is already in map
	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %s", item)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	// turn value to floating point number
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %s", item)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "added %s with price %q \n", item, db[item])

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400

		fmt.Fprintf(w, "invalid price: %q\n", price)
	} else {
		db[item] = dollars(f64)

		fmt.Fprintf(w, "new price %s for %s\n", dollars(f64), item)
	}
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])
}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)

	fmt.Fprintf(w, "dropped %s\n", item)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	// use curl in command line like
	// curl 'localhost:8080/create?item=hat&price=22'
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)
	http.HandleFunc("/read", db.fetch)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
