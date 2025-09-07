package keygen

import (
	"math/rand"
	"time"
)

// RandomKeyGenerator - Generates random keys
type RandomKeyGenerator struct {
	Characters string
}

func NewRandomKeyGenerator(chars ...string) *RandomKeyGenerator {
	defaultChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	if len(chars) > 0 && chars[0] != "" {
		defaultChars = chars[0]
	}

	return &RandomKeyGenerator{
		Characters: defaultChars,
	}
}

func (rr *RandomKeyGenerator) Generate(length int) string {
	result := make([]byte, length)
	charactersLength := len(rr.Characters)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(charactersLength)
		result[i] = rr.Characters[randomIndex]
	}

	return string(result)
}
