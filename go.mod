module github.com/vahaponur/repl-pokedex

go 1.21.0
require internal/commands v1.0.0
replace internal/commands => ./internal/commands
require internal/pokeapi v1.0.0
replace internal/pokeapi => ./internal/pokeapi
