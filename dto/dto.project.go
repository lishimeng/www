package dto

type Project struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	MenuCategory string `json:"menuCategory,omitempty"`
}
