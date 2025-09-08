package structs

// typecover:MyStruct -excludeMethods
func testBlockWithExcludesMethodsPass() {
	m := MyStruct{}
	m.MyField1 = "cool"
	m.MyField2 = "hello"
	if true {
		m.MyField3 = "world"
	}
}

// typecover:MyStruct -excludeMethods -exclude MyField1, MyField2
func testBlockWithExcludesMethodsAndExcludePass() {
	m := MyStruct{}
	if true {
		m.MyField3 = "cool"
	}
}
