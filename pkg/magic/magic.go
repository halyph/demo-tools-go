package magic

//go:generate mockery --exported --all --dir ./ --case=snake --outpkg=mocks
type IFoo interface {
	Process(input int) int
}

type Foo struct {
}

func (Foo) Process(input int) int {
	return input * 10
}
