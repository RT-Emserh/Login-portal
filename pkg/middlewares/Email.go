package middlewares

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
)

var verifier = emailverifier.
	NewVerifier().
	EnableSMTPCheck()

func EmailVerifier(name, email string) error {
	ret, err := verifier.CheckSMTP(email, "")
	if err != nil {
		fmt.Println("check smtp failed: ", err)
		return nil
	}

	fmt.Println("smtp validation result: ", ret)
	return nil
}
