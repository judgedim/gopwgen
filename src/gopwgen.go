package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strconv"
	"generator"
)

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

		passInfo := generator.PassInfo{c.String("set"), length}

		for i := 0; i < c.Int("num"); i++ {
			println(generator.RandStr(passInfo))
		}
	}

	app.Run(os.Args)
}
