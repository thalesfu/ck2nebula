package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Trait_OppositeTrait struct {
	Trait         *Trait `nebulaedgename:"trait_oppositetrait" nebulaedgecomment:"Trait -> Opposite Trait" nebulakey:"edgefrom" json:"from,omitempty"`
	OppositeTrait *Trait `nebulakey:"edgeto" json:"to,omitempty"`
}

func NewTrait_OppositeTrait(t *Trait, ot *Trait) *Trait_OppositeTrait {
	return &Trait_OppositeTrait{
		Trait:         t,
		OppositeTrait: ot,
	}
}

func (tot *Trait_OppositeTrait) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, tot)
}

func (tot *Trait_OppositeTrait) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, tot)
}

func (tot *Trait_OppositeTrait) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, tot)
}

func (tot *Trait_OppositeTrait) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tot)
}

func (tot *Trait_OppositeTrait) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, tot)
}

func InsertTrait_OppositeTraits(space *nebulagolang.Space, tots ...*Trait_OppositeTrait) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, tots)
}

func UpdateTrait_OppositeTraits(space *nebulagolang.Space, tots ...*Trait_OppositeTrait) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, tots)
}

func UpsertTrait_OppositeTraits(space *nebulagolang.Space, tots ...*Trait_OppositeTrait) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, tots)
}

func DeleteTrait_OppositeTraits(space *nebulagolang.Space, tots ...*Trait_OppositeTrait) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tots...)
}

func DeleteAllTrait_OppositeTraits(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Trait_OppositeTrait](space)
}

func GetTrait_OppositeTraitByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Trait_OppositeTrait] {
	return nebulagolang.GetEdgeByEid[*Trait_OppositeTrait](space, eid)
}

func GetTrait_OppositeTraitByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Trait_OppositeTrait] {
	return nebulagolang.GetEdgeByEid[*Trait_OppositeTrait](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Trait_OppositeTrait]()))
}

func GetAllTrait_OppositeTraitsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Trait_OppositeTrait](space, "")
}

func GetAllTrait_OppositeTraits(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Trait_OppositeTrait] {
	return nebulagolang.GetAllEdgesByEdgeType[*Trait_OppositeTrait](space)
}
