package ck2nebula

import "github.com/thalesfu/nebulagolang"

type ReligionGroup_Religion struct {
	ReligionGroup *ReligionGroup `nebulaedgename:"religiongroup_religion" nebulaedgecomment:"religion group -> religion" nebulakey:"edgefrom" json:"from,omitempty"`
	Religion      *Religion      `nebulakey:"edgeto" json:"to,omitempty"`
}

func NewReligionGroup_Religion(cg *ReligionGroup, c *Religion) *ReligionGroup_Religion {
	return &ReligionGroup_Religion{
		ReligionGroup: cg,
		Religion:      c,
	}
}

func (rgr *ReligionGroup_Religion) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertEdges(space, rgr)
}

func (rgr *ReligionGroup_Religion) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateEdges(space, rgr)
}

func (rgr *ReligionGroup_Religion) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertEdges(space, rgr)
}

func (rgr *ReligionGroup_Religion) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, rgr)
}

func (rgr *ReligionGroup_Religion) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadEdge(space, rgr)
}

func InsertReligionGroup_Religions(space *nebulagolang.Space, rgrs ...*ReligionGroup_Religion) *nebulagolang.Result {
	return nebulagolang.BatchInsertEdges(space, 10000, rgrs)
}

func UpdateReligionGroup_Religions(space *nebulagolang.Space, rgrs ...*ReligionGroup_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpdateEdges(space, 5000, rgrs)
}

func UpsertReligionGroup_Religions(space *nebulagolang.Space, rgrs ...*ReligionGroup_Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpsertEdges(space, 5000, rgrs)
}

func DeleteReligionGroup_Religions(space *nebulagolang.Space, rgrs ...*ReligionGroup_Religion) *nebulagolang.Result {
	return nebulagolang.DeleteEdges(space, rgrs...)
}

func DeleteAllReligionGroup_Religions(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllEdgesByEdgeType[ReligionGroup_Religion](space)
}

func GetReligionGroup_ReligionByEid(space *nebulagolang.Space, eid *nebulagolang.EID) *nebulagolang.ResultT[*ReligionGroup_Religion] {
	return nebulagolang.GetEdgeByEid[*ReligionGroup_Religion](space, eid)
}

func GetReligionGroup_ReligionByFromAndToIds(space *nebulagolang.Space, fromId string, toId string) *nebulagolang.ResultT[*ReligionGroup_Religion] {
	return nebulagolang.GetEdgeByEid[*ReligionGroup_Religion](space, nebulagolang.NewEID(fromId, toId, nebulagolang.GetEdgeName[ReligionGroup_Religion]()))
}

func GetAllReligionGroup_ReligionsEids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllEdgesEIDsByQuery[*ReligionGroup_Religion](space, "")
}

func GetAllReligionGroup_Religions(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*ReligionGroup_Religion] {
	return nebulagolang.GetAllEdgesByEdgeType[*ReligionGroup_Religion](space)
}
