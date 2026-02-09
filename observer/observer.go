package observer

type Observer interface {
	Update(string, bool) bool
	GetID() string
}
