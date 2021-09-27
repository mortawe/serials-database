package search

type Person struct {
	Name      string      `json:"name"`
	Birthdate BetweenDate `json:"birthdate"`
	Awards    string      `json:"awards"`
}

func (p *Person) ToSQL() string {
	query := ""
	if p.Name != "" {
		query += "name ILIKE :name "
	}
	if p.Birthdate.After != p.Birthdate.Before {
		if query != "" {
			query += " AND "
		}
		query += "release BETWEEN :after AND :before "
	}
	if p.Awards != "" {
		if query != "" {
			query += " AND "
		}
		query += "awards ILIKE :awards"
	}
	if query != "" {
		query = "WHERE " + query
	}

	return query
}
