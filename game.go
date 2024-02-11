package main

type Game struct {
	Tags  map[string]string `json:"tags"`
	Moves string            `json:"moves"`
}
