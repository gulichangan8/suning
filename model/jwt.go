package model

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Aud string `json:"aud"`
	Iat string `json:"iat"`
	Exp string `json:"exp"`
	Sub string `json:"sub"`
}
