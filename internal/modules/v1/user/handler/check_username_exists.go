package handler

type CheckUsernameExistsRequest struct {
	Username string `json:"username"`
}

func (handler *UserHandler) CheckUsernameExists(username string) (bool, error) {
	//exists, err := CheckUsernameExists(req.Username)
	return false, nil
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
