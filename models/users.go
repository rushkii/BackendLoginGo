package models

type Users struct {
	ID        uint64 `gorm:"ID" json:"id"`
	Username  string `gorm:"username" json:"username"`
	Email     string `gorm:"email" json:"email"`
	Password  string `gorm:"password" json:"-"`
	Name      string `gorm:"name" json:"name"`
	Picture   string `gorm:"picture" json:"picture"`
	CreatedAt int64  `gorm:"created_at" json:"created_at"`
}

type JSONData struct {
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ToJSON(status uint16, message string, data interface{}) *JSONData {
	return &JSONData{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
