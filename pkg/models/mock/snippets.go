package mock

import (
	"avgustavgust/snippetbox/pkg/models"
	"time"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetModel for testing
type SnippetModel struct{}

// Insert is used to mock data insertion to db
func (m *SnippetModel) Insert(string, string, string) (int, error) {
	return 2, nil
}

// Get is used to mock db response
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest is used to mock db response
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
