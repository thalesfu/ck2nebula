package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"github.com/thalesfu/paradoxtools/utils"
)

type Province struct {
	VID                            string  `nebulakey:"vid" nebulatagname:"province" nebulatagcomment:"province" json:"vid,omitempty"`
	ID                             int     `nebulaproperty:"id" description:"id" nebulaindexes:"id" json:"id,omitempty"`
	Code                           string  `nebulaproperty:"code" description:"code" nebulaindexes:"code" json:"code,omitempty"`
	Name                           string  `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	MaxSettlements                 int     `nebulaproperty:"max_settlements" description:"max_settlements" nebulaindexes:"max_settlements" json:"max_settlements,omitempty"`
	Culture                        string  `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	CultureName                    string  `nebulaproperty:"culture_name" description:"culture name" nebulaindexes:"culture_name" json:"culture_name,omitempty"`
	Religion                       string  `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	ReligionName                   string  `nebulaproperty:"religion_name" description:"religion name" nebulaindexes:"religion_name" json:"religion_name,omitempty"`
	TechInfantry                   float32 `nebulaproperty:"tech_infantry" description:"tech_infantry" nebulaindexes:"tech_infantry" json:"tech_infantry,omitempty"`
	TechCavalry                    float32 `nebulaproperty:"tech_cavalry" description:"tech_cavalry" nebulaindexes:"tech_cavalry" json:"tech_cavalry,omitempty"`
	TechSkirmish                   float32 `nebulaproperty:"tech_skirmish" description:"tech_skirmish" nebulaindexes:"tech_skirmish" json:"tech_skirmish,omitempty"`
	TechMelee                      float32 `nebulaproperty:"tech_melee" description:"tech_melee" nebulaindexes:"tech_melee" json:"tech_melee,omitempty"`
	TechSiegeEquipment             float32 `nebulaproperty:"tech_siege_equipment" description:"tech_siege_equipment" nebulaindexes:"tech_siege_equipment" json:"tech_siege_equipment,omitempty"`
	TechRecruitment                float32 `nebulaproperty:"tech_recruitment" description:"tech_recruitment" nebulaindexes:"tech_recruitment" json:"tech_recruitment,omitempty"`
	TechCastleConstruction         float32 `nebulaproperty:"tech_castle_construction" description:"tech_castle_construction" nebulaindexes:"tech_castle_construction" json:"tech_castle_construction,omitempty"`
	TechCityConstruction           float32 `nebulaproperty:"tech_city_construction" description:"tech_city_construction" nebulaindexes:"tech_city_construction" json:"tech_city_construction,omitempty"`
	TechTempleConstruction         float32 `nebulaproperty:"tech_temple_construction" description:"tech_temple_construction" nebulaindexes:"tech_temple_construction" json:"tech_temple_construction,omitempty"`
	TechFortificationsConstruction float32 `nebulaproperty:"tech_fortifications_construction" description:"tech_fortifications_construction" nebulaindexes:"tech_fortifications_construction" json:"tech_fortifications_construction,omitempty"`
	TechTradePractices             float32 `nebulaproperty:"tech_trade_practices" description:"tech_trade_practices" nebulaindexes:"tech_trade_practices" json:"tech_trade_practices,omitempty"`
	TechConstruction               float32 `nebulaproperty:"tech_construction" description:"tech_construction" nebulaindexes:"tech_construction" json:"tech_construction,omitempty"`
	TechNaval                      float32 `nebulaproperty:"tech_naval" description:"tech_naval" nebulaindexes:"tech_naval" json:"tech_naval,omitempty"`
	TechNobleCustoms               float32 `nebulaproperty:"tech_noble_customs" description:"tech_noble_customs" nebulaindexes:"tech_noble_customs" json:"tech_noble_customs,omitempty"`
	TechPopularCustoms             float32 `nebulaproperty:"tech_popular_customs" description:"tech_popular_customs" nebulaindexes:"tech_popular_customs" json:"tech_popular_customs,omitempty"`
	TechReligiousCustoms           float32 `nebulaproperty:"tech_religious_customs" description:"tech_religious_customs" nebulaindexes:"tech_religious_customs" json:"tech_religious_customs,omitempty"`
	TechMajesty                    float32 `nebulaproperty:"tech_majesty" description:"tech_majesty" nebulaindexes:"tech_majesty" json:"tech_majesty,omitempty"`
	TechCultureFlex                float32 `nebulaproperty:"tech_culture_flex" description:"tech_culture_flex" nebulaindexes:"tech_culture_flex" json:"tech_culture_flex,omitempty"`
	TechLegalism                   float32 `nebulaproperty:"tech_legalism" description:"tech_legalism" nebulaindexes:"tech_legalism" json:"tech_legalism,omitempty"`
	PrimarySettlement              string  `nebulaproperty:"primary_settlement" description:"primary_settlement" nebulaindexes:"primary_settlement" json:"primary_settlement,omitempty"`
	PlayID                         int     `nebulaproperty:"play_id" description:"game play id" nebulaindexes:"play_id" json:"play_id,omitempty"`
}

func NewProvince(playId int, provinceId int) *Province {
	nebulaProvince := Province{
		VID:    getProvinceVid(playId, provinceId),
		ID:     provinceId,
		PlayID: playId,
	}
	return &nebulaProvince
}

func NewProvinceByData(province *save.Province) *Province {
	nebulaProvince := utils.Mapping[Province](province)
	nebulaProvince.VID = getProvinceVid(nebulaProvince.PlayID, nebulaProvince.ID)

	if province.Technology != nil && len(province.Technology.TechLevels) > 0 {
		for i, v := range province.Technology.TechLevels {
			switch i {
			case 0:
				nebulaProvince.TechInfantry = v
			case 1:
				nebulaProvince.TechCavalry = v
			case 2:
				nebulaProvince.TechSkirmish = v
			case 3:
				nebulaProvince.TechMelee = v
			case 4:
				nebulaProvince.TechSiegeEquipment = v
			case 5:
				nebulaProvince.TechRecruitment = v
			case 6:
				nebulaProvince.TechCastleConstruction = v
			case 7:
				nebulaProvince.TechCityConstruction = v
			case 8:
				nebulaProvince.TechTempleConstruction = v
			case 9:
				nebulaProvince.TechFortificationsConstruction = v
			case 10:
				nebulaProvince.TechTradePractices = v
			case 11:
				nebulaProvince.TechConstruction = v
			case 12:
				nebulaProvince.TechNaval = v
			case 13:
				nebulaProvince.TechNobleCustoms = v
			case 14:
				nebulaProvince.TechPopularCustoms = v
			case 15:
				nebulaProvince.TechReligiousCustoms = v
			case 16:
				nebulaProvince.TechMajesty = v
			case 17:
				nebulaProvince.TechCultureFlex = v
			case 18:
				nebulaProvince.TechLegalism = v
			}
		}
	}

	return &nebulaProvince
}

func getProvinceVid(playId int, provinceId int) string {
	return fmt.Sprintf("province.%d.%d", provinceId, playId)
}

func (p *Province) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, p)
}

func (p *Province) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, p)
}

func (p *Province) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, p)
}

func (p *Province) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, p)
}

func (p *Province) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, p)
}

func (p *Province) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, p)
}

func InsertProvinces(space *nebulagolang.Space, ps ...*Province) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ps)
}

func UpdateProvinces(space *nebulagolang.Space, ps ...*Province) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ps)
}

func UpsertProvinces(space *nebulagolang.Space, ps ...*Province) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ps)
}

func DeleteProvinces(space *nebulagolang.Space, playId int, ids ...int) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getProvinceVid(playId, c)
	}
	return nebulagolang.DeleteVertexesByVids(space, ts...)
}

func DeleteProvinceWithEdge(space *nebulagolang.Space, playId int, ids ...int) *nebulagolang.Result {
	ts := make([]string, len(ids))
	for i, c := range ids {
		ts[i] = getProvinceVid(playId, c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, ts...)
}

func DeleteAllProvince(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Province](space)
}

func DeleteAllProvinceByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByQuery[Province](space, getPlayIdQuery[Province](playId))
}

func DeleteAllProvinceWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Province](space)
}

func DeleteAllProvinceWithEdgesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByQuery[Province](space, getPlayIdQuery[Province](playId))
}

func GetProvinceByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Province] {
	return nebulagolang.GetVertexByVid[*Province](space, vid)
}

func GetProvinceByID(space *nebulagolang.Space, playId int, id int) *nebulagolang.ResultT[*Province] {
	return nebulagolang.GetVertexByVid[*Province](space, getProvinceVid(playId, id))
}

func GetAllProvinces(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Province] {
	return nebulagolang.GetAllVertexesByQuery[*Province](space, "")
}

func GetAllProvincesByPlayIdAndName(space *nebulagolang.Space, playId int, name string) *nebulagolang.ResultT[map[string]*Province] {
	return nebulagolang.GetAllVertexesByQuery[*Province](space, getPlayIdAndPropertyQuery[Province](playId, "name", name))
}

func GetAllProvincesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]*Province] {
	return nebulagolang.GetAllVertexesByQuery[*Province](space, getPlayIdQuery[Province](playId))
}

func GetAllProvincesVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Province](space, "")
}

func GetAllProvincesVidsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Province](space, getPlayIdQuery[Province](playId))
}

func GetAllProvincesIds(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Province](space, "", "id", "")
}

func GetAllProvincesIdsByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Province](space, getPlayIdQuery[Province](playId), "id", "")
}

func GetAllProvincesNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Province](space, "", "name", "")
}

func GetAllProvinceNamesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Province](space, getPlayIdQuery[Province](playId), "name", "")
}

func GetAllProvincesCodes(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Province](space, "", "code", "")
}

func GetAllProvinceCodesByPlayId(space *nebulagolang.Space, playId int) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Province](space, getPlayIdQuery[Province](playId), "code", "")
}
