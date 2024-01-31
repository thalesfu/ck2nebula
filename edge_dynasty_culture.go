package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
)

type Dynasty_Culture struct {
	Dynasty *Dynasty `nebulaedgename:"dynasty_culture" nebulaedgecomment:"Dynasty -> Culture" nebulakey:"edgefrom" json:"from,omitempty"`
	Culture *Culture `nebulakey:"edgeto" json:"to,omitempty"`
	StoryID int      `nebulaproperty:"story_id" mappingalias:"PlayID" description:"game play id" nebulaindexes:"story_id" json:"story_id,omitempty"`
}

func NewDynasty_Culture(d *Dynasty, c *Culture) *Dynasty_Culture {
	return &Dynasty_Culture{
		Dynasty: d,
		Culture: c,
		StoryID: d.StoryID,
	}
}

func (dc *Dynasty_Culture) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, dc)
}

func (dc *Dynasty_Culture) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, dc)
}

func (dc *Dynasty_Culture) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, dc)
}

func (dc *Dynasty_Culture) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, dc)
}

func (dc *Dynasty_Culture) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, dc)
}

func InsertDynasty_Cultures(space *nebulagolang.Space, dcs ...*Dynasty_Culture) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 250, dcs)
}

func UpdateDynasty_Cultures(space *nebulagolang.Space, dcs ...*Dynasty_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 250, dcs)
}

func UpsertDynasty_Cultures(space *nebulagolang.Space, dcs ...*Dynasty_Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 250, dcs)
}

func DeleteDynasty_Cultures(space *nebulagolang.Space, dcs ...*Dynasty_Culture) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, dcs...)
}

func DeleteAllDynasty_Cultures(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[Dynasty_Culture](space)
}

func DeleteAllDynasty_CulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByQuery[Dynasty_Culture](space, getPlayIdQuery[Dynasty_Culture](playId))
}

func GetDynasty_CultureByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*Dynasty_Culture] {
	return nebulagolang.GetEdgeByEid[*Dynasty_Culture](space, eid)
}

func GetDynasty_CultureByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*Dynasty_Culture] {
	return nebulagolang.GetEdgeByEid[*Dynasty_Culture](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[Dynasty_Culture]()))
}

func GetAllDynasty_CulturesEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Dynasty_Culture](space, "")
}

func GetAllDynasty_CulturesEidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*Dynasty_Culture](space, getPlayIdQuery[Dynasty_Culture](playId))
}

func GetAllDynasty_Cultures(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Dynasty_Culture] {
	return nebulagolang.GetAllEdgesByEdgeType[*Dynasty_Culture](space)
}

func GetAllDynasty_CulturesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Dynasty_Culture] {
	return nebulagolang.GetAllEdgesByQuery[*Dynasty_Culture](space, getPlayIdQuery[Dynasty_Culture](playId))
}
