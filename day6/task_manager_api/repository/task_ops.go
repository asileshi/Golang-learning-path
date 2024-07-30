package repository

import (
	"context"
	"github.com/asileshi/task_manager_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "reflect"
)


func IsZeroValue(value interface{}) bool {
	return reflect.ValueOf(value).IsZero()
}

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

    
    //insertedTaskID := result.InsertedID.(primitive.ObjectID)

    var insertedTask model.Task
    err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&insertedTask)
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

    v := reflect.ValueOf(updatedTask)
	typeOfTask := v.Type()

	for i := 0; i < v.NumField(); i++ {
        field := typeOfTask.Field(i).Name
        value := v.Field(i).Interface()
        
        if !IsZeroValue(value){
            updateFields[field] = value
        }
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