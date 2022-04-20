package dto

type (
	FindPostRequest struct {
		Offset, Limit int64
	}
	FindPostResponse struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	FindPostByIdRequest struct {
		Id int64 `param:"id"`
	}
	FindPostByIdResponse struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	CreatePostRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	CreatePostResponse struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	UpdatePostRequest struct {
		Id      int64  `param:"id" json:"-"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	UpdatePostResponse struct {
		Id      int64  `param:"id" json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	DeletePostRequest struct {
		Id int64 `param:"id"`
	}
	DeletePostResponse struct{}
)
