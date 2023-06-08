package models

import "cloud.google.com/go/firestore"

type Fellow struct {
	Img      string                   `json:"img" firestore:"img"`
	Name     string                   `json:"name" firestore:"name"`
	Title    string                   `json:"title" firestore:"title" binding:"required"`
	BioText  string                   `json:"bio_text" firestore:"bio_text" binding:"required"`
	Linkedin string                   `json:"linkedin" firestore:"linkedin"`
	Projects []*firestore.DocumentRef `json:"projects" firestore:"projects" binding:"required"`
}
