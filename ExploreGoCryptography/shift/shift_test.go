package shift_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/AbnormalAndy/shift"
)

var testKey = bytes.Repeat([]byte{1}, shift.BlockSize)

var cipherCases = []struct {
	plaintext, ciphertext []byte
}{
	{
		plaintext:  []byte{0, 1, 2, 3, 4, 5},
		ciphertext: []byte{1, 2, 3, 4, 5, 6},
	},
}

func TestEncrypt(t *testing.T) {
	t.Parallel()
	block, err := shift.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range cipherCases {
		name := fmt.Sprintf("%x + %x = %x", tc.plaintext, testKey, tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got := make([]byte, len(tc.plaintext))
			block.Encrypt(got, tc.plaintext)
			if !bytes.Equal(tc.ciphertext, got) {
				t.Errorf("want %x, got %x", tc.ciphertext, got)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	t.Parallel()
	block, err := shift.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range cipherCases {
		name := fmt.Sprintf("%x - %x = %x", tc.ciphertext, testKey, tc.plaintext)
		t.Run(name, func(t *testing.T) {
			got := make([]byte, len(tc.ciphertext))
			block.Decrypt(got, tc.ciphertext)
			if !bytes.Equal(tc.plaintext, got) {
				t.Errorf("want %x, got %x", tc.plaintext, got)
			}
		})
	}
}

func TestNewCipher_GivesErrKeySizeForInvalidKey(t *testing.T) {
	t.Parallel()
	_, err := shift.NewCipher([]byte{})
	if !errors.Is(err, shift.ErrKeySize) {
		t.Errorf("want ErrKeySize, got %v", err)
	}
}

func TestNewCipher_GivesNoErrorForValidKey(t *testing.T) {
	t.Parallel()
	_, err := shift.NewCipher(make([]byte, shift.BlockSize))
	if err != nil {
		t.Fatalf("want no error, got %v", err)
	}
}

func TestBlockSize_ReturnsBlockSize(t *testing.T) {
	t.Parallel()
	block, err := shift.NewCipher(make([]byte, shift.BlockSize))
	if err != nil {
		t.Fatal(err)
	}
	want := shift.BlockSize
	got := block.BlockSize()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestEncrypterEnciphersBlockAlignedMessage(t *testing.T) {
	t.Parallel()
	plaintext := []byte("This message is exactly 32 bytes")
	block, err := shift.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	enc := shift.NewEncrypter(block)
	want := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")
	got := make([]byte, 32)
	enc.CryptBlocks(got, plaintext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDecrypterCryptBlocks(t *testing.T) {
	t.Parallel()
	ciphertext := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")
	block, err := shift.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	dec := shift.NewDecrypter(block)
	want := []byte("This message is exactly 32 bytes")
	got := make([]byte, 32)
	dec.CryptBlocks(got, ciphertext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCrack(t *testing.T) {
	t.Parallel()
	plaintext := []byte("This message is exactly 32 bytes")
	ciphertext := []byte("Ujls message is exactly 32 bytes")
	want := append([]byte{1, 2, 3}, bytes.Repeat([]byte{0}, 29)...)
	got, err := shift.Crack(ciphertext, plaintext)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(want, got) {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func BenchmarkCrack(b *testing.B) {
	plaintext := []byte("This message is exactly 32 bytes")
	ciphertext := []byte("Ujlw message is exactly 32 bytes")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = shift.Crack(ciphertext, plaintext)
	}
}

var padCases = []struct {
	name           string
	actual, padded []byte
}{
	{
		name:   "1 short of full block",
		actual: []byte{0, 0, 0},
		padded: []byte{0, 0, 0, 1},
	},
	{
		name:   "2 short of full block",
		actual: []byte{0, 0},
		padded: []byte{0, 0, 2, 2},
	},
	{
		name:   "3 short of full block",
		actual: []byte{0},
		padded: []byte{0, 3, 3, 3},
	},
	{
		name:   "full block",
		actual: []byte{0, 0, 0, 0},
		padded: []byte{0, 0, 0, 0, 4, 4, 4, 4},
	},
	{
		name:   "empty block",
		actual: []byte{},
		padded: []byte{4, 4, 4, 4},
	},
}

func TestPad(t *testing.T) {
	t.Parallel()
	blockSize := 4
	for _, tc := range padCases {
		t.Run(tc.name, func(t *testing.T) {
			got := shift.Pad(tc.actual, blockSize)
			if !bytes.Equal(tc.padded, got) {
				t.Errorf("want %v, got %v", tc.padded, got)
			}
		})
	}
}

func TestUnpad(t *testing.T) {
	t.Parallel()
	blockSize := 4
	for _, tc := range padCases {
		t.Run(tc.name, func(t *testing.T) {
			got := shift.Unpad(tc.padded, blockSize)
			if !bytes.Equal(tc.actual, got) {
				t.Errorf("want %v, got %v", tc.actual, got)
			}
		})
	}
}
