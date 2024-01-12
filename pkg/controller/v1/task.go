package v1

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/biz"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type TaskController struct {
}

func (t *TaskController) Create(c *gin.Context) {
	r := &task.CreateRequest{}
	if err := c.ShouldBindJSON(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.TaskBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if m, err := b.Create(r, userId); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	} else {
		okResult := &response.OkResult{}
		_ = copier.Copy(okResult, m)
		response.Write(c, &response.Response{
			HTTP:   http.StatusOK,
			Result: okResult,
		})
		return
	}
}

func (t *TaskController) Edit(c *gin.Context) {
	r := &task.EditRequest{}
	if err := c.ShouldBindUri(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	if err := c.ShouldBindJSON(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.TaskBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if m, err := b.Edit(r, userId); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	} else {
		okResult := &response.OkResult{}
		_ = copier.Copy(okResult, m)
		response.Write(c, &response.Response{
			HTTP:   http.StatusOK,
			Result: okResult,
		})
		return
	}
}

func (t *TaskController) Delete(c *gin.Context) {
	r := &task.DeleteRequest{}
	if err := c.ShouldBindUri(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.TaskBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if err := b.Delete(r, userId); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	}
	okResult := make(response.OkResult)
	if r.Id != 0 {
		okResult["id"] = r.Id
	}
	response.Write(c, &response.Response{
		HTTP:   http.StatusOK,
		Result: okResult,
	})
}

func (t *TaskController) Read(c *gin.Context) {
	r := &task.ReadRequest{}
	if err := c.ShouldBindUri(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.TaskBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if m, err := b.Read(r, userId); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	} else {
		okResult := &response.OkResult{}
		_ = copier.Copy(okResult, m)
		response.Write(c, &response.Response{
			HTTP:   http.StatusOK,
			Result: okResult,
		})
		return
	}
}

func (t *TaskController) ReadList(c *gin.Context) {
	r := &task.ReadRequest{}
	if err := c.ShouldBindQuery(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.TaskBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if m, err := b.ReadList(r, userId); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	} else {
		okResultList := &response.OkResultList{}
		_ = copier.Copy(okResultList, m)
		response.Write(c, &response.Response{
			HTTP:   http.StatusOK,
			Result: okResultList,
		})
		return
	}
}
