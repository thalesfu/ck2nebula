package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_SecretReligion struct {
	People         *People   `nebulaedgename:"people_secretReligion" nebulaedgecomment:"People -> SecretReligion" nebulakey:"edgefrom" json:"from,omitempty"`
	SecretReligion *Religion `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID        int       `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_SecretReligion(p *People, sr *Religion) *People_SecretReligion {
	return &People_SecretReligion{
		People:         p,
		SecretReligion: sr,
		StoryID:        p.StoryID,
	}
}

func (psr *People_SecretReligion) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, psr)
}

func (psr *People_SecretReligion) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, psr)
}

func (psr *People_SecretReligion) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, psr)
}

func (psr *People_SecretReligion) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, psr)
}

func (psr *People_SecretReligion) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, psr)
}

func InsertPeople_SecretReligions(space *nebulagolang.Space, psrs ...*People_SecretReligion) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, psrs)
}

func UpdatePeople_SecretReligions(space *nebulagolang.Space, psrs ...*People_SecretReligion) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, psrs)
}

func UpsertPeople_SecretReligions(space *nebulagolang.Space, psrs ...*People_SecretReligion) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, psrs)
}

func DeletePeople_SecretReligions(space *nebulagolang.Space, psrs ...*People_SecretReligion) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, psrs...)
}

func DeleteAllPeople_SecretReligions(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_SecretReligion](space)
}

func DeleteAllPeople_SecretReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_SecretReligion](space, getPlayIdQuery[People_SecretReligion](playId))
}

func GetPeople_SecretReligionByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_SecretReligion] {
	return nebulagolang.GetEdgeByEid[*People_SecretReligion](space, eid)
}

func GetPeople_SecretReligionByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_SecretReligion] {
	return nebulagolang.GetEdgeByEid[*People_SecretReligion](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_SecretReligion]()))
}

func GetAllPeople_SecretReligionsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_SecretReligion](space, "")
}

func GetAllPeople_SecretReligionsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_SecretReligion](space, getPlayIdQuery[People_SecretReligion](playId))
}

func GetAllPeople_SecretReligions(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_SecretReligion] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_SecretReligion](space)
}

func GetAllPeople_SecretReligionsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_SecretReligion] {
	return nebulagolang.GetAllEdgesByQuery[*People_SecretReligion](space, getPlayIdQuery[People_SecretReligion](playId))
}
