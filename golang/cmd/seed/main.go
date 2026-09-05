package main

import (
	"golang_load_test/internal/config"
	"log"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	tickets := []struct {
		name        string
		price       float64
		bibPrefix   string
		bibPadding  int
		bibIncrement int
	}{
		{"Early Bird", 25.00, "EB", 5, 1},
		{"Regular", 50.00, "RG", 5, 1},
		{"VIP", 100.00, "VIP", 4, 1},
	}

	for _, t := range tickets {
		_, err := config.DB.Exec(
			"INSERT INTO tickets (name, price, bib_prefix, bib_padding, bib_increment, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())",
			t.name, t.price, t.bibPrefix, t.bibPadding, t.bibIncrement,
		)
		if err != nil {
			log.Fatalf("Failed to seed ticket %s: %v", t.name, err)
		}
		log.Printf("Seeded ticket: %s", t.name)
	}

	log.Println("Seeding complete")
}
