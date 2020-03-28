// Package jmdict helps with decoding JMdict files
package jmdict

import (
	"encoding/xml"
	"io"
)

// Decoder is the object to decode Dicts
type Decoder struct {
	d *xml.Decoder
}

// NewDecoder for decoding JMdict files
func NewDecoder(r io.Reader) *Decoder {
	// use Decoder instead of Unmarshal so we can disable strict
	var decoder Decoder
	decoder.d = xml.NewDecoder(r)
	decoder.d.Strict = false

	return &decoder
}

// dict decodes the entire dictionary at once
func (decoder *Decoder) dict(dict *Dict) error {
	return decoder.d.Decode(dict)
}

// Decode decodes the entire dictionary at once
func Decode(r io.Reader, dict *Dict) error {
	return NewDecoder(r).dict(dict)
}

// Entry parses a single Entry from the Dict
func (decoder *Decoder) Entry(entry *Entry) error {
	for {
		token, err := decoder.d.Token()

		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			if elem.Name.Local == "entry" {
				return decoder.d.DecodeElement(entry, &elem)
			}
		}
	}
}
