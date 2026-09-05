package models

import (
	"golang_load_test/internal/config"
	"time"
)

type Participant struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateParticipant(name, email, phone string) (*Participant, error) {
	result, err := config.DB.Exec("INSERT INTO participants (name, email, phone, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())", name, email, phone)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetParticipantByID(id)
}

func GetParticipantByID(id int64) (*Participant, error) {
	var p Participant
	err := config.DB.QueryRow("SELECT id, name, email, phone, created_at, updated_at FROM participants WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Email, &p.Phone, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetParticipantByEmail(email string) (*Participant, error) {
	var p Participant
	err := config.DB.QueryRow("SELECT id, name, email, phone, created_at, updated_at FROM participants WHERE email = ?", email).
		Scan(&p.ID, &p.Name, &p.Email, &p.Phone, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
