package repository

import (
	"context"
	//"fmt"
	//"go/token"
	"time"

	"github.com/asileshi/task_manager_api/model"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

	}

	
func CheckHashedPassword(password,hashed string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil

}

func GenerateToken(user model.User) (string, error){
	expirationTime := time.Now().Add(24*time.Hour)
	claims := &model.Claim{
		Email: user.Email,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(model.Secretkey)
	if err != nil {
		return "", err
	}
	
	return tokenString, err
}
	



func FindUserByEmail(email string) (model.User,error){

	var curUser model.User
	err := UserCollection.FindOne(context.TODO(), bson.D{{"email",email}}).Decode(&curUser)

	return curUser, err
}

func CreateUser(user model.User) (model.User, string) {
	_, err := FindUserByEmail(user.Email)
	if err == nil {
		return model.User{}, "user name already taken"
	}
	hashedPassword, err := HashPassword(user.Password)

	if err != nil{
		return model.User{}, "failed to hash the password"
	}

	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()
	_, err = UserCollection.InsertOne(context.TODO(),user)
	return user,""
}

func Login(user model.User) (model.User,string){


	existingUser, err := FindUserByEmail(user.Email)
	if err != nil {
		return model.User{}, "User doesn't exist"
	}

	if !CheckHashedPassword(user.Password, existingUser.Password){
		return model.User{}, "invalid user name or password"
	}

	signedToken, error := GenerateToken(existingUser)

	if error != nil{
		return model.User{}, "failed to generate token"
	}

	return existingUser, signedToken




}