package db

import (
	"context"
	"time"

	"github.com/chetan177/oneclick/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) CreateProject(project *models.Project) error {
	collection := db.Client.Database("oneclick").Collection("projects")
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	_, err := collection.InsertOne(context.TODO(), project)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetProjects() ([]models.Project, error) {
	collection := db.Client.Database("oneclick").Collection("projects")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var projects []models.Project
	for cursor.Next(context.TODO()) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (db *DB) GetProject(id string) (models.Project, error) {
	collection := db.Client.Database("oneclick").Collection("projects")
	var project models.Project
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&project)
	return project, err
}

func (db *DB) UpdateProject(id string, project models.Project) error {
	collection := db.Client.Database("oneclick").Collection("projects")
	project.UpdatedAt = time.Now()
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": project})
	return err
}

func (db *DB) DeleteProject(id string) error {
	collection := db.Client.Database("oneclick").Collection("projects")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (db *DB) GetProjectByName(name string) (*models.Project, error) {
	collection := db.Client.Database("oneclick").Collection("projects")
	var project models.Project
	err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&project)
	return &project, err
}
