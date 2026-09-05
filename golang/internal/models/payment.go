package models

import (
	"golang_load_test/internal/config"
	"time"
)

type Payment struct {
	ID            int64      `json:"id"`
	ParticipantID int64      `json:"participant_id"`
	TicketID      int64      `json:"ticket_id"`
	OrderID       string     `json:"order_id"`
	Subtotal      float64    `json:"subtotal"`
	Status        string     `json:"status"`
	BibNumber     *string    `json:"bib_number,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func CreatePayment(participantID, ticketID int64, orderID string, subtotal float64) (*Payment, error) {
	result, err := config.DB.Exec(
		"INSERT INTO payments (participant_id, ticket_id, order_id, subtotal, status, created_at, updated_at) VALUES (?, ?, ?, ?, 'pending', NOW(), NOW())",
		participantID, ticketID, orderID, subtotal,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetPaymentByID(id)
}

func GetPaymentByID(id int64) (*Payment, error) {
	var p Payment
	err := config.DB.QueryRow("SELECT id, participant_id, ticket_id, order_id, subtotal, status, bib_number, created_at, updated_at FROM payments WHERE id = ?", id).
		Scan(&p.ID, &p.ParticipantID, &p.TicketID, &p.OrderID, &p.Subtotal, &p.Status, &p.BibNumber, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetPaymentByOrderID(orderID string) (*Payment, error) {
	var p Payment
	err := config.DB.QueryRow("SELECT id, participant_id, ticket_id, order_id, subtotal, status, bib_number, created_at, updated_at FROM payments WHERE order_id = ?", orderID).
		Scan(&p.ID, &p.ParticipantID, &p.TicketID, &p.OrderID, &p.Subtotal, &p.Status, &p.BibNumber, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func MarkPaymentPaid(id int64, bibNumber string) error {
	_, err := config.DB.Exec("UPDATE payments SET status = 'paid', bib_number = ?, updated_at = NOW() WHERE id = ?", bibNumber, id)
	return err
}

func GetPaymentsByParticipantID(participantID int64) ([]Payment, error) {
	rows, err := config.DB.Query("SELECT id, participant_id, ticket_id, order_id, subtotal, status, bib_number, created_at, updated_at FROM payments WHERE participant_id = ?", participantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var p Payment
		if err := rows.Scan(&p.ID, &p.ParticipantID, &p.TicketID, &p.OrderID, &p.Subtotal, &p.Status, &p.BibNumber, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}
