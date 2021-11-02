package lr

type LangSpecs struct {
	Go *GoSpec `yaml:"go"`
}

type GoSpec struct {
	Package    string            `yaml:"package"`
	Prefix     string            `yaml:"prefix"`
	ValueTypes map[string]string `yaml:"value_types"`
}
