package shared

import "time"

type Filter[F any] interface {
	GetFilter() F
}

type FilterItem[F any] interface {
	FilterAnd() []F
	FilterOr() []F
	AddAnd(f F)
	AddOr(f F)
}

type StringExp interface {
	IdenticalExp[string]
	LikeExp[string]
}

type LikeExp[T any] interface {
	Like() *T
	Nlike() *T
	Ilike() *T
	Nilike() *T
}

type IdenticalExp[T any] interface {
	IsNullExp
	EqExp[T]
	InExp[T]
}

type EqExp[T any] interface {
	Eq() *T
	Neq() *T
}

type InExp[T any] interface {
	In() []T
	Nin() []T
}

type IsNullExp interface {
	IsNull() *bool
}

type QuantifyExp[T any] interface {
	Gt() *T
	Gte() *T
	Lt() *T
	Lte() *T
}

type ComparableExp[T any] interface {
	IdenticalExp[T]
	QuantifyExp[T]
}

type TimeExp interface {
	ComparableExp[time.Time]
	InRange(duration time.Duration) bool
	Since() time.Duration
	Until() time.Duration
}

type BoolExp interface {
	IsNullExp
	EqExp[bool]
}
