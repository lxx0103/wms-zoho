package user

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"wms.com/core/database"
	"wms.com/core/queue"
)

type NewAuthCreated struct {
	AuthID     int64  `json:"auth_id"`
	AuthType   int    `json:"auth_type"`
	Identifier string `json:"identifier"`
	Credential string `json:"credential"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type NewProfileCreated struct {
	AuthID int64 `json:"auth_id"`
	UserID int64 `json:"user_id"`
}

func Subscribe(conn *queue.Conn) {
	conn.StartConsumer("CreateUserProfile", "NewAuthCreated", CreateUserProfile)
}

func CreateUserProfile(d amqp.Delivery) bool {
	if d.Body == nil {
		return false
	}
	var newAuthCreated NewAuthCreated
	err := json.Unmarshal(d.Body, &newAuthCreated)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var userInfo UserProfile
	userInfo.Name = newAuthCreated.Name
	userInfo.Email = newAuthCreated.Email
	db := database.InitMySQL()
	repo := NewUserRepository(db)
	userID, err := repo.CreateUser(userInfo)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var newEvent NewProfileCreated
	newEvent.AuthID = newAuthCreated.AuthID
	newEvent.UserID = userID
	rabbit, _ := queue.GetConn()
	msg, _ := json.Marshal(newEvent)
	err = rabbit.Publish("NewProfileCreated", msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
