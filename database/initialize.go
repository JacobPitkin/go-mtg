package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	log.Println("Initializing sql database")
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	cardsTableStatement := `
		CREATE TABLE IF NOT EXISTS cards (
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			cmc REAL,
			defense TEXT,
			description TEXT,
			digital INTEGER,
			edhrec_rank INTEGER,
			game_changer INTEGER,
			layout TEXT,
			mana_cost TEXT,
			power TEXT,
			rarity TEXT,
			set TEXT,
			set_name TEXT,
			toughness TEXT,
			type_line TEXT
		);
	`

	_, err = db.Exec(cardsTableStatement)
	if err != nil {
		log.Fatal(err)
	}

	cardFacesTableStatement := `
		CREATE TABLE IF NOT EXISTS card_faces (
			id TEXT NOT NULL FOREIGN KEY,
			cmc REAL,
			defense TEXT,
			description TEXT,
			layout TEXT,
			mana_cost TEXT,
			name TEXT,
			oracle_id TEXT,
			power TEXT,
			toughness TEXT
		);
	`

	_, err = db.Exec(cardFacesTableStatement)
	if err != nil {
		log.Fatal(err)
	}
}
