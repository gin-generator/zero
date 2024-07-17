package demo

import (
	"github.com/gin-generator/zero/app/http/demo/logic"
	http "github.com/gin-generator/zero/package/respond"
	"github.com/gin-gonic/gin"
)

type Logic struct{}

func NewLogic() *Logic {
	return &Logic{}
}

// Ping Ping
func (l *Logic) Ping(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Ping](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: Your logic.

	// TODO: Replace your return struct.
	http.SuccessWithData(c, params.Data())
}
