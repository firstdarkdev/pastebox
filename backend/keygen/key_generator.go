package keygen

// KeyGenerator Generic interface for key generators
type KeyGenerator interface {
	// Generate a new key
	Generate(length int) string
}
