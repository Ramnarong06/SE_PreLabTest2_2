package entity
import ( "gorm.io/gorm" )

// ชื่อต้องนำหน้าด้วยตัวใหญ่
type User struct {
	gorm.Model
	Username	string		`gorm:"uniqueIndex" valid:"required~Username is required"`
	Email		string		`gorm:"uniqueIndex" valid:"email~Email is Invalid"`
	Password	string		
}
