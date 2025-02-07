package snippets

import (
	"reflect"
	"testing"
	"time"
)

func TestSuccessfullyAddingNewSnippet(t *testing.T) {
	var mockSnippets Snippets = Snippets{
		Version: "1.0.0",
		Snippets: map[string]Snippet{
			"Snippet1": {
				Name:        "Snippet1",
				Description: "Description1",
				ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	var snippetToAdd Snippet = Snippet{
		Name:        "Snippet2",
		Description: "Description2",
		ModifiedOn:  time.Date(2025, 5, 2, 2, 0, 0, 0, time.UTC),
	}

	mockSnippets.AddSnippet(snippetToAdd)

	if addedSnippet, ok := mockSnippets.Snippets[snippetToAdd.Name]; ok == false {
		t.Errorf("Failed to add the snippet to the map.")

		if !reflect.DeepEqual(snippetToAdd, addedSnippet) {
			t.Errorf("Snippet has been added incorrectly to the map.")
		}
	}
}
