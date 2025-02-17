package snippets

import (
	"reflect"
	"testing"
	"time"
)

func TestSuccessfullyAddingNewSnippet(t *testing.T) {
	var mockSnippets Snippets = Snippets{
		SnippetsMap: SnippetsMap{
			"Snippet1": {
				Name:        "Snippet1",
				Value:       "Value1",
				Description: "Description1",
				ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	var snippetToAdd Snippet = Snippet{
		Name:        "Snippet2",
		Value:       "Value2",
		Description: "Description2",
		ModifiedOn:  time.Date(2025, 5, 2, 2, 0, 0, 0, time.UTC),
	}

	err := mockSnippets.SnippetsMap.AddSnippet(snippetToAdd)
	if err != nil {
		t.Errorf("Adding a snippet throws an error when it shouldn't.")
	}

	if addedSnippet, ok := mockSnippets.SnippetsMap[snippetToAdd.Name]; ok == false {
		t.Errorf("Failed to add the snippet to the map.")

		if !reflect.DeepEqual(snippetToAdd, addedSnippet) {
			t.Errorf("Snippet has been added incorrectly to the map.")
		}
	}
}

func TestAddingDuplicateSnippet(t *testing.T) {
	var mockSnippets Snippets = Snippets{
		SnippetsMap: SnippetsMap{
			"Snippet1": {
				Name:        "Snippet1",
				Value:       "Value1",
				Description: "Description1",
				ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	var snippetToAdd Snippet = Snippet{
		Name:        "Snippet1",
		Value:       "Value2",
		Description: "Description2",
		ModifiedOn:  time.Date(2025, 5, 2, 2, 0, 0, 0, time.UTC),
	}

	err := mockSnippets.SnippetsMap.AddSnippet(snippetToAdd)
	if err == nil {
		t.Errorf("Snippet was added while there was already a snippet with a similar key present.")
	}
}

func TestDeletingExistingSnippet(t *testing.T) {
	var mockSnippets Snippets = Snippets{
		SnippetsMap: SnippetsMap{
			"Snippet1": {
				Name:        "Snippet1",
				Value:       "Value1",
				Description: "Description1",
				ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			"Snippet2": {
				Name:        "Snippet2",
				Value:       "Value2",
				Description: "Description2",
				ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	err := mockSnippets.SnippetsMap.DeleteSnippet("Snippet2")
	if err != nil {
		t.Errorf("Removing a snippet throws an error when it shouldn't.")
	}

	if _, ok := mockSnippets.SnippetsMap["Snippet2"]; ok {
		t.Errorf("Failed to remove the snippet to the map.")
	}

	if len(mockSnippets.SnippetsMap) > 1 {
		t.Errorf("Too many snippets in the resulting snippets map")
	}
}

func TestDeletingNonExistingSnippet(t *testing.T) {
	var mockSnippets Snippets = Snippets{
		SnippetsMap: SnippetsMap{
			"Snippet1": {
				Name:        "Snippet1",
				Value:       "Value1",
				Description: "Description1",
				ModifiedOn:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	if err := mockSnippets.SnippetsMap.DeleteSnippet("Snippet2"); err == nil {
		t.Errorf("Expected an error since the snippet that's requested to be deleted, shouldn't exist.")
	}
}
