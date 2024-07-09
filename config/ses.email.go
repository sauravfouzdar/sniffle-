// configure SES email service

package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var (
	Region    = goDotEnvVariable("AWS_SES_REGION")
	AccessKey = goDotEnvVariable("AWS_ACCESS_KEY_ID")
	SecretKey = goDotEnvVariable("AWS_SECRET_ACCESS_KEY")
	Email     = goDotEnvVariable("EMAIL")
)

var svc *ses.SES

func InitSES() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	svc = ses.New(sess)
	log.Println("SES service connection established")
}

func GetSES() *ses.SES {
	return svc
}

// send email
func SendEmail(to, subject, body string) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(Email),
	}
	_, err := svc.SendEmail(input)
	if err != nil {
		println("Email not sent to: ", to)
		println("Error: ", err.Error())
		return err
	}
	println("Email sent to: ", to)
	return nil
}
