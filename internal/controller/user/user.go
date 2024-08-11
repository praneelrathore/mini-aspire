package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/admin"
	"github.com/personal/mini-aspire/internal/service/loan"
	"github.com/personal/mini-aspire/internal/service/mappers"
	"github.com/personal/mini-aspire/internal/service/user"
)

func RegisterUser(c *gin.Context) {
	var userSignUpRequest *mappers.UserRegistrationRequest
	if err := c.ShouldBindJSON(&userSignUpRequest); err != nil {
		log.Printf("Error in binding user signup request: %v", err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - JSON binding error in incoming request"),
		)
		return
	}

	repository := model.NewDatabase(c)
	userService := user.NewService(repository)
	err := userService.AddUser(c, userSignUpRequest)
	if err != nil {
		log.Printf("Error in signing up user: %v", err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	c.JSON(http.StatusOK, mappers.StatusSuccess("User signed up successfully"))
}

func SubmitLoanRequest(ctx *gin.Context) {
	var loanRequest *mappers.LoanRequest
	if err := ctx.ShouldBindJSON(&loanRequest); err != nil {
		log.Printf("Error in binding loan request: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - JSON binding error in incoming request"),
		)
		return
	}

	repository := model.NewDatabase(ctx)
	loanService := loan.NewService(repository, user.NewService(repository), admin.NewService(repository))
	err := loanService.SubmitLoanRequest(ctx, loanRequest)
	if err != nil {
		log.Printf("Error in submitting loan request: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	ctx.JSON(http.StatusOK, mappers.StatusSuccess("Loan request submitted successfully"))
}

func GetLoanStatus(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	parsedUserID, parsingError := strconv.ParseUint(userID, 10, 64)
	if parsingError != nil {
		log.Printf("Error in parsing user ID: %v", parsingError)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - Invalid user ID"),
		)
		return
	}

	repository := model.NewDatabase(ctx)
	loanService := loan.NewService(repository, user.NewService(repository), admin.NewService(repository))
	loanStatus, err := loanService.GetLoans(ctx, parsedUserID)
	if err != nil {
		log.Printf("Error in getting loan status: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	ctx.JSON(http.StatusOK, loanStatus)
}

func RepayLoanInstallment(ctx *gin.Context) {
	var loanRequest *mappers.RepaymentRequest
	if err := ctx.ShouldBindJSON(&loanRequest); err != nil {
		log.Printf("Error in binding loan request: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - JSON binding error in incoming request"),
		)
		return
	}

	repository := model.NewDatabase(ctx)
	loanService := loan.NewService(repository, user.NewService(repository), admin.NewService(repository))
	err := loanService.RepayLoanInstallment(ctx, loanRequest)
	if err != nil {
		log.Printf("Error in submitting loan request: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	ctx.JSON(http.StatusOK, mappers.StatusSuccess("Loan installment repaid successfully"))
}
