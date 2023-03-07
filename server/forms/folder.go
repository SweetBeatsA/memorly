package forms

type CreateFolderForm struct {
	Title string `json:"title" binding:"required"`
}
