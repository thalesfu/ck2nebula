package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
	"time"
)

type Baron struct {
	VID          string    `nebulakey:"vid" nebulatagname:"baron" nebulatagcomment:"baron" json:"vid,omitempty"`
	Code         string    `nebulaproperty:"code" description:"code" nebulaindexes:"code" json:"code,omitempty"`
	Name         string    `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Type         string    `nebulaproperty:"type" description:"type" nebulaindexes:"type" json:"type,omitempty"`
	BuiltDate    time.Time `mappingalias:"Date" nebulaproperty:"built_date" nebulatype:"Date" description:"end build date" nebulaindexes:"built_date" json:"built_date,omitempty"`
	BuildingDate time.Time `mappingalias:"BuildTime" nebulaproperty:"building_date" nebulatype:"Date" description:"start build date" nebulaindexes:"building_time" json:"building_time,omitempty"`
	Leader       int       `nebulaproperty:"leader" description:"leader" nebulaindexes:"leader" json:"leader,omitempty"`
	PlayID       int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewBaron(playId int, code string) *Baron {
	nebulaBaron := Baron{
		VID:    getBaronVid(playId, code),
		Code:   code,
		PlayID: playId,
	}
	return &nebulaBaron
}

func NewBaronByData(baron *save.Baron) *Baron {
	nebulaBaron := utils.Mapping[Baron](baron)
	nebulaBaron.VID = getBaronVid(nebulaBaron.PlayID, nebulaBaron.Code)
	return &nebulaBaron
}

func getBaronVid(playId int, code string) string {
	return fmt.Sprintf("baron.%s.%d", code, playId)
}

func (b *Baron) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, b)
}

func (b *Baron) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, b)
}

func (b *Baron) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, b)
}

func (b *Baron) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, b)
}

func (b *Baron) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, b)
}

func (b *Baron) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, b)
}

func InsertBarons(space *nebulagolang.Space, bs ...*Baron) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, bs)
}

func UpdateBarons(space *nebulagolang.Space, bs ...*Baron) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, bs)
}

func UpsertBarons(space *nebulagolang.Space, bs ...*Baron) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, bs)
}

func DeleteBarons(space *nebulagolang.Space, playId int, codes ...string) *nebulagolang.Result {
	ts := make([]string, len(codes))
	for i, c := range codes {
		ts[i] = getBaronVid(playId, c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ts...)
}

func DeleteBaronWithEdge(space *nebulagolang.Space, playId int, codes ...string) *nebulagolang.Result {
	ts := make([]string, len(codes))
	for i, c := range codes {
		ts[i] = getBaronVid(playId, c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ts...)
}

func DeleteAllBaron(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Baron](space)
}

func DeleteAllBaronByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByQuery[Baron](space, getPlayIdQuery[Baron](playId))
}

func DeleteAllBaronWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Baron](space)
}

func DeleteAllBaronWithEdgesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByQuery[Baron](space, getPlayIdQuery[Baron](playId))
}

func GetBaronByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Baron] {
	return nebulagolang.GetVertexByVid[*Baron](space, vid)
}

func GetBaronByCode(space *nebulagolang.Space, playId int, code string) *nebulagolang.ResultT[*Baron] {
	return nebulagolang.GetVertexByVid[*Baron](space, getBaronVid(playId, code))
}

func GetAllBaron(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Baron] {
	return nebulagolang.GetAllVertexesByQuery[*Baron](space, "")
}

func GetAllBaronByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Baron] {
	return nebulagolang.GetAllVertexesByQuery[*Baron](space, getPlayIdQuery[Baron](playId))
}

func GetAllBaronsByPlayIdAndName(space *nebulagolang.Space, playId int, name string) *nebulagolang.ResultT[map[string]*Baron] {
	return nebulagolang.GetAllVertexesByQuery[*Baron](space, getPlayIdAndPropertyQuery[Baron](playId, "name", name))
}

func GetAllBaronVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Baron](space, "")
}

func GetAllBaronVidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Baron](space, getPlayIdQuery[Baron](playId))
}

func GetAllBaronCodes(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Baron](space, "", "code", "")
}

func GetAllBaronCodesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Baron](space, getPlayIdQuery[Baron](playId), "code", "")
}

func GetAllBaronNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Baron](space, "", "name", "")
}

func GetAllBaronNamesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Baron](space, getPlayIdQuery[Baron](playId), "name", "")
}
