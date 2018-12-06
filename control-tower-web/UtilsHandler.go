package main

import (
	"github.com/julienschmidt/httprouter"
)

func RouteUtils(router *httprouter.Router, a *AuthMiddleware) {
	router.GET("/utils/isAdminUser", Handle(UtilsIsAdminUser, a))
	router.GET("/utils/isValidId", Handle(UtilsIsValidId, a))
}

func UtilsIsAdminUser(r RequestContext) (interface{}, int, error) {
	if r.UserData.IsAdmin {
		return true, 200, nil
	} else {
		return false, 200, nil
	}
}

func UtilsIsValidId(r RequestContext) (interface{}, int, error) {
	//TODO: Implement
	return true, 200, nil
}
