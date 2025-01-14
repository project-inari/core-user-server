package dto

// UserEntity represents the user entity in the database
type UserEntity struct {
	Username         string `sql:"username"`
	UID              string `sql:"uid"`
	FirstName        string `sql:"first_name"`
	LastName         string `sql:"last_name"`
	PhoneNo          string `sql:"phone_no"`
	Email            string `sql:"email"`
	SelectedLocale   string `sql:"selected_locale"`
	UserActive       bool   `sql:"user_active"`
	StripeCustomerID string `sql:"stripe_customer_id"`
	CreatedAt        string `sql:"created_at"`
	UpdatedAt        string `sql:"updated_at"`
}
