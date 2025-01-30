package shared

type SortByDirection string

const (
	SortByDirectionUnspecified SortByDirection = "Unspecified"
	SortByDirectionAsc         SortByDirection = "Asc"
	SortByDirectionDesc        SortByDirection = "Desc"
)

var SortByDirections = []SortByDirection{
	SortByDirectionUnspecified,
	SortByDirectionAsc,
	SortByDirectionDesc,
}

type SortBy[Field ~string] interface {
	GetSortBy() []SortByItem[Field]
}

type SortByItem[Field ~string] interface {
	GetDirection() SortByDirection
	GetField() Field
	AddSortBy(field Field, direction SortByDirection) error
	ClearSortBy()
}
