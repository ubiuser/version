package version

const Foo = "v2"

func Version() string {
	return Foo
}

func Submodule() string {
	return "submodule"
}
