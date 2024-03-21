package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/objectives"
	"log"
)

func GenerateObjectivesData(path string) []*Objective {
	objs := objectives.LoadAllObjective(path)

	os := make([]*Objective, 0)

	for _, o := range objs {
		no := NewObjectiveByData(o)
		os = append(os, no)
	}

	return os
}

func BuildObjectives(path string) {
	os := GenerateObjectivesData(path)

	ur, cr := nebulagolang.CompareAndUpdateNebulaEntityBySliceAndQuery(SPACE, os, "", false)

	if !ur.Ok {
		log.Fatalf("Load and update objectives error: %s", ur.Err.Error())
	} else {
		log.Println("Objectives added:", cr.AddedCount)
		log.Println("Objectives updated:", len(cr.Updated))
		log.Println("Objectives deleted:", len(cr.Deleted))
		log.Println("Objectives kept:", len(cr.Kept))
	}
}
