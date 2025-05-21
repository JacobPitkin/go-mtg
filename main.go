package main

import (
	"encoding/json"
	"fmt"
	"log"

	"jacobpitkin.com/go-mtg/cards"
)

func main() {
	cardsList := cards.NewCards()

	flipCards, ok := cardsList.IsLayout(cards.Saga)
	if !ok {
		log.Fatal("Couldn't filter on multiple faces")
	}

	// cmcFiltered, ok := cards.WithCmc(2.0)
	// if !ok {
	// 	log.Fatal("Couldn't filter cards by cmc")
	// }

	// identity := []string{"G", "W"}
	// identityFiltered, ok := cmcFiltered.HasIdentity(identity)
	// if !ok {
	// 	log.Fatal("Couldn't filter cards by identity")
	// }

	uniqueFiltered, ok := flipCards.Unique()
	if !ok {
		log.Fatal("Couldn't filter unique card names")
	}

	// digitalFiltered, ok := uniqueFiltered.IncludeDigital(false)
	// if !ok {
	// 	log.Fatal("Couldn't filter digital cards")
	// }

	// for _, card := range uniqueFiltered {
	// 	fmt.Println(card.Name)
	// }

	result, err := json.MarshalIndent(uniqueFiltered, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
	fmt.Printf("%d / %d\n", len(uniqueFiltered), len(*cardsList))
}
