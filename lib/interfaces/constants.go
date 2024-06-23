package interfaces

type Enum interface {
	String() string
}

type Model interface {
	String() string
	Dict() map[string]interface{}
}
