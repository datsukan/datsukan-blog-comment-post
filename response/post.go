package response

// Post は、コメント投稿のレスポンスの構造体。
type PostResponse struct {
	ID        string `json:"id"`
	ArticleID string `json:"article_id"`
	ParentID  string `json:"parent_id"`
	UserName  string `json:"user_name"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
