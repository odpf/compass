package star

//go:generate mockery --name Repository --outpkg mocks --output ../lib/mocks/ --with-expecter  --structname StarRepository --filename star_repository.go

import (
	"context"

	"github.com/odpf/columbus/asset"
	"github.com/odpf/columbus/user"
)

type Repository interface {
	Create(ctx context.Context, userID string, assetID string) (string, error)
	GetStargazers(ctx context.Context, flt Filter, assetID string) ([]user.User, error)
	GetAllAssetsByUserID(ctx context.Context, flt Filter, userID string) ([]asset.Asset, error)
	GetAllAssetsByUserEmail(ctx context.Context, flt Filter, userEmail string) ([]asset.Asset, error)
	GetAssetByUserID(ctx context.Context, userID string, assetID string) (asset.Asset, error)
	Delete(ctx context.Context, userID string, assetID string) error
}
