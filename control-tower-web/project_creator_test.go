package main

import (
	"testing"
	"time"
)

func TestGenerateProjectDefinition(t *testing.T) {
	request := ProjectRequest{
		ProjectId:      "sfeir-test-project-0001",
		ProjectName:    "Test Project",
		RequesterEmail: "pujo.j@gmail.com",
		RequestStatus:  GRANTED,
		Tags:           []string{"toto", "tata", "titi"},
		Folder:         "794931302318",
		Creation:       time.Now(),
	}
	yaml, err := genProjectYAML("control-tower-project", "billingAccounts/TEST_BILLING", request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(yaml)
}
