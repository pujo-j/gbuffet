package main

import (
	"github.com/pujo-j/gogae"
	"golang.org/x/oauth2/google"
	deploymentsAPI "google.golang.org/api/deploymentmanager/v2"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var project_template *template.Template

func init() {
	templateFile, err := os.Open("template.yaml")
	if err != nil {
		panic(err)
	}
	templateText, err := ioutil.ReadAll(templateFile)
	if err != nil {
		panic(err)
	}
	project_template, err = template.New("project").Parse(string(templateText))
	if err != nil {
		panic(err)
	}
}

func createProject(context gogae.RequestContext, request *ProjectRequest) (string, error) {
	yaml, err := genProjectYAML(config.AuthConfig.Project, config.BillingAccount, request)
	if err != nil {
		return "", err
	}
	client, err := google.DefaultClient(context.Request.Context(), deploymentsAPI.CloudPlatformScope)
	if err != nil {
		context.Log.WithError(err).Error("Creating a google API client")
		return "", err
	}
	context.Log.WithField("client", client).WithField("transport", client.Transport).Info("Creating deployment manager client")
	deployments, err := deploymentsAPI.New(client)
	if err != nil {
		return "", err
	}
	deployement := deploymentsAPI.Deployment{
		Name: request.ProjectId + "-create",
		Target: &deploymentsAPI.TargetConfiguration{
			Config: &deploymentsAPI.ConfigFile{
				Content: yaml,
			},
		},
	}
	operation, err := deployments.Deployments.Insert(config.AuthConfig.Project, &deployement).Do()
	if err != nil {
		return "", err
	}
	context.Log.WithField("result", operation).Info("Starting deployment manager")
	return operation.SelfLink, nil
}

func genProjectYAML(controlTower string, billing string, request *ProjectRequest) (string, error) {
	type vars struct {
		Project        *ProjectRequest
		BillingAccount string
		ControlTower   string
	}
	w := strings.Builder{}
	data := vars{
		Project:        request,
		BillingAccount: billing,
		ControlTower:   controlTower,
	}
	err := project_template.Execute(&w, data)
	if err != nil {
		return "", err
	} else {
		return w.String(), nil
	}
}
