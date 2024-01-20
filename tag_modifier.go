package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/modifier"
	"github.com/thalesfu/paradoxtools/utils"
)

type Modifier struct {
	VID                                   string  `nebulakey:"vid" nebulatagname:"modifier" nebulatagcomment:"Event modifier" json:"vid,omitempty"`
	Code                                  string  `nebulaproperty:"code" description:"modifier code" nebulaindexes:"code" json:"code,omitempty"`
	Name                                  string  `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Description                           string  `nebulaproperty:"description" description:"description" nebulaindexes:"description" json:"description,omitempty"`
	Major                                 bool    `nebulaproperty:"major" description:"major" nebulaindexes:"major" json:"major,omitempty"`
	LocalRevoltRisk                       float32 `nebulaproperty:"local_revolt_risk" description:"local_revolt_risk" nebulaindexes:"local_revolt_risk" json:"local_revolt_risk,omitempty"`
	LandMorale                            float32 `nebulaproperty:"land_morale" description:"land_morale" nebulaindexes:"land_morale" json:"land_morale,omitempty"`
	MonthlyCharacterPrestige              float32 `nebulaproperty:"monthly_character_prestige" description:"monthly_character_prestige" nebulaindexes:"monthly_character_prestige" json:"monthly_character_prestige,omitempty"`
	MonthlyCharacterPiety                 float32 `nebulaproperty:"monthly_character_piety" description:"monthly_character_piety" nebulaindexes:"monthly_character_piety" json:"monthly_character_piety,omitempty"`
	MonthlyCharacterWealth                float32 `nebulaproperty:"monthly_character_wealth" description:"monthly_character_wealth" nebulaindexes:"monthly_character_wealth" json:"monthly_character_wealth,omitempty"`
	Fertility                             float32 `nebulaproperty:"fertility" description:"fertility" nebulaindexes:"fertility" json:"fertility,omitempty"`
	BuildCostModifier                     float32 `nebulaproperty:"build_cost_modifier" description:"build_cost_modifier" nebulaindexes:"build_cost_modifier" json:"build_cost_modifier,omitempty"`
	BuildTimeModifier                     float32 `nebulaproperty:"build_time_modifier" description:"build_time_modifier" nebulaindexes:"build_time_modifier" json:"build_time_modifier,omitempty"`
	CastleLevySize                        float32 `nebulaproperty:"castle_levy_size" description:"castle_levy_size" nebulaindexes:"castle_levy_size" json:"castle_levy_size,omitempty"`
	GlobalTaxModifier                     float32 `nebulaproperty:"global_tax_modifier" description:"global_tax_modifier" nebulaindexes:"global_tax_modifier" json:"global_tax_modifier,omitempty"`
	LocalTaxModifier                      float32 `nebulaproperty:"local_tax_modifier" description:"local_tax_modifier" nebulaindexes:"local_tax_modifier" json:"local_tax_modifier,omitempty"`
	LocalBuildTimeModifier                float32 `nebulaproperty:"local_build_time_modifier" description:"local_build_time_modifier" nebulaindexes:"local_build_time_modifier" json:"local_build_time_modifier,omitempty"`
	TechGrowthModifierCulture             float32 `nebulaproperty:"tech_growth_modifier_culture" description:"tech_growth_modifier_culture" nebulaindexes:"tech_growth_modifier_culture" json:"tech_growth_modifier_culture,omitempty"`
	TechGrowthModifierMilitary            float32 `nebulaproperty:"tech_growth_modifier_military" description:"tech_growth_modifier_military" nebulaindexes:"tech_growth_modifier_military" json:"tech_growth_modifier_military,omitempty"`
	CastleTaxModifier                     float32 `nebulaproperty:"castle_tax_modifier" description:"castle_tax_modifier" nebulaindexes:"castle_tax_modifier" json:"castle_tax_modifier,omitempty"`
	LevyReinforceRate                     float32 `nebulaproperty:"levy_reinforce_rate" description:"levy_reinforce_rate" nebulaindexes:"levy_reinforce_rate" json:"levy_reinforce_rate,omitempty"`
	CityTaxModifier                       float32 `nebulaproperty:"city_tax_modifier" description:"city_tax_modifier" nebulaindexes:"city_tax_modifier" json:"city_tax_modifier,omitempty"`
	TempleTaxModifier                     float32 `nebulaproperty:"temple_tax_modifier" description:"temple_tax_modifier" nebulaindexes:"temple_tax_modifier" json:"temple_tax_modifier,omitempty"`
	TribalTaxModifier                     float32 `nebulaproperty:"tribal_tax_modifier" description:"tribal_tax_modifier" nebulaindexes:"tribal_tax_modifier" json:"tribal_tax_modifier,omitempty"`
	NomadTaxModifier                      float32 `nebulaproperty:"nomad_tax_modifier" description:"nomad_tax_modifier" nebulaindexes:"nomad_tax_modifier" json:"nomad_tax_modifier,omitempty"`
	DiseaseDefence                        float32 `nebulaproperty:"disease_defence" description:"disease_defence" nebulaindexes:"disease_defence" json:"disease_defence,omitempty"`
	LocalBuildCostModifier                float32 `nebulaproperty:"local_build_cost_modifier" description:"local_build_cost_modifier" nebulaindexes:"local_build_cost_modifier" json:"local_build_cost_modifier,omitempty"`
	LevySize                              float32 `nebulaproperty:"levy_size" description:"levy_size" nebulaindexes:"levy_size" json:"levy_size,omitempty"`
	TribalLevySize                        float32 `nebulaproperty:"tribal_levy_size" description:"tribal_levy_size" nebulaindexes:"tribal_levy_size" json:"tribal_levy_size,omitempty"`
	ManpowerGrowth                        float32 `nebulaproperty:"manpower_growth" description:"manpower_growth" nebulaindexes:"manpower_growth" json:"manpower_growth,omitempty"`
	DefensivePlotPowerModifier            float32 `nebulaproperty:"defensive_plot_power_modifier" description:"defensive_plot_power_modifier" nebulaindexes:"defensive_plot_power_modifier" json:"defensive_plot_power_modifier,omitempty"`
	GlobalRevoltRisk                      float32 `nebulaproperty:"global_revolt_risk" description:"global_revolt_risk" nebulaindexes:"global_revolt_risk" json:"global_revolt_risk,omitempty"`
	PopulationGrowth                      float32 `nebulaproperty:"population_growth" description:"population_growth" nebulaindexes:"population_growth" json:"population_growth,omitempty"`
	TaxIncome                             float32 `nebulaproperty:"tax_income" description:"tax_income" nebulaindexes:"tax_income" json:"tax_income,omitempty"`
	ArrestChanceModifier                  float32 `nebulaproperty:"arrest_chance_modifier" description:"arrest_chance_modifier" nebulaindexes:"arrest_chance_modifier" json:"arrest_chance_modifier,omitempty"`
	PlotPowerModifier                     float32 `nebulaproperty:"plot_power_modifier" description:"plot_power_modifier" nebulaindexes:"plot_power_modifier" json:"plot_power_modifier,omitempty"`
	PlotDiscoveryChance                   float32 `nebulaproperty:"plot_discovery_chance" description:"plot_discovery_chance" nebulaindexes:"plot_discovery_chance" json:"plot_discovery_chance,omitempty"`
	Health                                float32 `nebulaproperty:"health" description:"health" nebulaindexes:"health" json:"health,omitempty"`
	MaxManpowerMult                       float32 `nebulaproperty:"max_manpower_mult" description:"max_manpower_mult" nebulaindexes:"max_manpower_mult" json:"max_manpower_mult,omitempty"`
	MaxPopulationMult                     float32 `nebulaproperty:"max_population_mult" description:"max_population_mult" nebulaindexes:"max_population_mult" json:"max_population_mult,omitempty"`
	GarrisonSize                          float32 `nebulaproperty:"garrison_size" description:"garrison_size" nebulaindexes:"garrison_size" json:"garrison_size,omitempty"`
	HealthPenalty                         float32 `nebulaproperty:"health_penalty" description:"health_penalty" nebulaindexes:"health_penalty" json:"health_penalty,omitempty"`
	GalleysPerc                           float32 `nebulaproperty:"galleys_perc" description:"galleys_perc" nebulaindexes:"galleys_perc" json:"galleys_perc,omitempty"`
	MilitaryTechpoints                    float32 `nebulaproperty:"military_techpoints" description:"military_techpoints" nebulaindexes:"military_techpoints" json:"military_techpoints,omitempty"`
	CultureTechpoints                     float32 `nebulaproperty:"culture_techpoints" description:"culture_techpoints" nebulaindexes:"culture_techpoints" json:"culture_techpoints,omitempty"`
	EconomyTechpoints                     float32 `nebulaproperty:"economy_techpoints" description:"economy_techpoints" nebulaindexes:"economy_techpoints" json:"economy_techpoints,omitempty"`
	TechGrowthModifier                    float32 `nebulaproperty:"tech_growth_modifier" description:"tech_growth_modifier" nebulaindexes:"tech_growth_modifier" json:"tech_growth_modifier,omitempty"`
	AddPietyModifier                      float32 `nebulaproperty:"add_piety_modifier" description:"add_piety_modifier" nebulaindexes:"add_piety_modifier" json:"add_piety_modifier,omitempty"`
	AddPrestigeModifier                   float32 `nebulaproperty:"add_prestige_modifier" description:"add_prestige_modifier" nebulaindexes:"add_prestige_modifier" json:"add_prestige_modifier,omitempty"`
	GlobalLevySize                        float32 `nebulaproperty:"global_levy_size" description:"global_levy_size" nebulaindexes:"global_levy_size" json:"global_levy_size,omitempty"`
	CityVassalTaxModifier                 float32 `nebulaproperty:"city_vassal_tax_modifier" description:"city_vassal_tax_modifier" nebulaindexes:"city_vassal_tax_modifier" json:"city_vassal_tax_modifier,omitempty"`
	SiegeDefence                          float32 `nebulaproperty:"siege_defence" description:"siege_defence" nebulaindexes:"siege_defence" json:"siege_defence,omitempty"`
	MultiplicativeTradePostIncomeModifier float32 `nebulaproperty:"multiplicative_trade_post_income_modifier" description:"multiplicative_trade_post_income_modifier" nebulaindexes:"multiplicative_trade_post_income_modifier" json:"multiplicative_trade_post_income_modifier,omitempty"`
	LocalBuildTimeTempleModifier          float32 `nebulaproperty:"local_build_time_temple_modifier" description:"local_build_time_temple_modifier" nebulaindexes:"local_build_time_temple_modifier" json:"local_build_time_temple_modifier,omitempty"`
	GlobalMovementSpeed                   float32 `nebulaproperty:"global_movement_speed" description:"global_movement_speed" nebulaindexes:"global_movement_speed" json:"global_movement_speed,omitempty"`
	SiegeSpeed                            float32 `nebulaproperty:"siege_speed" description:"siege_speed" nebulaindexes:"siege_speed" json:"siege_speed,omitempty"`
	SupplyLimit                           float32 `nebulaproperty:"supply_limit" description:"supply_limit" nebulaindexes:"supply_limit" json:"supply_limit,omitempty"`
	Attrition                             float32 `nebulaproperty:"attrition" description:"attrition" nebulaindexes:"attrition" json:"attrition,omitempty"`
	ReligionFlex                          float32 `nebulaproperty:"religion_flex" description:"religion_flex" nebulaindexes:"religion_flex" json:"religion_flex,omitempty"`
	CultureFlex                           float32 `nebulaproperty:"culture_flex" description:"culture_flex" nebulaindexes:"culture_flex" json:"culture_flex,omitempty"`
	LandOrganisation                      float32 `nebulaproperty:"land_organisation" description:"land_organisation" nebulaindexes:"land_organisation" json:"land_organisation,omitempty"`
	MurderPlotPowerModifier               float32 `nebulaproperty:"murder_plot_power_modifier" description:"murder_plot_power_modifier" nebulaindexes:"murder_plot_power_modifier" json:"murder_plot_power_modifier,omitempty"`
	Icon                                  int     `nebulaproperty:"icon" description:"icon" nebulaindexes:"icon" json:"icon,omitempty"`
	Intrigue                              int     `nebulaproperty:"intrigue" description:"intrigue" nebulaindexes:"intrigue" json:"intrigue,omitempty"`
	Martial                               int     `nebulaproperty:"martial" description:"martial" nebulaindexes:"martial" json:"martial,omitempty"`
	ChurchOpinion                         int     `nebulaproperty:"church_opinion" description:"church_opinion" nebulaindexes:"church_opinion" json:"church_opinion,omitempty"`
	Stewardship                           int     `nebulaproperty:"stewardship" description:"stewardship" nebulaindexes:"stewardship" json:"stewardship,omitempty"`
	Learning                              int     `nebulaproperty:"learning" description:"learning" nebulaindexes:"learning" json:"learning,omitempty"`
	GeneralOpinion                        int     `nebulaproperty:"general_opinion" description:"general_opinion" nebulaindexes:"general_opinion" json:"general_opinion,omitempty"`
	CombatRating                          int     `nebulaproperty:"combat_rating" description:"combat_rating" nebulaindexes:"combat_rating" json:"combat_rating,omitempty"`
	Tradevalue                            int     `nebulaproperty:"tradevalue" description:"tradevalue" nebulaindexes:"tradevalue" json:"tradevalue,omitempty"`
	TradeRouteWealth                      int     `nebulaproperty:"trade_route_wealth" description:"trade_route_wealth" nebulaindexes:"trade_route_wealth" json:"trade_route_wealth,omitempty"`
	TradeRouteValue                       int     `nebulaproperty:"trade_route_value" description:"trade_route_value" nebulaindexes:"trade_route_value" json:"trade_route_value,omitempty"`
	TownOpinion                           int     `nebulaproperty:"town_opinion" description:"town_opinion" nebulaindexes:"town_opinion" json:"town_opinion,omitempty"`
	SexAppealOpinion                      int     `nebulaproperty:"sex_appeal_opinion" description:"sex_appeal_opinion" nebulaindexes:"sex_appeal_opinion" json:"sex_appeal_opinion,omitempty"`
	CastleOpinion                         int     `nebulaproperty:"castle_opinion" description:"castle_opinion" nebulaindexes:"castle_opinion" json:"castle_opinion,omitempty"`
	TechGrowthModifierEconomy             int     `nebulaproperty:"tech_growth_modifier_economy" description:"tech_growth_modifier_economy" nebulaindexes:"tech_growth_modifier_economy" json:"tech_growth_modifier_economy,omitempty"`
	SameReligionOpinion                   int     `nebulaproperty:"same_religion_opinion" description:"same_religion_opinion" nebulaindexes:"same_religion_opinion" json:"same_religion_opinion,omitempty"`
	VassalOpinion                         int     `nebulaproperty:"vassal_opinion" description:"vassal_opinion" nebulaindexes:"vassal_opinion" json:"vassal_opinion,omitempty"`
	TribalOpinion                         int     `nebulaproperty:"tribal_opinion" description:"tribal_opinion" nebulaindexes:"tribal_opinion" json:"tribal_opinion,omitempty"`
	ClanSentiment                         int     `nebulaproperty:"clan_sentiment" description:"clan_sentiment" nebulaindexes:"clan_sentiment" json:"clan_sentiment,omitempty"`
	DiplomacyPenalty                      int     `nebulaproperty:"diplomacy_penalty" description:"diplomacy_penalty" nebulaindexes:"diplomacy_penalty" json:"diplomacy_penalty,omitempty"`
	MartialPenalty                        int     `nebulaproperty:"martial_penalty" description:"martial_penalty" nebulaindexes:"martial_penalty" json:"martial_penalty,omitempty"`
	StewardshipPenalty                    int     `nebulaproperty:"stewardship_penalty" description:"stewardship_penalty" nebulaindexes:"stewardship_penalty" json:"stewardship_penalty,omitempty"`
	IntriguePenalty                       int     `nebulaproperty:"intrigue_penalty" description:"intrigue_penalty" nebulaindexes:"intrigue_penalty" json:"intrigue_penalty,omitempty"`
	LearningPenalty                       int     `nebulaproperty:"learning_penalty" description:"learning_penalty" nebulaindexes:"learning_penalty" json:"learning_penalty,omitempty"`
	LiegeOpinion                          int     `nebulaproperty:"liege_opinion" description:"liege_opinion" nebulaindexes:"liege_opinion" json:"liege_opinion,omitempty"`
	DynastyOpinion                        int     `nebulaproperty:"dynasty_opinion" description:"dynasty_opinion" nebulaindexes:"dynasty_opinion" json:"dynasty_opinion,omitempty"`
	SocietyInfluence                      int     `nebulaproperty:"society_influence" description:"society_influence" nebulaindexes:"society_influence" json:"society_influence,omitempty"`
	MonthlyGrace                          int     `nebulaproperty:"monthly_grace" description:"monthly_grace" nebulaindexes:"monthly_grace" json:"monthly_grace,omitempty"`
	DuelistOpinion                        int     `nebulaproperty:"duelist_opinion" description:"duelist_opinion" nebulaindexes:"duelist_opinion" json:"duelist_opinion,omitempty"`
	CruelOpinion                          int     `nebulaproperty:"cruel_opinion" description:"cruel_opinion" nebulaindexes:"cruel_opinion" json:"cruel_opinion,omitempty"`
	StrategistOpinion                     int     `nebulaproperty:"strategist_opinion" description:"strategist_opinion" nebulaindexes:"strategist_opinion" json:"strategist_opinion,omitempty"`
	KindOpinion                           int     `nebulaproperty:"kind_opinion" description:"kind_opinion" nebulaindexes:"kind_opinion" json:"kind_opinion,omitempty"`
	Diplomacy                             int     `nebulaproperty:"Diplomacy" description:"Diplomacy" nebulaindexes:"Diplomacy" json:"Diplomacy,omitempty"`
	WrothOpinion                          int     `nebulaproperty:"wroth_opinion" description:"wroth_opinion" nebulaindexes:"wroth_opinion" json:"wroth_opinion,omitempty"`
	ZealousOpinion                        int     `nebulaproperty:"zealous_opinion" description:"zealous_opinion" nebulaindexes:"zealous_opinion" json:"zealous_opinion,omitempty"`
	DeceitfulOpinion                      int     `nebulaproperty:"deceitful_opinion" description:"deceitful_opinion" nebulaindexes:"deceitful_opinion" json:"deceitful_opinion,omitempty"`
	AmbitiousOpinion                      int     `nebulaproperty:"ambitious_opinion" description:"ambitious_opinion" nebulaindexes:"ambitious_opinion" json:"ambitious_opinion,omitempty"`
	DaysOfSupply                          int     `nebulaproperty:"days_of_supply" description:"days_of_supply" nebulaindexes:"days_of_supply" json:"days_of_supply,omitempty"`
	BraveOpinion                          int     `nebulaproperty:"brave_opinion" description:"brave_opinion" nebulaindexes:"brave_opinion" json:"brave_opinion,omitempty"`
	PatientOpinion                        int     `nebulaproperty:"patient_opinion" description:"patient_opinion" nebulaindexes:"patient_opinion" json:"patient_opinion,omitempty"`
	HumbleOpinion                         int     `nebulaproperty:"humble_opinion" description:"humble_opinion" nebulaindexes:"humble_opinion" json:"humble_opinion,omitempty"`
	ProudOpinion                          int     `nebulaproperty:"proud_opinion" description:"proud_opinion" nebulaindexes:"proud_opinion" json:"proud_opinion,omitempty"`
	CharitableOpinion                     int     `nebulaproperty:"charitable_opinion" description:"charitable_opinion" nebulaindexes:"charitable_opinion" json:"charitable_opinion,omitempty"`
}

func NewModifier(code string) *Modifier {
	nebulaModifier := Modifier{
		VID:  getModifierVid(code),
		Code: code,
	}
	return &nebulaModifier
}

func NewModifierByData(modifier *modifier.Modifier) *Modifier {
	nebulaModifier := utils.Mapping[Modifier](modifier)
	nebulaModifier.VID = getModifierVid(nebulaModifier.Code)
	return &nebulaModifier
}

func getModifierVid(code string) string {
	return "modifier." + code
}

func (m *Modifier) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, m)
}

func (m *Modifier) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, m)
}

func (m *Modifier) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, m)
}

func (m *Modifier) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, m)
}

func (m *Modifier) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, m)
}

func (m *Modifier) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, m)
}

func InsertModifiers(space *nebulagolang.Space, ms ...*Modifier) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ms)
}

func UpdateModifiers(space *nebulagolang.Space, ms ...*Modifier) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ms)
}

func UpsertModifiers(space *nebulagolang.Space, ms ...*Modifier) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ms)
}

func DeleteModifiers(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getModifierVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteModifierWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getModifierVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllModifier(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Modifier](space)
}

func DeleteAllModifierWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Modifier](space)
}

func GetModifierByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Modifier] {
	return nebulagolang.GetVertexByVid[*Modifier](space, vid)
}

func GetModifierByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*Modifier] {
	return nebulagolang.GetVertexByVid[*Modifier](space, getModifierVid(code))
}

func GetAllModifier(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Modifier] {
	return nebulagolang.GetAllVertexesByQuery[*Modifier](space, "")
}

func GetAllModifierVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Modifier](space, "")
}

func GetAllModifierCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Modifier](space, "", "code", "")
}

func GetAllModifierName(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Modifier](space, "", "name", "")
}
