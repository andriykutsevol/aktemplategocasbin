package menuaction

import (
	"context"

	"github.com/andriykutsevol/WeatherServer/internal/domain/pagination"
)

type Repository interface {
	Query(ctx context.Context, params QueryParam) (MenuActions, *pagination.Pagination, error)
	Get(ctx context.Context, id string) (*MenuAction, error)
	GetByIdString(ctx context.Context, id string) (*MenuAction, error)
	Create(ctx context.Context, item *MenuAction) error
	Update(ctx context.Context, id string, item *MenuAction) error
	Delete(ctx context.Context, id string) error
	DeleteByMenuID(ctx context.Context, menuID string) error
	Purge(ctx context.Context) error
}
