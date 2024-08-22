package utils

type Service[T any, N any] interface {
	Create(dto *T) error
	GetByID(id uint) (*N, error)
	Update(dto *T) error
	Delete(id uint) error
}
