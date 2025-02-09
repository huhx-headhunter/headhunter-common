package base

import (
	"github.com/huhx/common-go/times"
	"gorm.io/plugin/optimisticlock"
)

type Entity struct {
	Id        int64                  `gorm:"primaryKey" json:"id" example:"1800783867820785152"`
	CreatedAt times.LocalDateTime    `json:"createdAt" example:"2019-10-03T12:00:00"`
	CreatedBy string                 `json:"-"`
	UpdatedAt times.LocalDateTime    `json:"-"`
	UpdatedBy string                 `json:"-"`
	Version   optimisticlock.Version `json:"-"`
	DeletedAt int64                  `json:"-"`
}

type ImmutableEntity struct {
	Id        int64                  `gorm:"primaryKey" json:"id" example:"1800783867820785152"`
	CreatedAt times.LocalDateTime    `json:"createdAt" example:"2019-10-03T12:00:00"`
	Version   optimisticlock.Version `json:"-"`
	DeletedAt int64                  `json:"-"`
}

type RelationEntity struct {
	CreatedAt times.LocalDateTime    `json:"createdAt" example:"2019-10-03T12:00:00"`
	Version   optimisticlock.Version `json:"-"`
	DeletedAt int64                  `json:"-"`
}
