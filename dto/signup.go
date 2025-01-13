package dto

// SignUpReq represents the request to sign up a user
type SignUpReq struct {
	Username       string `json:"username"`
	UID            string `json:"uid"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	PhoneNo        string `json:"phoneNo"`
	Email          string `json:"email"`
	SelectedLocale string `json:"selectedLocale"`
}

// SignUpRes represents the response to sign up a user
type SignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Success  bool   `json:"success"`
}
