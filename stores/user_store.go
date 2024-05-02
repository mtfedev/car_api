package stores

import (
	"context"

	"github.com/mtfedev/hotel-one/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserStore interface {
	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpateUserParams) error
	Dropper
}
type Dropper interface {
	Drop(context.Context) error
}
