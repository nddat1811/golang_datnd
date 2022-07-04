package testing

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/mock"
)

type AuthService interface {
	Auth(email, pw string) (*jwt.Token, error)
}

func TestAuth(t *testing.T){
	fmt.Printf("mock.Anything: %v\n", mock.Anything)
	fmt.Printf("t: %v\n", t)
}