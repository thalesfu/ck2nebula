package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
	"time"
)

func GenerateProvinces(provinces map[int]*save.Province, cm map[string]string, rm map[string]string) ([]*Province, []*Province_Modifier, []*Province_Culture, []*Province_Religion, []*Province_Title) {
	rps := make([]*Province, len(provinces))
	rpms := make([]*Province_Modifier, 0)
	rpcs := make([]*Province_Culture, 0)
	rprs := make([]*Province_Religion, 0)
	rpts := make([]*Province_Title, 0)

	i := 0
	for _, province := range provinces {
		rps[i] = NewProvinceByData(province)
		rps[i].ReligionName = rm[rps[i].Religion]
		rps[i].CultureName = cm[rps[i].Culture]

		if rps[i].Culture != "" {
			rpcs = append(rpcs, NewProvince_Culture(rps[i], NewCulture(rps[i].Culture)))
		}

		if rps[i].Religion != "" {
			rprs = append(rprs, NewProvince_Religion(rps[i], NewReligion(rps[i].Religion)))
		}

		if rps[i].Code != "" {
			rpts = append(rpts, NewProvince_Title(rps[i], NewTitle(province.PlayID, rps[i].Code)))
		}

		for _, modifier := range province.Modifiers {
			pm := NewProvince_Modifier(rps[i], NewModifier(modifier.Modifier))
			pm.Abate = time.Time(modifier.Date)
			pm.Hidden = bool(modifier.Hidden)
			rpms = append(rpms, pm)
		}

		i++
	}

	return rps, rpms, rpcs, rprs, rpts
}
