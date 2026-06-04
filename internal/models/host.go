package models

const (
	TipoTCP   = "tcp"
	TipoHTTP  = "http"
	TipoHTTPS = "https"
)

type Host struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Tipo     string `json:"tipo"`
	Porta    int    `json:"porta"`
	Ativo    bool   `json:"ativo"`
}
