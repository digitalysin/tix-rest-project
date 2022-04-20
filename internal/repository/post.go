package repository

import (
	"time"
	"tix-rest-project/internal/entity"

	"github.com/digitalysin/goblog/db"
)

type (
	Post interface {
		Find(orm db.ORM, offset, limit int64) ([]entity.Post, error)
		FindById(orm db.ORM, id int64) (*entity.Post, error)
		Create(orm db.ORM, post *entity.Post) (*entity.Post, error)
		Update(orm db.ORM, post *entity.Post) (*entity.Post, error)
		Delete(orm db.ORM, id int64) error
	}

	postRepo struct{}
)

func (impl *postRepo) Find(orm db.ORM, offset, limit int64) ([]entity.Post, error) {
	var (
		posts = make([]entity.Post, 0)
	)

	if err := orm.Offset(offset).Limit(limit).Find(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (impl *postRepo) FindById(orm db.ORM, id int64) (*entity.Post, error) {
	var (
		post = entity.Post{}
	)

	if err := orm.Where("id = ?", id).First(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (impl *postRepo) Create(orm db.ORM, post *entity.Post) (*entity.Post, error) {
	if err := orm.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (impl *postRepo) Update(orm db.ORM, post *entity.Post) (*entity.Post, error) {
	post.UpdatedAt = time.Now()

	if err := orm.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (impl *postRepo) Delete(orm db.ORM, id int64) error {
	var (
		post = entity.Post{}
	)

	if err := orm.Where("id = ?", id).First(&post); err != nil {
		return err
	}

	return orm.Delete(&post)
}

func NewPost() (Post, error) {
	return &postRepo{}, nil
}
