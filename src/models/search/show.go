package search

type Show struct {
	Title      string      `json:"title" db:"title"`
	Release    BetweenDate `json:"release" db:""`
	Genre      string      `json:"genre" db:"genre"`
	EpisodeNum int         `json:"episode_num" db:"episode_num"`
}

func (s *Show) ToSQL() string {
	query := ""
	if s.Title != "" {
		s.Title = "%" + s.Title + "%"
		query += " title ILIKE :title "
	}
	if s.Release.Before != s.Release.After {
		if query != "" {
			query += " AND "
		}
		query += "release BETWEEN :after AND :before "
	}
	if s.Genre != "" {
		if query != "" {
			query += " AND "
		}
		query += "genre ILIKE :genre "
	}
	if s.EpisodeNum > 0 {
		if query != "" {
			query += " AND "
		}
		query += "episode_num = :episode_num"
	}
	if query != "" {
		query = "WHERE " + query
	}
	return query
}
