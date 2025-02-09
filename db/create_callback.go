package db

import (
	"github.com/huhx-headhunter/headhunter-common/base"
	"github.com/huhx-headhunter/headhunter-common/store"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
	"reflect"
	"sync"
)

var ThreadLocal sync.Map

func BeforeCreateCallback(db *gorm.DB) {
	if db.Statement.SkipHooks {
		return
	}
	entity := db.Statement.Model
	val := reflect.ValueOf(entity)

	switch val.Type().Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			element := val.Index(i).Interface()
			processEntityCreation(element)
		}
	default:
		processEntityCreation(entity)
	}
	db.Model(entity)
}

func processEntityCreation(input interface{}) {
	val := reflect.ValueOf(input).Elem()
	entity := val.FieldByName("Entity")
	immutableEntity := val.FieldByName("ImmutableEntity")
	relationEntity := val.FieldByName("RelationEntity")

	if entity.IsValid() {
		baseEntity := entity.Addr().Interface().(*base.Entity)
		username := store.LoadDefault("username", "unknown").(string)
		baseEntity.CreatedBy = username
		baseEntity.UpdatedBy = username

		baseEntity.Version = optimisticlock.Version{Valid: true}
		baseEntity.DeletedAt = -1
	} else if immutableEntity.IsValid() {
		baseEntity := immutableEntity.Addr().Interface().(*base.ImmutableEntity)
		baseEntity.Version = optimisticlock.Version{Valid: true}
		baseEntity.DeletedAt = -1
	} else if relationEntity.IsValid() {
		baseEntity := relationEntity.Addr().Interface().(*base.RelationEntity)
		baseEntity.Version = optimisticlock.Version{Valid: true}
		baseEntity.DeletedAt = -1
	}
}
