package pushover

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const endpoint string = "https://api.pushover.net/1/messages.json"

var PushoverError = errors.New("PushoverError")

type Pushover struct {
	UserKey, AppKey string
}

type Response struct {
	Status  int
	Errors  []interface{}
	Message string
}

type Notification struct {
	Message, Title, Url, UrlTitle, Sound, Device, Callback string
	Timestamp time.Time
	Priority, Retry, Expire int
}

func (n Notification) toValues(p Pushover) url.Values {
	return url.Values{
		"user":      {p.UserKey},
		"token":     {p.AppKey},
		"message":   {n.Message},
		"title":     {n.Title},
		"url":       {n.Url},
		"url_title": {n.UrlTitle},
		"sound":     {n.Sound},
		"device":    {n.Device},
		"timestamp": {fmt.Sprintf("%d", n.Timestamp.Unix())},
		"priority":  {fmt.Sprintf("%d", n.Priority)},
		"retry":     {fmt.Sprintf("%d", n.Retry)},
		"expire":    {fmt.Sprintf("%d", n.Expire)},
		"callback":  {n.Callback},
	}
}

func (p Pushover) Notify(n Notification) (*Response, error) {
	resp, err := http.PostForm(endpoint, n.toValues(p))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(Response)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response, PushoverError
}
