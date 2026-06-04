package services

import (
	"errors"
	"fmt"
	"projeto_host_monitor/internal/models"
	"projeto_host_monitor/internal/storage"
	"strings"

	"github.com/google/uuid" // usar o go get github.com/google/uuid no terminal
)

const caminho = "data/hosts.json"

var (
	ErrHostNaoEncontrado   = errors.New("host não encontrado")
	ErrNomeObrigatorio     = errors.New("nome não informado")
	ErrEnderecoObrigatorio = errors.New("endereço não informado")
	ErrHostJaExiste        = errors.New("host já existe")
)

func BuscarHosts() ([]models.Host, error) {
	hosts, err := storage.CarregarHosts(caminho)

	if err != nil {
		return []models.Host{}, err
	}

	return hosts, nil
}

func BuscarHostPorID(id string, hosts []models.Host) (int, error) {
	for i := range hosts {
		if id == hosts[i].ID {
			return i, nil
		}
	}

	return -1, ErrHostNaoEncontrado
}

func CriarHost(host models.Host) error {
	hosts, err := BuscarHosts()
	if err != nil {
		return err
	}

	if err := validarHost(host); err != nil {
		return err
	}

	if err := hostJaExiste(host, hosts); err != nil {
		return err
	}

	host.ID = uuid.NewString()

	hosts = append(hosts, host)

	err = storage.SalvarHosts(caminho, hosts)
	if err != nil {
		return err
	}

	return nil
}

func AtualizarHost(host models.Host) error {
	hosts, err := BuscarHosts()
	if err != nil {
		return err
	}

	if err := validarHost(host); err != nil {
		return err
	}

	index, err := BuscarHostPorID(host.ID, hosts)

	if errors.Is(err, ErrHostNaoEncontrado) {
		return err
	}

	hosts[index].Nome = host.Nome
	hosts[index].Endereco = host.Endereco
	hosts[index].Tipo = host.Tipo
	hosts[index].Porta = host.Porta
	hosts[index].Ativo = host.Ativo

	err = storage.SalvarHosts(caminho, hosts)

	if err != nil {
		return err
	}

	return nil
}

func ExcluirHost(id string) error {
	hosts, err := BuscarHosts()
	if err != nil {
		return err
	}

	index, err := BuscarHostPorID(id, hosts)

	if err != nil {
		return err
	}

	hosts = append(hosts[:index], hosts[index+1:]...)

	err = storage.SalvarHosts(caminho, hosts)

	if err != nil {
		return err
	}

	return nil

}

func validarHost(host models.Host) error {
	if strings.TrimSpace(host.Nome) == "" {
		return ErrNomeObrigatorio
	}

	if strings.TrimSpace(host.Endereco) == "" {
		return ErrEnderecoObrigatorio
	}

	if host.Porta <= 0 || host.Porta > 65535 {
		return fmt.Errorf("porta %d informada é inválida.", host.Porta)
	}

	if host.Tipo != models.TipoHTTP && host.Tipo != models.TipoHTTPS && host.Tipo != models.TipoTCP {
		return fmt.Errorf("tipo %s informado é inválido.", host.Tipo)
	}

	return nil
}

func hostJaExiste(host models.Host, hosts []models.Host) error {
	for i := range hosts {
		if hosts[i].Nome == host.Nome ||
			(hosts[i].Endereco == host.Endereco &&
				hosts[i].Porta == host.Porta &&
				hosts[i].Tipo == host.Tipo) {
			return ErrHostJaExiste
		}
	}

	return nil
}
