package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

type Room struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Price       int        `json:"price"`
	IsAvailable bool       `json:"is_available"`
	Facilities  []Facility `json:"facilities"`
}

type Reservation struct {
	ID                int    `json:"id"`
	CustomerName      string `json:"customer_name"`
	CustomerPhone     string `json:"customer_phone"`
	RoomID            int    `json:"room_id"`
	ReservationDate   string `json:"reservation_date"`
	TotalPayment      int    `json:"total_payment"`
	PaymentStatus     string `json:"payment_status"`
	ReservationStatus string `json:"reservation_status"`
	CreatedAt         string `json:"created_at"`
}

type Facility struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
