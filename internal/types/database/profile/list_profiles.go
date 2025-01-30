package database

import (
	"time"

	"github.com/Madinab99999/blogging-platform/internal/types/shared"
)

type ProfileSortByField string

const (
	ProfilesSortByFieldUnspecified  ProfileSortByField = "unspecified"
	ProfilesSortByFieldFullName     ProfileSortByField = "full_name"
	ProfilesSortByFieldDateCreate   ProfileSortByField = "created_at"
	ProfilesSortByFieldLocation     ProfileSortByField = "location"
	ProfilesSortByFieldCountOfPosts ProfileSortByField = "count_of_posts"
)

var ProfileSortByFields = []ProfileSortByField{
	ProfilesSortByFieldUnspecified,
	ProfilesSortByFieldFullName,
	ProfilesSortByFieldLocation,
	ProfilesSortByFieldCountOfPosts,
}

type ListProfilesReqFilterItem interface {
	shared.FilterItem[ListProfilesReqFilterItem]
	FilterFullName() shared.StringExp
	FilterCreatedAt() shared.TimeExp
	FilterLocation() shared.StringExp
	FilterBio() shared.StringExp
}

type ListProfilesReq interface {
	shared.Pagination
	shared.SortBy[ProfileSortByField]
	shared.Filter[ListProfilesReqFilterItem]
}

type ListProfilesResp interface {
	GetList() []ItemProfileResp
}

type ItemProfileResp interface {
	GetUserID() string
	GetFullName() string
	GetLocation() string
	GetBio() string
	GetCountOfPosts() int
	GetCreatedAt() time.Time
}
