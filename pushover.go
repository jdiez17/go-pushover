package pushover

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const endpoint string = "https://api.pushover.net/1/messages.json"

type Pushover struct {
	User, Apikey string
}

type Response struct {
	Status  float64
	Errors  []interface{}
	Message string
}

type Notification struct {
	Message, Title, Url, UrlTitle, Sound, Device string
	Timestamp                                    time.Time
	Priority                                     float64
}

func (n Notification) toValues(p Pushover) url.Values {
	return url.Values{
		"token":     {p.Apikey},
		"user":      {p.User},
		"message":   {n.Message},
		"title":     {n.Title},
		"url":       {n.Url},
		"url_title": {n.UrlTitle},
		"sound":     {n.Sound},
		"device":    {n.Device},
		"timestamp": {fmt.Sprintf("%d", n.Timestamp.Unix())},
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
	return response, err
}
