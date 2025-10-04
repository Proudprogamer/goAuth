package handlers

import (
	"net/http"

	"github.com/Proudprogamer/goAuth/http/types"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"github.com/Proudprogamer/goAuth/prisma/db"
	"github.com/gin-gonic/gin"
)



type Handler struct {
	client *db.PrismaClient
}

func NewHandler(client *db.PrismaClient) *Handler {
	return &Handler {
		client : client,
	}
}

func (h *Handler) Home(ctx *gin.Context){
	ctx.String(200, "This is the home page")
}

func (h *Handler) SignUp(ctx *gin.Context){
	var UserReq types.User

	if err := ctx.ShouldBindJSON(&UserReq); err!=nil {
		verr:= err.(validator.ValidationErrors)
		ctx.JSON(500, gin.H{
			"error" :verr.Error(),
		})
		return
	}

	existingUser, err := h.client.Users.FindFirst(
		db.Users.Email.Equals(UserReq.Email),
	).Exec(ctx)

	if err==nil && existingUser!=nil {
		ctx.JSON(500, gin.H{
			"message" : "Email already exists",
		})
		return 
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(UserReq.Password), 10)

	if err!=nil {
		ctx.JSON(500, gin.H{
			"message" : "Failed to hash password",
		})
		return 
	}

	user, err:=h.client.Users.CreateOne(
		db.Users.Name.Set(UserReq.Name),
		db.Users.Email.Set(UserReq.Email),
		db.Users.Password.Set(string(hash)),
	).Exec(ctx)

	if err!=nil {
		ctx.JSON(500, gin.H{
			"message" : "Error in creating the user",
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "user created successfully", 
		"user" : gin.H{
			"id" : user.ID, 
			"name": user.Name, 
			"email": user.Email,
		},
	})
}