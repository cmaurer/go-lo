package types

// Operations struct
type Operations struct {
	Operations []Operation `json:"operations"`
}

// Datatype struct
type Datatype struct {
	DataTypeCap string `json:"data_type_cap"`
	DataType    string `json:"data_type"`
}

// Operation struct
type Operation struct {
	Name         string     `json:"name"`
	Template     string     `json:"template"`
	TestTemplate string     `json:"test_template"`
	Datatypes    []Datatype `json:"datatypes"`
}
