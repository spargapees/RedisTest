package model

type Product struct {
	Id        int     `json:"id" db:"id"`
	Title     string  `json:"title" db:"title"`
	CompanyId int     `json:"company_id" db:"company_id"`
	Date      string  `json:"date" db:"date"`
	Rating    float32 `json:"rating" db:"rating"`
}

type Company struct {
	Name        string `json:"name" db:"name"`
	CreatedDate string `json:"created_date" db:"created_date"`
}
