package models

import (
	"golang_load_test/internal/config"
	"time"
)

type Ticket struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	BibPrefix   string    `json:"bib_prefix"`
	BibPadding  int       `json:"bib_padding"`
	BibIncrement int      `json:"bib_increment"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetAllTickets() ([]Ticket, error) {
	rows, err := config.DB.Query("SELECT id, name, price, bib_prefix, bib_padding, bib_increment, created_at, updated_at FROM tickets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []Ticket
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Name, &t.Price, &t.BibPrefix, &t.BibPadding, &t.BibIncrement, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func GetTicketByID(id int64) (*Ticket, error) {
	var t Ticket
	err := config.DB.QueryRow("SELECT id, name, price, bib_prefix, bib_padding, bib_increment, created_at, updated_at FROM tickets WHERE id = ?", id).
		Scan(&t.ID, &t.Name, &t.Price, &t.BibPrefix, &t.BibPadding, &t.BibIncrement, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func IncrementBib(ticketID int64) error {
	_, err := config.DB.Exec("UPDATE tickets SET bib_increment = bib_increment + 1, updated_at = NOW() WHERE id = ?", ticketID)
	return err
}
