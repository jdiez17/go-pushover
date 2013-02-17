package pushover

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
)

const endpoint string = "https://api.pushover.net/1/messages.json"

type Pushover struct { 
    User, Apikey string
}

type Response struct {
    Status float64
    Errors []interface{} 
    Message string
}

func (p Pushover) Notify(title, message string) (*Response, error) {
    resp, err := http.PostForm(endpoint, 
        url.Values{"token": {p.Apikey}, "user": {p.User}, "message": {message}})

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
