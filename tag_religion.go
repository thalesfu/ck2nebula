package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/religion"
	"github.com/thalesfu/paradoxtools/utils"
)

type Religion struct {
	VID  string `nebulakey:"vid" nebulatagname:"religion" nebulatagcomment:"religion" json:"vid,omitempty"`
	Name string `nebulaproperty:"name" description:"religion name" nebulaindexes:"name" json:"name,omitempty"`
	Code string `nebulaproperty:"code" description:"religion code" nebulaindexes:"code" json:"code,omitempty"`
}

func NewReligion(code string) *Religion {
	nebulaReligion := Religion{
		VID:  getReligionVid(code),
		Code: code,
	}
	return &nebulaReligion
}

func NewReligionByData(religion *religion.Religion) *Religion {
	nebulaReligion := utils.Mapping[Religion](religion)
	nebulaReligion.VID = getReligionVid(nebulaReligion.Code)
	return &nebulaReligion
}

func getReligionVid(code string) string {
	return "religion." + code
}

func (r *Religion) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, r)
}

func (r *Religion) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, r)
}

func (r *Religion) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, r)
}

func (r *Religion) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, r)
}

func (r *Religion) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, r)
}

func (r *Religion) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, r)
}

func InsertReligions(space *nebulagolang.Space, cs ...*Religion) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, cs)
}

func UpdateReligions(space *nebulagolang.Space, cs ...*Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, cs)
}

func UpsertReligions(space *nebulagolang.Space, cs ...*Religion) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, cs)
}

func DeleteReligions(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getReligionVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteReligionWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getReligionVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllReligion(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Religion](space)
}

func DeleteAllReligionWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Religion](space)
}

func GetReligionByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Religion] {
	return nebulagolang.GetVertexByVid[*Religion](space, vid)
}

func GetReligionByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*Religion] {
	return nebulagolang.GetVertexByVid[*Religion](space, getReligionVid(code))
}

func GetAllReligion(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Religion] {
	return nebulagolang.GetAllVertexesByQuery[*Religion](space, "")
}

func GetAllReligionVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Religion](space, "")
}

func GetAllReligionCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Religion](space, "", "code", "")
}
