package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/server/cmd"

	"github.com/xiao-hub-create/book/controller"
	"github.com/xiao-hub-create/book/models"
	"gorm.io/gorm"

	// 引入Gin Root Router: *gin.Engine
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	ioc_gin "github.com/infraboard/mcube/v2/ioc/config/gin"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

func main() {
	// 注册业务对象到 ioc区域
	ioc.Controller().Registry(&BookController{})
	ioc.Api().Registry(&BookApiHandler{})

	// 启动应用
	cmd.Start()
}

type BookApiHandler struct {
	// 继承自Ioc对象
	ioc.ObjectImpl
}

// BookApiHandler Api 模块, 约定好了的URL逻辑
// 模块的名称, 会作为路径的一部分比如: /mcube_service/api/v1/books/
// 路径构成规则 <service_name>/<path_prefix>/<service_version>/<module_name>
func (h *BookApiHandler) Name() string {
	return "books"
}

// API路由
func (h *BookApiHandler) Init() error {
	r := ioc_gin.ObjectRouter(h)
	r.GET("/:isbn", h.GetHook)
	return nil
}

func (h *BookApiHandler) GetHook(ctx *gin.Context) {
	strId := ctx.Param("isbn")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	// 传递HTTP请求的上下文
	ins, err := ioc.Controller().Get("books").(*BookController).GetBook(ctx.Request.Context(), &controller.GetBookRequest{
		Isbn: id,
	})
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

type BookController struct {
	// 继承自Ioc对象
	ioc.ObjectImpl
}

// Hook
func (c *BookController) Init() error {
	log.L().Debug().Msgf("BookController Init ...")
	return nil
}

func (c *BookController) Name() string {
	return "books"
}

func (c *BookController) GetBook(ctx context.Context, req *controller.GetBookRequest) (*models.Book, error) {
	ins := &models.Book{}

	if err := datasource.DBFromCtx(ctx).Where("isbn = ?", req.Isbn).Take(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("%d not found", req.Isbn)
		}
		return nil, fmt.Errorf("get book error, %s", err)
	}

	return ins, nil
}
