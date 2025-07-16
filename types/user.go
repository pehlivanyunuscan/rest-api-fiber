package types

type User struct {
	ID       uint   `json:"id"  gorm:"primarykey"` // Kullanıcı ID'si
	Email    string `json:"email" gorm:"unique"`   // Kullanıcı e-posta adresi, benzersiz olmalı
	Password string `json:"password"`              // Kullanıcı şifresi
}
