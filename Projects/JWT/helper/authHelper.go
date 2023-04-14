package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauth to access this")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")

	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errors.New("Unauth to access this")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
