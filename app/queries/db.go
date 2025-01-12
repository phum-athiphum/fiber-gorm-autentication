package queries

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// ฟังก์ชันสำหรับการกำหนดค่า db
func SetDB(database *gorm.DB) {
	db = database
}
