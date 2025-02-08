package parse

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/yentlvandamme/Scribe/snippets"
)

var mockSnippets snippets.Snippets = snippets.Snippets{
	Version: "1.0.0",
	SnippetsMap: snippets.SnippetsMap{
		"Snippet1": {
			Name:        "Snippet1",
			Description: "Description1",
			ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"Snippet2": {
			Name:        "Snippet2",
			Description: "Description2",
			ModifiedOn:  time.Date(2025, 5, 2, 2, 0, 0, 0, time.UTC),
		},
	},
}

func TestParseFromJson(t *testing.T) {
	snippetsInBytes, err := ParseToBytes(mockSnippets)
	if err != nil {
		t.Errorf("Could not parse snippets during the setup stage.")
	}
	bufferReader := bytes.NewBuffer(snippetsInBytes)
	result, err := ParseJson(bufferReader)
	if err != nil {
		t.Errorf("Could not parse snippets: %s", err)
	}
	if !reflect.DeepEqual(mockSnippets, result) {
		t.Errorf("Resulting struct does not match expected struct:\n%s\nVS\n%s", mockSnippets, result)
	}
}

func TestFailingParseFromJson(t *testing.T) {
	snippetsInBytes, err := ParseToBytes("Unparseable string")
	if err != nil {
		t.Errorf("Could not parse snippets during the setup stage.")
	}
	bufferReader := bytes.NewBuffer(snippetsInBytes)
	result, err := ParseJson(bufferReader)
	if err == nil {
		t.Errorf("Value was parsed while it shouldn't have been: %s", err)
	}
	if reflect.DeepEqual(mockSnippets, result) {
		t.Errorf("Values should not match:\n%s\nVS\n%s", mockSnippets, result)
	}
}

func TestParseToBytes(t *testing.T) {
	result, err := ParseToBytes(mockSnippets)
	if err != nil {
		t.Errorf("Could not parse snippets: %s", mockSnippets)
	}

	expected, err := json.MarshalIndent(mockSnippets, "", "\t")
	if err != nil {
		t.Errorf("Failed to marshal struct in order to prepare unit test: %s", t.Name())
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("Resulting bytes do not match expected bytes:\n%s\nVS\n%s", result, expected)
	}
}
