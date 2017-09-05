package main

import (
	"crypto/aes"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
)

var (
	filename = kingpin.Arg("filename", "plaintext or ciphertext").ExistingFile()
)

func main() {

	kingpin.Parse()

	plainText, err := ioutil.ReadFile(*filename)

	key := []byte{0xFF, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	cipherText := make([]byte, len(plainText))
	block.Encrypt(cipherText, plainText)
	fmt.Printf("Cipher text: %02X\n", cipherText)
}
