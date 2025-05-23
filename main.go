package main

import (
	"fmt"

	"jacobpitkin.com/go-mtg/cards"
)

func main() {
	cardsList := cards.NewCards()

	commanders := cardsList.IsEligibleCommander()
	commanders = commanders.Unique()
	commanders.SortByName()

	for _, card := range commanders {
		fmt.Println(card.Name)
	}

	// result, err := json.MarshalIndent(uniqueFiltered, "", "    ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(result))
	fmt.Printf("%d / %d\n", len(commanders), len(cardsList))
}
