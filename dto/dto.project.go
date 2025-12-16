package dto

// Project 每个是 platform (App、Saas、Admin、Manager) 之一
type Project struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	MenuCategory string `json:"menuCategory,omitempty"`
}
