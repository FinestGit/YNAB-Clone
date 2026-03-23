package budget

import (
	"errors"
	"testing"
)

func TestNewCategory_Success(t *testing.T) {
	categoryID, _ := NewCategoryID("test")
	var categoryName string = "test"
	newCategory, err := NewCategory(categoryID, categoryName)
	if err != nil {
		t.Fatal(err)
	}
	if newCategory.ID.String() != categoryID.String() {
		t.Fatalf("category id is %s, should be %s", newCategory.ID.String(), categoryID.String())
	}
	if newCategory.Name != categoryName {
		t.Fatalf("category name is %s, should be %s", newCategory.Name, categoryName)
	}
}

func TestNewCategory_EmptyName(t *testing.T) {
	categoryID, _ := NewCategoryID("test")
	newCategory, err := NewCategory(categoryID, " ")
	if !errors.Is(err, ErrEmptyCategoryName) {
		t.Fatalf("Category name is %s, should have received ErrEmptyCategoryName", newCategory.Name)
	}
}
