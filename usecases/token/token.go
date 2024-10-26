package token

type OTPTokenPayload struct {
	Email string `json:"email"`
}

type LoginTokenPayload struct {
	UserId string `json:"userId"`
}
