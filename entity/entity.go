package entity

type Product struct {
	ID        int64   `gorm:"primary_key" json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Rating    float64 `json:"rating"`
	TotalSold int64   `json:"total_sold"`
	Location  string  `json:"location"`
	Stock     int64   `json:"stock"`
	ImageURL  string  `json:"image_string"`
}

type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type Product struct {
// 	ID         int64    `gorm:"primary_key" json:"id"`
// 	Name       string   `json:"name"`
// 	Brand      string   `json:"brand"`
// 	Price      float64  `json:"price"`
// 	Seller     string   `json:"seller"`
// 	ImageURL   string   `json:"image_url"`
// 	CategoryID uint     `json:"-"`
// 	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
// }

// type Category struct {
// 	ID   int64  `gorm:"primary_key" json:"id"`
// 	Name string `json:"name"`
// }

// type Shopping_Cart struct {
// 	ID        int64   `gorm:"primary_key" json:"id"`
// 	UserID    uint    `json:"-"`
// 	ProductID uint    `json:"-"`
// 	User      User    `gorm:"foreignKey:UserID" json:"user"`
// 	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
// }
