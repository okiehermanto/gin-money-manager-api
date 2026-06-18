package options

type Filter struct {
	Field    string
	Operator string
	Value    any
}

type SearchOptions struct {
	Keyword string
	Fields  []string
}

type FindAllOptions struct {
	Relationships []string

	Filters []Filter

	Search  *SearchOptions
	OrderBy string

	Paginate bool
	Limit    int
	Page     int
}
