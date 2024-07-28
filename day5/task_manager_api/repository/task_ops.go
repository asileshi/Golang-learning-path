package repository

import (
    "context"
    "github.com/asileshi/model"
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

func CreateTask(task model.Task) (model.Task, error){
    _, err := TaskCollection.InsertOne(context.TODO(), task)
    if err != nil{
        return model.Task{}, err
    }
    return task, nil
}
func UpdateTask(id string, updatedTask model.Task) (model.Task, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return model.Task{}, err
    }

    filter := bson.D{{"_id", objectID}}

    update := bson.D{{"$set", bson.D{}}}
    if updatedTask.Title != "" {
        update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "title", Value: updatedTask.Title})
    }
    if updatedTask.Description != "" {
        update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "description", Value: updatedTask.Description})
    }
    if !updatedTask.DueDate.IsZero() {
        update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "due_date", Value: updatedTask.DueDate})
    }
    if updatedTask.Status != "" {
        update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "status", Value: updatedTask.Status})
    }

    _, err = TaskCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return model.Task{}, err
    }

    updatedTask.ID = objectID
    return updatedTask, nil
}

func DeleteTask(id string) error{
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