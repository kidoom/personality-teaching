package middle

import (
	"net/http"
	"personality-teaching/src/code"
	"personality-teaching/src/logger"
	"personality-teaching/src/logic"
	"personality-teaching/src/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// VerifyTeacher 负责验证用户是否有教师权限，若有则在上下文中存入teacher_id
func VerifyTeacher(c *gin.Context) {
	key, err := c.Cookie(utils.SessionKey)
	if err == http.ErrNoCookie {
		code.CommonResp(c, http.StatusOK, code.NeedLogin, code.EmptyData)
		return
	}
	teacherID, err := logic.NewTeacherService().CheckTeacherPermission(key)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.ServerBusy, code.EmptyData)
		logger.L.Error("check teacher permission error: ", zap.Error(err))
		return
	}
	if teacherID == "" {
		code.CommonResp(c, http.StatusOK, code.InvalidPermission, code.EmptyData)
		return
	}
	c.Set(utils.TeacherID, teacherID)
}