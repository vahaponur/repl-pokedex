# REPL Pokedex
This is a REPL app that uses Pokeapi (https://pokeapi.co)
It helps you explore the Pokemon World

## Installation
This is a public repo. You can clone it.
gh repo clone vahaponur/repl-pokedex

## Running
In your CLI, while you are at in your clone
```
go build && ./repl-pokedex
```
## Usage
```
"help"  will help you with commands
"map" will get the next location area (next 20 after each call)
"mapb" will get previous 20 location area (this effects the "map" call)
"explore <area_name>" will get the pokemons in that area, you can learn area names with "map" or "mapb"
EXAMPLE USAGE
"explore canalave-city-area" will get all the pokemons in canavale city area
"exit" will stop the program
```

