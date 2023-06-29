package repogen

type Generator interface {
	Entity(entity any, opts ...func(*entityStruct)) (Entity, error)
}

type Entity interface {
	Method(name string) Method
}

type Method interface {
	Params() Returner
}

type Returner interface {
	Returns() SQLBuilder
}

type SQLBuilder interface {
	Select() SelectQuery
	Insert() InsertQuery
	Update() UpdateQuery
	Delete() DeleteQuery
}

type SelectQuery interface {
	Where() WhereClause
	Limit() LimitClause
	Offset() OffsetClause
	OrderBy() OrderByClause
}

type InsertQuery interface {
	Columns(columns ...string) InsertColumnsClause
	Values(values ...any) InsertValuesClause
}

type InsertColumnsClause interface {
	Values(values ...any) InsertValuesClause
}

type InsertValuesClause interface {
	Generate()
}

type DeleteQuery interface {
	Where() WhereClause
	Limit() LimitClause
	Offset() OffsetClause
}

type UpdateQuery interface {
	Where() WhereClause
	Limit() LimitClause
	Offset() OffsetClause
	Generate()
}

type WhereClause interface {
	Limit() LimitClause
	Offset() OffsetClause
	OrderBy() OrderByClause
	Generate()
}

type LimitClause interface {
	Offset() OffsetClause
	OrderBy() OrderByClause
	Generate()
}

type OffsetClause interface {
	OrderBy() OrderByClause
	Generate()
}

type OrderByClause interface {
	Asc(column string) OrderByClause
	Desc(column string) OrderByClause
	Generate()
}
