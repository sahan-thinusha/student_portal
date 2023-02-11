package student

import (
	"context"
	"github.com/rs/xid"
	"student_portal/pkg/entity"
	"student_portal/pkg/env"
)

func AddStudent(student *entity.Student) (*entity.Student, error) {
	id := xid.New()
	student.Id = id.String()
	db := env.MongoDBConnection
	_, err := db.Collection("Student").InsertOne(context.Background(), student)
	if err != nil {
		return nil, err
	}
	return student, nil
}
