package dal

type Todo struct {
	ID    int
	Title string
	Done  bool `gorm:"default:false"` // Default value for the 'Done' field
}
