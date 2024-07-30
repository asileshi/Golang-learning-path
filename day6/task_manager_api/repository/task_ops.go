package repository

import (
	"context"
	"fmt"

	"github.com/asileshi/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks() ([]model.Task, error) {
    var tasks []model.Task
    cursor, err := TaskCollection.Find(context.TODO(), bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    err = cursor.All(context.Background(), &tasks)
    if err != nil {
        return nil, err
    }
    return tasks, nil
}

func GetTaskByID(id string) (model.Task, error){
    var task model.Task
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil{
        return task, err
    }

    filter := bson.D{{"_id",objectID}}

    err = TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
    if err != nil{
        return model.Task{}, err
    }

    return task, nil

}

func CreateTask(task model.Task) (model.Task, error) {
    result, err := TaskCollection.InsertOne(context.TODO(), task)
    if err != nil {
        return model.Task{}, err
    }

    insertedTaskID := result.InsertedID.(primitive.ObjectID)

    var insertedTask model.Task
    err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": insertedTaskID}).Decode(&insertedTask)
    if err != nil {
        return model.Task{}, err
    }

    return insertedTask, nil
}

func UpdateTask(id string, updatedTask model.Task) (model.Task, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return model.Task{}, err
    }

    filter := bson.D{{"_id", objectID}}

    updateFields := bson.M{}

    if updatedTask.Title != "" {
        updateFields["title"] = updatedTask.Title
    }
    if updatedTask.Description != "" {
        updateFields["description"] = updatedTask.Description
    }
    if !updatedTask.DueDate.IsZero() {
        updateFields["due_date"] = updatedTask.DueDate
    }
    if updatedTask.Status != "" {
        updateFields["status"] = updatedTask.Status
    }

    update := bson.D{{"$set", updateFields}}

    _, err = TaskCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return model.Task{}, err
    }

    // Fetch the updated document from the database
    var result model.Task
    err = TaskCollection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        return model.Task{}, err
    }

    return result, nil
}

func DeleteTask(id string) error{

    fmt.Print("completed")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil{

        return err
    }

    filter := bson.D{{"_id",objectID}}

    _, err = TaskCollection.DeleteOne(context.TODO(), filter)
    if err != nil{


        return err
    } 

    return nil
}