package keygen

import (
	"math/rand"
	"time"
)

var vowels = []rune("aeiou")
var consonants = []rune("bcdfghjklmnpqrstvwxyz")

func randOf(collection []rune) rune {
	return collection[rand.Intn(len(collection))]
}

// PhoneticKeyGenerator - Attempts to generate phonetic keys, similar to pwgen
type PhoneticKeyGenerator struct{}

func (p *PhoneticKeyGenerator) Generate(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]rune, length)

	start := rnd.Intn(2)

	for i := 0; i < length; i++ {
		if i%2 == start {
			result[i] = randOf(consonants)
		} else {
			result[i] = randOf(vowels)
		}
	}

	return string(result)
}
