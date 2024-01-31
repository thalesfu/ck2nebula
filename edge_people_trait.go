package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_Trait struct {
	People  *People `nebulaedgename:"people_trait" nebulaedgecomment:"People -> Trait" nebulakey:"edgefrom" json:"from,omitempty"`
	Trait   *Trait  `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_Trait(p *People, t *Trait) *People_Trait {
	return &People_Trait{
		People:  p,
		Trait:   t,
		StoryID: p.StoryID,
	}
}

func (pt *People_Trait) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pt)
}

func (pt *People_Trait) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pt)
}

func (pt *People_Trait) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pt)
}

func (pt *People_Trait) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pt)
}

func (pt *People_Trait) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pt)
}

func InsertPeople_Traits(space *nebulagolang.Space, pts ...*People_Trait) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pts)
}

func UpdatePeople_Traits(space *nebulagolang.Space, pts ...*People_Trait) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pts)
}

func UpsertPeople_Traits(space *nebulagolang.Space, pts ...*People_Trait) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pts)
}

func DeletePeople_Traits(space *nebulagolang.Space, pts ...*People_Trait) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pts...)
}

func DeleteAllPeople_Traits(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Trait](space)
}

func DeleteAllPeople_TraitsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Trait](space, getPlayIdQuery[People_Trait](playId))
}

func GetPeople_TraitByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Trait] {
	return nebulagolang.GetEdgeByEid[*People_Trait](space, eid)
}

func GetPeople_TraitByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Trait] {
	return nebulagolang.GetEdgeByEid[*People_Trait](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Trait]()))
}

func GetAllPeople_TraitsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Trait](space, "")
}

func GetAllPeople_TraitsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Trait](space, getPlayIdQuery[People_Trait](playId))
}

func GetAllPeople_Traits(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Trait] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Trait](space)
}

func GetAllPeople_TraitsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Trait] {
	return nebulagolang.GetAllEdgesByQuery[*People_Trait](space, getPlayIdQuery[People_Trait](playId))
}
