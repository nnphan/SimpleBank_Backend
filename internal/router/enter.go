package router

import (
	"simplebank/internal/router/manage"
	"simplebank/internal/router/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)