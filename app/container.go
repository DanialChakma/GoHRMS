package app

import (
	"go.mod/services"
	"gorm.io/gorm"
)

type Container struct {
	DB    *gorm.DB
	Cache *services.Cache
	IDGen *services.IDGenerator
}

// type AppContainer struct {
// 	DB    *gorm.DB
// 	Cache *services.Cache
// 	IDGen *services.IDGenerator
// }
