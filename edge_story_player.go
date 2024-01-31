package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Story_Player struct {
	Story   *Story  `nebulaedgename:"story_player" nebulaedgecomment:"Story -> Player" nebulakey:"edgefrom" json:"from,omitempty"`
	Player  *People `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewStory_Player(s *Story, p *People) *Story_Player {
	return &Story_Player{
		Story:   s,
		Player:  p,
		StoryID: s.StoryId,
	}
}

func (sp *Story_Player) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, sp)
}

func (sp *Story_Player) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, sp)
}

func (sp *Story_Player) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, sp)
}

func (sp *Story_Player) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sp)
}

func (sp *Story_Player) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, sp)
}

func InsertStory_Players(space *nebulagolang.Space, sps ...*Story_Player) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, sps)
}

func UpdateStory_Players(space *nebulagolang.Space, sps ...*Story_Player) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, sps)
}

func UpsertStory_Players(space *nebulagolang.Space, sps ...*Story_Player) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, sps)
}

func DeleteStory_Players(space *nebulagolang.Space, sps ...*Story_Player) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, sps...)
}

func DeleteAllStory_Players(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Story_Player](space)
}

func DeleteAllStory_PlayerByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Story_Player](space, getPlayIdQuery[Story_Player](playId))
}

func GetStory_PlayerByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Story_Player] {
	return nebulagolang.GetEdgeByEid[*Story_Player](space, eid)
}

func GetStory_PlayerByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Story_Player] {
	return nebulagolang.GetEdgeByEid[*Story_Player](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Story_Player]()))
}

func GetAllStory_PlayersEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Player](space, "")
}

func GetAllStory_PlayersEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Story_Player](space, getPlayIdQuery[Story_Player](playId))
}

func GetAllStory_Players(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Story_Player] {
	return nebulagolang.GetAllEdgesByEdgeType[*Story_Player](space)
}

func GetAllStory_PlayersByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Story_Player] {
	return nebulagolang.GetAllEdgesByQuery[*Story_Player](space, getPlayIdQuery[Story_Player](playId))
}
