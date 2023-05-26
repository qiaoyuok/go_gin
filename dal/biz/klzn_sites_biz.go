package biz

import (
	"fmt"
	"go_gin/dal/model"
	"go_gin/dal/query"
)

func SaveKlznSite(site *model.KlznSite) {
	ks := query.KlznSite
	err := ks.Create(site)
	if err != nil {
		return
	}
	fmt.Println(site)
}
