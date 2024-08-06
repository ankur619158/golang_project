package definition

import "gopkg.in/guregu/null.v4"

type Person struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

type GetPersonInfoResponse struct {
	Error  ErrorMsg `json:"error"`
	Person Person   `json:"topics"`
	Total  int64    `json:"total"`
}

type ErrorMsg struct {
	Message null.String `json:"message"`
}

type CreatePersonRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

type CreatePersonResponse struct {
	Error    ErrorMsg   `json:"error"`
	Success  SuccessMsg `json:"success"`
	PersonID int64      `json:"user_id"`
}

type SuccessMsg struct {
	Message null.String `json:"message"`
}
