package student

import (
	"context"
	"github.com/rs/xid"
	"student_portal/pkg/entity"
	"student_portal/pkg/env"
	"time"
)

func AddStudent(student *entity.Student) (*entity.Student, error) {
	id := xid.New()
	currentTime := time.Now()
	student.CreatedAt = &currentTime
	student.UpdatedAt = &currentTime
	student.Id = id.String()
	db := env.MongoDBConnection
	_, err := db.Collection("Student").InsertOne(context.Background(), student)
	if err != nil {
		return nil, err
	}
	return student, nil
}
