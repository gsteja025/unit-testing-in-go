package schema

import "github.com/jinzhu/gorm"

type Variant struct {
	gorm.Model
	ProductID uint   `json:"productid"`
	Color     string `json:"color"`
	Image     string `sql:"type:VARCHAR(255)" json:"image"`
}

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	Image       string `sql:"type:VARCHAR(255)" json:"image"`
	Variants    []Variant
	Myratings   []Rating
}

type Rating struct {
	gorm.Model
	ProductID uint   `json:"productid"`
	Review    string `json:"review"`
	Ratin     uint
}

func createDB(db *gorm.DB) *gorm.DB {

	db.DropTable(&Product{})
	db.CreateTable(&Product{})
	db.DropTable(&Rating{})
	db.CreateTable(&Rating{})
	db.DropTable(&Variant{})
	db.CreateTable(&Variant{})

	db.Save(&Product{
		Name:        "Guitar",
		Description: "Water resistent and made with fine wood",
		Category:    "Music Instrument",
		Quantity:    100,
		Price:       5430,
		Variants: []Variant{
			{Color: "Black"},
			{Color: "Brown"},
		},
		Myratings: []Rating{
			{Name: "GST",
				Review: "This guitar strings are so good",
				Ratin:  6,
			},
		},
	})
	return db

}
