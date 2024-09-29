package models

type Book struct {
        id              int64   `gorm:"PrimaryKey" json:"id"`
        Title           string  `gorm:"type:varchar(255)" json:"title"`
        Description     string  `gorm:"type:text" json:"description"`
        Autor           string  `gorm:"type:varchar(255)" json:"autor"`
        PublishDate     string  `gorm:"type:date" json:"publish_date"`
}

