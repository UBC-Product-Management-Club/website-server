package database

import (
	"context"

	"github.com/UBC-Product-Management-Club/website-server/models"
	"google.golang.org/api/iterator"
)

// Add a project to database
func AddProject(project models.Project) error {
	_, _, err := client.Collection("project").Add(context.Background(), project)
	return err
}

// Get all projects from database
func GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	iter := client.Collection("project").Documents(context.Background())
	for doc, err := iter.Next(); err != iterator.Done; doc, err = iter.Next() {
		if err != nil {
			return []models.Project{}, err
		}

		var tempProject models.Project
		if err := doc.DataTo(&tempProject); err != nil {
			return []models.Project{}, err
		}
		projects = append(projects, tempProject)
	}
	return projects, nil
}
