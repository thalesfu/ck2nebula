package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/building"
	"log"
)

func GenerateBuildingsData(path string) []*Building {
	bs := make([]*Building, 0)

	data := building.LoadAllBuildings(path)

	for _, bg := range data {
		for _, b := range bg.Buildings {
			nb := NewBuildingByData(b)
			bs = append(bs, nb)
		}
	}

	return bs
}

func BuildBuildings(path string) {
	ms := GenerateBuildingsData(path)

	ur, cr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, ms, "")

	if !ur.Ok {
		log.Fatalf("Load and update buildings error: %s", ur.Err.Error())
	} else {
		log.Println("Buildings added:", len(cr.Added))
		log.Println("Buildings updated:", len(cr.Updated))
		log.Println("Buildings deleted:", len(cr.Deleted))
		log.Println("Buildings kept:", len(cr.Kept))
	}
}
