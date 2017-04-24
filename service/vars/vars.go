package vars

import (
	"github.com/gin-gonic/gin"
)

type Err struct {
	Code int64
	Msg  string
}

var (
	ErrBindJSON        = &Err{100, "bind json error"}
	ErrUserNotFound    = &Err{101, "user not found"}
	ErrLoginParams     = &Err{106, "user  params err"}
	ErrNeedToken       = &Err{301, "need auth token"}
	ErrInvalidToken    = &Err{301, "invalid token"}
	ErrIncompleteToken = &Err{301, "token incomplete"}
)

func FailReturn(c *gin.Context, err *Err) {
	c.JSON(400, err)
}
