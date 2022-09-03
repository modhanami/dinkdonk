package domain

type Product struct {
	ID          uint
	Name        string
	Price       float64
	Description string
	ImageURL    string
	Categories  []Category
}

type Category string
