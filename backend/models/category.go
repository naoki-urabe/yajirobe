package models

type Category struct {
	CategoryCode string `db:"category_code" json:"category_code"`
	CategoryName string `db:"category_name" json:"category_name"`
}

var insertCategoryQuery = `
INSERT INTO categories VALUES(?, ?)`

var getAllCategoriesQuery = `
SELECT * FROM categories`

func AddCategory(category *Category) {
	Db.MustExec(insertCategoryQuery, category.CategoryCode, category.CategoryName)
}

func GetAllCategories(category *[]Category) {
	Db.Select(category, getAllCategoriesQuery)
}
