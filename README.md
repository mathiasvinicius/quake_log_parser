# Quake Log Parser

[English Version](#english-version)

## Introdução

Este projeto foi desenvolvido para analisar e gerar relatórios a partir dos logs do servidor Quake 3 Arena. A implementação foi feita em Go, utilizando um parser para agrupar os dados de cada partida e coletar informações detalhadas sobre as mortes e o desempenho dos jogadores.

## Estrutura do Projeto

```
quake_log_parser/
├── src/
│ ├── config.go # Arquivo para configurações ou funções auxiliares
│ ├── go.mod # Arquivo de dependências do Go
│ ├── quake_log_parser.go # Arquivo principal para parsing do log e geração de relatórios
│ ├── quake_server.log # Arquivo de log do servidor Quake
├── README.md
```


## Funcionalidades

### Parser de Log

- Leitura do arquivo de log `quake_server.log`.
- Agrupamento dos dados de cada partida.
- Coleta de informações sobre mortes, incluindo o causador da morte.

### Relatórios Gerados

- Ranking de jogadores por partida, com número de kills e mortes.
- Ranking das causas das mortes.


### Regras Implementadas

- Quando `<world>` mata um jogador, esse jogador perde 1 ponto de kill.
- Como `<world>` não é um jogador, ele nãoaparece na lista de jogadores nem no dicionário de kills.
- O contador `total_kills` inclui tanto as mortes dos jogadores quanto as mortes causadas pelo mundo.


### Menu Interativo

Ao executar o script, você verá um menu interativo com as seguintes opções:

```
1. Ranking de Jogadores por Partida
2. Ranking de Mortes
3. Sair
```


### Exemplo de Saída

#### Exemplo de saída agrupado por jogadores :

```json
"game_4": {
"total_kills": 4,
"players": ["Isgalamido","Mocinha","Zeh","Dono da Bola"],
"kills": {
  "Isgalamido": 1,
  "Mocinha": 0,
  "Zeh": 0,
  "Dono da Bola": 0
  }
}

```


#### Exemplo de saída agrupada por causa de mortes :

```json

"game_4": {
"kills_by_means": {
  "MOD_ROCKET": 1,
  "MOD_TRIGGER_HURT": 2,
  "MOD_FALLING": 1
  }
}

```


## English Version

## Introduction
This project was developed to analyze and generate reports from Quake 3 Arena server logs. The implementation was done in Go, using a parser to group the data from each match and collect detailed information about deaths and player performance.

## Project Structure

```
quake_log_parser/
├── src/
│ ├── config.go # Configuration or utility functions file
│ ├── go.mod # Go dependencies file
│ ├── quake_log_parser.go # Main file for log parsing and report generation
│ ├── quake_server.log # Arquivo de log do servidor Quake
├── README.md
```


## Features

### Log Parser

- Reads the quake_server.log file.
- Groups data from each match.
- Collects information about deaths, including the cause of death.

### Generated Reports

- Player ranking per match, including the number of kills and deaths.
- Ranking of death causes.


### Implemented Rules

- When <world> kills a player, that player loses 1 kill point.
- Since <world> is not a player, it does not appear in the player list or the kill dictionary.
- The total_kills counter includes both player deaths and deaths caused by the world.


### Interactive Menu

When you run the script, you will see an interactive menu with the following options:

```
1. Ranking de Jogadores por Partida
2. Ranking de Mortes
3. Sair
```


### Example Output

#### Example output grouped by players:

```json
"game_4": {
"total_kills": 4,
"players": ["Isgalamido","Mocinha","Zeh","Dono da Bola"],
"kills": {
  "Isgalamido": 1,
  "Mocinha": 0,
  "Zeh": 0,
  "Dono da Bola": 0
  }
}

```


#### Example output grouped by cause of death:

```json

"game_4": {
"kills_by_means": {
  "MOD_ROCKET": 1,
  "MOD_TRIGGER_HURT": 2,
  "MOD_FALLING": 1
  }
}

```
