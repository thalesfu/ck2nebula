package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Title_Dynasty struct {
	Title   *Title   `nebulaedgename:"title_dynasty" nebulaedgecomment:"Title -> Dynasty" nebulakey:"edgefrom" json:"from,omitempty"`
	Dynasty *Dynasty `nebulakey:"edgeto" json:"to,omitempty"`
	PlayID  int      `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewTitle_Dynasty(t *Title, d *Dynasty) *Title_Dynasty {
	return &Title_Dynasty{
		Title:   t,
		Dynasty: d,
		PlayID:  t.PlayID,
	}
}

func (td *Title_Dynasty) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, td)
}

func (td *Title_Dynasty) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, td)
}

func (td *Title_Dynasty) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, td)
}

func (td *Title_Dynasty) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, td)
}

func (td *Title_Dynasty) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, td)
}

func InsertTitle_Dynastys(space *nebulagolang.Space, tds ...*Title_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, tds)
}

func UpdateTitle_Dynastys(space *nebulagolang.Space, tds ...*Title_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, tds)
}

func UpsertTitle_Dynastys(space *nebulagolang.Space, tds ...*Title_Dynasty) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, tds)
}

func DeleteTitle_Dynastys(space *nebulagolang.Space, tds ...*Title_Dynasty) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, tds...)
}

func DeleteAllTitle_Dynastys(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Title_Dynasty](space)
}

func DeleteAllTitle_DynastysByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Title_Dynasty](space, getPlayIdQuery[Title_Dynasty](playId))
}

func GetTitle_DynastyByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Title_Dynasty] {
	return nebulagolang.GetEdgeByEid[*Title_Dynasty](space, eid)
}

func GetTitle_DynastyByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Title_Dynasty] {
	return nebulagolang.GetEdgeByEid[*Title_Dynasty](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Title_Dynasty]()))
}

func GetAllTitle_DynastysEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_Dynasty](space, "")
}

func GetAllTitle_DynastysEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Title_Dynasty](space, getPlayIdQuery[Title_Dynasty](playId))
}

func GetAllTitle_Dynastys(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title_Dynasty] {
	return nebulagolang.GetAllEdgesByEdgeType[*Title_Dynasty](space)
}

func GetAllTitle_DynastysByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Title_Dynasty] {
	return nebulagolang.GetAllEdgesByQuery[*Title_Dynasty](space, getPlayIdQuery[Title_Dynasty](playId))
}
