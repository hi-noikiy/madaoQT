package utils

import "testing"

func TestEncryptDecrypt(t *testing.T) {

	file := "madaoQT.exe"
	encrypt := FileEncrypt{
		File:  file,
		Key:   "hello, world",
		Nonce: "cooldcool",
	}

	encrypt.Encrypt()

	encrypt = FileEncrypt{
		File:  file + "-encrypted",
		Key:   "hello, world",
		Nonce: "cooldcool",
	}

	encrypt.Decrypt()
}

func TestSample(t *testing.T) {
	ExampleNewGCM_encrypt()
}