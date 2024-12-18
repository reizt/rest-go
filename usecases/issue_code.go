package usecases

import (
	"context"
	"fmt"
	"time"

	"reij.uno/entities"
	"reij.uno/iservices"
	"reij.uno/iservices/idatabase"
	"reij.uno/iservices/imailer"
	i "reij.uno/iusecases"
	"reij.uno/usecases/id"
)

func issueCode(s *iservices.All) i.IssueCode {
	return func(input i.IssueCodeInput, ctx context.Context) (*i.IssueCodeOutput, error) {
		if input.Action == entities.CodeActionCreateUser {
			user, _ := s.Database.User.GetByEmail(input.Email, ctx)
			if user != nil {
				return nil, i.ErrUserAlreadyExists
			}
		}

		// Generate code hash
		codeValue, err := id.GenerateCode()
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}
		codeValueHash, err := s.Hasher.Hash(codeValue)
		if err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
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
		if err := s.Database.Code.Create(code, ctx); err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}

		// Send email
		mailerInput := imailer.CodeInput{
			To:      input.Email,
			CodeId:  code.Id,
			Code:    codeValue,
			Expires: int64(code.ExpiresAt),
		}
		if err := s.Mailer.Code(mailerInput); err != nil {
			fmt.Println(err)
			return nil, i.ErrUnexpected
		}

		// Return
		output := i.IssueCodeOutput{
			CodeId: code.Id,
		}
		return &output, nil
	}
}
