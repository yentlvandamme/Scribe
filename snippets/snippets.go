package snippets

import "time"

type Snippet struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"ModifiedOn"`
}

type Snippets struct {
	Version  string             `json:"version"`
	Snippets map[string]Snippet `json:"Snippets"`
}

func (snippetCollection *Snippets) AddSnippet(newSnippet Snippet) {
	snippetCollection.Snippets[newSnippet.Name] = newSnippet
}
