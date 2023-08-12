package request

import "github.com/google/uuid"

type PlayerDeposit struct {
	PlayerID uuid.UUID `json:"player_id"`
	Type     string    `json:"type"`
	Nominal  float64   `json:"nominal"`
}
