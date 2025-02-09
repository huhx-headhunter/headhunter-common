package db

import (
	"github.com/huhx-headhunter/headhunter-common/base"
	"github.com/huhx-headhunter/headhunter-common/store"
	"gorm.io/gorm"
	"reflect"
)

func BeforeUpdateCallback(db *gorm.DB) {
	if db.Statement.SkipHooks {
		return
	}
	entity := db.Statement.Model
	val := reflect.ValueOf(entity)

	switch val.Type().Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			element := val.Index(i).Interface()
			processEntityUpdate(element)
		}
	default:
		processEntityUpdate(entity)
	}
	db.Model(entity)
}

func processEntityUpdate(input interface{}) {
	val := reflect.ValueOf(input).Elem()
	baseEntityField := val.FieldByName("Entity")

	if baseEntityField.IsValid() {
		baseEntity := baseEntityField.Addr().Interface().(*base.Entity)
		baseEntity.UpdatedBy = store.LoadDefault("username", "unknown").(string)
	}
}
