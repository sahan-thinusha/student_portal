package student

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"student_portal/pkg/entity"
	"student_portal/pkg/env"
)

func GetAllStudents(index, limit int) ([]*entity.Student, error) {
	var students []*entity.Student
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find().SetSort(bson.D{{"updated_at", 1}})
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("Student").Find(context.Background(), bson.M{}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), students); err != nil {
			return nil, err
		}

		return students, nil
	} else {
		cursor, err := db.Collection("Student").Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), students); err != nil {
			return nil, err
		}

		return students, nil
	}
}
