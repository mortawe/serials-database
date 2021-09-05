package search

type Show struct {
	Title    string      `json:"title"`
	Release  BetweenDate `json:"release"`
	GenreID  []int       `json:"genres"`
	PersonID []int       `json:"persons"`
}


func (s *Show) ToSQL() string {
	query := ""
	if s.Title != "" {
		s.Title = "%" + s.Title + "%"
		query += " title ILIKE :title "
	}
	if s.Release.Before != s.Release.After {
		query += "release BETWEEN :after AND :before "
	}
	if s.GenreID != nil && len(s.GenreID) > 0 {
		query += "genreID "
	}

	return query
}