package snippet

import (
	"Creata21/snippetbox/pkg/models"
	"database/sql"
)

type SnippetStorage struct {
	db *sql.DB
}


func NewSnippetStorage(db *sql.DB) *SnippetStorage {
	return &SnippetStorage{
		db: db,
	}
}

func (s *SnippetStorage) Insert(title, content string) (int, error) {
	query := `INSERT INTO snippets (title, content, created) VALUES ($1, $2, CURRENT_TIMESTAMP) RETURNING id`

	var id int
	err := s.db.QueryRow(query, title, content).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SnippetStorage) Get(id int64) (*models.Snippet, error) {
	query := `SELECT id, title, content, created FROM snippets WHERE id = $1`

	row := s.db.QueryRow(query, id)
	snippet := &models.Snippet{}

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (s *SnippetStorage) Latest() ([]*models.Snippet, error) {
	query := `SELECT * FROM snippets ORDER BY created DESC LIMIT 10`

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
