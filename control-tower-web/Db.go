package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RequestStatus string

const (
	NEW      RequestStatus = "NEW"
	GRANTED  RequestStatus = "GRANTED"
	REJECTED RequestStatus = "REJECTED"
)

type ProjectRequest struct {
	ProjectId        string        `json:"project_id"`
	ProjectName      string        `json:"project_name"`
	ProjectNumber    string        `json:"project_number"`
	RequesterEmail   string        `json:"requester_email"`
	RequesterGroup   *string       `json:"requester_group"`
	ExpectedLifetime uint64        `json:"expected_lifetime"`
	RequestStatus    RequestStatus `json:"request_status"`
	Creation         time.Time     `json:"creation"`
	ProjectCreation  *time.Time    `json:"project_creation"`
	ProjectDeletion  *time.Time    `json:"project_deletion"`
	RequesterComment string        `json:"requester_comment"`
	AdminComment     string        `json:"admin_comment"`
	Folder           string        `json:"folder"`
	Tags             []string      `json:"tags"`
	Allocations      []Allocation  `json:"allocations"`
}

func SelectProjectRequests(db *sql.DB, clause string, args ...interface{}) ([]ProjectRequest, error) {
	rs, err := db.Query("SELECT "+pr_Column_List+" FROM project_request "+clause, args...)
	if err != nil {
		return nil, err
	}
	return scanProjectRequest(rs)
}

func (p *ProjectRequest) Save(db *sql.DB) error {
	//TODO: Save Allocations
	tagsJson, _ := json.Marshal(p.Tags)
	_, err := db.Exec(
		"REPLACE INTO project_request ("+pr_Column_List+") VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,CAST(? AS JSON))",
		p.ProjectId,
		p.ProjectName,
		p.ProjectNumber,
		p.RequesterEmail,
		p.RequesterGroup,
		p.ExpectedLifetime,
		p.RequestStatus,
		p.Creation,
		p.ProjectCreation,
		p.ProjectDeletion,
		p.RequesterComment,
		p.AdminComment,
		p.Folder,
		tagsJson)
	return err
}

func scanProjectRequest(rs *sql.Rows) ([]ProjectRequest, error) {
	result := make([]ProjectRequest, 0, 16)
	for rs.Next() {
		pr := ProjectRequest{}
		tagsJson := ""
		if err := rs.Scan(
			&pr.ProjectId,
			&pr.ProjectName,
			&pr.ProjectNumber,
			&pr.RequesterEmail,
			&pr.RequesterGroup,
			&pr.ExpectedLifetime,
			&pr.RequestStatus,
			&pr.Creation,
			&pr.ProjectCreation,
			&pr.ProjectDeletion,
			&pr.RequesterComment,
			&pr.AdminComment,
			&pr.Folder,
			&tagsJson,
		); err != nil {
			return nil, err
		}
		err := json.Unmarshal([]byte(tagsJson), &pr.Tags)
		if err != nil {
			return nil, err
		}
		result = append(result, pr)
	}
	return result, nil
}

type Allocation struct {
}

const pr_Column_List = `project_id,
project_name,
project_number,
requester_email,
requester_group,
expected_lifetime,
request_status,
creation,
project_creation,
project_deletion,
requester_comment,
admin_comment,
folder,
tags`
