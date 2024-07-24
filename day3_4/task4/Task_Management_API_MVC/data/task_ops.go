package data

import (
	"errors"
	model "github.com/asileshi/task_manager_api/Model"
	
)

func GetTasks() []model.Task {
	return Tasks
}
func GetTaskById(id string) (model.Task,error){
	for _, task:= range Tasks{
		if task.ID == id{
			return task,nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func CreateTask(newTask model.Task){
	Tasks = append(Tasks, newTask)
}

func UpdateTask(id string, updatedTask model.Task) (model.Task,error){
	for i:= range(Tasks){
		if Tasks[i].ID == id{
			if updatedTask.Title != ""{
				Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != ""{
				Tasks[i].Description = updatedTask.Description

			}
			if !updatedTask.DueDate.IsZero(){
				Tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != ""{
				Tasks[i].Status = updatedTask.Status
			}
			return Tasks[i],nil
		}

	}
	return model.Task{},errors.New("task not found")
}

func DeleteTask(id string) error{
	for i, task:= range Tasks{
		if task.ID == id{
			Tasks = append(Tasks[:i],Tasks[i+1:]... )
			return nil
		}
	}
	return errors.New("task not found")
}