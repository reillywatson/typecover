package builder

type Builder struct{}

func New() *Builder { return &Builder{} }

func (b *Builder) SetA(string) *Builder { return b }

func (b *Builder) SetB(string) *Builder { return b }

type Pair struct {
	First  string
	Second string
}
