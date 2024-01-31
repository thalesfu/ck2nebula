package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Baron_Building struct {
	Baron    *Baron    `nebulaedgename:"baron_building" nebulaedgecomment:"Baron -> Building" nebulakey:"edgefrom" json:"from,omitempty"`
	Building *Building `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID  int       `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewBaron_Building(baron *Baron, building *Building) *Baron_Building {
	return &Baron_Building{
		Baron:    baron,
		Building: building,
		StoryID:  baron.StoryID,
	}
}

func (bb *Baron_Building) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, bb)
}

func (bb *Baron_Building) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, bb)
}

func (bb *Baron_Building) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, bb)
}

func (bb *Baron_Building) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, bb)
}

func (bb *Baron_Building) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, bb)
}

func InsertBaron_Buildings(space *nebulagolang.Space, bbs ...*Baron_Building) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, bbs)
}

func UpdateBaron_Buildings(space *nebulagolang.Space, bbs ...*Baron_Building) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, bbs)
}

func UpsertBaron_Buildings(space *nebulagolang.Space, bbs ...*Baron_Building) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, bbs)
}

func DeleteBaron_Buildings(space *nebulagolang.Space, bbs ...*Baron_Building) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, bbs...)
}

func DeleteAllBaron_Buildings(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Baron_Building](space)
}

func DeleteAllBaron_BuildingsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Baron_Building](space, getPlayIdQuery[Baron_Building](playId))
}

func GetBaron_BuildingByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Baron_Building] {
	return nebulagolang.GetEdgeByEid[*Baron_Building](space, eid)
}

func GetBaron_BuildingByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Baron_Building] {
	return nebulagolang.GetEdgeByEid[*Baron_Building](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Baron_Building]()))
}

func GetAllBaron_BuildingsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Baron_Building](space, "")
}

func GetAllBaron_BuildingsEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Baron_Building](space, getPlayIdQuery[Baron_Building](playId))
}

func GetAllBaron_Buildings(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Baron_Building] {
	return nebulagolang.GetAllEdgesByEdgeType[*Baron_Building](space)
}

func GetAllBaron_BuildingsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Baron_Building] {
	return nebulagolang.GetAllEdgesByQuery[*Baron_Building](space, getPlayIdQuery[Baron_Building](playId))
}
