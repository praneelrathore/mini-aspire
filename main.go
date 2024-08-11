package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/personal/mini-aspire/internal/appInit"
	"github.com/personal/mini-aspire/internal/controller/admin"
	"github.com/personal/mini-aspire/internal/controller/user"
)

func main() {
	ctx := context.TODO()
	gin.SetMode(gin.ReleaseMode)
	appInit.InitializeConfig()
	db := appInit.InitializeDatabase(ctx)
	env := appInit.NewEnv(
		appInit.WithDatabaseConnection(db),
	)
	nCtx := env.WithContext(ctx)
	// setup router and start server
	r := Initialize(nCtx)
	srv := &http.Server{
		Addr:         ":8081",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start with error : %v", err)
	}
}

// Initialize sets up the router and returns the gin engine
func Initialize(ctx context.Context) *gin.Engine {
	r := gin.New()
	r.ContextWithFallback = true
	v1 := r.Group("/v1")
	setupUserRoutes(v1)
	setupAdminRoutes(v1)
	return r
}

// setupUserRoutes sets up the user routes
func setupUserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/user")
	userRoutes.POST("/register", user.RegisterUser)
	userLoanRoutes := userRoutes.Group("/loan")
	userLoanRequestRoutes := userLoanRoutes.Group("/request")
	userLoanRequestRoutes.POST("/submit", user.SubmitLoanRequest)
	userLoanRequestRoutes.GET("/get", user.GetLoanStatus)
	userLoanRequestRoutes.POST("/repay", user.RepayLoanInstallment)
}

// setupAdminRoutes sets up the admin routes
func setupAdminRoutes(r *gin.RouterGroup) {
	adminRoutes := r.Group("/admin")
	adminRoutes.POST("/register", admin.RegisterAdmin)
	loanRoutes := adminRoutes.Group("/loan")
	loanRoutes.GET("/requests", admin.GetLoanRequests)
	loanRoutes.POST("/approve", admin.ApproveLoanRequest)
}
