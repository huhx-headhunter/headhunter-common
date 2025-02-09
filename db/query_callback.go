package db

import "gorm.io/gorm"

func BeforeQueryCallback(db *gorm.DB) {
	if !db.Statement.SkipHooks {
		//db.Where("deleted_at = -1")
	}
}
