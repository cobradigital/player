package models

import (
	"time"

	"github.com/google/uuid"
)

type Players struct {
	ID           uuid.UUID
	Username     string
	Password     string
	Email        string
	Bank         string
	NamaRekening string
	NoRekening   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PlayersList struct {
	ID           uuid.UUID
	Username     string
	Password     string
	Email        string
	Bank         string
	NamaRekening string
	NoRekening   string
	Debit        float64
	Credit       float64
	Deposit      float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
