package hashpass

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) []byte {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hashedPass
}
