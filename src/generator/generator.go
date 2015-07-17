package generator

import (
	"crypto/rand"
	"log"
)

const char_spec = "(){}[]<>?!`~*#$^%;:'\"\\/"
const char_alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const char_number = "0123456789"

type (
	PassInfo struct {
		CharacterSet string
		Length int64
	}
)

func getDictionary(characterSet string) string {
	switch characterSet {
	default:
	case "alphanumspec":
		return char_alpha + char_number + char_spec

	case "alphanum":
		return char_alpha + char_number

	case "alpha":
		return char_alpha

	case "number":
		return char_number
	}

	log.Fatal("Character set is not exist")
	return ""
}

func RandStr(passInfo PassInfo) string {
	dictionary := getDictionary(passInfo.CharacterSet)

	var bytes = make([]byte, passInfo.Length)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
