package label

import (
	db "CoreSystemBase/service/db"
)

type Label struct {
	Id       int64  `json:"id"`
	TypeId   int64  `josn:"type_id"`
	Name     string `json:"name"`
	OrderNum int64  `json:"order_num"`
	Abbr     string `json:"abbr"`
}
