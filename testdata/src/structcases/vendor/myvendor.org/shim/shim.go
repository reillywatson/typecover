// Package shim re-exports builder's types as aliases, the way a package that moved leaves a forwarding shim behind.
package shim

import "myvendor.org/builder"

type (
	Builder = builder.Builder
	Pair    = builder.Pair
)

func New() *Builder { return builder.New() }
