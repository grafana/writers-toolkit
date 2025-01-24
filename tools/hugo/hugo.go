// Package hugo has functions for working with Hugo source files.
package hugo

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
	"gopkg.in/yaml.v3"
)

var errUnsupportedFrontMatterFormat = fmt.Errorf("unsupported front matter format")

// Unparse reverses the Hugo page parsing process.
// Front matter is not round tripped exactly.
// It's expected that some formatting tool is run on the output.
func Unparse(cfm pageparser.ContentFrontMatter) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	switch cfm.FrontMatterFormat {
	case metadecoders.YAML:
		if _, err := buf.Write([]byte("---\n")); err != nil {
			return buf.Bytes(), err
		}
		fm, err := yaml.Marshal(cfm.FrontMatter)
		if err != nil {
			return buf.Bytes(), err
		}
		if _, err := buf.Write(fm); err != nil {
			return buf.Bytes(), err
		}
		if _, err := buf.Write([]byte("---\n")); err != nil {
			return buf.Bytes(), err
		}
	case metadecoders.TOML:
		if _, err := buf.Write([]byte("+++\n")); err != nil {
			return buf.Bytes(), err
		}
		fm := bytes.NewBuffer([]byte{})
		if err := toml.NewEncoder(fm).Encode(cfm.FrontMatter); err != nil {
			return buf.Bytes(), err
		}
		if _, err := buf.Write(fm.Bytes()); err != nil {
			return buf.Bytes(), err
		}
		if _, err := buf.Write([]byte("+++\n")); err != nil {
			return buf.Bytes(), err
		}
	default:
		return buf.Bytes(), fmt.Errorf("%w: %s", errUnsupportedFrontMatterFormat, cfm.FrontMatterFormat)
	}

	if _, err := buf.Write(cfm.Content); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}
