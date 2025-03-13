package dto

type UserEmailResponse struct {
	RegisterNo  string `json:"register_no,omitempty"`
	IdNo        string `json:"id_no"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Suffix      string `json:"suffix,omitempty"`
	FullName    string `json:"full_name"`
	Department  string `json:"department"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	EmailAction string `json:"email_action,omitempty"`
	SrsNo       string `json:"srs_no"`
}

type NewUserEmailResponse struct {
	RegisterNo string `json:"register_no"`
	Email      string `json:"email"`
}

type DeleteUserEmailResponse struct {
	IdNo       string `json:"id_no"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Suffix     string `json:"suffix,omitempty"`
	FullName   string `json:"full_name"`
	Department string `json:"department"`
}

type ByIdNoUserEmailResponse struct {
	IdNo        string `json:"id_no"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Suffix      string `json:"suffix"`
	Department  string `json:"department"`
	Email       string `json:"email"`
	SrsNo       string `json:"srs_no"`
	EmailAction string `json:"email_action"`
	DateCreated string `json:"date_created"`
	DateDeleted string `json:"deleted_at"`
}
