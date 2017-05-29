package slacker

import (
	"net/http"
	"bytes"
	"errors"
	"encoding/json"
)

type SlackBody struct {
	Text string `json:"text"`
}
// make json-formatted body for slack webhook post request
func MakeReqBody(message string) (*bytes.Buffer, error) {
	bodyStruct := SlackBody{
		Text: message,
	}
	bodyJson, err := json.Marshal(bodyStruct)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bodyJson), nil
}


func PostSlackMessage(message string, slackTarget string) error {
	// @TODO - how to add slackTarget to some config so that
		// calling library doesn't have to include it every time?

	reqBody, err := MakeReqBody(message)
	if err != nil {
		return err
	}
	response, err := http.Post(slackTarget, "application/json", reqBody)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return errors.New(response.Status)
	}
	return nil
}
