package models

type Usuario struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Nombre string `json:"nombre"`
    Email  string `json:"email" gorm:"unique"`
}
