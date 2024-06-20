package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
)

func GenerateDynasties(dynasties map[int]*save.Dynasty) (
	[]*Dynasty,
	[]*Dynasty_Culture,
	[]*Dynasty_Religion) {
	rds := make([]*Dynasty, len(dynasties))
	rdcs := make([]*Dynasty_Culture, 0)
	rdrs := make([]*Dynasty_Religion, 0)

	i := 0
	for _, dynasty := range dynasties {
		rds[i] = NewDynastyByData(dynasty)

		if rds[i].Culture != "" {
			rdcs = append(rdcs, NewDynasty_Culture(rds[i], NewCulture(rds[i].Culture)))
		}

		if rds[i].Religion != "" {
			rdrs = append(rdrs, NewDynasty_Religion(rds[i], NewReligion(rds[i].Religion)))
		}

		i++
	}

	return rds, rdcs, rdrs
}
