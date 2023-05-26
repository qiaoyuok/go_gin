package biz

import (
	"fmt"
	"go_gin/dal/model"
	"go_gin/dal/query"
)

func SaveKlznMovie(move *model.KlznMovie) {
	ks := query.KlznMovie
	err := ks.Create(move)
	if err != nil {
		return
	}
	fmt.Println(move)
}
