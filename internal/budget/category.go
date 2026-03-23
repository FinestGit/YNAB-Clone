package budget

import (
	"errors"
	"strings"
)

type Category struct {
	ID   CategoryID
	Name string
}

var ErrEmptyCategoryName = errors.New("category name cannot be empty")

func NewCategory(id CategoryID, name string) (Category, error) {
	var categoryName string = strings.TrimSpace(name)
	if categoryName == "" {
		return Category{}, ErrEmptyCategoryName
	}
	return Category{
		ID:   CategoryID(id),
		Name: categoryName,
	}, nil
}
