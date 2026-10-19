package structcases

import (
	"myvendor.org/builder"
	"myvendor.org/shim"
)

// Chained calls return *builder.Builder, while the annotation names the alias.
//typecover:shim.Builder
func testAliasTargetChainedCalls() {
	shim.New().SetA("a").SetB("b")
}

//typecover:shim.Builder
func testAliasTargetMissingMethod() { // want `Type myvendor.org/shim.Builder is missing SetB`
	shim.New().SetA("a")
}

// The variable has the alias type, while the annotation names the aliased type.
//typecover:builder.Builder
func testAliasedVariable() {
	b := shim.New()
	b.SetA("a")
	b.SetB("b")
}

//typecover:builder.Pair
func testCompositeLiteralThroughAlias() {
	_ = shim.Pair{
		First:  "",
		Second: "",
	}
}

//typecover:shim.Pair
func testCompositeLiteralOfAliasedType() { // want `Type myvendor.org/shim.Pair is missing Second`
	_ = builder.Pair{
		First: "",
	}
}
