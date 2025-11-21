package controller

import (
	"belajar-golang-dasar/internal/module/member/entity"
	"belajar-golang-dasar/internal/module/member/interfaces"
	"belajar-golang-dasar/internal/module/member/repository"
	"belajar-golang-dasar/internal/module/member/service"
	"belajar-golang-dasar/pkg/handler"
	"belajar-golang-dasar/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type memberController struct {
	service interfaces.MemberService
}

func NewMemberController(db *gorm.DB) *memberController {
	var (
		controller = new(memberController)
		repo       = repository.NewMemberRepository(db)
		service    = service.NewMemberService(repo)
	)
	controller.service = service

	return controller
}

func (c *memberController) Register(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/users/:uuid", func(ctx *gin.Context) {
		var req entity.MemberReqByID
		if !validator.BindUri(ctx, &req) {
			return
		}

		res, err := c.service.GetMember(&req)
		if err != nil {
			handler.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		handler.Success(ctx, http.StatusOK, "Success getting users", &res)
	})
}
