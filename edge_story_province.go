package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Story_Province struct {
	Story    *Story    `nebulaedgename:"story_province" nebulaedgecomment:"Story -> Province" nebulakey:"edgefrom" json:"from,omitempty"`
	Province *Province `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID   int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewStory_Province(s *Story, p *Province) *Story_Province {
	return &Story_Province{
		Story:    s,
		Province: p,
		PlayID:   s.PlayID,
	}
}

func (sp *Story_Province) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, sp)
}

func (sp *Story_Province) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, sp)
}

func (sp *Story_Province) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, sp)
}

func (sp *Story_Province) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sp)
}

func (sp *Story_Province) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, sp)
}

func InsertStory_Provinces(space *nebulagolang.Space, sps ...*Story_Province) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, sps)
}

func UpdateStory_Provinces(space *nebulagolang.Space, sps ...*Story_Province) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, sps)
}

func UpsertStory_Provinces(space *nebulagolang.Space, sps ...*Story_Province) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, sps)
}

func DeleteStory_Provinces(space *nebulagolang.Space, sps ...*Story_Province) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sps...)
}

func DeleteAllStory_Provinces(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Story_Province](space)
}

func DeleteAllStory_ProvincesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Story_Province](space, getPlayIdQuery[Story_Province](playId))
}

func GetStory_ProvinceByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Story_Province] {
	return nebulagolang.GetEdgeByEid[*Story_Province](space, eid)
}

func GetStory_ProvinceByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Story_Province] {
	return nebulagolang.GetEdgeByEid[*Story_Province](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Story_Province]()))
}

func GetAllStory_ProvincesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Province](space, "")
}

func GetAllStory_ProvincesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Province](space, getPlayIdQuery[Story_Province](playId))
}

func GetAllStory_Provinces(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story_Province] {
	return nebulagolang.GetAllEdgesByEdgeType[*Story_Province](space)
}

func GetAllStory_ProvincesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Story_Province] {
	return nebulagolang.GetAllEdgesByQuery[*Story_Province](space, getPlayIdQuery[Story_Province](playId))
}
