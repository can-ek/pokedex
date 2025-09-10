module github.com/can-ek/pokedex

go 1.24.4

replace github.com/can-ek/pokedex/pokeapi v0.0.0 => ./internal/pokeapi

replace github.com/can-ek/pokedex/pokecache v0.0.0 => ./internal/pokecache

require github.com/can-ek/pokedex/pokeapi v0.0.0

require github.com/can-ek/pokedex/pokecache v0.0.0 // indirect
