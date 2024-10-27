package token

type OTPTokenPayload struct {
	Email  string `json:"email"`
	Action string `json:"action"`
}

type LoginTokenPayload struct {
	UserId string `json:"userId"`
}
