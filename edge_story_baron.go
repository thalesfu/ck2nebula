package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Story_Baron struct {
	Story   *Story `nebulaedgename:"story_baron" nebulaedgecomment:"Story -> Baron" nebulakey:"edgefrom" json:"from,omitempty"`
	Baron   *Baron `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int    `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewStory_Baron(s *Story, b *Baron) *Story_Baron {
	return &Story_Baron{
		Story:   s,
		Baron:   b,
		StoryID: s.StoryId,
	}
}

func (sb *Story_Baron) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, sb)
}

func (sb *Story_Baron) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, sb)
}

func (sb *Story_Baron) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, sb)
}

func (sb *Story_Baron) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sb)
}

func (sb *Story_Baron) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, sb)
}

func InsertStory_Barons(space *nebulagolang.Space, sbs ...*Story_Baron) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, sbs)
}

func UpdateStory_Barons(space *nebulagolang.Space, sbs ...*Story_Baron) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, sbs)
}

func UpsertStory_Barons(space *nebulagolang.Space, sbs ...*Story_Baron) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, sbs)
}

func DeleteStory_Barons(space *nebulagolang.Space, sbs ...*Story_Baron) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sbs...)
}

func DeleteAllStory_Barons(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Story_Baron](space)
}

func DeleteAllStory_BaronsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Story_Baron](space, getPlayIdQuery[Story_Baron](playId))
}

func GetStory_BaronByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Story_Baron] {
	return nebulagolang.GetEdgeByEid[*Story_Baron](space, eid)
}

func GetStory_BaronByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Story_Baron] {
	return nebulagolang.GetEdgeByEid[*Story_Baron](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Story_Baron]()))
}

func GetAllStory_BaronsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Baron](space, "")
}

func GetAllStory_BaronsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Baron](space, getPlayIdQuery[Story_Baron](playId))
}

func GetAllStory_Barons(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story_Baron] {
	return nebulagolang.GetAllEdgesByEdgeType[*Story_Baron](space)
}

func GetAllStory_BaronsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Story_Baron] {
	return nebulagolang.GetAllEdgesByQuery[*Story_Baron](space, getPlayIdQuery[Story_Baron](playId))
}
