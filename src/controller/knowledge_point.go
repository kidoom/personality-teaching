package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"personality-teaching/src/code"
	"personality-teaching/src/logger"
	"personality-teaching/src/logic"
	"personality-teaching/src/model"
)

type KnowledgePointController struct{}

func KnowledgePointRegister(group *gin.RouterGroup) {
	point := &KnowledgePointController{}
	group.GET("/point_list", point.PointList)
	group.DELETE("/point_delete", point.PointDelete)
	group.GET("/point_detail", point.PointDetail)
	group.POST("/point_add", point.PointAdd)
	group.PUT("/point_update", point.PointUpdate)

}

var knowledgePointService = logic.NewKnowledgePointService()

// PointList godoc
// @Summary 知识点列表
// @Description 知识点列表
// @Tags 知识点管理
// @ID /point/point_list
// @Accept  json
// @Produce  json
// @Param info query string false "知识点关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} code.RespMsg{data=model.KnowledgePointListOutput} "success"
// @Router /point/point_list [get]
func (point *KnowledgePointController) PointList(c *gin.Context) {
	//从上下文获取参数并校验
	params := &model.KnowledgePointListInput{}
	if err := c.ShouldBind(params); err != nil {
		logger.L.Error("Input params error:", zap.Error(err))
		code.CommonResp(c, http.StatusBadRequest, code.InvalidParam, code.EmptyData)
		return
	}
	out, err := knowledgePointService.KnowledgePointListService(c, params)
	if err == gorm.ErrRecordNotFound {
		code.CommonResp(c, http.StatusInternalServerError, code.RecordNotFound, code.EmptyData)
	} else if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.ServerBusy, code.EmptyData)
	}
	code.CommonResp(c, http.StatusOK, code.Success, out)
}

// PointDelete godoc
// @Summary 知识点删除
// @Description 知识点删除
// @Tags 知识点管理
// @ID /point/point_delete
// @Accept  json
// @Produce  json
// @Param knp_id query string true "知识点编号"
// @Success 200 {object} code.RespMsg{data=string} "success"
// @Router /point/point_delete [delete]
func (point *KnowledgePointController) PointDelete(c *gin.Context) {
	params := &model.KnowledgePointDeleteInput{}
	if err := c.ShouldBind(params); err != nil {
		logger.L.Error("Input params error:", zap.Error(err))
		code.CommonResp(c, http.StatusBadRequest, code.InvalidParam, code.EmptyData)
		return
	}
	err := knowledgePointService.KnowledgePointDeleteService(c, params)
	if err == gorm.ErrRecordNotFound {
		code.CommonResp(c, http.StatusInternalServerError, code.RecordNotFound, code.EmptyData)
	} else if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.ServerBusy, code.EmptyData)
	}
	code.CommonResp(c, http.StatusOK, code.Success, code.EmptyData)
}

// PointAdd godoc
// @Summary 添加知识点
// @Description 添加知识点
// @Tags 知识点管理
// @ID /point/point_add
// @Accept  json
// @Produce  json
// @Param body body model.KnowledgePointAddInput true "body"
// @Success 200 {object} code.RespMsg{data=string} "success"
// @Router /point/point_add [post]
func (point *KnowledgePointController) PointAdd(c *gin.Context) {
	params := &model.KnowledgePointAddInput{}
	if err := c.ShouldBind(params); err != nil {
		logger.L.Error("Input params error:", zap.Error(err))
		code.CommonResp(c, http.StatusBadRequest, code.InvalidParam, code.EmptyData)
		return
	}
	err := knowledgePointService.KnowledgePointAddService(c, params)
	if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.ServerBusy, code.EmptyData)
	}
	code.CommonResp(c, http.StatusOK, code.Success, code.EmptyData)
}

// PointDetail godoc
// @Summary 知识点详情
// @Description 知识点详情
// @Tags 知识点管理
// @ID /point/point_detail
// @Accept  json
// @Produce  json
// @Param knp_id query string true "知识点编号"
// @Success 200 {object} code.RespMsg{data=mysql.TKnowledgePoint} "success"
// @Router /point/point_detail [get]
func (point *KnowledgePointController) PointDetail(c *gin.Context) {
	params := &model.KnowledgePointDetailInput{}
	if err := c.ShouldBind(params); err != nil {
		logger.L.Error("Input params error:", zap.Error(err))
		code.CommonResp(c, http.StatusBadRequest, code.InvalidParam, code.EmptyData)
		return
	}
	pointDetail, err := knowledgePointService.KnowledgePointDetailService(c, params)
	if err == gorm.ErrRecordNotFound {
		code.CommonResp(c, http.StatusInternalServerError, code.RecordNotFound, code.EmptyData)
	} else if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.ServerBusy, code.EmptyData)
	}
	code.CommonResp(c, http.StatusOK, code.Success, pointDetail)
}

// PointUpdate godoc
// @Summary 修改知识点
// @Description 修改知识点
// @Tags 知识点管理
// @ID /point/point_update
// @Accept  json
// @Produce  json
// @Param body body model.KnowledgePointUpdateInput true "body"
// @Success 200 {object} code.RespMsg{data=string} "success"
// @Router /point/point_update [put]
func (point *KnowledgePointController) PointUpdate(c *gin.Context) {
	params := &model.KnowledgePointUpdateInput{}
	if err := c.ShouldBind(params); err != nil {
		logger.L.Error("Input params error:", zap.Error(err))
		code.CommonResp(c, http.StatusBadRequest, code.InvalidParam, code.EmptyData)
		return
	}
	err := knowledgePointService.KnowledgePointUpdateService(c, params)
	if err == gorm.ErrRecordNotFound {
		code.CommonResp(c, http.StatusInternalServerError, code.RecordNotFound, code.EmptyData)
	} else if err != nil {
		code.CommonResp(c, http.StatusInternalServerError, code.ServerBusy, code.EmptyData)
	}
	code.CommonResp(c, http.StatusOK, code.Success, code.EmptyData)
}
