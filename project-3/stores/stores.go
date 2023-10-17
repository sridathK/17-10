package stores

type Storer interface {
	Create(usr User) error
	Update(usr User) error
	Delete(usr User) error
}

var StorerInterface Storer

type Service struct {
	Storer
}

// instaed of global variable.

func NewService(storer Storer) Service {
	s := Service{Storer: storer}
	return s
}
