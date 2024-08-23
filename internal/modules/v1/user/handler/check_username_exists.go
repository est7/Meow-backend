package handler

import (
	"Meow-backend/internal/models"
	"Meow-backend/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type CheckUsernameExistsRequest struct {
	Username string `json:"username"`
}

func CheckUsernameExists(c *gin.Context) {
	var req CheckUsernameExistsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.FailResult(c, errcode.ErrInvalidParam)
		return
	}

	exists, err := CheckUsernameExists(req.Username)

}

//func (d *repository) UserIsExist(user *models.UserBaseModel) (bool, error) {
//err := d.orm.Where("username = ? or email = ?", user.Username, user.Email).First(&model.UserBaseModel{}).Error
//if errors.Is(err, gorm.ErrRecordNotFound) {
//	return false, nil
//} else if err != nil {
//	return false, err
//}
//return true, nil
//}
