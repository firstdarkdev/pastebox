package keygen

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// DictionaryKeyGenerator - This is currently NOT implemented
type DictionaryKeyGenerator struct {
	Dictionary []string
}

func NewDictionaryKeyGenerator(path string) (*DictionaryKeyGenerator, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open dictionary file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	words := make([]string, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read dictionary file: %w", err)
	}

	return &DictionaryKeyGenerator{Dictionary: words}, nil
}

func (d *DictionaryKeyGenerator) Generate(keyLength int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]string, keyLength)

	for i := 0; i < keyLength; i++ {
		idx := rnd.Intn(len(d.Dictionary))
		result[i] = d.Dictionary[idx]
	}

	return strings.Join(result, "")
}
