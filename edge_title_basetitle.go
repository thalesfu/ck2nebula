package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Title_BaseTitle struct {
	Title     *Title `nebulaedgename:"title_basetitle" nebulaedgecomment:"Title -> Base Title" nebulakey:"edgefrom" json:"from,omitempty"`
	BaseTitle *Title `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID    int    `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewTitle_BaseTitle(t *Title, bt *Title) *Title_BaseTitle {
	return &Title_BaseTitle{
		Title:     t,
		BaseTitle: bt,
	}
}

func (tbt *Title_BaseTitle) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, tbt)
}

func (tbt *Title_BaseTitle) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, tbt)
}

func (tbt *Title_BaseTitle) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, tbt)
}

func (tbt *Title_BaseTitle) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tbt)
}

func (tbt *Title_BaseTitle) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, tbt)
}

func InsertTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, tbts)
}

func UpdateTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, tbts)
}

func UpsertTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, tbts)
}

func DeleteTitle_BaseTitles(space *nebulagolang.Space, tbts ...*Title_BaseTitle) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tbts...)
}

func DeleteAllTitle_BaseTitles(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_BaseTitle](space)
}

func DeleteAllTitle_BaseTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Title_BaseTitle](space, getPlayIdQuery[Title_BaseTitle](playId))
}

func GetTitle_BaseTitleByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_BaseTitle] {
	return nebulagolang.GetEdgeByEid[*Title_BaseTitle](space, eid)
}

func GetTitle_BaseTitleByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_BaseTitle] {
	return nebulagolang.GetEdgeByEid[*Title_BaseTitle](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_BaseTitle]()))
}

func GetAllTitle_BaseTitlesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_BaseTitle](space, "")
}

func GetAllTitle_BaseTitlesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_BaseTitle](space, getPlayIdQuery[Title_BaseTitle](playId))
}

func GetAllTitle_BaseTitles(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_BaseTitle] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_BaseTitle](space)
}

func GetAllTitle_BaseTitlesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Title_BaseTitle] {
	return nebulagolang.GetAllEdgesByQuery[*Title_BaseTitle](space, getPlayIdQuery[Title_BaseTitle](playId))
}
