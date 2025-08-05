package blog

type IDatabaseRepository interface {
	CreateOne()
	UpdateOne()
	DeleteOne()
	FindOne(id string)
	FindByProperty()
	DeleteByProperty()
	UpdateByProperty()
}
