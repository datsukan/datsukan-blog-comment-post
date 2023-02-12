package request

import "fmt"

// Post は、コメント投稿のリクエストの構造体。
type PostRequest struct {
	ArticleID string `json:"article_id"`
	ParentID  string `json:"parent_id"`
	UserName  string `json:"user_name"`
	Content   string `json:"content"`
}

func (r *PostRequest) Validate() error {
	if r.ArticleID == "" {
		return fmt.Errorf("article_id is empty")
	}

	if r.UserName == "" {
		return fmt.Errorf("user_name is empty")
	}

	if r.Content == "" {
		return fmt.Errorf("content is empty")
	}

	return nil
}
