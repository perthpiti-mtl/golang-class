package repository

type Database interface {
	Save(data []byte) error
}
