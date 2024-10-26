package imailer

type Service interface {
	Code(input CodeInput) error
}

type CodeInput struct {
	To      string
	CodeId  string
	Code    string
	Expires int64
}
