package json

import (
	"encoding/json"
	"stash.tcsbank.ru/a.krutyakov/gen_parse/internal/card"
)

func New() *JSON {
	return &JSON{}
}

type JSON struct{}

func (j *JSON) Generate(c card.Card) ([]byte, error) {
	return json.Marshal(c)
}

func (j *JSON) Parse(b []byte, c *card.Card) error {
	return json.Unmarshal(b, c)
}

func (j *JSON) Ext() string {
	return "json"
}
