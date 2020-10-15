package xml

import (
	"encoding/xml"

	"stash.tcsbank.ru/a.krutyakov/gen_parse/internal/card"
)

func New() *XML {
	return &XML{}
}

type XML struct{}

func (x *XML) Generate(c card.Card) ([]byte, error) {
	return xml.Marshal(c)
}

func (x *XML) Parse(b []byte, c *card.Card) error {
	return xml.Unmarshal(b, c)
}

func (x *XML) Ext() string {
	return "xml"
}
