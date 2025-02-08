package snippets

import (
	"errors"
	"time"
)

type Snippet struct {
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modifiedOn"`
}

type SnippetsMap map[string]Snippet

type Snippets struct {
	Version     string      `json:"version"`
	SnippetsMap SnippetsMap `json:"snippets"`
}

func (snippetCollection *SnippetsMap) AddSnippet(newSnippet Snippet) error {
	if _, ok := (*snippetCollection)[newSnippet.Name]; ok {
		return errors.New("Snippet already exists")
	}

	(*snippetCollection)[newSnippet.Name] = newSnippet
	return nil
}
