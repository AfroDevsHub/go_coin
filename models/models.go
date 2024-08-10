package models

type Model interface {
	TableName() string
	String() string
	Dict() map[string]interface{}
}
