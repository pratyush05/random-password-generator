package param

// Environment variable names
const (
	Port       = "PORT"
	PrivateKey = "PRIVATE_KEY"
	PublicKey  = "PUBLIC_KEY"
)

// HTTP headers
const (
	ContentType = "Content-Type"
	Status      = "Status"
)

// Acceptable letters for generating passwords
const (
	AcceptableLetters         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#$&"
	AcceptableLettersAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)
