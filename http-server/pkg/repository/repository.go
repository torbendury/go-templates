package repository

type Repository interface {
	Get()
	Create()
	Update()
	Delete()
}
