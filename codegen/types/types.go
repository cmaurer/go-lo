package types

// Operations struct
type Operations struct {
	Operations []Operation `json:"operations"`
}

// Datatype struct
type Datatype struct {
	DataTypeCap string `json:"data_type_cap"`
	DataType    string `json:"data_type"`
	TestDataP1  string `json:"test_data_p1"`
	TestDataP2  string `json:"test_data_p2"`
	IsPtrTest   bool   `json:"is_ptr_test"`
	PackageName string `json:"package_name,omitempty"`
}

// Operation struct
type Operation struct {
	Name            string     `json:"name"`
	Template        string     `json:"template"`
	TestTemplate    string     `json:"test_template"`
	BenchTemplate   string     `json:"bench_template"`
	ExampleTemplate string     `json:"example_template"`
	Datatypes       []Datatype `json:"datatypes"`
	Imports         []string   `json:"imports"`
}
