package sms

import (
    "log"
    "os"

    "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// SendSms sends an SMS message to a list of phone numbers
func SendSms(to []string, body string) {
    // Create a new Twilio client
    client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
    params.SetBody(body)

    // Loop through the list of phone numbers and send the SMS message to each one
    for _, number := range to {
        params.SetTo(number)
        _, err := client.Api.CreateMessage(params)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Message sent to:", number)
    }
}
