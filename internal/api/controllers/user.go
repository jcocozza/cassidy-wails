package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	userrepo "github.com/jcocozza/cassidy-wails/internal/repository/userRepo"
	ctxutil "github.com/jcocozza/cassidy-wails/internal/utils/ctxUtil"
	"github.com/jcocozza/cassidy-wails/internal/utils/dateutil"
	"github.com/jcocozza/cassidy-wails/internal/utils/uuidgen"
)

type UserHandler struct {
	UserRepository userrepo.UserRepository
}

func NewUserHandler(db database.DbOperations) *UserHandler {
	return &UserHandler{
		UserRepository: userrepo.NewIUserRespository(db),
	}
}

// Create a user
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var createRequest model.User
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := uh.UserRepository.Read(createRequest.Username)

	if err != nil {
		// if there is no user
		if errors.Is(err, sql.ErrNoRows) {
			uuid := uuidgen.GenerateUUID()
			createRequest.SetUuid(uuid)
			err2 := createRequest.Validate()
			if err2 != nil {
				ctx.JSON(http.StatusBadRequest, "user is invalid")
				return
			}
			uh.UserRepository.Create(&createRequest)
			ctx.JSON(http.StatusAccepted, createRequest)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong creating the user: " + err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusConflict, "user already exists")
		return
	}
}

// the frontend will send this struct to us
type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// authenticate a user
func (uh *UserHandler) AuthenticateUser(ctx *gin.Context) {
	var authRequest authRequest
	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr, err1 := uh.UserRepository.Read(authRequest.Username)

	if err1 != nil {
		ctx.JSON(http.StatusNotFound, err1.Error())
		return
	}

	if usr.Password != authRequest.Password {
		ctx.JSON(http.StatusUnauthorized, "incorrect password")
	} else if usr.Password == authRequest.Password {
		ctx.JSON(http.StatusOK, usr)
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, "unknown authorization failure")
		return
	}
}

// Return the start date and end date of the current microcycle
func (uh *UserHandler) GetMicrocycleCurrentDates(ctx *gin.Context) {
	currentDate := time.Now().Format(dateutil.Layout)
	uuid := ctxutil.GetUserUuidFromHeader(ctx)

	cycleStart, cycleDays, initialStartDate, err := uh.UserRepository.ReadPreferences(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var mc []*dateutil.DateObject
	if cycleDays == 7 {
		mc = dateutil.GetDateMicrocycle(currentDate, cycleStart, cycleDays)
	} else {
		mc = dateutil.GetCurrentCycleFromInitialDate(initialStartDate, cycleDays)
	}

	d := gin.H{
		"start_date": mc[0].Date,
		"end_date":   mc[len(mc)-1].Date,
	}

	ctx.JSON(http.StatusOK, d)
}

// Update the user preferences
func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var updateRequest model.User
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uh.UserRepository.Update(&updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, updateRequest)
}
