package interfaces

type Enum interface {
	String() string
}

type Model interface {
	TableName() string
	String() string
	Dict() map[string]interface{}
}
