package request

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Player struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Bank         string `json:"bank"`
	NamaRekening string `json:"nama_rekening"`
	NoRekening   string `json:"no_rekening"`
}
