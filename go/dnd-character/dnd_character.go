package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(16) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	result := new(Character)
	result.Strength = Ability()
	result.Dexterity = Ability()
	result.Constitution = Ability()
	result.Intelligence = Ability()
	result.Wisdom = Ability()
	result.Charisma = Ability()
	result.Hitpoints = 10 + Modifier(result.Constitution)
	return *result
}
