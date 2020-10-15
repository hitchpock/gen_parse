package html

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"stash.tcsbank.ru/a.krutyakov/gen_parse/internal/card"
)

func New() *HTML {
	return &HTML{}
}

type HTML struct{}

func (x *HTML) Generate(c card.Card) ([]byte, error) {
	template := fmt.Sprintf("<p>ID<span>%s</span></p>", c.ID)
	return []byte(template), nil
}

func (x *HTML) Parse(b []byte, c *card.Card) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("%w: can't create new document from bytes", err)
	}

	c.ID = doc.Find("span").Text()

	return nil
}

func (x *HTML) Ext() string {
	return "html"
}
