package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserType(c *gin.Context, id string) (err error) {

	userType := c.GetString("user_Type")

	uid := c.GetString("uid")

	err = nil

	if userType == "USER" && uid != id {
		err = errors.New("unauthorize to access this resources")
		return err
	}

	err = CheckUser(c, userType)

	return err

}

func CheckUser(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unauthorize to access this resources")
		return err
	}
	return err
}
