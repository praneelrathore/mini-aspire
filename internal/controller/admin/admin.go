package admin

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/personal/mini-aspire/internal/model"
	"github.com/personal/mini-aspire/internal/service/admin"
	"github.com/personal/mini-aspire/internal/service/mappers"
)

func RegisterAdmin(c *gin.Context) {
	var adminRegistrationRequest *mappers.AdminRegistrationRequest
	if err := c.ShouldBindJSON(&adminRegistrationRequest); err != nil {
		log.Printf("Error in binding user signup request: %v", err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - JSON binding error in incoming request"),
		)
		return
	}

	repository := model.NewDatabase(c)
	adminService := admin.NewService(repository)
	err := adminService.AddAdmin(c, adminRegistrationRequest)
	if err != nil {
		log.Printf("Error in signing up admin: %v", err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	c.JSON(http.StatusOK, mappers.StatusSuccess("Admin signed up successfully"))
}

func GetLoanRequests(c *gin.Context) {
	adminID := c.Query("admin_id")
	parsedAdminID, parsingError := strconv.ParseUint(adminID, 10, 64)
	if parsingError != nil {
		log.Printf("Error in parsing admin ID: %v", parsingError)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - Invalid admin ID"),
		)
		return
	}

	repository := model.NewDatabase(c)
	adminService := admin.NewService(repository)
	loanRequests, err := adminService.GetSubmittedLoanRequests(c, &mappers.AdminLoanRequest{AdminID: parsedAdminID})
	if err != nil {
		log.Printf("Error in getting loan requests: %v", err)
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	c.JSON(http.StatusOK, loanRequests)
}

func ApproveLoanRequest(ctx *gin.Context) {
	var loanUpdateRequest *mappers.AdminLoanUpdateRequest
	if err := ctx.ShouldBindJSON(&loanUpdateRequest); err != nil {
		log.Printf("Error in binding loan request: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			mappers.StatusFailed("Bad request - JSON binding error in incoming request"),
		)
		return
	}

	repository := model.NewDatabase(ctx)
	adminService := admin.NewService(repository)
	err := adminService.UpdateLoan(ctx, loanUpdateRequest)
	if err != nil {
		log.Printf("Error in updating loan request: %v", err)
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			mappers.StatusFailed("Unable to process the request."),
		)
		return
	}

	ctx.JSON(http.StatusOK, mappers.StatusSuccess("Loan request updated successfully"))
}
