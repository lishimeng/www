package www

type PasswordHandler interface {
	Encoding(h func(plaintext string) string) string
	Verify(h func(plaintext string) bool) bool
}
