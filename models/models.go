package models

type Model interface {
	TableName() string
	String() string
}
