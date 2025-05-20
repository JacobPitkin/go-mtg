package data

import (
	"encoding/json"
	"log"
	"os"
)

type Cards []struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"oracle_text"`
	ManaCost    string   `json:"mana_cost"`
	CMC         float32  `json:"cmc"`
	Identity    []string `json:"color_identity"`
	Keywords    []string `json:"keywords"`
	// Legalities []struct {
	// 	Commander string `json:"commander"`
	// } `json:"legalities"`
	Set         string `json:"set"`
	SetName     string `json:"set_name"`
	GameChanger bool   `json:"game_changer"`
	EdhrecRank  int    `json:"edhrec_rank"`
}

func NewCards() *Cards {
	cardFilepath := "/home/groggy/mtg-cards/cards.json"

	fileData, err := os.ReadFile(cardFilepath)
	if err != nil {
		log.Fatal(err)
	}

	var cards Cards
	if err := json.Unmarshal(fileData, &cards); err != nil {
		log.Fatal(err)
	}

	return &cards
}

func (cards *Cards) HasIdentity(identity []string) (Cards, bool) {
	var result Cards

	for _, card := range *cards {
		cardMatch := true

		for _, identityColor := range identity {
			identityMatch := false

			for _, cardColor := range card.Identity {
				if identityColor == cardColor {
					identityMatch = true
				}
			}

			if !identityMatch {
				cardMatch = false
			}
		}

		if cardMatch {
			result = append(result, card)
		}
	}

	return result, true
}

func (cards *Cards) Unique() (Cards, bool) {
	var result Cards
	names := []string{}

	for _, card := range *cards {
		unique := true
		for _, name := range names {
			if name == card.Name {
				unique = false
			}
		}

		if unique {
			names = append(names, card.Name)
			result = append(result, card)
		}
	}

	return result, true
}

func (cards *Cards) WithCmc(cmc float32) (Cards, bool) {
	var result Cards

	for _, card := range *cards {
		if card.CMC == cmc {
			result = append(result, card)
		}
	}

	return result, true
}
