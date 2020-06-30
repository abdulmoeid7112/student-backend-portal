package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/abdulmoeid7112/student-backend-portal/config"
	"github.com/abdulmoeid7112/student-backend-portal/db"
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

const (
	collectionName = "Student"
)

func init() {
	db.Register("mongo", NewMongoClient)
}

type client struct {
	conn *mongo.Client
}

// NewMongoClient initializes a mysql database connection.
func NewMongoClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	return &client{conn: cli}, nil
}

func (m client) AddStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "id is not empty", nil
	}
	student.ID = guuid.NewV4().String()
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.InsertOne(context.TODO(), student); err != nil {
		return "", errors.Wrap(err, "failed to add student")
	}
	return student.ID, nil
}

func (m client) UpdateStudent(id string, student *models.Student) error {
	student.ID = id
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": student.ID}}, bson.M{"$set": student}); err != nil {
		return errors.Wrap(err, "failed to update student")
	}
	return nil
}

func (m client) GetStudentByID(id string) (*models.Student, error) {
	var stu *models.Student
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&stu); err != nil {
		if mongo.ErrNoDocuments != nil {
			return nil, errors.Wrap(err, "failed to fetch student....not found")
		}
		return nil, err
	}
	return stu, nil
}

func (m client) ListStudent(filter map[string]interface{}, lim int64, off int64) ([]*models.Student, error) {
	var std []*models.Student
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Skip:  &off,
		Limit: &lim,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve student")
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var em *models.Student
		if err = cursor.Decode(&em); err != nil {
			return nil, errors.Wrap(err, "failed to scan retrieved rows")
		}
		std = append(std, em)
	}
	return std, nil
}

func (m client) RemoveStudentByID(id string) error {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(collectionName)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete student")
	}

	return nil
}
