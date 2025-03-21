package menu

import (
	"context"

	"github.com/andriykutsevol/WeatherServer/internal/domain/pagination"
)

type Repository interface {
	Query(ctx context.Context, params QueryParam) (Menus, *pagination.Pagination, error)
	Get(ctx context.Context, id string) (*Menu, error)
	GetByIdString(ctx context.Context, id string) (*Menu, error)
	Create(ctx context.Context, item *Menu) error
	Update(ctx context.Context, id string, item *Menu) error
	UpdateParentPath(ctx context.Context, id, parentPath string) error
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string, status int) error
	Purge(ctx context.Context) error
}
