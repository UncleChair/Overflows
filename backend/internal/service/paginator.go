// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IPaginator interface {
		Paginate(ctx context.Context, model *gdb.Model, structure interface{}, page int, pageSize int) (totalPage int, err error)
	}
)

var (
	localPaginator IPaginator
)

func Paginator() IPaginator {
	if localPaginator == nil {
		panic("implement not found for interface IPaginator, forgot register?")
	}
	return localPaginator
}

func RegisterPaginator(i IPaginator) {
	localPaginator = i
}
