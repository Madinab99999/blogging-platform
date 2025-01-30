package shared

type Pagination interface {
	GetCursor() *int
	GetLimit() int
}
