package interfaces

// typecover:MyInterface -excludeMethods
func excludeMethodsPass() {
	_ = &myStruct{}
}

// typecover:MyInterface -excludeMethods -exclude MyFunc1
func excludeMethodsWithExcludePass() {
	_ = &myStruct{}
}
