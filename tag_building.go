package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/building"
	"github.com/thalesfu/paradoxtools/utils"
)

type Building struct {
	VID                      string  `nebulakey:"vid" nebulatagname:"building" nebulatagcomment:"building" json:"vid,omitempty"`
	Code                     string  `nebulaproperty:"code" description:"code" nebulaindexes:"code" json:"code,omitempty"`
	Name                     string  `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Description              string  `nebulaproperty:"description" description:"description" json:"description,omitempty"`
	Group                    string  `nebulaproperty:"group" description:"group" nebulaindexes:"group" json:"group,omitempty"`
	GroupName                string  `nebulaproperty:"group_name" description:"group_name" nebulaindexes:"group_name" json:"group_name,omitempty"`
	GroupDescription         string  `nebulaproperty:"group_description" description:"group_description" json:"group_description,omitempty"`
	Port                     bool    `nebulaproperty:"port" description:"port" nebulaindexes:"port" json:"port,omitempty"`
	AddNumberToName          bool    `nebulaproperty:"add_number_to_name" description:"add_number_to_name" nebulaindexes:"add_number_to_name" json:"add_number_to_name,omitempty"`
	Scouting                 bool    `nebulaproperty:"scouting" description:"scouting" nebulaindexes:"scouting" json:"scouting,omitempty"`
	Desc                     string  `nebulaproperty:"hasdesc" description:"has desc" nebulaindexes:"hasdesc" json:"hasdesc,omitempty"`
	UpgradesFrom             string  `nebulaproperty:"upgrades_from" description:"upgrades_from" nebulaindexes:"upgrades_from" json:"upgrades_from,omitempty"`
	Replaces                 string  `nebulaproperty:"replaces" description:"replaces" nebulaindexes:"replaces" json:"replaces,omitempty"`
	ConvertToTribal          string  `nebulaproperty:"convert_to_tribal" description:"convert_to_tribal" nebulaindexes:"convert_to_tribal" json:"convert_to_tribal,omitempty"`
	ConvertToCity            string  `nebulaproperty:"convert_to_city" description:"convert_to_city" nebulaindexes:"convert_to_city" json:"convert_to_city,omitempty"`
	ConvertToCastle          string  `nebulaproperty:"convert_to_castle" description:"convert_to_castle" nebulaindexes:"convert_to_castle" json:"convert_to_castle,omitempty"`
	FortLevel                float32 `nebulaproperty:"fort_level" description:"fort_level" nebulaindexes:"fort_level" json:"fort_level,omitempty"`
	ExtraTechBuildingStart   float32 `nebulaproperty:"extra_tech_building_start" description:"extra_tech_building_start" nebulaindexes:"extra_tech_building_start" json:"extra_tech_building_start,omitempty"`
	LevySize                 float32 `nebulaproperty:"levy_size" description:"levy_size" nebulaindexes:"levy_size" json:"levy_size,omitempty"`
	TaxIncome                float32 `nebulaproperty:"tax_income" description:"tax_income" nebulaindexes:"tax_income" json:"tax_income,omitempty"`
	GarrisonSize             float32 `nebulaproperty:"garrison_size" description:"garrison_size" nebulaindexes:"garrison_size" json:"garrison_size,omitempty"`
	LandMorale               float32 `nebulaproperty:"land_morale" description:"land_morale" nebulaindexes:"land_morale" json:"land_morale,omitempty"`
	LevyReinforceRate        float32 `nebulaproperty:"levy_reinforce_rate" description:"levy_reinforce_rate" nebulaindexes:"levy_reinforce_rate" json:"levy_reinforce_rate,omitempty"`
	LocalRevoltRisk          float32 `nebulaproperty:"local_revolt_risk" description:"local_revolt_risk" nebulaindexes:"local_revolt_risk" json:"local_revolt_risk,omitempty"`
	LiegePrestige            float32 `nebulaproperty:"liege_prestige" description:"liege_prestige" nebulaindexes:"liege_prestige" json:"liege_prestige,omitempty"`
	ArchersOffensive         float32 `nebulaproperty:"archers_offensive" description:"archers_offensive" nebulaindexes:"archers_offensive" json:"archers_offensive,omitempty"`
	PikemenDefensive         float32 `nebulaproperty:"pikemen_defensive" description:"pikemen_defensive" nebulaindexes:"pikemen_defensive" json:"pikemen_defensive,omitempty"`
	HeavyInfantryDefensive   float32 `nebulaproperty:"heavy_infantry_defensive" description:"heavy_infantry_defensive" nebulaindexes:"heavy_infantry_defensive" json:"heavy_infantry_defensive,omitempty"`
	HorseArchersOffensive    float32 `nebulaproperty:"horse_archers_offensive" description:"horse_archers_offensive" nebulaindexes:"horse_archers_offensive" json:"horse_archers_offensive,omitempty"`
	HorseArchersMorale       float32 `nebulaproperty:"horse_archers_morale" description:"horse_archers_morale" nebulaindexes:"horse_archers_morale" json:"horse_archers_morale,omitempty"`
	KnightsOffensive         float32 `nebulaproperty:"knights_offensive" description:"knights_offensive" nebulaindexes:"knights_offensive" json:"knights_offensive,omitempty"`
	LightCavalryOffensive    float32 `nebulaproperty:"light_cavalry_offensive" description:"light_cavalry_offensive" nebulaindexes:"light_cavalry_offensive" json:"light_cavalry_offensive,omitempty"`
	PikemenMorale            float32 `nebulaproperty:"pikemen_morale" description:"pikemen_morale" nebulaindexes:"pikemen_morale" json:"pikemen_morale,omitempty"`
	HeavyInfantryOffensive   float32 `nebulaproperty:"heavy_infantry_offensive" description:"heavy_infantry_offensive" nebulaindexes:"heavy_infantry_offensive" json:"heavy_infantry_offensive,omitempty"`
	LightCavalryDefensive    float32 `nebulaproperty:"light_cavalry_defensive" description:"light_cavalry_defensive" nebulaindexes:"light_cavalry_defensive" json:"light_cavalry_defensive,omitempty"`
	CamelCavalryDefensive    float32 `nebulaproperty:"camel_cavalry_defensive" description:"camel_cavalry_defensive" nebulaindexes:"camel_cavalry_defensive" json:"camel_cavalry_defensive,omitempty"`
	LightInfantryOffensive   float32 `nebulaproperty:"light_infantry_offensive" description:"light_infantry_offensive" nebulaindexes:"light_infantry_offensive" json:"light_infantry_offensive,omitempty"`
	LightInfantryDefensive   float32 `nebulaproperty:"light_infantry_defensive" description:"light_infantry_defensive" nebulaindexes:"light_infantry_defensive" json:"light_infantry_defensive,omitempty"`
	KnightsDefensive         float32 `nebulaproperty:"knights_defensive" description:"knights_defensive" nebulaindexes:"knights_defensive" json:"knights_defensive,omitempty"`
	PikemenOffensive         float32 `nebulaproperty:"pikemen_offensive" description:"pikemen_offensive" nebulaindexes:"pikemen_offensive" json:"pikemen_offensive,omitempty"`
	HeavyInfantryMorale      float32 `nebulaproperty:"heavy_infantry_morale" description:"heavy_infantry_morale" nebulaindexes:"heavy_infantry_morale" json:"heavy_infantry_morale,omitempty"`
	WarElephantsOffensive    float32 `nebulaproperty:"war_elephants_offensive" description:"war_elephants_offensive" nebulaindexes:"war_elephants_offensive" json:"war_elephants_offensive,omitempty"`
	WarElephantsDefensive    float32 `nebulaproperty:"war_elephants_defensive" description:"war_elephants_defensive" nebulaindexes:"war_elephants_defensive" json:"war_elephants_defensive,omitempty"`
	ArchersDefensive         float32 `nebulaproperty:"archers_defensive" description:"archers_defensive" nebulaindexes:"archers_defensive" json:"archers_defensive,omitempty"`
	LightInfantryMorale      float32 `nebulaproperty:"light_infantry_morale" description:"light_infantry_morale" nebulaindexes:"light_infantry_morale" json:"light_infantry_morale,omitempty"`
	LightCavalryMorale       float32 `nebulaproperty:"light_cavalry_morale" description:"light_cavalry_morale" nebulaindexes:"light_cavalry_morale" json:"light_cavalry_morale,omitempty"`
	TechGrowthModifier       float32 `nebulaproperty:"tech_growth_modifier" description:"tech_growth_modifier" nebulaindexes:"tech_growth_modifier" json:"tech_growth_modifier,omitempty"`
	BuildingTechpoints       float32 `nebulaproperty:"building_techpoints" description:"building_techpoints" nebulaindexes:"building_techpoints" json:"building_techpoints,omitempty"`
	LiegePiety               float32 `nebulaproperty:"liege_piety" description:"liege_piety" nebulaindexes:"liege_piety" json:"liege_piety,omitempty"`
	EconomyTechpoints        float32 `nebulaproperty:"economy_techpoints" description:"economy_techpoints" nebulaindexes:"economy_techpoints" json:"economy_techpoints,omitempty"`
	DiseaseDefence           float32 `nebulaproperty:"disease_defence" description:"disease_defence" nebulaindexes:"disease_defence" json:"disease_defence,omitempty"`
	LocalBuildTimeModifier   float32 `nebulaproperty:"local_build_time_modifier" description:"local_build_time_modifier" nebulaindexes:"local_build_time_modifier" json:"local_build_time_modifier,omitempty"`
	LocalBuildCostModifier   float32 `nebulaproperty:"local_build_cost_modifier" description:"local_build_cost_modifier" nebulaindexes:"local_build_cost_modifier" json:"local_build_cost_modifier,omitempty"`
	MilitaryTechpoints       float32 `nebulaproperty:"military_techpoints" description:"military_techpoints" nebulaindexes:"military_techpoints" json:"military_techpoints,omitempty"`
	MonthlyCharacterPrestige float32 `nebulaproperty:"monthly_character_prestige" description:"monthly_character_prestige" nebulaindexes:"monthly_character_prestige" json:"monthly_character_prestige,omitempty"`
	Fertility                float32 `nebulaproperty:"fertility" description:"fertility" nebulaindexes:"fertility" json:"fertility,omitempty"`
	MonthlyCharacterPiety    float32 `nebulaproperty:"monthly_character_piety" description:"monthly_character_piety" nebulaindexes:"monthly_character_piety" json:"monthly_character_piety,omitempty"`
	HorseArchersDefensive    float32 `nebulaproperty:"horse_archers_defensive" description:"horse_archers_defensive" nebulaindexes:"horse_archers_defensive" json:"horse_archers_defensive,omitempty"`
	GlobalMovementSpeed      float32 `nebulaproperty:"global_movement_speed" description:"global_movement_speed" nebulaindexes:"global_movement_speed" json:"global_movement_speed,omitempty"`
	PopulationGrowth         float32 `nebulaproperty:"population_growth" description:"population_growth" nebulaindexes:"population_growth" json:"population_growth,omitempty"`
	NomadTaxModifier         float32 `nebulaproperty:"nomad_tax_modifier" description:"nomad_tax_modifier" nebulaindexes:"nomad_tax_modifier" json:"nomad_tax_modifier,omitempty"`
	GlobalSupplyLimit        float32 `nebulaproperty:"global_supply_limit" description:"global_supply_limit" nebulaindexes:"global_supply_limit" json:"global_supply_limit,omitempty"`
	MovedCapitalMonthsMult   float32 `nebulaproperty:"moved_capital_months_mult" description:"moved_capital_months_mult" nebulaindexes:"moved_capital_months_mult" json:"moved_capital_months_mult,omitempty"`
	HordeMaintenenceCost     float32 `nebulaproperty:"horde_maintenence_cost" description:"horde_maintenence_cost" nebulaindexes:"horde_maintenence_cost" json:"horde_maintenence_cost,omitempty"`
	GlobalWinterSupply       float32 `nebulaproperty:"global_winter_supply" description:"global_winter_supply" nebulaindexes:"global_winter_supply" json:"global_winter_supply,omitempty"`
	MaxPopulationMult        float32 `nebulaproperty:"max_population_mult" description:"max_population_mult" nebulaindexes:"max_population_mult" json:"max_population_mult,omitempty"`
	GlobalTradeRouteValue    float32 `nebulaproperty:"global_trade_route_value" description:"global_trade_route_value" nebulaindexes:"global_trade_route_value" json:"global_trade_route_value,omitempty"`
	CamelCavalryOffensive    float32 `nebulaproperty:"camel_cavalry_offensive" description:"camel_cavalry_offensive" nebulaindexes:"camel_cavalry_offensive" json:"camel_cavalry_offensive,omitempty"`
	TradeRouteWealth         float32 `nebulaproperty:"trade_route_wealth" description:"trade_route_wealth" nebulaindexes:"trade_route_wealth" json:"trade_route_wealth,omitempty"`
	GoldCost                 int     `nebulaproperty:"gold_cost" description:"gold_cost" nebulaindexes:"gold_cost" json:"gold_cost,omitempty"`
	BuildTime                int     `nebulaproperty:"build_time" description:"build_time" nebulaindexes:"build_time" json:"build_time,omitempty"`
	AiCreationFactor         int     `nebulaproperty:"ai_creation_factor" description:"ai_creation_factor" nebulaindexes:"ai_creation_factor" json:"ai_creation_factor,omitempty"`
	LightInfantry            int     `nebulaproperty:"light_infantry" description:"light_infantry" nebulaindexes:"light_infantry" json:"light_infantry,omitempty"`
	Archers                  int     `nebulaproperty:"archers" description:"archers" nebulaindexes:"archers" json:"archers,omitempty"`
	Retinuesize              int     `nebulaproperty:"retinuesize" description:"retinuesize" nebulaindexes:"retinuesize" json:"retinuesize,omitempty"`
	HeavyInfantry            int     `nebulaproperty:"heavy_infantry" description:"heavy_infantry" nebulaindexes:"heavy_infantry" json:"heavy_infantry,omitempty"`
	Pikemen                  int     `nebulaproperty:"pikemen" description:"pikemen" nebulaindexes:"pikemen" json:"pikemen,omitempty"`
	LightCavalry             int     `nebulaproperty:"light_cavalry" description:"light_cavalry" nebulaindexes:"light_cavalry" json:"light_cavalry,omitempty"`
	Knights                  int     `nebulaproperty:"knights" description:"knights" nebulaindexes:"knights" json:"knights,omitempty"`
	CourtSizeModifier        int     `nebulaproperty:"court_size_modifier" description:"court_size_modifier" nebulaindexes:"court_size_modifier" json:"court_size_modifier,omitempty"`
	Galleys                  int     `nebulaproperty:"galleys" description:"galleys" nebulaindexes:"galleys" json:"galleys,omitempty"`
	HorseArchers             int     `nebulaproperty:"horse_archers" description:"horse_archers" nebulaindexes:"horse_archers" json:"horse_archers,omitempty"`
	CamelCavalry             int     `nebulaproperty:"camel_cavalry" description:"camel_cavalry" nebulaindexes:"camel_cavalry" json:"camel_cavalry,omitempty"`
	WarElephants             int     `nebulaproperty:"war_elephants" description:"war_elephants" nebulaindexes:"war_elephants" json:"war_elephants,omitempty"`
	Tradevalue               int     `nebulaproperty:"tradevalue" description:"tradevalue" nebulaindexes:"tradevalue" json:"tradevalue,omitempty"`
	MaxTradeposts            int     `nebulaproperty:"max_tradeposts" description:"max_tradeposts" nebulaindexes:"max_tradeposts" json:"max_tradeposts,omitempty"`
	Stewardship              int     `nebulaproperty:"stewardship" description:"stewardship" nebulaindexes:"stewardship" json:"stewardship,omitempty"`
	Diplomacy                int     `nebulaproperty:"diplomacy" description:"diplomacy" nebulaindexes:"diplomacy" json:"diplomacy,omitempty"`
	Martial                  int     `nebulaproperty:"martial" description:"martial" nebulaindexes:"martial" json:"martial,omitempty"`
	ChurchOpinion            int     `nebulaproperty:"church_opinion" description:"church_opinion" nebulaindexes:"church_opinion" json:"church_opinion,omitempty"`
	Learning                 int     `nebulaproperty:"learning" description:"learning" nebulaindexes:"learning" json:"learning,omitempty"`
	Intrigue                 int     `nebulaproperty:"intrigue" description:"intrigue" nebulaindexes:"intrigue" json:"intrigue,omitempty"`
	HospitalLevel            int     `nebulaproperty:"hospital_level" description:"hospital_level" nebulaindexes:"hospital_level" json:"hospital_level,omitempty"`
	MaxPopulation            int     `nebulaproperty:"max_population" description:"max_population" nebulaindexes:"max_population" json:"max_population,omitempty"`
	CommanderLimit           int     `nebulaproperty:"commander_limit" description:"commander_limit" nebulaindexes:"commander_limit" json:"commander_limit,omitempty"`
	ClanSentiment            int     `nebulaproperty:"clan_sentiment" description:"clan_sentiment" nebulaindexes:"clan_sentiment" json:"clan_sentiment,omitempty"`
	GlobalTradeRouteWealth   int     `nebulaproperty:"global_trade_route_wealth" description:"global_trade_route_wealth" nebulaindexes:"global_trade_route_wealth" json:"global_trade_route_wealth,omitempty"`
	GlobalTradevalue         int     `nebulaproperty:"global_tradevalue" description:"global_tradevalue" nebulaindexes:"global_tradevalue" json:"global_tradevalue,omitempty"`
	AiFeudalModifier         int     `nebulaproperty:"ai_feudal_modifier" description:"ai_feudal_modifier" nebulaindexes:"ai_feudal_modifier" json:"ai_feudal_modifier,omitempty"`
	AiRepublicModifier       int     `nebulaproperty:"ai_republic_modifier" description:"ai_republic_modifier" nebulaindexes:"ai_republic_modifier" json:"ai_republic_modifier,omitempty"`
	PietyCost                int     `nebulaproperty:"piety_cost" description:"piety_cost" nebulaindexes:"piety_cost" json:"piety_cost,omitempty"`
	PrestigeCost             int     `nebulaproperty:"prestige_cost" description:"prestige_cost" nebulaindexes:"prestige_cost" json:"prestige_cost,omitempty"`
}

func NewBuilding(code string) *Building {
	nebulaBuilding := Building{
		VID:  getBuildingVid(code),
		Code: code,
	}
	return &nebulaBuilding
}

func NewBuildingByData(building *building.Building) *Building {
	nebulaBuilding := utils.Mapping[Building](building)
	nebulaBuilding.VID = getBuildingVid(nebulaBuilding.Code)
	return &nebulaBuilding
}

func getBuildingVid(code string) string {
	return "building." + code
}

func (b *Building) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, b)
}

func (b *Building) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, b)
}

func (b *Building) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, b)
}

func (b *Building) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, b)
}

func (b *Building) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, b)
}

func (b *Building) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, b)
}

func InsertBuildings(space *nebulagolang.Space, bs ...*Building) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, bs)
}

func UpdateBuildings(space *nebulagolang.Space, bs ...*Building) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, bs)
}

func UpsertBuildings(space *nebulagolang.Space, bs ...*Building) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, bs)
}

func DeleteBuildings(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getBuildingVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteBuildingWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getBuildingVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllBuilding(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Building](space)
}

func DeleteAllBuildingWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Building](space)
}

func GetBuildingByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Building] {
	return nebulagolang.GetVertexByVid[*Building](space, vid)
}

func GetBuildingByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*Building] {
	return nebulagolang.GetVertexByVid[*Building](space, getBuildingVid(code))
}

func GetAllBuildings(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Building] {
	return nebulagolang.GetAllVertexesByQuery[*Building](space, "")
}

func GetAllBuildingVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Building](space, "")
}

func GetAllBuildingCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Building](space, "", "code", "")
}
