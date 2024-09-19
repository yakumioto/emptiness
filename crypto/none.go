package crypto

type None struct {
}

func (n *None) Encrypt(plaintext []byte) ([]byte, error) {
	return plaintext, nil
}

func (n *None) Decrypt(ciphertext []byte) ([]byte, error) {
	return ciphertext, nil
}
