package paginator

import (
	"context"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

type sPaginator struct{}

func init() {
	service.RegisterPaginator(New())
}

func New() service.IPaginator {
	return &sPaginator{}
}

func (s *sPaginator) Paginate(ctx context.Context, model *gdb.Model, structure interface{}, page int, pageSize int) (totalPage int, err error) {
	total := 0
	if page <= 0 || pageSize <= 0 {
		err = model.Scan(&structure)
	} else {
		err = model.Limit((page-1)*pageSize, pageSize).ScanAndCount(&structure, &total, true)
	}
	if err != nil {
		return
	}
	if page > 0 && pageSize > 0 {
		remainder := total % pageSize
		totalPage := total / pageSize
		if remainder != 0 {
			totalPage++
		}
	}

	return
}
