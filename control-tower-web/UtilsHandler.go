package main

import (
	"github.com/pujo-j/gogae"
)

func RouteUtils(g gogae.Gogae) {
	g.Router.GET("/utils/isAdminUser", g.Handle(UtilsIsAdminUser))
	g.Router.GET("/utils/isValidId", g.Handle(UtilsIsValidId))
}

func UtilsIsAdminUser(r gogae.RequestContext) (interface{}, int, error) {
	ud, err := UserDataParse(r.UserDataJson)
	if err != nil {
		r.Log.WithError(err).Error("Parsing user data")
		return nil, 500, err
	}
	if ud.IsAdmin {
		return true, 200, nil
	} else {
		return false, 200, nil
	}
}

func UtilsIsValidId(r gogae.RequestContext) (interface{}, int, error) {
	//TODO: Implement
	return true, 200, nil
}
