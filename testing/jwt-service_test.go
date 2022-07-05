package testing

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetSecretKey(t *testing.T){
	expSecretKey := os.Getenv("JWT_SECRET")  //return null
	actual := "nddat1811"
	exp2 := "nddat18113"
	fmt.Println("SOS: ", len(expSecretKey))
	fmt.Println("SOS2: ", os.Getenv("JWT_SECRET"))
	assert.Equal(t, expSecretKey, actual, "they should similar")
	assert.Equal(t, exp2, actual, "they should similar 222")
}

func TestGenerateToken(t *testing.T) {

}
// func (j *jwtService) GenerateToken(UserID string, UserName string, UserEmail string) string {
// 	claims := &jwtCustomClaim{
// 		UserID,
// 		UserName,
// 		UserEmail,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
// 			Issuer:    j.issuer,
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	t, err := token.SignedString([]byte(j.secretKey))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return t
// }
// func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
// 	jwtString := strings.Split(token, "Bearer ")[1]
// 	return jwt.Parse(jwtString, func(t_ *jwt.Token) (interface{}, error) {
// 		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
// 		}
// 		return []byte(j.secretKey), nil
// 	})
// }
