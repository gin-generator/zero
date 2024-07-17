package user

import (
	"github.com/gin-generator/ginctl/package/database"
	"github.com/gin-generator/zero/app/http/demo/logic"
	model "github.com/gin-generator/zero/model/davari"
	http "github.com/gin-generator/zero/package/respond"
	"github.com/gin-gonic/gin"
)

type Logic struct{}

func NewLogic() *Logic {
	return &Logic{}
}

// Index Get page list
func (l *Logic) Index(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Index](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}

	var count int64
	db := database.DB.Model(model.NewUser())
	if params.Data().Keywords != nil {
		db = db.Where("name LIKE ?", *params.Data().Keywords)
	}
	err = db.Count(&count).Error
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}

	var list []model.User
	err = db.Limit(params.Data().Size).
		Offset((params.Data().Page - 1) * params.Data().Size).
		Order("id DESC").
		Find(&list).Error
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}

	data := make(map[string]interface{})
	data["count"] = count
	data["list"] = list
	http.SuccessWithData(c, data)
}

// Show Get info
func (l *Logic) Show(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Show](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: Your logic.

	// TODO: Replace your return struct.
	http.SuccessWithData(c, params.Data())
}

// Create Save one source
func (l *Logic) Create(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Create](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: Your logic.

	// TODO: Replace your return struct.
	http.SuccessWithData(c, params.Data())
}

// Update Modifying a resource
func (l *Logic) Update(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Update](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: Your logic.

	// TODO: Replace your return struct.
	http.SuccessWithData(c, params.Data())
}

// Destroy Delete a resource
func (l *Logic) Destroy(c *gin.Context) {
	params, err := logic.ParseAndCheckParams[Destroy](c)
	if err != nil {
		http.Alert400(c, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: Your logic.

	// TODO: Replace your return struct.
	http.SuccessWithData(c, params.Data())
}
