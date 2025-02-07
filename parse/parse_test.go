package parse

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var snippets Snippets = Snippets{
	Version: "1.0.0",
	Snippets: map[string]Snippet{
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
	snippetsInBytes, err := ParseToBytes(snippets)
	if err != nil {
		t.Errorf("Could not parse snippets during the setup stage.")
	}
	bufferReader := bytes.NewBuffer(snippetsInBytes)
	result, err := ParseJson(bufferReader)
	if err != nil {
		t.Errorf("Could not parse snippets: %s", err)
	}
	if !reflect.DeepEqual(snippets, result) {
		t.Errorf("Resulting struct does not match expected struct:\n%s\nVS\n%s", snippets, result)
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
	if reflect.DeepEqual(snippets, result) {
		t.Errorf("Values should not match:\n%s\nVS\n%s", snippets, result)
	}
}

func TestParseToBytes(t *testing.T) {
	result, err := ParseToBytes(snippets)
	if err != nil {
		t.Errorf("Could not parse snippets: %s", snippets)
	}

	expected, err := json.MarshalIndent(snippets, "", "\t")
	if err != nil {
		t.Errorf("Failed to marshal struct in order to prepare unit test: %s", t.Name())
	}

	if !bytes.Equal(result, expected) {
		t.Errorf("Resulting bytes do not match expected bytes:\n%s\nVS\n%s", result, expected)
	}
}
