package api

import (
	"TodoList/pkg/utils"
	"TodoList/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var taskService service.TaskService
	// 验证身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&taskService)

	// 处理错误
	if err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		return
	}

	res := taskService.Create(claim.Id)
	c.JSON(200, res)
}

func UpdateTask(c *gin.Context) {
	var taskService service.TaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&taskService)

	// 处理错误
	if err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		return
	}

	res := taskService.Update(claim.Id, c.Param("id"))
	c.JSON(200, res)
}

func ShowTask(c *gin.Context) {
	var taskService service.TaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := taskService.Show(claim.Id, c.Param("id"))
	c.JSON(200, res)
}

func DeleteTask(c *gin.Context) {
	var taskService service.TaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := taskService.Delete(claim.Id, c.Param("id"))
	c.JSON(200, res)
}

func SearchTask(c *gin.Context) {
	var listService service.ListService
	err := c.ShouldBind(&listService)
	// 处理错误
	if err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		return
	}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := listService.Search(claim.Id)
	c.JSON(200, res)
}

func TaskList(c *gin.Context) {
	var listService service.ListService
	err := c.ShouldBind(&listService)
	// 处理错误
	if err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		return
	}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	res := listService.TaskList(claim.Id)
	c.JSON(200, res)
}
