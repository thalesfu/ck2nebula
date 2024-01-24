package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Title_LiegeTitle struct {
	Title      *Title `nebulaedgename:"title_liegetitle" nebulaedgecomment:"Title -> Liege Title" nebulakey:"edgefrom" json:"from,omitempty"`
	LiegeTitle *Title `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID     int    `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewTitle_LiegeTitle(t *Title, lt *Title) *Title_LiegeTitle {
	return &Title_LiegeTitle{
		Title:      t,
		LiegeTitle: lt,
		PlayID:     t.PlayID,
	}
}

func (tlt *Title_LiegeTitle) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, tlt)
}

func (tlt *Title_LiegeTitle) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, tlt)
}

func (tlt *Title_LiegeTitle) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, tlt)
}

func (tlt *Title_LiegeTitle) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tlt)
}

func (tlt *Title_LiegeTitle) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, tlt)
}

func InsertTitle_LiegeTitles(space *nebulagolang.Space, tlts ...*Title_LiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, tlts)
}

func UpdateTitle_LiegeTitles(space *nebulagolang.Space, tlts ...*Title_LiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, tlts)
}

func UpsertTitle_LiegeTitles(space *nebulagolang.Space, tlts ...*Title_LiegeTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, tlts)
}

func DeleteTitle_LiegeTitles(space *nebulagolang.Space, tlts ...*Title_LiegeTitle) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tlts...)
}

func DeleteAllTitle_LiegeTitles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_LiegeTitle](space)
}

func DeleteAllTitle_LiegeTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Title_LiegeTitle](space, getPlayIdQuery[Title_LiegeTitle](playId))
}

func GetTitle_LiegeTitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_LiegeTitle] {
	return nebulagolang.GetEdgeByEid[*Title_LiegeTitle](space, eid)
}

func GetTitle_LiegeTitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_LiegeTitle] {
	return nebulagolang.GetEdgeByEid[*Title_LiegeTitle](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_LiegeTitle]()))
}

func GetAllTitle_LiegeTitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_LiegeTitle](space, "")
}

func GetAllTitle_LiegeTitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_LiegeTitle](space, getPlayIdQuery[Title_LiegeTitle](playId))
}

func GetAllTitle_LiegeTitles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_LiegeTitle] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_LiegeTitle](space)
}

func GetAllTitle_LiegeTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Title_LiegeTitle] {
	return nebulagolang.GetAllEdgesByQuery[*Title_LiegeTitle](space, getPlayIdQuery[Title_LiegeTitle](playId))
}
