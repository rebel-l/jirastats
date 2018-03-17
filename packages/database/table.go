package database

type Table interface {
	CreateStructure() error
	Truncate() error
}
