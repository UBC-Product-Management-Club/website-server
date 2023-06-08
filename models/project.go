package models

type Project struct {
	DetailId   string `json:"detailId" firestore:"detailId"`
	Image      string `json:"image" firestore:"image"`
	IsFinished bool   `json:"isFinished" firestore:"isFinished" binding:"required"`
	Link       string `json:"link" firestore:"link" binding:"required"`
	Text       string `json:"text" firestore:"text" binding:"required"`
	Title      string `json:"title" firestore:"title" binding:"required"`
}
