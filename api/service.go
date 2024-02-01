package api

import (
	"errors"

	"github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envAccountSID(),
	Password: envAuthToken(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}

	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	response, err := client.VerifyV2.CreateVerification(envServiceSID(), params)
	if err != nil {
		return "", err
	}

	return *response.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}

	params.SetTo(phoneNumber)
	params.SetCode(code)

	response, err := client.VerifyV2.CreateVerificationCheck(envServiceSID(), params)
	if err != nil {
		return err
	}

	if *response.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}
