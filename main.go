package main

import (
	"encoding/json"
	"fmt"
	"log"

	"jacobpitkin.com/go-mtg/data"
)

func main() {
	cards := data.NewCards()

	cmcFiltered, ok := cards.WithCmc(2.0)
	if !ok {
		log.Fatal("Couldn't filter cards by cmc")
	}

	identity := []string{"G", "W"}
	identityFiltered, ok := cmcFiltered.HasIdentity(identity)
	if !ok {
		log.Fatal("Couldn't filter cards by identity")
	}

	uniqueFiltered, ok := identityFiltered.Unique()
	if !ok {
		log.Fatal("Couldn't filter unique card names")
	}

	digitalFiltered, ok := uniqueFiltered.IncludeDigital(false)
	if !ok {
		log.Fatal("Couldn't filter digital cards")
	}

	result, err := json.MarshalIndent(digitalFiltered, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
	fmt.Printf("%d / %d\n", len(digitalFiltered), len(*cards))
}
