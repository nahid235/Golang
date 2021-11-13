package api

import (
	helper "Golang/api/helpers"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// GetAllTodosFromDB - Mongo
func GetAllPersonDetailsFromDB() error {
	res := []PersonDetails{}
	if err := helper.Collection().Find(nil).All(&res); err != nil {
		return err
	}

	PersonList = res
	return nil
}

// SaveTodoToDB - Mongo
func SavePersonDetailsToDB(data PersonDetails) error {
	uuidVar, _ := uuid.NewV4() // 47891-sskjsknfdsf-89624-dhnkskfhdkf
	data.Guid = uuidVar.String()
	//data.CreatedAt, data.UpdatedAt = time.Now(), time.Now()
	return helper.Collection().Insert(data)
}

// DeleteTodoFromDB - Mongo
func DeletePersonDetailsFromDB(id string) error {

	return helper.Collection().Remove(bson.M{"guid": id})
}

// // UpdateTodoFromDB - Mongo
// // You have to implement it
func UpdatePersonDetailsFromDB(data PersonDetails) error {
	//accept id to match
	userID := bson.M{"guid": data.Guid}
	//udata := bson.M{"$set": bson.M{"task": todo.Task}}
	//data.UpdatedAt = time.Now()
	return helper.Collection().Update(userID, data)

	// err := helper.Collection().Update(tskID, udata)
	// if err != nil {
	// 	panic(err)
	// }
}

// //DeleteTodoFromMYSQLDB  - MYSQL
// func DeleteTodoFromMYSQLDB(id string) {
// 	db := helper.DbConn()
// 	delForm, err := db.Prepare("DELETE FROM Todo WHERE Id=?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	delForm.Exec(id)
// }
