package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Appointment struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	UserID      int64  `json:"user_id"`
}

type Payment struct {
	ID            int64   `json:"id"`
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
	PaymentMethod string  `json:"payment_method"`
	UserID        int64   `json:"user_id"`
}
