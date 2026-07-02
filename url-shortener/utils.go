package main

import "math/rand"

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(characters))
		code += string(characters[randomIndex])
	}
	return code
}
