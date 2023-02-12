package student

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"student_portal/pkg/entity"
	"student_portal/pkg/env"
	"time"
)

func UpdateStudent(student *entity.Student) (*entity.Student, error) {
	currentTime := time.Now()
	student.UpdatedAt = &currentTime
	db := env.MongoDBConnection
	_, err := db.Collection("Student").UpdateOne(context.Background(), bson.M{"_id": student.Id}, bson.D{{"$set", student}})
	if err != nil {
		return nil, err
	}
	return student, nil
}
