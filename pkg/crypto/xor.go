package crypto

type XOR struct {
	key []byte
}

func NewXOR(key []byte) *XOR {
	return &XOR{key: key}
}

func (X *XOR) Encrypt(plaintext []byte) ([]byte, error) {
	return X.xor(plaintext), nil
}

func (X *XOR) Decrypt(ciphertext []byte) ([]byte, error) {
	return X.xor(ciphertext), nil
}

func (x *XOR) xor(data []byte) []byte {
	if len(x.key) == 0 {
		return data
	}
	j := 0
	for i := 0; i < len(data); i++ {
		data[i] ^= x.key[j]
		j += 1
		if j >= len(x.key) {
			j = 0
		}
	}

	return data
}
