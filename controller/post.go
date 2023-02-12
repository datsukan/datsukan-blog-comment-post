package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	coreResponse "github.com/datsukan/datsukan-blog-comment-core/response"
	"github.com/datsukan/datsukan-blog-comment-core/usecase"
	"github.com/datsukan/datsukan-blog-comment-post/request"
	"github.com/datsukan/datsukan-blog-comment-post/response"
)

func Post(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req request.PostRequest
	err := json.Unmarshal([]byte(r.Body), &req)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}

	if err := req.Validate(); err != nil {
		return coreResponse.ResponseBadRequestError(err)
	}

	c, err := usecase.Post(req.ArticleID, req.ParentID, req.UserName, req.Content)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}

	pr := response.PostResponse{
		ID:        c.ID,
		ArticleID: c.ArticleID,
		ParentID:  c.ParentID,
		UserName:  c.UserName,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}

	j, err := json.Marshal(pr)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}
	js := string(j)

	return coreResponse.ResponseSuccess(js)
}
