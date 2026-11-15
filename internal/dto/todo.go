package dto

type CreateTodoRequest struct {
	Title   string `json:"title" binding:"required"`
	Details string `json:"details" binding:"required"`
}

type UpdateTodoRequest struct {
	Title   *string `json:"title"`
	Details *string `json:"details"`
}
