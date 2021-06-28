// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"Gf-Vben/app/dao/internal"
)

// routerDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type routerDao struct {
	*internal.RouterDao
}

var (
	// Router is globally public accessible object for table router operations.
	Router routerDao
)

func init() {
	Router = routerDao{
		internal.NewRouterDao(),
	}
}

// Fill with you ideas below.