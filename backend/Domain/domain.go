package domain

import "time"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type Patient struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	MedicalCase string    `json:"medical_case"`
	Note        string    `json:"note,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Braces      bool      `json:"braces"`
}

type Appointment struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	PatientID   int64     `json:"patient_id"`
	Status      string    `json:"status"`
	Reason      string    `json:"reason,omitempty"`
}

type Payment struct {
	ID            int64     `json:"id"`
	Amount        float64   `json:"amount"`
	Date          time.Time `json:"date"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	PatientID     int64     `json:"patient_id"`
}
