package entities

type Node struct {
	Id    string `json:"id"`
	Label string `json:"label"`

	LinkedRepository *Repository `json:"-"`
}
