package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
	"time"
)

type Title struct {
	VID         string    `nebulakey:"vid" nebulatagname:"title" nebulatagcomment:"title" json:"vid,omitempty"`
	Name        string    `nebulaproperty:"name" description:"title name" nebulaindexes:"name" json:"name,omitempty"`
	ID          string    `nebulaproperty:"id" description:"title code" nebulaindexes:"code" json:"code,omitempty"`
	Foa         string    `nebulaproperty:"foa" description:"title foa" nebulaindexes:"foa" son:"foa,omitempty"`
	Title       string    `nebulaproperty:"title" description:"title" nebulaindexes:"title" json:"title,omitempty"`
	TitleFemale string    `nebulaproperty:"title_female" description:"female title" nebulaindexes:"title_female" json:"title_female,omitempty"`
	Gender      string    `nebulaproperty:"gender" description:"title gender" nebulaindexes:"gender" json:"gender,omitempty"`
	Active      bool      `nebulaproperty:"active" description:"title is active" nebulaindexes:"active" json:"active,omitempty"`
	IsCustom    bool      `nebulaproperty:"is_custom" description:"is custom title" nebulaindexes:"is_custom" json:"is_custom,omitempty"`
	IsDynamic   bool      `nebulaproperty:"is_dynamic" mappingalias:"Dynamic" description:"is dynamic title" nebulaindexes:"is_dynamic" json:"is_dynamic,omitempty"`
	Nomad       bool      `nebulaproperty:"nomad" description:"title is nomad" nebulaindexes:"nomad" json:"nomad,omitempty"`
	Landless    bool      `nebulaproperty:"landless" description:"title is land less" nebulaindexes:"landless" json:"landless,omitempty"`
	MajorRevolt bool      `nebulaproperty:"major_revolt" description:"title is major revolt" nebulaindexes:"major_revolt" json:"major_revolt,omitempty"`
	Mercenary   bool      `nebulaproperty:"mercenary" description:"title is mercenary" nebulaindexes:"mercenary" json:"mercenary,omitempty"`
	Rebels      bool      `nebulaproperty:"rebels" description:"title is rebels" nebulaindexes:"rebels" json:"rebels,omitempty"`
	Temporary   bool      `nebulaproperty:"temporary" description:"title is temporary" nebulaindexes:"temporary" json:"temporary,omitempty"`
	PlayID      int       `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
	PlayDate    time.Time `nebulaproperty:"play_date" nebulatype:"Date" description:"game play date" nebulaindexes:"play_date" json:"play_date,omitempty"`

	//Holder            int              `paradox_field:"holder" json:"holder,omitempty"`
	//Dynasty           int
}

func NewTitle(id string) *Title {
	nebulaTitle := Title{
		VID: getTitleVid(id),
		ID:  id,
	}
	return &nebulaTitle
}

func NewTitleByData(title *save.Title) *Title {
	nebulaTitle := utils.Mapping[Title](title)
	nebulaTitle.VID = getTitleVid(nebulaTitle.ID)
	return &nebulaTitle
}

func getTitleVid(code string) string {
	return "title." + code
}

func (t *Title) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, t)
}

func (t *Title) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, t)
}

func (t *Title) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, t)
}

func (t *Title) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, t)
}

func (t *Title) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, t)
}

func (t *Title) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, t)
}

func InsertTitles(space *nebulagolang.Space, ts ...*Title) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 10000, ts)
}

func UpdateTitles(space *nebulagolang.Space, ts ...*Title) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 5000, ts)
}

func UpsertTitles(space *nebulagolang.Space, ts ...*Title) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 5000, ts)
}

func DeleteTitles(space *nebulagolang.Space, ids ...string) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getTitleVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ts...)
}

func DeleteTitleWithEdge(space *nebulagolang.Space, ids ...string) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getTitleVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ts...)
}

func DeleteAllTitle(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Title](space)
}

func DeleteAllTitleWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Title](space)
}

func GetTitleByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Title] {
	return nebulagolang.GetVertexByVid[*Title](space, vid)
}

func GetTitleByID(space *nebulagolang.Space, id string) *nebulagolang.ResultT[*Title] {
	return nebulagolang.GetVertexByVid[*Title](space, getTitleVid(id))
}

func GetAllTitle(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Title] {
	return nebulagolang.GetAllVertexesByQuery[*Title](space, "")
}

func GetAllTitleVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Title](space, "")
}

func GetAllTitleCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Title](space, "", "id", "")
}

func GetAllTitleNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Title](space, "", "name", "")
}
