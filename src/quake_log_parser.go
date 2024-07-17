package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Exibe o menu e retorna a escolha do usuário
func showMenu() int {
	fmt.Println("Menu:")
	fmt.Println("1. Ranking de Jogadores por Partida")
	fmt.Println("2. Ranking de Mortes")
	fmt.Println("4. Sair")
	fmt.Print("Escolha uma opção: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))

	return choice
}

// Função para capturar grupos de regex
func extractGroups(input string) []string {
	// Compilar a expressão regular
	regex := regexp.MustCompile(`:\s*(<?\w+>?)\s+killed (.*) by (.*)$`)

	// Encontrar os grupos na string de entrada
	matches := regex.FindStringSubmatch(input)

	// Verificar se houve correspondência
	if len(matches) > 0 {
		return matches
	}

	return nil
}

// Formata uma string JSON
func jsonFormat(key string, value interface{}) string {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("\"%s\": %s", key, string(jsonData))
}

func main() {

	// Abrir o arquivo de log
	file, err := os.Open("quake_server.log")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Leitor do arquivo
	scanner := bufio.NewScanner(file)

	// Inicializar variáveis
	players := make(Set)
	deathCauses := make(Set)
	gameCounter := 0
	total_kills := 0

	// Variáveis para armazenar o relatório
	var reportPlayers strings.Builder
	var reportDeathCauses strings.Builder

	// Iterar sobre cada linha do arquivo
	for scanner.Scan() {
		line := scanner.Text()

		// Verificar se a linha contém informações de morte
		if strings.Contains(line, "killed") {
			total_kills++ // Incrementa o contador geral de kills

			// Separa os Players
			result := extractGroups(line)

			if result != nil {

				killer := result[1]
				killed := result[2]
				meansOfDeath := result[3]

				// Atualizar dados do jogador
				if killer != "<world>" {
					players.Add(killer, 1)
				}
				if killed != "<world>" {
					players.Add(killed, 0)
					if killer == "<world>" {
						players.RemovePoints(killed, 1)
					}
				}

				// Atualizar causas de morte
				deathCauses.Add(meansOfDeath, 1)

			}
		}

		// Checar por nova partida
		if strings.Contains(line, "InitGame") {
			if gameCounter > 0 {

				// Adicionar dados ao relatório
				reportPlayers.WriteString(fmt.Sprintf("\"game_%d\": {\n\"total_kills\": %d,\n", gameCounter, total_kills))
				reportDeathCauses.WriteString(fmt.Sprintf("\"game_%d\": {\n\"kills_by_means\": { \n", gameCounter))

			}

			// Adicionar causas de morte ao relatório
			for i, name := range deathCauses.Values() {
				deaths, _ := deathCauses.Get(name)
				points := deaths.Points
				if i == len(deathCauses.Values())-1 {
					reportDeathCauses.WriteString(fmt.Sprintf("  \"%s\": %d\n", name, points))
				} else {
					reportDeathCauses.WriteString(fmt.Sprintf("  \"%s\": %d,\n", name, points))
				}
			}

			// Adicionar jogadores e pontos ao relatório
			reportPlayers.WriteString(jsonFormat("players", players.Values()) + ",\n")
			reportPlayers.WriteString("\"kills\": { \n")
			for i, name := range players.Values() {
				player, _ := players.Get(name)
				points := 0
				if player.Points < 0 {
					points = 0
				} else {
					points = player.Points
				}
				if i == len(players.Values())-1 {
					reportPlayers.WriteString(fmt.Sprintf("  \"%s\": %d\n", name, points))
				} else {
					reportPlayers.WriteString(fmt.Sprintf("  \"%s\": %d,\n", name, points))
				}
			}

			// Fechar a formatação JSON
			reportPlayers.WriteString("  } \n} \n")
			reportDeathCauses.WriteString("  } \n} \n")

			// Limpar dados para a próxima partida
			players = make(Set)
			deathCauses = make(Set)
			total_kills = 0
			gameCounter++
		}

	}

	// Menu para exibir os relatórios
	for {
		switch showMenu() {
		case 1:
			fmt.Println("Ranking de Jogadores por Partida:")
			fmt.Println(reportPlayers.String())

		case 2:
			fmt.Println("Ranking Causas da Morte:")
			fmt.Println(reportDeathCauses.String())

		case 4:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}

}
