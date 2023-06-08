package models

type Executive struct {
	Img   string `json:"img" firestore:"img"`
	Name  string `json:"name" firestore:"name" binding:"required"`
	Title string `json:"title" firestore:"title" binding:"required"`
}
