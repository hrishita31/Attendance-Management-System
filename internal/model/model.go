package model

const ConstUserDB = "users_db"

type User struct {
	Name       string `bson:"name"`
	Username   string `bson:"username"`
	Email      string `bson:"Email"`
	Rollno     string `bson:"rollno"`
	Department string `bson:"department"`
	Year       int    `bson:"year"`
}

type Attendance struct {
	Date     string `bson:"date"`
	Username string `bson:"username"`
	CC       bool   `bson:"CC"`
	GT       bool   `bson:"GT"`
	DM       bool   `bson:"DM"`
}

type TotalAttendance struct {
	Username string `bson:"username"`
	CC       bool   `bson:"CC"`
	GT       bool   `bson:"GT"`
	DM       bool   `bson:"DM"`
}

type Counter struct {
	ID  string `bson:"_id"`
	Cnt int    `bson: "cnt"`
}
