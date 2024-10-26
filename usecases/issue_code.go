package usecases

import (
	"fmt"
	"time"

	"github.com/reizt/rest-go/iservices"
	"github.com/reizt/rest-go/iservices/idatabase"
	"github.com/reizt/rest-go/iservices/imailer"
	i "github.com/reizt/rest-go/iusecases"
	"github.com/reizt/rest-go/usecases/id"
)

func IssueCode(s *iservices.All) i.IssueCode {
	return func(input i.IssueCodeInput) (*i.IssueCodeOutput, error) {
		// Generate code hash
		codeValue := id.GenerateCode()
		codeValueHash, err := s.Hasher.Hash(codeValue)
		if err != nil {
			return nil, err
		}

		// Save code to database
		now := time.Now().Unix()
		expiresAt := now + 60*60*24 // 1 day
		code := idatabase.Code{
			Id:        id.GenerateId(),
			Email:     input.Email,
			Action:    input.Action,
			ValueHash: codeValueHash,
			ExpiresAt: expiresAt,
			CreatedAt: now,
		}
		if err := s.Database.Code.Create(code); err != nil {
			return nil, err
		}

		// Send email
		mailerSendInput := imailer.SendInput{
			To:      input.Email,
			Subject: "Your code",
			Text:    codeValue,
			Html:    fmt.Sprintf("Your code is <code>%s</code>", codeValue),
		}
		s.Mailer.Send(mailerSendInput)

		// Return
		output := i.IssueCodeOutput{
			CodeId: code.Id,
		}
		return &output, nil
	}
}
