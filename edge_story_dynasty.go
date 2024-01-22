package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Story_Dynasty struct {
	Story   *Story   `nebulaedgename:"story_dynasty" nebulaedgecomment:"Story -> Dynasty" nebulakey:"edgefrom" json:"from,omitempty"`
	Dynasty *Dynasty `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID  int      `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewStory_Dynasty(s *Story, d *Dynasty) *Story_Dynasty {
	return &Story_Dynasty{
		Story:   s,
		Dynasty: d,
		PlayID:  s.PlayID,
	}
}

func (sd *Story_Dynasty) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, sd)
}

func (sd *Story_Dynasty) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, sd)
}

func (sd *Story_Dynasty) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, sd)
}

func (sd *Story_Dynasty) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sd)
}

func (sd *Story_Dynasty) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, sd)
}

func InsertStory_Dynastys(space *nebulagolang.Space, sds ...*Story_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, sds)
}

func UpdateStory_Dynastys(space *nebulagolang.Space, sds ...*Story_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, sds)
}

func UpsertStory_Dynastys(space *nebulagolang.Space, sds ...*Story_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, sds)
}

func DeleteStory_Dynastys(space *nebulagolang.Space, sds ...*Story_Dynasty) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sds...)
}

func DeleteAllStory_Dynastys(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Story_Dynasty](space)
}

func DeleteAllStory_DynastysByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Story_Dynasty](space, getPlayIdQuery[Story_Dynasty](playId))
}

func GetStory_DynastyByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Story_Dynasty] {
	return nebulagolang.GetEdgeByEid[*Story_Dynasty](space, eid)
}

func GetStory_DynastyByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Story_Dynasty] {
	return nebulagolang.GetEdgeByEid[*Story_Dynasty](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Story_Dynasty]()))
}

func GetAllStory_DynastysEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Dynasty](space, "")
}

func GetAllStory_DynastysEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Dynasty](space, getPlayIdQuery[Story_Dynasty](playId))
}

func GetAllStory_Dynastys(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story_Dynasty] {
	return nebulagolang.GetAllEdgesByEdgeType[*Story_Dynasty](space)
}

func GetAllStory_DynastysByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Story_Dynasty] {
	return nebulagolang.GetAllEdgesByQuery[*Story_Dynasty](space, getPlayIdQuery[Story_Dynasty](playId))
}
