package service

import (
	"TodoList/model"
	"TodoList/serializer"
	"fmt"
)

type TaskService struct {
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Status    int    `json:"status" form:"status"` // 0 ：未做 1：已做
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

type ListService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

func (service *TaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0, // 默认是未完成 0
		Content:   service.Content,
		StartTime: service.StartTime,
		EndTime:   service.EndTime,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建备忘录成功",
	}
}

func (service *TaskService) Update(uid uint, tid string) serializer.Response {
	var task model.Task
	err := model.DB.First(&task, tid).Error
	if err != nil {
		fmt.Printf("update error: %s\n", err.Error())
	}
	if task.Uid != uid {
		return serializer.Response{
			Status: 500,
			Msg:    "无权限修改此备忘录",
		}
	}
	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status
	task.StartTime = service.StartTime
	task.EndTime = service.EndTime
	fmt.Printf("tid: %s\n", tid)
	fmt.Printf("Title: %s\nContent: %s\nStatus: %d\nStartTime: %d\nEndTime: %d\n", service.Title, service.Content, service.Status, service.StartTime, service.EndTime)

	err = model.DB.Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "更新失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "更新成功",
		Data:   serializer.BuildTask(&task),
	}
}

func (service *TaskService) Show(uid uint, tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
			Error:  err.Error(),
		}
	}
	if task.Uid != uid {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "无权限查看此备忘录",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(&task),
	}
}

func (service *TaskService) Delete(uid uint, tid string) serializer.Response {
	var task model.Task
	err := model.DB.Find(&task, tid).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    fmt.Sprintf("备忘录不存在：%d", tid),
			Error:  err.Error(),
		}
	}

	if task.Uid != uid {
		return serializer.Response{
			Status: 500,
			Msg:    "用户无权限删除该备忘录",
		}
	}

	err = model.DB.Delete(&task, tid).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "备忘录删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

func (service *ListService) Search(uid uint) serializer.Response {
	var taskList []*model.Task
	var count int
	keyword := "%" + service.Info + "%"
	fmt.Printf("num: %d; size: %d;", service.PageNum, service.PageSize)
	model.DB.Model(&model.Task{}).
		Preload("User").
		Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", keyword, keyword).
		Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).
		Count(&count).
		Find(&taskList)

	return serializer.Response{
		Status: 200,
		Msg:    "搜索成功",
		Data:   serializer.BuildTasks(taskList),
		Count:  count,
	}
}

func (service *ListService) TaskList(uid uint) serializer.Response {
	var taskList []*model.Task
	count := 0
	model.DB.Model(&model.Task{}).
		Preload("User"). // 预加载 User 表
		Where("uid=?", uid).
		Count(&count).
		Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).
		Find(&taskList)

	return serializer.Response{
		Status: 200,
		Msg:    "获取成功",
		Data:   serializer.BuildTasks(taskList),
		Count:  count,
	}
}
