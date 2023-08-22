package protocol

type Generator interface {
	AddEntity(name, src string) error
	AddProcedure(procedure Procedure) error
}

type Parameter struct {
	Name string
	Type string
}

type Procedure struct {
	Name    string
	Inputs  []Parameter
	Outputs []Parameter
}
