package routers

import (
	"github.com/anle/codebase/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
