package gopher

type Repository interface {
	GetGophers() ([]Gopher, error)
	GetGopher(id string) (Gopher, error)
}
