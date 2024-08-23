package handler

//func GenerateToken(userID uint64) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"user_id": userID,
//	})
//}
//
//
//func RefreshToken(token string) (string, error) {
//	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
//		return []byte(config.Get().JWT.AccessSecret), nil
//	})
//	if err != nil {
//		return "", err
//	}
