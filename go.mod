module github.com/vahaponur/repl-pokedex

go 1.21.0

require internal/commands v1.0.0

require github.com/vahaponur/pokeapi v0.0.2 // indirect

replace internal/commands => ./internal/commands
