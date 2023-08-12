package response

import "time"

type Player struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Bank         string    `json:"bank"`
	NamaRekening string    `json:"nama_rekening"`
	NoRekening   string    `json:"no_rekening"`
	Debit        float64   `json:"debit"`
	Credit       float64   `json:"credit"`
	Deposit      float64   `json:"deposit"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ListPlayer struct {
	Player []Player `json:"data"`
	Meta   Meta     `json:"meta"`
}

type Meta struct {
	Total int   `json:"total"`
	Skip  int64 `json:"skip"`
	Limit int8  `json:"limit"`
}

type Login struct {
	Header map[string]string `json:"header"`
	Player Player            `json:"data"`
}
