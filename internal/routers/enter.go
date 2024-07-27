package routers

import (
	"github.com/tyri0n11/Muffin/internal/routers/manage"
	"github.com/tyri0n11/Muffin/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
