package slacklogging

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/nlopes/slack"
)

type SlackLogging struct {
	channel string
	slackAPI *slack.Client
	IsTesting bool
}

func New(slackChannel string, slackKey string) SlackLogging {

	if len(slackChannel) <= 0 {
		log.Fatal("Missing slack channel")
	}

	//slackKey := os.Getenv("SLACK_KEY")

	if len(slackKey) <= 0 && slackChannel != "test" {
		log.Fatal("Missing slack key")
	}

	return SlackLogging{
		slackAPI:slack.New(slackKey),
		channel:slackChannel,
		IsTesting:slackChannel == "test",
	}
}

func NewTest() SlackLogging {
	return SlackLogging{
		IsTesting:true,
	}
}

// LogErrorSlack Basic slack loggers
func (service SlackLogging) LogErrorSlack(message string, request *http.Request) {

	// Ignore if testing
	if service.IsTesting {
		return
	}

	attachment := slack.MsgOptionAttachments(slack.Attachment{
		Color:"danger",
		Text:formatRequest(request),
		Pretext:message,
	})

	_, _, err := service.slackAPI.PostMessage(service.channel, attachment)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (service SlackLogging) LogError(message string) {

	// Ignore if testing
	if service.IsTesting {
		return
	}

	attachment := slack.MsgOptionAttachments(slack.Attachment{
		Color:"danger",
		Text:message,
	})

	_, _, err := service.slackAPI.PostMessage(service.channel, attachment)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (service SlackLogging) LogWarningSlack(message string, request *http.Request) {

	// Ignore if testing
	if service.IsTesting {
		return
	}

	attachment := slack.MsgOptionAttachments(slack.Attachment{
		Color:"warning",
		Text:formatRequest(request),
		Pretext:message,
	})

	_, _, err := service.slackAPI.PostMessage(service.channel, attachment)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (service SlackLogging) LogInfoSlack(message string, request *http.Request) {

	// Ignore if testing
	if service.IsTesting {
		return
	}

	attachment := slack.MsgOptionAttachments(slack.Attachment{
		Color:"good",
		Text:formatRequest(request),
		Pretext:message,
	})

	_, _, err := service.slackAPI.PostMessage(service.channel, attachment)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func (service SlackLogging) SlackMessage(message string) {

	// Ignore if testing
	if service.IsTesting {
		return
	}

	attachment := slack.MsgOptionAttachments(slack.Attachment{
		Color:"good",
		Text:message,
	})

	_, _, err := service.slackAPI.PostMessage(service.channel, attachment)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

// LogHTTPRequestWarningSlack HTTP request slack loggers
func (service SlackLogging) LogHTTPRequestWarningSlack(reqType string, request *http.Request) {

	// Ignore if testing
	if service.IsTesting {
		return
	}

	if request == nil || len(reqType) <= 0{
		return
	}

	attachment := slack.MsgOptionAttachments(slack.Attachment{
		Color:"warning",
		Text:formatRequest(request),
		Pretext:reqType,
	})

	_, _, err := service.slackAPI.PostMessage(service.channel, attachment)

	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {

	// Create return string
	var request []string

	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	request = append(request, fmt.Sprintf("Client: %v", r.RemoteAddr))

	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	// Return the request as a string
	return strings.Join(request, "\n")
}