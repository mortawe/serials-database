package search

type Person struct {
	Name      string      `json:"name"`
	Birthdate BetweenDate `json:"birthdate"`
}

func (p *Person) ToSQL() string {
	query := ""
	if p.Name != "" {
		query += "name == :name "
	}
	if p.Birthdate.After != p.Birthdate.Before {
		query += "release BETWEEN :after AND :before "
	}
	if query != "" {
			query = "WHERE " + query
	}

	return query
}
