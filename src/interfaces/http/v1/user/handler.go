package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hieuphq/califit/src/domain/model"
	"github.com/hieuphq/califit/src/usecase"
)

// Handler handler for user
type Handler struct {
	userUC usecase.UserUsecase
}

// CreateHandler ..
func CreateHandler(userUC usecase.UserUsecase) *Handler {
	return &Handler{
		userUC: userUC,
	}
}

// CreateUserHandler create a new user handler
func (h *Handler) CreateUserHandler(c *gin.Context) {
	var json CreateUserRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userUC.RegisterUser(model.User{Name: json.Name, Email: json.Email, HashedPassword: json.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, CreateUserResponse{
		ID:    string(user.ID),
		Email: user.Email,
		Name:  user.Name,
	})
}
