package resp

import (
	"net/http"

	"github.com/DarwinWDEWEN/ranxingo/types"
	"github.com/gin-gonic/gin"
)

func SUCCESS(c *gin.Context, values ...interface{}) {
	if values != nil {
		c.JSON(http.StatusOK, types.BizVo{Code: types.Success, Data: values[0]})
	} else {
		c.JSON(http.StatusOK, types.BizVo{Code: types.Success})
	}

}

func ERROR(c *gin.Context, messages ...string) {
	if messages != nil {
		c.JSON(http.StatusBadRequest, types.BizVo{Code: types.Failed, Message: messages[0]})
	} else {
		c.JSON(http.StatusBadRequest, types.BizVo{Code: types.Failed})
	}
}

func HACKER(c *gin.Context) {
	c.JSON(http.StatusBadRequest, types.BizVo{Code: types.Failed, Message: "Hacker attempt!!!"})
}

func NotAuth(c *gin.Context, messages ...string) {
	if messages != nil {
		c.JSON(http.StatusUnauthorized, types.BizVo{Code: types.NotAuthorized, Message: messages[0]})
	} else {
		c.JSON(http.StatusUnauthorized, types.BizVo{Code: types.NotAuthorized, Message: "Not Authorized"})
	}
}
