package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"projeto_host_monitor/internal/models"
)

// 1. Carregar JSON
func CarregarHosts(caminho string) ([]models.Host, error) {
	// 1. Ler o conteúdo do arquivo hosts.json
	dadosJSON, err := os.ReadFile(caminho)
	if err != nil {

		return nil, fmt.Errorf("erro ao ler o arquivo: %w", err)
	}

	// 2. Decodificar o json em struct
	var hosts []models.Host
	if err := json.Unmarshal(dadosJSON, &hosts); err != nil {
		return nil, fmt.Errorf("erro ao decodificar o Json: %w", err)
	}

	return hosts, nil
}

// 2. Salvar JSON
func SalvarHosts(caminho string, hosts []models.Host) error {
	// 1. Converte a struct para JSON (Pretty Print para ficar legível)
	hostsJSON, err := json.MarshalIndent(hosts, "", " ")
	if err != nil {
		return fmt.Errorf("erro ao codificar o Json: %w", err)
	}

	// 2. Escreve no arquivo (0644 são permissões padrão de leitura/escrita)
	if err := os.WriteFile(caminho, hostsJSON, 0644); err != nil {
		return fmt.Errorf("erro ao persistir o arquivo: %w", err)
	}

	return nil
}
