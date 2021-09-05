package search

type Sort struct {
	Field string `json:"field" db:"field"`
	Order Order  `json:"order" db:"order"`
}

type Order string

var ASC Order = "ASC"
var DESC Order = "DESC"

func (s *Sort) Validate() {
	switch s.Order {
	case ASC, DESC:
	default:
		s.Order = ASC
	}
}

func (s *Sort) ToSQL() string {
	if s.Field == "" {
		return ""
	}
	return " ORDER BY " + s.Field + " " + string(s.Order) + " "
}
