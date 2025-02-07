package parse

import (
	"encoding/json"
	"io"

	"github.com/yentlvandamme/Scribe/snippets"
)

func ParseJson(reader io.Reader) (snippets.Snippets, error) {
	var snippetCollection snippets.Snippets

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&snippetCollection)
	if err != nil {
		return snippetCollection, err
	}

	return snippetCollection, nil
}

func ParseToBytes(data interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	return jsonData, err
}
