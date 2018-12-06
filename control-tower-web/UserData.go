package main

import (
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"strings"
)

type permissions struct {
	Permissions []string
}

type UserData struct {
	IsAdmin bool
}

func GetUserDataFromSession(session UserSession) (UserData, error) {
	res := UserData{}
	err := json.Unmarshal([]byte(session.UserData), &res)
	return res, err
}

func MakeUserDataFunction(project string) func(token *oauth2.Token) (string, error) {
	return func(token *oauth2.Token) (string, error) {
		result := UserData{}
		req2, _ := http.NewRequest("POST", "https://cloudresourcemanager.googleapis.com/v1/projects/"+project+":testIamPermissions", strings.NewReader("{\"permissions\":  [\"appengine.applications.get\"]}"))
		req2.Header.Set("Authorization", "Bearer "+token.AccessToken)
		response2, err := http.DefaultClient.Do(req2)
		defer response2.Body.Close()
		if err != nil {
			return "", err
		}
		if response2.StatusCode != 200 {
			log.WithField("status", response2.Status).Error("Testpermission request")
			data, err := ioutil.ReadAll(response2.Body)
			if err != nil {
				log.WithField("request", req2).WithError(err).Error("Reading response")
			}
			log.WithField("status", response2.Status).WithField("data", data).Error("Testpermission request")
			return "", errors.New(response2.Status)
		} else {
			data := permissions{}
			buffer, err := ioutil.ReadAll(response2.Body)
			if err != nil {
				log.WithError(err).Error("Reading permission response")
				return "", err
			}
			err = json.Unmarshal(buffer, &data)
			if err != nil {
				log.WithField("data", buffer).WithError(err).Error("Reading permission response")
				return "", err
			}
			if len(data.Permissions) > 0 {
				result.IsAdmin = true
			}
		}
		res, err := json.Marshal(result)
		return string(res), err
	}
}
