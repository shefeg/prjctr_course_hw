package types

type Child struct {
	Name string
	Age  int
}

type Couple struct {
	Child *Child
}
