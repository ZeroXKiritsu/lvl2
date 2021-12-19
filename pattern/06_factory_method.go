package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type creater interface {
	createProduct(owner string) User
	registerProduct(User)
}

type User interface {
	Use() string
}

type Factory struct {
}

func (self *Factory) Create(factory creater, owner string) User {
	user := factory.createProduct(owner)
	factory.registerProduct(user)
	return user
}

type IDCard struct {
	owner string
}

func (self *IDCard) Use() string {
	return self.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (self *IDCardFactory) createProduct(owner string) User {
	return &IDCard{owner}
}

func (self *IDCardFactory) registerProduct(product User) {
	owner := product.Use()
	self.owners = append(self.owners, &owner)
}