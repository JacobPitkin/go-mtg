package cards

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Cards []struct {
	Id           string            `json:"id"`
	Name         string            `json:"name"`
	CardFaces    []CardFace        `json:"card_faces"`
	CMC          float32           `json:"cmc"`
	Defense      string            `json:"defense"`
	Description  string            `json:"oracle_text"`
	Digital      bool              `json:"digital"`
	EdhrecRank   int               `json:"edhrec_rank"`
	GameChanger  bool              `json:"game_changer"`
	Games        []string          `json:"games"`
	Identity     []string          `json:"color_identity"`
	Keywords     []string          `json:"keywords"`
	Layout       string            `json:"layout"`
	Legalities   map[string]string `json:"legalities"`
	ManaCost     string            `json:"mana_cost"`
	Power        string            `json:"power"`
	ProducedMana []string          `json:"produced_mana"`
	Rarity       string            `json:"rarity"`
	Set          string            `json:"set"`
	SetName      string            `json:"set_name"`
	Toughness    string            `json:"toughness"`
	TypeLine     string            `json:"type_line"`
}

type CardFace struct {
	CMC         float32 `json:"cmc"`
	Defense     string  `json:"defense"`
	Description string  `json:"oracle_text"`
	Layout      string  `json:"layout"`
	ManaCost    string  `json:"mana_cost"`
	Name        string  `json:"name"`
	OracleId    string  `json:"oracle_id"`
	Power       string  `json:"power"`
	Toughness   string  `json:"toughness"`
}

type CardLayout int

const (
	Normal CardLayout = iota
	Split
	Flip
	Transform
	ModalDFC
	Meld
	Leveler
	Class
	Case
	Saga
	Adventure
	Mutate
	Prototype
	Battle
	Planar
	Scheme
	Vanguard
	Token
	DoubleFacedToken
	Emblem
	Augment
	Host
	ArtSeries
	Reversible
)

var layoutName = map[CardLayout]string{
	Normal:           "normal",
	Split:            "split",
	Flip:             "flip",
	Transform:        "transform",
	ModalDFC:         "modal_dfc",
	Meld:             "meld",
	Leveler:          "leveler",
	Class:            "class",
	Case:             "case",
	Saga:             "saga",
	Adventure:        "adventure",
	Mutate:           "mutate",
	Prototype:        "prototype",
	Battle:           "battle",
	Planar:           "planar",
	Scheme:           "scheme",
	Vanguard:         "vanguard",
	Token:            "token",
	DoubleFacedToken: "double_faced_token",
	Emblem:           "emblem",
	Augment:          "augment",
	Host:             "host",
	ArtSeries:        "art_series",
	Reversible:       "reversible_card",
}

func (cl CardLayout) String() string {
	return layoutName[cl]
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

func (cards *Cards) IsLayout(layout CardLayout) (Cards, bool) {
	var result Cards

	for _, card := range *cards {
		if card.Layout == layout.String() {
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
