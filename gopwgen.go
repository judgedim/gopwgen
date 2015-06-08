package main

import (
	"crypto/rand"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strconv"
)

const char_spec = "(){}[]<>?!`~*#$^%;:'\"\\/"
const char_alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const char_number = "0123456789"

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

func randStr(strSize int64, characterSet string) string {
	dictionary := getDictionary(characterSet)

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func check(s string) (i int64) {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func main() {
	app := cli.NewApp()
	app.Name = "passgener"
	app.Usage = "create highly secure passwords"
	app.Version = "0.0.1"
	app.Author = "judgedim"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "set",
			Value: "alphanumspec",
			Usage: "character set, available: alphanumspec, alphanum, alpha, number",
		},
		cli.IntFlag{
			Name:  "num",
			Value: 1,
			Usage: "Num. Passwords",
		},
	}
	app.Action = func(c *cli.Context) {
		var length int64
		length = 16
		if len(c.Args()) > 0 {
			length = check(c.Args().First())
		}

		for i := 0; i < c.Int("num"); i++ {
			println(randStr(length, c.String("set")))
		}
	}

	app.Run(os.Args)
}
