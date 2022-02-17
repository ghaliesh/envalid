package envalidator

type InputStuct struct {
	A string
	B int16
	C int64
	D string
}

type TestCase struct {
	input InputStuct
	keys  []string
	valus []interface{}
	types []string
}

var testcases []TestCase = []TestCase{
	{
		input: InputStuct{
			A: "1",
			B: 2,
			C: 20,
			D: "a"},
		valus: []interface{}{"1", 2, 20, "a"},
		keys:  []string{"A", "B", "C", "D"},
		types: []string{"string", "int16", "int64", "string"},
	},
}
