package util

import (
	"github.com/huhx/common-go/base"
)

func GetPageable(pageQuery func(string, string) string) base.Pageable {
	pageIndex := pageQuery("pageIndex", "0")
	pageSize := pageQuery("pageSize", "20")
	return base.NewPageable(pageIndex, pageSize)
}
