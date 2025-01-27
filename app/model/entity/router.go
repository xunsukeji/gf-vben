// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Router is the golang structure for table router.
type Router struct {
	Id        int64       `json:"id"        ` //
	Path      string      `json:"path"      ` //
	Name      string      `json:"name"      ` //
	Redirect  string      `json:"redirect"  ` //
	Title     string      `json:"title"     ` //
	Icon      string      `json:"icon"      ` //
	Component string      `json:"component" ` //
	Parent    int64       `json:"parent"    ` //
	OrderNo   int64       `json:"orderNo"   ` //
	Status    int         `json:"status"    ` //
	CreateAt  *gtime.Time `json:"createAt"  ` //
	UpdateAt  *gtime.Time `json:"updateAt"  ` //
	DeleteAt  *gtime.Time `json:"deleteAt"  ` //
}
