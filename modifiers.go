package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/modifier"
	"log"
)

func GenerateModifiersData(path string) []*Modifier {
	data := modifier.LoadAllModifier(path)

	ms := make([]*Modifier, 0)

	for _, m := range data {
		nm := NewModifierByData(m)
		ms = append(ms, nm)
	}

	return ms
}

func BuildModifiers(path string) {
	ms := GenerateModifiersData(path)

	ur, cr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, ms, "")

	if !ur.Ok {
		log.Fatalf("Load and update modifiers error: %s", ur.Err.Error())
	} else {
		log.Println("Modifiers added:", len(cr.Added))
		log.Println("Modifiers updated:", len(cr.Updated))
		log.Println("Modifiers deleted:", len(cr.Deleted))
		log.Println("Modifiers kept:", len(cr.Kept))
	}
}
