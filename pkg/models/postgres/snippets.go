package postgres

import (
	"Creata21/snippetbox/pkg/models"
	"database/sql"
	"fmt"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	query := `INSERT INTO snippets (title, content, created) VALUES ($1, $2, CURRENT_TIMESTAMP)`

	var id int
	err := m.DB.QueryRow(query, title, content).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Println(id)
	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	query := `SELECT id, title, content, created FROM snippets WHERE id = $1`

	row := m.DB.QueryRow(query, id)
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	query := `SELECT * FROM snippets ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(query)

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
