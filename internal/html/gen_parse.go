package html

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"stash.tcsbank.ru/a.krutyakov/gen_parse/internal/card"
	"time"
)

func New() *HTML {
	return &HTML{}
}

type HTML struct{}

func (x *HTML) Generate(c card.Card) ([]byte, error) {
	template := fmt.Sprintf( "<span>%s</span><time>%s</time><b>%s</b>", c.ID, c.T.Format(time.RFC3339), c.Name)
	return []byte(template), nil
}

func (x *HTML) Parse(b []byte, c *card.Card) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("%w: can't create new document from bytes", err)
	}

	c.ID = doc.Find("span").Text()
	c.T, _ = time.Parse(time.RFC3339, doc.Find("time").Text())
	c.Name = doc.Find("b").Text()

	return nil
}

func (x *HTML) Ext() string {
	return "html"
}
