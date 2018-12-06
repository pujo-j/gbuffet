package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"time"
)

func RouteProjectRequests(router *httprouter.Router, a *AuthMiddleware) {
	router.GET("/project_requests", Handle(ProjectRequestsGet, a))
	router.POST("/project_requests", Handle(ProjectRequestsPost, a))
	router.GET("/project_requests/:id", Handle(ProjectRequestGet, a))
	router.DELETE("/project_requests/:id", Handle(ProjectRequestDelete, a))
	router.PUT("/project_requests/:id", Handle(ProjectRequestPut, a))
}

func ProjectRequestsGet(r RequestContext) (interface{}, int, error) {
	if r.UserData.IsAdmin {
		pr, err := SelectProjectRequests(db, "")
		if err != nil {
			log.WithError(err).Error("Selecting project requests")
			return nil, 500, err
		}
		return pr, 200, nil
	} else {
		user, _ := r.JWTToken.Claims.Subject()
		pr, err := SelectProjectRequests(db, "WHERE requester_email=?", user)
		if err != nil {
			log.WithError(err).Error("Selecting project requests")
			return nil, 500, err
		}
		return pr, 200, nil
	}
}

func ProjectRequestsPost(r RequestContext) (interface{}, int, error) {
	if r.Request.Header.Get("Content-Type") != "application/json" {
		log.Error("Invalid content type in request")
		return nil, 400, nil
	}
	defer r.Request.Body.Close()
	pr := ProjectRequest{}
	err := json.NewDecoder(r.Request.Body).Decode(&pr)
	if err != nil {
		log.WithError(err).Error("Decoding json")
		return nil, 400, nil
	}
	requester, _ := r.JWTToken.Claims.Subject()
	// Copy only permitted fields
	pr2 := ProjectRequest{
		ProjectId:        pr.ProjectId,
		ProjectNumber:    pr.ProjectName,
		Tags:             pr.Tags,
		ExpectedLifetime: pr.ExpectedLifetime,
		RequesterComment: pr.RequesterComment,
		RequesterGroup:   pr.RequesterGroup,
		Creation:         time.Now(),
		RequestStatus:    NEW,
		RequesterEmail:   requester,
	}
	err = pr2.Save(db)
	if err != nil {
		log.WithField("project_request", pr2).WithError(err).Error("Saving new project request")
		return nil, 500, err
	}
	return r.redirect("/project_requests/"+pr2.ProjectId, 201)
}

func ProjectRequestGet(r RequestContext) (interface{}, int, error) {
	user, _ := r.JWTToken.Claims.Subject()
	id := r.Params.ByName("id")
	prs, err := SelectProjectRequests(db, "WHERE project_id=? LIMIT 1", id)
	if err != nil {
		log.WithField("id", id).WithError(err).Error("Selecting project request")
		return nil, 500, err
	}
	if len(prs) == 0 {
		return nil, 404, err
	}
	pr := prs[0]
	// Ok, we got the project request, check that the user can read it
	if !r.UserData.IsAdmin {
		//TODO: maybe check that user is in the requesting group if he is not the requesting user
		if !(pr.RequesterEmail == user) {
			return nil, 403, nil
		}
	}
	return prs[0], 200, nil
}

func ProjectRequestDelete(r RequestContext) (interface{}, int, error) {
	id := r.Params.ByName("id")
	if r.UserData.IsAdmin {
		prs, err := SelectProjectRequests(db, "WHERE project_id=?", id)
		if err != nil {
			log.WithError(err).Error("Loading project request")
			return nil, 500, err
		}
		if len(prs) == 0 {
			// No project for id...
			return nil, 404, nil
		}
		pr := prs[0]
		//TODO: Do something useful, like... deleting the GCP project before modifying the db
		deletion := time.Now()
		pr.ProjectDeletion = &deletion
		err = pr.Save(db)
		if err != nil {
			log.WithError(err).Error("Saving deleted project request")
			return nil, 500, err
		} else {
			return nil, 204, nil
		}
	} else {
		// No deletion by non admins, verboten
		log.Error("Non admin user tried to delete project request")
		return nil, 403, nil
	}
}

func ProjectRequestPut(r RequestContext) (interface{}, int, error) {
	id := r.Params.ByName("id")
	if r.UserData.IsAdmin {
		pr := ProjectRequest{}
		defer r.Request.Body.Close()
		err := json.NewDecoder(r.Request.Body).Decode(&pr)
		if err != nil {
			log.WithError(err).Error("Error decoding request")
			return nil, 400, err
		}
		// Check that projectId is consistent with URL...
		if !(pr.ProjectId == id) {
			log.Error("Trying to modify a project_id")
			return "Invalid project_id", 400, nil
		}
		err = pr.Save(db)
		if err != nil {
			log.WithError(err).Error("Saving project")
			return nil, 500, err
		}
		return nil, 202, nil
	} else {
		// No modification by non admins, verboten
		log.Error("Non admin user tried to update project request")
		return nil, 403, nil
	}
}
