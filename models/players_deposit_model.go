package models

import (
	"time"

	"github.com/google/uuid"
)

type PlayersDeposit struct {
	ID        uuid.UUID
	PlayerID  uuid.UUID
	Type      string
	Nominal   float64
	CreatedAt time.Time
}
