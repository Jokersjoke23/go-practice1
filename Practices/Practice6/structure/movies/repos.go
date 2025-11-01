package movies

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Movie, error) {
	rows, err := r.db.Query("SELECT id, title, year, actor_count FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Year, &m.ActorCount); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}
