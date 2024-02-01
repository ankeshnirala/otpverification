package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ankeshnirala/otpverification/data"
	"github.com/gin-gonic/gin"
)

var appTimeout = 10 * time.Second

func (app *Config) SendOTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			fmt.Println("Error: ", err)
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, "OTP sent succesfully")
	}
}

func (app *Config) VerifyOTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			fmt.Println("Error: ", err)
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, "OTP verified succesfully")

	}
}
