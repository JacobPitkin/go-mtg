package cards

import (
	"sort"
	"strings"
)

type CardList []Card

type Card struct {
	Id           string            `json:"id" bson:"id"`
	Name         string            `json:"name" bson:"name"`
	CardFaces    []CardFace        `json:"card_faces" bson:"card_faces"`
	CMC          float32           `json:"cmc" bson:"cmc"`
	Defense      string            `json:"defense" bson:"defense"`
	Description  string            `json:"oracle_text" bson:"oracle_text"`
	Digital      bool              `json:"digital" bson:"digital"`
	EdhrecRank   int               `json:"edhrec_rank" bson:"edhrec_rank"`
	GameChanger  bool              `json:"game_changer" bson:"game_changer"`
	Games        []string          `json:"games" bson:"games"`
	Identity     []string          `json:"color_identity" bson:"color_identity"`
	Keywords     []string          `json:"keywords" bson:"keywords"`
	Layout       string            `json:"layout" bson:"layout"`
	Legalities   map[string]string `json:"legalities" bson:"legalities"`
	ManaCost     string            `json:"mana_cost" bson:"mana_cost"`
	Power        string            `json:"power" bson:"power"`
	ProducedMana []string          `json:"produced_mana" bson:"produced_mana"`
	Rarity       string            `json:"rarity" bson:"rarity"`
	Set          string            `json:"set" bson:"set"`
	SetName      string            `json:"set_name" bson:"set_name"`
	Toughness    string            `json:"toughness" bson:"toughness"`
	TypeLine     string            `json:"type_line" bson:"type_line"`
}

type CardFace struct {
	CMC         float32 `json:"cmc" bson:"cmc"`
	Defense     string  `json:"defense" bson:"defense"`
	Description string  `json:"oracle_text" bson:"oracle_text"`
	Layout      string  `json:"layout" bson:"layout"`
	ManaCost    string  `json:"mana_cost" bson:"mana_cost"`
	Name        string  `json:"name" bson:"name"`
	OracleId    string  `json:"oracle_id" bson:"oracle_id"`
	Power       string  `json:"power" bson:"power"`
	Toughness   string  `json:"toughness" bson:"toughness"`
}

type CardLayout int

type ByName CardList

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

var cards CardList

func (cl CardLayout) String() string {
	return layoutName[cl]
}

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	cardFilepath := os.Getenv("CARDS_LOCATION")

// 	fileData, err := os.ReadFile(cardFilepath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := json.Unmarshal(fileData, &cards); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Unmarshaled card json to in-memroy store")
// }

func NewCards() CardList {
	var newCardList CardList
	return append(newCardList, cards...)
}

func (cards *CardList) HasIdentity(identity []string) CardList {
	var result CardList

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

	return result
}

func (cards *CardList) IncludeDigital(include bool) CardList {
	var result CardList

	for _, card := range *cards {
		if !include && card.Digital {
			continue
		}

		result = append(result, card)
	}

	return result
}

func (cards *CardList) IsEligibleCommander() CardList {
	var result CardList

	for _, card := range *cards {
		types := strings.Split(card.TypeLine, " ")
		isLegendary := false
		isCreature := false

		for _, cardType := range types {
			if cardType == "Legendary" {
				isLegendary = true
				continue
			}

			if cardType == "Creature" {
				isCreature = true
				continue
			}
		}

		if isLegendary && isCreature {
			result = append(result, card)
		}
	}

	return result
}

func (cards *CardList) IsLayout(layout CardLayout) CardList {
	var result CardList

	for _, card := range *cards {
		if card.Layout == layout.String() {
			result = append(result, card)
		}
	}

	return result
}

func (cards *CardList) SortByName() {
	sort.Sort(ByName(*cards))
}

func (cards *CardList) Unique() CardList {
	var result CardList
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

	return result
}

func (cards *CardList) WithCmc(cmc float32) CardList {
	var result CardList

	for _, card := range *cards {
		if card.CMC == cmc {
			result = append(result, card)
		}
	}

	return result
}

// Sorting
func (cards ByName) Len() int           { return len(cards) }
func (cards ByName) Less(i, j int) bool { return cards[i].Name < cards[j].Name }
func (cards ByName) Swap(i, j int)      { cards[i], cards[j] = cards[j], cards[i] }
