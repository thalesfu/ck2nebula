package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type People_ClaimTitle struct {
	People     *People `nebulaedgename:"people_claimtitle" nebulaedgecomment:"People -> ClaimTitle" nebulakey:"edgefrom" json:"from,omitempty"`
	ClaimTitle *Title  `nebulakey:"edgeto" json:"to,omitempty"`
	Pressed    bool    `nebulaproperty:"pressed" description:"pressed" nebulaindexes:"pressed" json:"pressed,omitempty"`
	Weak       bool    `nebulaproperty:"weak" description:"weak" nebulaindexes:"weak" json:"weak,omitempty"`
	StoryID    int     `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewPeople_ClaimTitle(p *People, ct *Title) *People_ClaimTitle {
	return &People_ClaimTitle{
		People:     p,
		ClaimTitle: ct,
		StoryID:    p.StoryID,
	}
}

func (pct *People_ClaimTitle) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, pct)
}

func (pct *People_ClaimTitle) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, pct)
}

func (pct *People_ClaimTitle) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, pct)
}

func (pct *People_ClaimTitle) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pct)
}

func (pct *People_ClaimTitle) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, pct)
}

func InsertPeople_ClaimTitles(space *nebulagolang.Space, pcts ...*People_ClaimTitle) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, pcts)
}

func UpdatePeople_ClaimTitles(space *nebulagolang.Space, pcts ...*People_ClaimTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, pcts)
}

func UpsertPeople_ClaimTitles(space *nebulagolang.Space, pcts ...*People_ClaimTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, pcts)
}

func DeletePeople_ClaimTitles(space *nebulagolang.Space, pcts ...*People_ClaimTitle) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, pcts...)
}

func DeleteAllPeople_ClaimTitles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[People_ClaimTitle](space)
}

func DeleteAllPeople_ClaimTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[People_ClaimTitle](space, getPlayIdQuery[People_ClaimTitle](playId))
}

func GetPeople_ClaimTitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*People_ClaimTitle] {
	return nebulagolang.GetEdgeByEid[*People_ClaimTitle](space, eid)
}

func GetPeople_ClaimTitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*People_ClaimTitle] {
	return nebulagolang.GetEdgeByEid[*People_ClaimTitle](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[People_ClaimTitle]()))
}

func GetAllPeople_ClaimTitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_ClaimTitle](space, "")
}

func GetAllPeople_ClaimTitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*People_ClaimTitle](space, getPlayIdQuery[People_ClaimTitle](playId))
}

func GetAllPeople_ClaimTitles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*People_ClaimTitle] {
	return nebulagolang.GetAllEdgesByEdgeType[*People_ClaimTitle](space)
}

func GetAllPeople_ClaimTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*People_ClaimTitle] {
	return nebulagolang.GetAllEdgesByQuery[*People_ClaimTitle](space, getPlayIdQuery[People_ClaimTitle](playId))
}
