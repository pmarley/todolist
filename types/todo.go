package types

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
