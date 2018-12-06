package main

import (
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"io/ioutil"
	"strings"
)

type permissions struct {
	Permissions []string
}

type UserData struct {
	IsAdmin bool
}

func UserDataParse(jsonString string) (UserData, error) {
	res := UserData{}
	err := json.Unmarshal([]byte(jsonString), &res)
	return res, err
}

func MakeUserDataFunction(project string) func(context context.Context, config *oauth2.Config, token *oauth2.Token) (string, error) {
	return func(context context.Context, config *oauth2.Config, token *oauth2.Token) (string, error) {
		result := UserData{}
		client := config.Client(context, token)
		response, err := client.Post("https://cloudresourcemanager.googleapis.com/v1/projects/"+project+":testIamPermissions", "application/json", strings.NewReader("{\"permissions\":  [\"appengine.applications.get\"]}"))
		if err != nil {
			return "", err
		}
		defer response.Body.Close()
		if err != nil {
			return "", err
		}
		if response.StatusCode != 200 {
			return "", errors.New(response.Status)
		} else {
			data := permissions{}
			buffer, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return "", err
			}
			err = json.Unmarshal(buffer, &data)
			if err != nil {
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
