package internal

import (
	"context"
	"errors"
	"fmt"
)

type (
	CDPImplementation struct {
	}

	IdentifyUser struct {
		UserID string
	}
)

func NewCDPImplementation() *CDPImplementation {
	return &CDPImplementation{}
}

func (*CDPImplementation) Identify(ctx context.Context, user IdentifyUser) error {
	fmt.Println("user sent to CDP")

	if user.UserID == "123" {
		return nil
	}

	return errors.New("invalid user_id")
}
