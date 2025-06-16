package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"jacobpitkin.com/go-mtg/cards"
	"jacobpitkin.com/go-mtg/datastore"
)

func main() {
	datastore.Connect()
	// http.HandleFunc("/", handleRoot)
	// http.HandleFunc("/cards", cardHandler)

	// http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "'Dis my house! Entry, please.")
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	cardsList := cards.NewCards()

	commanders := cardsList.IsEligibleCommander()
	commanders = commanders.Unique()
	commanders.SortByName()

	for _, card := range commanders {
		fmt.Println(card.Name)
	}

	result, err := json.MarshalIndent(commanders, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(result))
	// fmt.Fprintf(w, "%d / %d\n", len(commanders), len(cardsList))
}
