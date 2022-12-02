package sms

import (
    "fmt"
    "log"

    "github.com/twilio/twilio-go"
)

// SendSms sends an SMS message to a list of phone numbers
func SendSms(accountSid, authToken, from string, to []string, body string) {
    // Create a new Twilio client
    client := twilio.NewClient(accountSid, authToken, nil)

    // Loop through the list of phone numbers and send the SMS message to each one
    for _, number := range to {
        message, err := client.Messages.SendMessage(from, number, body, nil)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Message sent to %s: %s\n", number, message.Sid)
    }
}
