package user

import (
	"context"

	"main/app/user/api/entity"
	"main/app/user/internal/dao/user"
	"main/utility/injection"
	"main/utility/mongohelper"

	"github.com/gogf/gf/v2/util/gconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Service struct {
		user user.Dao
	}
)

func New() *Service {
	return &Service{
		user: user.New(injection.MustInvoke[*mongo.Database]()),
	}
}

func (s *Service) Create(ctx context.Context, name string) (string, error) {
	result, err := s.user.Create(ctx, user.CreateInput{
		Name: name,
	})
	if err != nil {
		return "", err
	}
	return result, err
}

func (s *Service) GetById(ctx context.Context, id string) (*entity.User, error) {
	var (
		item   *entity.User
		userId = mongohelper.MustObjectIDFromHex(id)
	)
	result, err := s.user.GetOne(ctx, userId)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(result, &item)
	return item, err
}

func (s *Service) GetList(ctx context.Context, ids []string) ([]*entity.User, error) {
	var (
		items   []*entity.User
		userIds = mongohelper.MustObjectIDsFromHexes(ids)
	)
	result, err := s.user.GetList(ctx, user.GetListInput{
		Ids: userIds,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(result, &items)
	return items, err
}

func (s *Service) DeleteById(ctx context.Context, id string) error {
	var (
		userId = mongohelper.MustObjectIDFromHex(id)
		err    = s.user.Delete(ctx, []primitive.ObjectID{userId})
	)
	return err
}
