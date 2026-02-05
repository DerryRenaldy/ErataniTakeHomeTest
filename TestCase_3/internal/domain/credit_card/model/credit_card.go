package model

type User struct {
	ID               int    `db:"id" json:"id"`
	Country          string `db:"country" json:"country"`
	CreditCardType   string `db:"credit_card_type" json:"credit_card_type"`
	CreditCardNumber string `db:"credit_card_number" json:"credit_card_number"`
	FirstName        string `db:"first_name" json:"first_name"`
	LastName         string `db:"last_name" json:"last_name"`
}

type UserTransaction struct {
	ID       int   `db:"id" json:"id"`
	IDUser   int   `db:"id_user" json:"id_user"`
	TotalBuy int64 `db:"total_buy" json:"total_buy"`
}

type TransactionWithUser struct {
	TransactionID int    `db:"id"`
	UserID        int    `db:"id_user"`
	TotalBuy      int64  `db:"total_buy"`
	UserFirstName string `db:"first_name"`
	UserLastName  string `db:"last_name"`
	UserCountry   string `db:"country"`
}

type CreateUserRequest struct {
	Country          string `json:"country" validate:"required"`
	CreditCardType   string `json:"credit_card_type" validate:"required"`
	CreditCardNumber string `json:"credit_card_number" validate:"required"`
	FirstName        string `json:"first_name" validate:"required"`
	LastName         string `json:"last_name" validate:"required"`
}

type CreateUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
