package service

import (
	"context"
	"tix-rest-project/internal/entity"
	"tix-rest-project/internal/repository"
	"tix-rest-project/internal/shared"
	"tix-rest-project/internal/shared/dto"
)

type (
	Post interface {
		Find(ctx context.Context, req *dto.FindPostRequest) ([]dto.FindPostResponse, error)
		FindById(ctx context.Context, req *dto.FindPostByIdRequest) (*dto.FindPostByIdResponse, error)
		Create(ctx context.Context, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error)
		Update(ctx context.Context, req *dto.UpdatePostRequest) (*dto.UpdatePostResponse, error)
		Delete(ctx context.Context, req *dto.DeletePostRequest) (*dto.DeletePostResponse, error)
	}

	postService struct {
		deps shared.Deps
		repo repository.Post
	}
)

func (impl *postService) Find(ctx context.Context, req *dto.FindPostRequest) ([]dto.FindPostResponse, error) {
	var (
		res = make([]dto.FindPostResponse, 0)
		orm = impl.deps.Database.WithContext(ctx)
	)

	posts, err := impl.repo.Find(orm, req.Offset, req.Limit)

	if err != nil {
		return nil, err
	}

	for _, p := range posts {
		res = append(res, dto.FindPostResponse{
			Id:      p.Id,
			Title:   p.Title,
			Content: p.Content,
		})
	}

	return res, nil
}

func (impl *postService) FindById(ctx context.Context, req *dto.FindPostByIdRequest) (*dto.FindPostByIdResponse, error) {
	var (
		orm = impl.deps.Database.WithContext(ctx)
	)

	post, err := impl.repo.FindById(orm, req.Id)

	if err != nil {
		return nil, err
	}

	return &dto.FindPostByIdResponse{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (impl *postService) Create(ctx context.Context, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	var (
		orm = impl.deps.Database.WithContext(ctx)
	)

	post, err := impl.repo.Create(orm, &entity.Post{
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &dto.CreatePostResponse{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (impl *postService) Update(ctx context.Context, req *dto.UpdatePostRequest) (*dto.UpdatePostResponse, error) {
	var (
		orm = impl.deps.Database.WithContext(ctx)
	)

	post, err := impl.repo.Update(orm, &entity.Post{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		return nil, err
	}
	return &dto.UpdatePostResponse{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (impl *postService) Delete(ctx context.Context, req *dto.DeletePostRequest) (*dto.DeletePostResponse, error) {
	var (
		orm = impl.deps.Database.WithContext(ctx)
	)

	if err := impl.repo.Delete(orm, req.Id); err != nil {
		return nil, err
	}

	return &dto.DeletePostResponse{}, nil
}

func NewPost(deps shared.Deps, repo repository.Post) (Post, error) {
	return &postService{deps, repo}, nil
}
