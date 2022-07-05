package testing

import (
	"testing"

	"github.com/nddat1811/simple_project_golang/entity"
	"github.com/stretchr/testify/mock"
	//"github.com/gin-gonic/gin"
)

// func SetUpRouter() *gin.Engine{
//     router := gin.Default()
//     return router
// }
type UserRepositoryMock struct {
	mock.Mock
}

func (r UserRepositoryMock) GetAllUser() []entity.User {
	//args := r.Called()
	users := []entity.User{
		{
			ID:       0,
			Name:     "tien",
			Email:    "tien@gmail.com",
			Password: "tien",
			Token:    "asdazxvvzxasdasdsad",
			Books:    &[]entity.Book{},
		},
		{
			ID:       1,
			Name:     "phu",
			Email:    "phu@gmail.com",
			Password: "phu",
			Token:    "asdazxvvzxasdasc1dasdsad",
			Books:    &[]entity.Book{},
		},
		{
			ID:       0,
			Name:     "quy",
			Email:    "quy@gmail.com",
			Password: "quy",
			Token:    "asdazxvvzxasdasdsad",
			Books:    &[]entity.Book{},
		},
	}
	return users
}

func TestInsertUser(t testing.T) {
	//r := SetUpRouter()
	arg := entity.User{
		ID:       1,
		Name:     "phu",
		Email:    "phu@gmail.com",
		Password: "phu",
		Token:    "asdazxvvzxasdasc1dasdsad",
		Books:    &[]entity.Book{},
	}
	arg.Email = "phu@gmail"
}
func TestProfileUser(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("GetAllUser").Return([]entity.User{})

	//service := service.UserService{UserID}

}
