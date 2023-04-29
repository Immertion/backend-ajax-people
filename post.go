package user

type Post struct {
	Id              int    `json:"id" db:"id"`
	UserId          int    `json:"userId" db:"user_id" binding:"required"`
	Text            string `json:"text" db:"text" binding:"required"`
	IsModerated     bool   `json:"isModerated" db:"is_moderated" binding:"required"`
	PublicationTime string `json:"publicationTime" db:"publication_time"`
}
