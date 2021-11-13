package models

import (
	"log"
)

type Category struct {
	CategoryCode string `db:"category_code" json:"category_code"`
	CategoryName string `db:"category_name" json:"category_name"`
}

var insertCategoryQuery = `
INSERT INTO categories VALUES(?, ?)`

var getAllCategoriesQuery = `
SELECT * FROM categories`

var updateCategoryQuery = `
UPDATE categories SET category_code = ?, category_name = ? WHERE category_code = ?;`

var deleteCategoryQuery = `
DELETE FROM categories WHERE category_code = ?;
`

func AddCategory(category *Category) bool {
	_, err := Db.Queryx(insertCategoryQuery, category.CategoryCode, category.CategoryName)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func EditCategory(category *Category, editId string) {
	_, err := Db.Queryx(updateCategoryQuery, category.CategoryCode, category.CategoryName, editId)
	if err != nil {
		log.Println(err)
	}
}

func DeleteCategory(deleteId string) {
	_, err := Db.Queryx(deleteCategoryQuery, deleteId)
	if err != nil {
		log.Println(err)
	}
}

func GetAllCategories(category *[]Category) {
	Db.Select(category, getAllCategoriesQuery)
}
