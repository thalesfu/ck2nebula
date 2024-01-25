package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/trait"
	"github.com/thalesfu/paradoxtools/utils"
)

type Trait struct {
	VID                             string  `nebulakey:"vid" nebulatagname:"trait" nebulatagcomment:"trait" json:"vid,omitempty"`
	Code                            string  `nebulaproperty:"code" description:"code" nebulaindexes:"code" json:"code,omitempty"`
	ID                              int     `nebulaproperty:"id" description:"id" nebulaindexes:"id" json:"id,omitempty"`
	Name                            string  `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Description                     string  `nebulaproperty:"description" description:"description" nebulaindexes:"description" json:"description,omitempty"`
	Education                       bool    `nebulaproperty:"education" description:"education" nebulaindexes:"education" json:"education,omitempty"`
	IsHealth                        bool    `nebulaproperty:"is_health" description:"is_health" nebulaindexes:"is_health" json:"is_health,omitempty"`
	IsIllness                       bool    `nebulaproperty:"is_illness" description:"is_illness" nebulaindexes:"is_illness" json:"is_illness,omitempty"`
	Customizer                      bool    `nebulaproperty:"customizer" description:"customizer" nebulaindexes:"customizer" json:"customizer,omitempty"`
	SuccessionGfx                   bool    `nebulaproperty:"succession_gfx" description:"succession_gfx" nebulaindexes:"succession_gfx" json:"succession_gfx,omitempty"`
	Incapacitating                  bool    `nebulaproperty:"incapacitating" description:"incapacitating" nebulaindexes:"incapacitating" json:"incapacitating,omitempty"`
	IsEpidemic                      bool    `nebulaproperty:"is_epidemic" description:"is_epidemic" nebulaindexes:"is_epidemic" json:"is_epidemic,omitempty"`
	Religious                       bool    `nebulaproperty:"religious" description:"religious" nebulaindexes:"religious" json:"religious,omitempty"`
	Random                          bool    `nebulaproperty:"random" description:"random" nebulaindexes:"random" json:"random,omitempty"`
	Inbred                          bool    `nebulaproperty:"inbred" description:"inbred" nebulaindexes:"inbred" json:"inbred,omitempty"`
	Lifestyle                       bool    `nebulaproperty:"lifestyle" description:"lifestyle" nebulaindexes:"lifestyle" json:"lifestyle,omitempty"`
	Personality                     bool    `nebulaproperty:"personality" description:"personality" nebulaindexes:"personality" json:"personality,omitempty"`
	Vice                            bool    `nebulaproperty:"vice" description:"vice" nebulaindexes:"vice" json:"vice,omitempty"`
	Virtue                          bool    `nebulaproperty:"virtue" description:"virtue" nebulaindexes:"virtue" json:"virtue,omitempty"`
	Leader                          bool    `nebulaproperty:"leader" description:"leader" nebulaindexes:"leader" json:"leader,omitempty"`
	Cached                          bool    `nebulaproperty:"cached" description:"cached" nebulaindexes:"cached" json:"cached,omitempty"`
	Pilgrimage                      bool    `nebulaproperty:"pilgrimage" description:"pilgrimage" nebulaindexes:"pilgrimage" json:"pilgrimage,omitempty"`
	Agnatic                         bool    `nebulaproperty:"agnatic" description:"agnatic" nebulaindexes:"agnatic" json:"agnatic,omitempty"`
	CannotMarry                     bool    `nebulaproperty:"cannot_marry" description:"cannot_marry" nebulaindexes:"cannot_marry" json:"cannot_marry,omitempty"`
	CannotInherit                   bool    `nebulaproperty:"cannot_inherit" description:"cannot_inherit" nebulaindexes:"cannot_inherit" json:"cannot_inherit,omitempty"`
	Blinding                        bool    `nebulaproperty:"blinding" description:"blinding" nebulaindexes:"blinding" json:"blinding,omitempty"`
	RebelInherited                  bool    `nebulaproperty:"rebel_inherited" description:"rebel_inherited" nebulaindexes:"rebel_inherited" json:"rebel_inherited,omitempty"`
	ToleratesChristian              bool    `nebulaproperty:"tolerates_christian" description:"tolerates_christian" nebulaindexes:"tolerates_christian" json:"tolerates_christian,omitempty"`
	ToleratesMuslim                 bool    `nebulaproperty:"tolerates_muslim" description:"tolerates_muslim" nebulaindexes:"tolerates_muslim" json:"tolerates_muslim,omitempty"`
	ToleratesPaganGroup             bool    `nebulaproperty:"tolerates_pagan_group" description:"tolerates_pagan_group" nebulaindexes:"tolerates_pagan_group" json:"tolerates_pagan_group,omitempty"`
	ToleratesZoroastrianGroup       bool    `nebulaproperty:"tolerates_zoroastrian_group" description:"tolerates_zoroastrian_group" nebulaindexes:"tolerates_zoroastrian_group" json:"tolerates_zoroastrian_group,omitempty"`
	ToleratesJewishGroup            bool    `nebulaproperty:"tolerates_jewish_group" description:"tolerates_jewish_group" nebulaindexes:"tolerates_jewish_group" json:"tolerates_jewish_group,omitempty"`
	ToleratesIndianGroup            bool    `nebulaproperty:"tolerates_indian_group" description:"tolerates_indian_group" nebulaindexes:"tolerates_indian_group" json:"tolerates_indian_group,omitempty"`
	InHiding                        bool    `nebulaproperty:"in_hiding" description:"in_hiding" nebulaindexes:"in_hiding" json:"in_hiding,omitempty"`
	Childhood                       bool    `nebulaproperty:"childhood" description:"childhood" nebulaindexes:"childhood" json:"childhood,omitempty"`
	CanHoldTitles                   bool    `nebulaproperty:"can_hold_titles" description:"can_hold_titles" nebulaindexes:"can_hold_titles" json:"can_hold_titles,omitempty"`
	IsSymptom                       bool    `nebulaproperty:"is_symptom" description:"is_symptom" nebulaindexes:"is_symptom" json:"is_symptom,omitempty"`
	Immortal                        bool    `nebulaproperty:"immortal" description:"immortal" nebulaindexes:"immortal" json:"immortal,omitempty"`
	Hidden                          bool    `nebulaproperty:"hidden" description:"hidden" nebulaindexes:"hidden" json:"hidden,omitempty"`
	HiddenFromOthers                bool    `nebulaproperty:"hidden_from_others" description:"hidden_from_others" nebulaindexes:"hidden_from_others" json:"hidden_from_others,omitempty"`
	Attribute                       string  `nebulaproperty:"attribute" description:"attribute" nebulaindexes:"attribute" json:"attribute,omitempty"`
	IsTribal                        string  `nebulaproperty:"is_tribal" description:"is_tribal" nebulaindexes:"is_tribal" json:"is_tribal,omitempty"`
	ReligionGroup                   string  `nebulaproperty:"religion_group" description:"religion_group" nebulaindexes:"religion_group" json:"religion_group,omitempty"`
	Terrain                         string  `nebulaproperty:"terrain" description:"terrain" nebulaindexes:"terrain" json:"terrain,omitempty"`
	Religion                        string  `nebulaproperty:"religion" description:"religion" nebulaindexes:"religion" json:"religion,omitempty"`
	HasBloodlineFlag                string  `nebulaproperty:"has_bloodline_flag" description:"has_bloodline_flag" nebulaindexes:"has_bloodline_flag" json:"has_bloodline_flag,omitempty"`
	IsRuler                         string  `nebulaproperty:"is_ruler" description:"is_ruler" nebulaindexes:"is_ruler" json:"is_ruler,omitempty"`
	IsFemale                        string  `nebulaproperty:"is_female" description:"is_female" nebulaindexes:"is_female" json:"is_female,omitempty"`
	IsTheocracy                     string  `nebulaproperty:"is_theocracy" description:"is_theocracy" nebulaindexes:"is_theocracy" json:"is_theocracy,omitempty"`
	ControlsReligion                string  `nebulaproperty:"controls_religion" description:"controls_religion" nebulaindexes:"controls_religion" json:"controls_religion,omitempty"`
	ReligiousBranch                 string  `nebulaproperty:"religious_branch" description:"religious_branch" nebulaindexes:"religious_branch" json:"religious_branch,omitempty"`
	Prisoner                        string  `nebulaproperty:"prisoner" description:"prisoner" nebulaindexes:"prisoner" json:"prisoner,omitempty"`
	Race                            string  `nebulaproperty:"race" description:"race" nebulaindexes:"race" json:"race,omitempty"`
	HasReligionFeature              string  `nebulaproperty:"has_religion_feature" description:"has_religion_feature" nebulaindexes:"has_religion_feature" json:"has_religion_feature,omitempty"`
	Character                       string  `nebulaproperty:"character" description:"character" nebulaindexes:"character" json:"character,omitempty"`
	SocietyMemberOf                 string  `nebulaproperty:"society_member_of" description:"society_member_of" nebulaindexes:"society_member_of" json:"society_member_of,omitempty"`
	IsCloseRelative                 string  `nebulaproperty:"is_close_relative" description:"is_close_relative" nebulaindexes:"is_close_relative" json:"is_close_relative,omitempty"`
	Trait                           string  `nebulaproperty:"trait" description:"trait" nebulaindexes:"trait" json:"trait,omitempty"`
	Ai                              string  `nebulaproperty:"ai" description:"ai" nebulaindexes:"ai" json:"ai,omitempty"`
	IsNomadic                       string  `nebulaproperty:"is_nomadic" description:"is_nomadic" nebulaindexes:"is_nomadic" json:"is_nomadic,omitempty"`
	HasDharmicReligionTrigger       string  `nebulaproperty:"has_dharmic_religion_trigger" description:"has_dharmic_religion_trigger" nebulaindexes:"has_dharmic_religion_trigger" json:"has_dharmic_religion_trigger,omitempty"`
	Culture                         string  `nebulaproperty:"culture" description:"culture" nebulaindexes:"culture" json:"culture,omitempty"`
	HasCharacterFlag                string  `nebulaproperty:"has_character_flag" description:"has_character_flag" nebulaindexes:"has_character_flag" json:"has_character_flag,omitempty"`
	GraphicalCulture                string  `nebulaproperty:"graphical_culture" description:"graphical_culture" nebulaindexes:"graphical_culture" json:"graphical_culture,omitempty"`
	MonthlyCharacterPiety           float32 `nebulaproperty:"monthly_character_piety" description:"monthly_character_piety" nebulaindexes:"monthly_character_piety" json:"monthly_character_piety,omitempty"`
	GlobalTaxModifier               float32 `nebulaproperty:"global_tax_modifier" description:"global_tax_modifier" nebulaindexes:"global_tax_modifier" json:"global_tax_modifier,omitempty"`
	MonthlyCharacterPrestige        float32 `nebulaproperty:"monthly_character_prestige" description:"monthly_character_prestige" nebulaindexes:"monthly_character_prestige" json:"monthly_character_prestige,omitempty"`
	MonthlyCharacterWealth          float32 `nebulaproperty:"monthly_character_wealth" description:"monthly_character_wealth" nebulaindexes:"monthly_character_wealth" json:"monthly_character_wealth,omitempty"`
	WonderBuildCostModifier         float32 `nebulaproperty:"wonder_build_cost_modifier" description:"wonder_build_cost_modifier" nebulaindexes:"wonder_build_cost_modifier" json:"wonder_build_cost_modifier,omitempty"`
	WonderBuildTimeModifier         float32 `nebulaproperty:"wonder_build_time_modifier" description:"wonder_build_time_modifier" nebulaindexes:"wonder_build_time_modifier" json:"wonder_build_time_modifier,omitempty"`
	GlobalLevySize                  float32 `nebulaproperty:"global_levy_size" description:"global_levy_size" nebulaindexes:"global_levy_size" json:"global_levy_size,omitempty"`
	MaxManpowerMult                 float32 `nebulaproperty:"max_manpower_mult" description:"max_manpower_mult" nebulaindexes:"max_manpower_mult" json:"max_manpower_mult,omitempty"`
	GlobalRevoltRisk                float32 `nebulaproperty:"global_revolt_risk" description:"global_revolt_risk" nebulaindexes:"global_revolt_risk" json:"global_revolt_risk,omitempty"`
	Attrition                       float32 `nebulaproperty:"attrition" description:"attrition" nebulaindexes:"attrition" json:"attrition,omitempty"`
	Diplomacy                       int     `nebulaproperty:"diplomacy" description:"diplomacy" nebulaindexes:"diplomacy" json:"diplomacy,omitempty"`
	Martial                         int     `nebulaproperty:"martial" description:"martial" nebulaindexes:"martial" json:"martial,omitempty"`
	Stewardship                     int     `nebulaproperty:"stewardship" description:"stewardship" nebulaindexes:"stewardship" json:"stewardship,omitempty"`
	Intrigue                        int     `nebulaproperty:"intrigue" description:"intrigue" nebulaindexes:"intrigue" json:"intrigue,omitempty"`
	Learning                        int     `nebulaproperty:"learning" description:"learning" nebulaindexes:"learning" json:"learning,omitempty"`
	Fertility                       float32 `nebulaproperty:"fertility" description:"fertility" nebulaindexes:"fertility" json:"fertility,omitempty"`
	Health                          float32 `nebulaproperty:"health" description:"health" nebulaindexes:"health" json:"health,omitempty"`
	SexAppealOpinion                int     `nebulaproperty:"sex_appeal_opinion" description:"sex_appeal_opinion" nebulaindexes:"sex_appeal_opinion" json:"sex_appeal_opinion,omitempty"`
	CombatRating                    int     `nebulaproperty:"combat_rating" description:"combat_rating" nebulaindexes:"combat_rating" json:"combat_rating,omitempty"`
	DiplomacyPenalty                int     `nebulaproperty:"diplomacy_penalty" description:"diplomacy_penalty" nebulaindexes:"diplomacy_penalty" json:"diplomacy_penalty,omitempty"`
	MartialPenalty                  int     `nebulaproperty:"martial_penalty" description:"martial_penalty" nebulaindexes:"martial_penalty" json:"martial_penalty,omitempty"`
	StewardshipPenalty              int     `nebulaproperty:"stewardship_penalty" description:"stewardship_penalty" nebulaindexes:"stewardship_penalty" json:"stewardship_penalty,omitempty"`
	IntriguePenalty                 int     `nebulaproperty:"intrigue_penalty" description:"intrigue_penalty" nebulaindexes:"intrigue_penalty" json:"intrigue_penalty,omitempty"`
	LearningPenalty                 int     `nebulaproperty:"learning_penalty" description:"learning_penalty" nebulaindexes:"learning_penalty" json:"learning_penalty,omitempty"`
	FertilityPenalty                float32 `nebulaproperty:"fertility_penalty" description:"fertility_penalty" nebulaindexes:"fertility_penalty" json:"fertility_penalty,omitempty"`
	HealthPenalty                   float32 `nebulaproperty:"health_penalty" description:"health_penalty" nebulaindexes:"health_penalty" json:"health_penalty,omitempty"`
	LeadershipTraits                int     `nebulaproperty:"leadership_traits" description:"leadership_traits" nebulaindexes:"leadership_traits" json:"leadership_traits,omitempty"`
	AiZeal                          int     `nebulaproperty:"ai_zeal" description:"ai_zeal" nebulaindexes:"ai_zeal" json:"ai_zeal,omitempty"`
	VassalOpinion                   int     `nebulaproperty:"vassal_opinion" description:"vassal_opinion" nebulaindexes:"vassal_opinion" json:"vassal_opinion,omitempty"`
	SameOpinion                     int     `nebulaproperty:"same_opinion" description:"same_opinion" nebulaindexes:"same_opinion" json:"same_opinion,omitempty"`
	AiRationality                   int     `nebulaproperty:"ai_rationality" description:"ai_rationality" nebulaindexes:"ai_rationality" json:"ai_rationality,omitempty"`
	InheritChance                   int     `nebulaproperty:"inherit_chance" description:"inherit_chance" nebulaindexes:"inherit_chance" json:"inherit_chance,omitempty"`
	GeneralOpinion                  int     `nebulaproperty:"general_opinion" description:"general_opinion" nebulaindexes:"general_opinion" json:"general_opinion,omitempty"`
	ChurchOpinion                   int     `nebulaproperty:"church_opinion" description:"church_opinion" nebulaindexes:"church_opinion" json:"church_opinion,omitempty"`
	SameOpinionIfSameReligion       int     `nebulaproperty:"same_opinion_if_same_religion" description:"same_opinion_if_same_religion" nebulaindexes:"same_opinion_if_same_religion" json:"same_opinion_if_same_religion,omitempty"`
	TwinOpinion                     int     `nebulaproperty:"twin_opinion" description:"twin_opinion" nebulaindexes:"twin_opinion" json:"twin_opinion,omitempty"`
	SpouseOpinion                   int     `nebulaproperty:"spouse_opinion" description:"spouse_opinion" nebulaindexes:"spouse_opinion" json:"spouse_opinion,omitempty"`
	SameReligionOpinion             int     `nebulaproperty:"same_religion_opinion" description:"same_religion_opinion" nebulaindexes:"same_religion_opinion" json:"same_religion_opinion,omitempty"`
	DynastyOpinion                  int     `nebulaproperty:"dynasty_opinion" description:"dynasty_opinion" nebulaindexes:"dynasty_opinion" json:"dynasty_opinion,omitempty"`
	RulerDesignerCost               int     `nebulaproperty:"ruler_designer_cost" description:"ruler_designer_cost" nebulaindexes:"ruler_designer_cost" json:"ruler_designer_cost,omitempty"`
	Birth                           int     `nebulaproperty:"birth" description:"birth" nebulaindexes:"birth" json:"birth,omitempty"`
	BothParentHasTraitInheritChance int     `nebulaproperty:"both_parent_has_trait_inherit_chance" description:"both_parent_has_trait_inherit_chance" nebulaindexes:"both_parent_has_trait_inherit_chance" json:"both_parent_has_trait_inherit_chance,omitempty"`
	TribalOpinion                   int     `nebulaproperty:"tribal_opinion" description:"tribal_opinion" nebulaindexes:"tribal_opinion" json:"tribal_opinion,omitempty"`
	ChristianChurchOpinion          int     `nebulaproperty:"christian_church_opinion" description:"christian_church_opinion" nebulaindexes:"christian_church_opinion" json:"christian_church_opinion,omitempty"`
	OppositeOpinion                 int     `nebulaproperty:"opposite_opinion" description:"opposite_opinion" nebulaindexes:"opposite_opinion" json:"opposite_opinion,omitempty"`
	AiHonor                         int     `nebulaproperty:"ai_honor" description:"ai_honor" nebulaindexes:"ai_honor" json:"ai_honor,omitempty"`
	AiGreed                         int     `nebulaproperty:"ai_greed" description:"ai_greed" nebulaindexes:"ai_greed" json:"ai_greed,omitempty"`
	AiAmbition                      int     `nebulaproperty:"ai_ambition" description:"ai_ambition" nebulaindexes:"ai_ambition" json:"ai_ambition,omitempty"`
	LiegeOpinion                    int     `nebulaproperty:"liege_opinion" description:"liege_opinion" nebulaindexes:"liege_opinion" json:"liege_opinion,omitempty"`
	AmbitionOpinion                 int     `nebulaproperty:"ambition_opinion" description:"ambition_opinion" nebulaindexes:"ambition_opinion" json:"ambition_opinion,omitempty"`
	InfidelOpinion                  int     `nebulaproperty:"infidel_opinion" description:"infidel_opinion" nebulaindexes:"infidel_opinion" json:"infidel_opinion,omitempty"`
	MuslimOpinion                   int     `nebulaproperty:"muslim_opinion" description:"muslim_opinion" nebulaindexes:"muslim_opinion" json:"muslim_opinion,omitempty"`
	ZoroastrianOpinion              int     `nebulaproperty:"zoroastrian_opinion" description:"zoroastrian_opinion" nebulaindexes:"zoroastrian_opinion" json:"zoroastrian_opinion,omitempty"`
	NorsePaganOpinion               int     `nebulaproperty:"norse_pagan_opinion" description:"norse_pagan_opinion" nebulaindexes:"norse_pagan_opinion" json:"norse_pagan_opinion,omitempty"`
	NorsePaganReformedOpinion       int     `nebulaproperty:"norse_pagan_reformed_opinion" description:"norse_pagan_reformed_opinion" nebulaindexes:"norse_pagan_reformed_opinion" json:"norse_pagan_reformed_opinion,omitempty"`
	CasteTier                       int     `nebulaproperty:"caste_tier" description:"caste_tier" nebulaindexes:"caste_tier" json:"caste_tier,omitempty"`
	PaganGroupOpinion               int     `nebulaproperty:"pagan_group_opinion" description:"pagan_group_opinion" nebulaindexes:"pagan_group_opinion" json:"pagan_group_opinion,omitempty"`
	TaoistOpinion                   int     `nebulaproperty:"taoist_opinion" description:"taoist_opinion" nebulaindexes:"taoist_opinion" json:"taoist_opinion,omitempty"`
	DaysOfSupply                    int     `nebulaproperty:"days_of_supply" description:"days_of_supply" nebulaindexes:"days_of_supply" json:"days_of_supply,omitempty"`
	MonthlyGrace                    int     `nebulaproperty:"monthly_grace" description:"monthly_grace" nebulaindexes:"monthly_grace" json:"monthly_grace,omitempty"`
	ChristianOpinion                int     `nebulaproperty:"christian_opinion" description:"christian_opinion" nebulaindexes:"christian_opinion" json:"christian_opinion,omitempty"`
	TraitEffectCaptureCommanders    int     `nebulaproperty:"trait_effect_capture_commanders" description:"trait_effect_capture_commanders" nebulaindexes:"trait_effect_capture_commanders" json:"trait_effect_capture_commanders,omitempty"`
	IndianGroupOpinion              int     `nebulaproperty:"indian_group_opinion" description:"indian_group_opinion" nebulaindexes:"indian_group_opinion" json:"indian_group_opinion,omitempty"`
	JewishGroupOpinion              int     `nebulaproperty:"jewish_group_opinion" description:"jewish_group_opinion" nebulaindexes:"jewish_group_opinion" json:"jewish_group_opinion,omitempty"`
	ZoroastrianGroupOpinion         int     `nebulaproperty:"zoroastrian_group_opinion" description:"zoroastrian_group_opinion" nebulaindexes:"zoroastrian_group_opinion" json:"zoroastrian_group_opinion,omitempty"`
	CastleOpinion                   int     `nebulaproperty:"castle_opinion" description:"castle_opinion" nebulaindexes:"castle_opinion" json:"castle_opinion,omitempty"`
	TownOpinion                     int     `nebulaproperty:"town_opinion" description:"town_opinion" nebulaindexes:"town_opinion" json:"town_opinion,omitempty"`
}

func NewTrait(id int) *Trait {
	nebulaTrait := Trait{
		VID: getTraitVid(id),
		ID:  id,
	}
	return &nebulaTrait
}

func NewTraitByData(trait *trait.Trait) *Trait {
	nebulaTrait := utils.Mapping[Trait](trait)
	nebulaTrait.VID = getTraitVid(nebulaTrait.ID)
	return &nebulaTrait
}

func getTraitVid(id int) string {
	return fmt.Sprintf("trait.%d", id)
}

func (t *Trait) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, t)
}

func (t *Trait) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, t)
}

func (t *Trait) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, t)
}

func (t *Trait) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, t)
}

func (t *Trait) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, t)
}

func (t *Trait) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, t)
}

func InsertTraits(space *nebulagolang.Space, ms ...*Trait) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ms)
}

func UpdateTraits(space *nebulagolang.Space, ms ...*Trait) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ms)
}

func UpsertTraits(space *nebulagolang.Space, ms ...*Trait) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ms)
}

func DeleteTraits(space *nebulagolang.Space, ids ...int) *nebulagolang.Result {
	is := make([]string, len(ids))
	for i, id := range ids {
		is[i] = getTraitVid(id)
	}
	return nebulagolang.DeleteVertexesByVids(space, is...)
}

func DeleteTraitWithEdge(space *nebulagolang.Space, ids ...int) *nebulagolang.Result {
	is := make([]string, len(ids))
	for i, id := range ids {
		is[i] = getTraitVid(id)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, is...)
}

func DeleteAllTraits(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Trait](space)
}

func DeleteAllTraitsWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Trait](space)
}

func GetTraitByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Trait] {
	return nebulagolang.GetVertexByVid[*Trait](space, vid)
}

func GetTraitById(space *nebulagolang.Space, id int) *nebulagolang.ResultT[*Trait] {
	return nebulagolang.GetVertexByVid[*Trait](space, getTraitVid(id))
}

func GetAllTraits(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Trait] {
	return nebulagolang.GetAllVertexesByQuery[*Trait](space, "")
}

func GetAllTraitsMap(space *nebulagolang.Space) *nebulagolang.ResultT[map[int]*Trait] {
	result := GetAllTraits(space)

	if !result.Ok {
		return nebulagolang.NewResultT[map[int]*Trait](result.Result)
	}

	traits := make(map[int]*Trait)

	for _, t := range result.Data {
		traits[t.ID] = t
	}

	return nebulagolang.NewResultTWithData[map[int]*Trait](result.Result, traits)
}

func GetTraitsByName(space *nebulagolang.Space, name string) *nebulagolang.ResultT[map[string]*Trait] {
	return nebulagolang.GetAllVertexesByQuery[*Trait](space, nebulagolang.GetPropertyQueryByPropertyNameAndValue[Trait]("name", name))
}

func GetTraitsByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[map[string]*Trait] {
	return nebulagolang.GetAllVertexesByQuery[*Trait](space, nebulagolang.GetPropertyQueryByPropertyNameAndValue[Trait]("code", code))
}

func GetTraitsVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Trait](space, "")
}

func GetAllTraitCodes(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Trait](space, "", "code", "")
}

func GetAllTraitNames(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Trait](space, "", "name", "")
}
