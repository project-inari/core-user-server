package dto

// InquiryRes is a struct to hold the response data for user inquiry.
type InquiryRes struct {
	Username  string `json:"username"`
	UID       string `json:"uid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhoneNo   string `json:"phoneNo"`
	Email     string `json:"email"`
}
