package mysql

import (
	"avgustavgust/snippetbox/pkg/models"
	"database/sql"
)

// SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert used to insert record to DB
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get used to get record by id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest return 10 most recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
