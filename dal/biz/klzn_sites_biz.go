package biz

import (
	"fmt"
	"go_gin/dal/model"
	"go_gin/dal/query"
)

func SaveKlznSite(site *model.KlznSites) {
	ks := query.KlznSites
	err := ks.Create(site)
	if err != nil {
		return
	}
	fmt.Println(site)
}
