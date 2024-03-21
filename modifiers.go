package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/modifier"
	"github.com/thalesfu/paradoxtools/CK2/traderoute"
	"log"
)

func GenerateModifiersData(path string) []*Modifier {
	routes := traderoute.LoadAllTradeRoutes(path)

	ms := make([]*Modifier, 0)

	for _, r := range routes {
		nr := NewModifierByData(r.Modifier)
		ms = append(ms, nr)
	}

	data := modifier.LoadAllModifier(path)

	for _, m := range data {
		nm := NewModifierByData(m)
		ms = append(ms, nm)
	}

	return ms
}

func BuildModifiers(path string) {
	ms := GenerateModifiersData(path)

	ur, cr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, ms, "", false)

	if !ur.Ok {
		log.Fatalf("Load and update modifiers error: %s", ur.Err.Error())
	} else {
		log.Println("Modifiers added:", cr.AddedCount)
		log.Println("Modifiers updated:", cr.UpdatedCount)
		log.Println("Modifiers deleted:", cr.DeletedCount)
		log.Println("Modifiers kept:", cr.KeptCount)
	}
}
