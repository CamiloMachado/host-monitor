package main

// Para chamar todos os arquivos go run .

import (
	"projeto_host_monitor/internal/models"
	"projeto_host_monitor/internal/services"
)

func main() {

	host := models.Host{
		Nome:     "Servidor Web",
		Endereco: "192.168.1.10",
		Tipo:     "tcp",
		Porta:    80,
		Ativo:    true,
	}

	//fmt.Print(services.CriarHost(host))

	err := services.CriarHost(host)
	//hosts, index, err := services.BuscarHostPorID("e7da13df-cf2d-4a6b-9975-d083fb0be124")
	if err != nil {
		panic(err)
	}

	//fmt.Println("Hostname:", hosts[index].Nome)
	//fmt.Println("Endereço:", hosts[index].Endereco)
	//fmt.Println("Tipo:", hosts[index].Tipo)
	//fmt.Println("Porta: ", hosts[index].Porta)
	//fmt.Println("Ativo: ", hosts[index].Ativo)
}
