package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"stash.tcsbank.ru/a.krutyakov/gen_parse/internal/card"
	"stash.tcsbank.ru/a.krutyakov/gen_parse/internal/file"
	internalHTML "stash.tcsbank.ru/a.krutyakov/gen_parse/internal/html"
	internalJSON "stash.tcsbank.ru/a.krutyakov/gen_parse/internal/json"
	internalXML "stash.tcsbank.ru/a.krutyakov/gen_parse/internal/xml"
)

type GenParser interface {
	Generate(c card.Card) ([]byte, error)
	Parse(b []byte, c *card.Card) error
	Ext() string
}

func main() {
	var gp GenParser
	c := card.Card{ID: "123456789", T: time.Now(), Name: "Anton"}

	gp = internalJSON.New()
	_ = internalXML.New()
	_ = internalHTML.New()

	out := file.NewFile(gp.Ext())
	//out := os.Stdout

	b, err := gp.Generate(c)
	if err != nil {
		log.Fatalf("can't generate object: %s", err)
	}

	_, err = io.WriteString(out, string(b))
	if err != nil {
		log.Fatalf("can't write object: %s", err)
	}

	var obj card.Card

	err = gp.Parse(b, &obj)
	if err != nil {
		log.Fatalf("can't parse object: %s", err)
	}

	fmt.Printf("%+v\n", obj)
}
