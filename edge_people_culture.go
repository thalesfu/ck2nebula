package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_Culture struct {
	People  *People  `nebulaedgename:"people_culture" nebulaedgecomment:"People -> Culture" nebulakey:"edgefrom" json:"from,omitempty"`
	Culture *Culture `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int      `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_Culture(p *People, c *Culture) *People_Culture {
	return &People_Culture{
		People:  p,
		Culture: c,
		StoryID: p.StoryID,
	}
}

func (pc *People_Culture) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pc)
}

func (pc *People_Culture) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pc)
}

func (pc *People_Culture) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pc)
}

func (pc *People_Culture) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pc)
}

func (pc *People_Culture) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pc)
}

func InsertPeople_Cultures(space *nebulagolang.Space, pcs ...*People_Culture) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pcs)
}

func UpdatePeople_Cultures(space *nebulagolang.Space, pcs ...*People_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pcs)
}

func UpsertPeople_Cultures(space *nebulagolang.Space, pcs ...*People_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pcs)
}

func DeletePeople_Cultures(space *nebulagolang.Space, pcs ...*People_Culture) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pcs...)
}

func DeleteAllPeople_Cultures(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_Culture](space)
}

func DeleteAllPeople_CulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_Culture](space, getPlayIdQuery[People_Culture](playId))
}

func GetPeople_CultureByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_Culture] {
	return nebulagolang.GetEdgeByEid[*People_Culture](space, eid)
}

func GetPeople_CultureByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_Culture] {
	return nebulagolang.GetEdgeByEid[*People_Culture](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_Culture]()))
}

func GetAllPeople_CulturesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Culture](space, "")
}

func GetAllPeople_CulturesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_Culture](space, getPlayIdQuery[People_Culture](playId))
}

func GetAllPeople_Cultures(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_Culture] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_Culture](space)
}

func GetAllPeople_CulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_Culture] {
	return nebulagolang.GetAllEdgesByQuery[*People_Culture](space, getPlayIdQuery[People_Culture](playId))
}
