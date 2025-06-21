package models

// Role merepresentasikan tabel 'roles'
type Role struct {
	Base
	RoleName string `gorm:"type:varchar(50);uniqueIndex;not null" json:"role_name"` // cth: 'admin', 'teacher', 'student'
	Users    []User `gorm:"foreignKey:RoleID" json:"-"`                             // Relasi Has Many (opsional, jika perlu query sebaliknya)
}
