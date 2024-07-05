package service

import (
	"attendance-system/internal/infra"
	"attendance-system/internal/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUser(name string, username string, password string, client *mongo.Client) (string, error) {

	coll := infra.GetDatabase().Collection("user")
	//u := model.User{}
	us := model.User{Name: name, Username: username}
	result, err := coll.InsertOne(context.Background(), us)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data entered successfully: ", result)

	return fmt.Sprintf("%v", result.InsertedID), err

}

func ValidateUser(name string, username string, password string, client *mongo.Client) error {
	coll := infra.GetDatabase().Collection("user")
	toFind := bson.D{{Key: "name", Value: name}, {Key: "username", Value: username}, {Key: "password", Value: password}}

	var result model.User
	err := coll.FindOne(context.Background(), toFind).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		fmt.Println(err)
	}

	return err

}

func CheckUsername(name string, username string, client *mongo.Client) error {
	coll := infra.GetDatabase().Collection("user")
	toFindUsername := bson.D{{Key: "name", Value: name}, {Key: "username", Value: username}}

	var resultUsername model.User
	err := coll.FindOne(context.Background(), toFindUsername).Decode(&resultUsername)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Valid username")

		}
	}
	return err
}

func EnterAttendance(date string, username string, cc bool, gt bool, dm bool, client *mongo.Client) (string, error) {
	att := infra.GetDatabase().Collection("attendance")
	coll := infra.GetDatabase().Collection("user")

	at := model.Attendance{Date: date, Username: username, CC: cc, GT: gt, DM: dm}

	toFindUser := bson.D{{Key: "username", Value: username}}

	var resultUser model.User
	err := coll.FindOne(context.Background(), toFindUser).Decode(&resultUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "no matching username", err
		}
		return "error occurred", err
	}

	enter, err := att.InsertOne(context.Background(), at)
	if err != nil {
		return "error occurred", err
	}
	return fmt.Sprintf("%v", enter.InsertedID), err
}

func CheckAttendanceByDate(date string, username string, client *mongo.Client) (*model.Attendance, error) {
	att := infra.GetDatabase().Collection("attendance")
	toFindDate := bson.D{{Key: "date", Value: date}, {Key: "username", Value: username}}

	var resultDate model.Attendance
	err := att.FindOne(context.Background(), toFindDate).Decode(&resultDate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		fmt.Println(err)
	}
	return &resultDate, err
}

func TotalAttendance(username string, subject string, client *mongo.Client) (int, error) {
	att := infra.GetDatabase().Collection("attendance")
	//coll := infra.GetDatabase().Collection("user")

	//at := model.Attendance{Date: date, Username: username, CC: cc, GT: gt, DM: dm}

	toFindUser := bson.D{{Key: "username", Value: username}}

	var resultUser model.User
	err := att.FindOne(context.Background(), toFindUser).Decode(&resultUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, err
		}
		return 0, err
	}

	// matchStage := bson.D{{Key: "$match", Value: bson.D{
	// 	{Key: "username", Value: username},
	// },
	// },
	// }

	// groupStage := bson.D{
	// 	{Key: "$group", Value: bson.D{
	// 		{Key: "_id", Value: nil},
	// 		{Key: "total", Value: bson.D{
	// 			{Key: "$count", Value: fmt.Sprintf("$%v", subject)},
	// 		}},
	// 	}}}

	filter := bson.D{{Key: "subject", Value: bson.D{{Key: "$gt", Value: 0}}}}
	totalAtt, err := att.CountDocuments(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	// cursor, err := att.Aggregate(context.Background(), mongo.Pipeline{groupStage, matchStage})
	// if err != nil {
	// 	return 0, err
	// }
	// defer cursor.Close(context.Background())

	// total := []bson.M{}
	// if err = cursor.All(context.Background(), &total); err != nil {
	// 	fmt.Println(err)
	// 	return 0, err
	// }
	// if len(total) > 0 {
	// 	totalAtt, ok := total[0]["total"].(int32)
	// 	if !ok {
	// 		return 0, fmt.Errorf("error occurred")
	// 	}
	// 	return int32(totalAtt), err
	// }
	return int(totalAtt), nil

}

func GenerateStudentID(year int, dept string, client *mongo.Client) (string, error) {
	coll := infra.GetDatabase().Collection("user")

	departmentCode := map[string]string{
		"CSE": "bce",
		"ECE": "ece",
	}
	var resultCounter model.Counter
	err := coll.FindOneAndUpdate(context.Background(), bson.M{"_id": nil}, bson.M{
		"$inc": bson.M{"count": 1}}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	).Decode(&resultCounter)

	if err != nil {
		return "", nil
	}

	newId := fmt.Sprintf("%d%s%03d", year, departmentCode, resultCounter.Cnt)
	return newId, nil

}

// func ValidateCreds(username string, password string, client *mongo.Client) error {
// 	coll := infra.GetDatabase().Collection("user")
// 	toCheckCreds := bson.D{{Key: "username", Value: username}, {Key: "password", Value: password}}

// 	var resultCreds model.User
// 	err := coll.FindOne(context.Background(), toCheckCreds).Decode(&resultCreds)
// 	if err != nil {
// 		fmt.Println("username and passwords do not match")
// 	}
// 	return err
// }
