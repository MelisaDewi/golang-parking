package parking

type Observer interface {
	Update(string, bool) bool
	GetID() string
}
