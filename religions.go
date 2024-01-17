package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/religion"
)

func GenerateReligionData(path string) ([]*ReligionGroup, []*Religion, []*ReligionGroup_Religion) {
	data := religion.LoadAllReligions(path)

	rgs := make([]*ReligionGroup, 0)
	rs := make([]*Religion, 0)
	rg_r_s := make([]*ReligionGroup_Religion, 0)

	for _, rg := range data {
		nrg := NewReligionGroupByData(rg)
		rgs = append(rgs, nrg)

		for _, r := range rg.Religions {
			nr, nrgr := nrg.NewReligionByData(r)
			rs = append(rs, nr)
			rg_r_s = append(rg_r_s, nrgr)
		}
	}

	return rgs, rs, rg_r_s
}
