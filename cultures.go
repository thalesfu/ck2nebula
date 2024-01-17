package ck2nebula

import "github.com/thalesfu/paradoxtools/CK2/culture"

func GenerateCultureData(path string) ([]*CultureGroup, []*Culture, []*CultureGroup_Culture) {
	cultureGroups := culture.LoadAllCultures(path)

	cgs := make([]*CultureGroup, 0)
	cs := make([]*Culture, 0)
	cg_c_s := make([]*CultureGroup_Culture, 0)

	for _, cultureGroup := range cultureGroups {
		nebulaCultureGroup := NewCultureGroupByData(cultureGroup)
		cgs = append(cgs, nebulaCultureGroup)

		for _, culture := range cultureGroup.Cultures {
			nebulaCulture, nebulaCultureGroup_Culture := nebulaCultureGroup.NewCultureByData(culture)
			cs = append(cs, nebulaCulture)
			cg_c_s = append(cg_c_s, nebulaCultureGroup_Culture)
		}
	}

	return cgs, cs, cg_c_s
}
