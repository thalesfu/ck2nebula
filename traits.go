package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/trait"
	"log"
)

func GenerateTraitsData(path string) (
	[]*Trait,
	[]*Trait_OppositeTrait) {
	traits := trait.LoadAllTraits(path)

	ts := make([]*Trait, 0)
	tots := make([]*Trait_OppositeTrait, 0)

	for _, t := range traits {
		nt := NewTraitByData(t)
		ts = append(ts, nt)

		for _, ot := range t.Opposites {
			otTrait := NewTraitByData(traits[ot])
			tot := NewTrait_OppositeTrait(nt, otTrait)
			tots = append(tots, tot)
		}
	}

	return ts, tots
}

func BuildTraits(path string) {
	ts, tots := GenerateTraitsData(path)

	ur, cr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, ts, "", false)

	if !ur.Ok {
		log.Fatalf("Load and update trait error: %s", ur.Err.Error())
	} else {
		log.Println("Traits added:", len(cr.Added))
		log.Println("Traits updated:", len(cr.Updated))
		log.Println("Traits deleted:", len(cr.Deleted))
		log.Println("Traits kept:", len(cr.Kept))
	}

	utotr, ctotr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, tots, "", false)

	if !utotr.Ok {
		log.Fatalf("Load and update trait_oppositetrait error: %s", utotr.Err.Error())
	} else {
		log.Println("Trait_OppositeTraits added:", len(ctotr.Added))
		log.Println("Trait_OppositeTraits updated:", len(ctotr.Updated))
		log.Println("Trait_OppositeTraits deleted:", len(ctotr.Deleted))
		log.Println("Trait_OppositeTraits kept:", len(ctotr.Kept))
	}
}
