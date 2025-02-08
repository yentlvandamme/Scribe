package snippets

import "time"

type Snippet struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"ModifiedOn"`
}

type SnippetsMap map[string]Snippet

type Snippets struct {
	Version     string      `json:"version"`
	SnippetsMap SnippetsMap `json:"Snippets"`
}

func (snippetCollection *SnippetsMap) AddSnippet(newSnippet Snippet) {
	(*snippetCollection)[newSnippet.Name] = newSnippet
}
