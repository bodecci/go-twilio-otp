package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioAPI.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms") // using sms here. you can use whatsapp and other forms

	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioAPI.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}
