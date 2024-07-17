// O Set é uma coleção de elementos que não permite duplicatas, porem armazena os elementos de forma desordenada.
package main

// Estrutura Player para armazenar dados do jogador
type Player struct {
	Points int
	// Adicione mais campos conforme necessário, por exemplo, Kills, Deaths, etc.
}

// Tipo Set para armazenar jogadores com nomes únicos
type Set map[string]Player

// Adiciona um jogador ou atualiza seus pontos
func (s Set) Add(name string, points int) {
	player, exists := s[name]
	if !exists {
		player = Player{}
	}
	player.Points += points
	s[name] = player
}

// Remove pontos de um jogador
func (s Set) RemovePoints(name string, points int) {
	player, exists := s[name]
	if exists {
		player.Points -= points
		s[name] = player
	}
}

// Obtém todos os nomes dos jogadores
func (s Set) Values() []string {
	names := make([]string, 0, len(s))
	for name := range s {
		names = append(names, name)
	}
	return names
}

// Obtém dados de um jogador
func (s Set) Get(name string) (Player, bool) {
	player, exists := s[name]
	return player, exists
}
