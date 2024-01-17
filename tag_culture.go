package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/culture"
	"github.com/thalesfu/paradoxtools/utils"
)

type Culture struct {
	VID                   string `nebulakey:"vid" nebulatagname:"culture" nebulatagcomment:"culture" json:"vid,omitempty"`
	Name                  string `nebulaproperty:"name" description:"culture name" nebulaindexes:"name" json:"name,omitempty"`
	Code                  string `nebulaproperty:"code" description:"culture code" nebulaindexes:"code" json:"code,omitempty"`
	FromDynastyPrefix     string `nebulaproperty:"from_dynasty_prefix" description:"from dynasty prefix" nebulaindexes:"from_dynasty_prefix" json:"from_dynasty_prefix,omitempty"`
	MaleParonym           string `nebulaproperty:"male_paron_ym" description:"male paron ym" nebulaindexes:"male_paron_ym" json:"male_paron_ym,omitempty"`
	FemaleParonym         string `nebulaproperty:"female_paron_ym" description:"female paron ym" nebulaindexes:"female_paron_ym" json:"female_paron_ym,omitempty"`
	DynastyTitleNames     bool   `nebulaproperty:"dynasty_title_names" description:"dynasty title names" nebulaindexes:"dynasty_title_names" json:"dynasty_title_names,omitempty"`
	FounderNamedDynasties bool   `nebulaproperty:"founder_named_dynasties" description:"founder named dynasties" nebulaindexes:"founder_named_dynasties" json:"founder_named_dynasties,omitempty"`
	Caste                 bool   `nebulaproperty:"caste" description:"caste" nebulaindexes:"caste" json:"caste,omitempty"`
	DynastyNameFirst      bool   `nebulaproperty:"dynasty_name_first" description:"dynasty name first" nebulaindexes:"dynasty_name_first" json:"dynasty_name_first,omitempty"`
	DukesCalledKings      bool   `nebulaproperty:"dukes_called_kings" description:"dukes called kings" nebulaindexes:"dukes_called_kings" json:"dukes_called_kings,omitempty"`
	CountTitlesHidden     bool   `nebulaproperty:"count_titles_hidden" description:"count titles hidden" nebulaindexes:"count_titles_hidden" json:"count_titles_hidden,omitempty"`
	BaronTitlesHidden     bool   `nebulaproperty:"baron_titles_hidden" description:"baron titles hidden" nebulaindexes:"baron_titles_hidden" json:"baron_titles_hidden,omitempty"`
	AllowLooting          bool   `nebulaproperty:"allow_looting" description:"allow looting" nebulaindexes:"allow_looting" json:"allow_looting,omitempty"`
}

func NewCulture(code string) *Culture {
	nebulaCulture := Culture{
		VID:  getCultureVid(code),
		Code: code,
	}
	return &nebulaCulture
}

func NewCultureByData(culture *culture.Culture) *Culture {
	nebulaCulture := utils.Mapping[Culture](culture)
	nebulaCulture.VID = getCultureVid(nebulaCulture.Code)
	return &nebulaCulture
}

func getCultureVid(code string) string {
	return "culture." + code
}

func (c *Culture) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, c)
}

func (c *Culture) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, c)
}

func (c *Culture) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, c)
}

func (c *Culture) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, c)
}

func (c *Culture) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, c)
}

func (c *Culture) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, c)
}

func InsertCultures(space *nebulagolang.Space, cs ...*Culture) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 10000, cs)
}

func UpdateCultures(space *nebulagolang.Space, cs ...*Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 5000, cs)
}

func UpsertCultures(space *nebulagolang.Space, cs ...*Culture) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 5000, cs)
}

func DeleteCultures(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getCultureVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteCultureWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getCultureVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllCulture(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Culture](space)
}

func DeleteAllCultureWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Culture](space)
}

func GetCultureByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Culture] {
	return nebulagolang.GetVertexByVid[*Culture](space, vid)
}

func GetCultureByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*Culture] {
	return nebulagolang.GetVertexByVid[*Culture](space, getCultureVid(code))
}

func GetAllCulture(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Culture] {
	return nebulagolang.GetAllVertexesByQuery[*Culture](space, "")
}

func GetAllCultureVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Culture](space, "")
}

func GetAllCultureCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Culture](space, "", "code", "")
}
