package data

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Cards []struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"oracle_text"`
	ManaCost    string            `json:"mana_cost"`
	CMC         float32           `json:"cmc"`
	Digital     bool              `json:"digital"`
	Identity    []string          `json:"color_identity"`
	Keywords    []string          `json:"keywords"`
	Legalities  map[string]string `json:"legalities"`
	Set         string            `json:"set"`
	SetName     string            `json:"set_name"`
	GameChanger bool              `json:"game_changer"`
	EdhrecRank  int               `json:"edhrec_rank"`
}

func NewCards() *Cards {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cardFilepath := os.Getenv("CARDS_LOCATION")

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

func (cards *Cards) IncludeDigital(include bool) (Cards, bool) {
	var result Cards

	for _, card := range *cards {
		if !include && card.Digital {
			continue
		}

		result = append(result, card)
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
