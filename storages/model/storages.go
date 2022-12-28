package model

type Storage interface {
	Add(data any) (index int64, err error)
	Delete(index int64) (ok bool)
	Print()
	Get(index int64) (data any)
	Sort(more func(i, j any) bool)
}

func Add(storage Storage, data any) (index int64, err error) {
	return storage.Add(data)
}

func Delete(storage Storage, index int64) {
	storage.Delete(index)
}

func Print(storage Storage) {
	storage.Print()
}

func Get(storage Storage, index int64) (data any) {
	return storage.Get(index)
}

func Sort(storage Storage, more func(i, j any) bool) {
	storage.Sort(more)
}
