package v1

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/biz"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
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
		response.Write(c, &response.Response{
			HTTP: http.StatusOK,
			Result: response.OkResult{
				"id":          m.Id,
				"timestamp":   m.Timestamp,
				"title":       m.Title,
				"description": m.Description,
				"type":        m.Type,
				"level":       m.Level,
			},
		})
		return
	}
}

func (t *TaskController) Edit(c *gin.Context) {

}

func (t *TaskController) Delete(c *gin.Context) {
	r := &task.DeleteRequest{}
	if err := c.ShouldBindUri(r); err != nil {
		fmt.Println(err.Error())
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.TaskBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if err := b.Delete(r, userId); err != nil {
		fmt.Println(err.Error())
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

}
