package imailer

type Service interface {
	Send(input SendInput) error
}

type SendInput struct {
	To      string
	Subject string
	Text    string
	Html    string
}
