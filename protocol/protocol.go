package protocol

type Parser interface {
	Parse() error
}

type Statement interface {
}
