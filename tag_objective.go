package ck2nebula

import (
	"github.com/thalesfu/nebulagolang"
	"github.com/thalesfu/paradoxtools/CK2/objectives"
	"github.com/thalesfu/paradoxtools/utils"
)

type Objective struct {
	VID                   string  `nebulakey:"vid" nebulatagname:"objective" nebulatagcomment:"Event objective" json:"vid,omitempty"`
	Code                  string  `nebulaproperty:"code" description:"objective code" nebulaindexes:"code" json:"code,omitempty"`
	Name                  string  `nebulaproperty:"name" description:"name" nebulaindexes:"name" json:"name,omitempty"`
	Description           string  `nebulaproperty:"description" description:"description" nebulaindexes:"description" json:"description,omitempty"`
	CanCancel             bool    `nebulaproperty:"can_cancel" description:"can_cancel" nebulaindexes:"can_cancel" json:"can_cancel,omitempty"`
	ExpectationOfLiege    bool    `nebulaproperty:"expectation_of_liege" description:"expectation_of_liege" nebulaindexes:"expectation_of_liege" json:"expectation_of_liege,omitempty"`
	AiCapitalKingdomFocus bool    `nebulaproperty:"ai_capital_kingdom_focus" description:"ai_capital_kingdom_focus" nebulaindexes:"ai_capital_kingdom_focus" json:"ai_capital_kingdom_focus,omitempty"`
	Exclusive             bool    `nebulaproperty:"exclusive" description:"exclusive" nebulaindexes:"exclusive" json:"exclusive,omitempty"`
	RelHeadLoyalist       bool    `nebulaproperty:"rel_head_loyalist" description:"rel_head_loyalist" nebulaindexes:"rel_head_loyalist" json:"rel_head_loyalist,omitempty"`
	CancelOnLeaderDeath   bool    `nebulaproperty:"cancel_on_leader_death" description:"cancel_on_leader_death" nebulaindexes:"cancel_on_leader_death" json:"cancel_on_leader_death,omitempty"`
	MilitaryPlot          bool    `nebulaproperty:"military_plot" description:"military_plot" nebulaindexes:"military_plot" json:"military_plot,omitempty"`
	VassalRankPlot        bool    `nebulaproperty:"vassal_rank_plot" description:"vassal_rank_plot" nebulaindexes:"vassal_rank_plot" json:"vassal_rank_plot,omitempty"`
	IntriguePlot          bool    `nebulaproperty:"intrigue_plot" description:"intrigue_plot" nebulaindexes:"intrigue_plot" json:"intrigue_plot,omitempty"`
	MurderPlot            bool    `nebulaproperty:"murder_plot" description:"murder_plot" nebulaindexes:"murder_plot" json:"murder_plot,omitempty"`
	VassalIntriguePlot    bool    `nebulaproperty:"vassal_intrigue_plot" description:"vassal_intrigue_plot" nebulaindexes:"vassal_intrigue_plot" json:"vassal_intrigue_plot,omitempty"`
	Type                  string  `nebulaproperty:"type" description:"type" nebulaindexes:"type" json:"type,omitempty"`
	Text                  string  `nebulaproperty:"text" description:"text" nebulaindexes:"text" json:"text,omitempty"`
	Diplomacy             int     `nebulaproperty:"diplomacy" description:"diplomacy" nebulaindexes:"diplomacy" json:"diplomacy,omitempty"`
	Martial               int     `nebulaproperty:"martial" description:"martial" nebulaindexes:"martial" json:"martial,omitempty"`
	Stewardship           int     `nebulaproperty:"stewardship" description:"stewardship" nebulaindexes:"stewardship" json:"stewardship,omitempty"`
	Intrigue              int     `nebulaproperty:"intrigue" description:"intrigue" nebulaindexes:"intrigue" json:"intrigue,omitempty"`
	Learning              int     `nebulaproperty:"learning" description:"learning" nebulaindexes:"learning" json:"learning,omitempty"`
	Fertility             float32 `nebulaproperty:"fertility" description:"fertility" nebulaindexes:"fertility" json:"fertility,omitempty"`
	Health                float32 `nebulaproperty:"health" description:"health" nebulaindexes:"health" json:"health,omitempty"`
	SexAppealOpinion      int     `nebulaproperty:"sex_appeal_opinion" description:"sex_appeal_opinion" nebulaindexes:"sex_appeal_opinion" json:"sex_appeal_opinion,omitempty"`
	CombatRating          int     `nebulaproperty:"combat_rating" description:"combat_rating" nebulaindexes:"combat_rating" json:"combat_rating,omitempty"`
	Aggression            float32 `nebulaproperty:"aggression" description:"aggression" nebulaindexes:"aggression" json:"aggression,omitempty"`
	GlobalRevoltRisk      float32 `nebulaproperty:"global_revolt_risk" description:"global_revolt_risk" nebulaindexes:"global_revolt_risk" json:"global_revolt_risk,omitempty"`
	WarningLevel          float32 `nebulaproperty:"warning_level" description:"warning_level" nebulaindexes:"warning_level" json:"warning_level,omitempty"`
	PlotPowerModifier     float32 `nebulaproperty:"plot_power_modifier" description:"plot_power_modifier" nebulaindexes:"plot_power_modifier" json:"plot_power_modifier,omitempty"`
	TownOpinion           int     `nebulaproperty:"town_opinion" description:"town_opinion" nebulaindexes:"town_opinion" json:"town_opinion,omitempty"`
	ChurchOpinion         int     `nebulaproperty:"church_opinion" description:"church_opinion" nebulaindexes:"church_opinion" json:"church_opinion,omitempty"`
}

func NewObjective(code string) *Objective {
	nebulaObjective := Objective{
		VID:  getObjectiveVid(code),
		Code: code,
	}
	return &nebulaObjective
}

func NewObjectiveByData(objective *objectives.Objective) *Objective {
	nebulaObjective := utils.Mapping[Objective](objective)
	nebulaObjective.VID = getObjectiveVid(nebulaObjective.Code)
	return &nebulaObjective
}

func getObjectiveVid(code string) string {
	return "objective." + code
}

func (m *Objective) InsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.InsertVertexes(space, m)
}

func (m *Objective) UpdateToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpdateVertexes(space, m)
}

func (m *Objective) UpsertToNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.UpsertVertexes(space, m)
}

func (m *Objective) DeleteFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexes(space, m)
}

func (m *Objective) DeleteFromNebulaWithEdge(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteVertexesWithEdges(space, m)
}

func (m *Objective) LoadFromNebula(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.LoadVertex(space, m)
}

func InsertObjectives(space *nebulagolang.Space, ms ...*Objective) *nebulagolang.Result {
	return nebulagolang.BatchInsertVertexes(space, 250, ms)
}

func UpdateObjectives(space *nebulagolang.Space, ms ...*Objective) *nebulagolang.Result {
	return nebulagolang.BatchUpdateVertexes(space, 250, ms)
}

func UpsertObjectives(space *nebulagolang.Space, ms ...*Objective) *nebulagolang.Result {
	return nebulagolang.BatchUpsertVertexes(space, 250, ms)
}

func DeleteObjectives(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getObjectiveVid(c)
	}
	return nebulagolang.DeleteVertexesByVids(space, cs...)
}

func DeleteObjectiveWithEdge(space *nebulagolang.Space, codes ...string) *nebulagolang.Result {
	cs := make([]string, len(codes))
	for i, c := range codes {
		cs[i] = getObjectiveVid(c)
	}
	return nebulagolang.DeleteVertexesWithEdgesByVids(space, cs...)
}

func DeleteAllObjective(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesByTag[Objective](space)
}

func DeleteAllObjectiveWithEdges(space *nebulagolang.Space) *nebulagolang.Result {
	return nebulagolang.DeleteAllVertexesWithEdgesByTag[Objective](space)
}

func GetObjectiveByVid(space *nebulagolang.Space, vid string) *nebulagolang.ResultT[*Objective] {
	return nebulagolang.GetVertexByVid[*Objective](space, vid)
}

func GetObjectiveByCode(space *nebulagolang.Space, code string) *nebulagolang.ResultT[*Objective] {
	return nebulagolang.GetVertexByVid[*Objective](space, getObjectiveVid(code))
}

func GetAllObjective(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]*Objective] {
	return nebulagolang.GetAllVertexesByQuery[*Objective](space, "")
}

func GetAllObjectiveVids(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesVIDsByQuery[*Objective](space, "")
}

func GetAllObjectiveCode(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Objective](space, "", "code", "")
}

func GetAllObjectiveName(space *nebulagolang.Space) *nebulagolang.ResultT[map[string]bool] {
	return nebulagolang.GetAllVertexesPropertyByQuery[*Objective](space, "", "name", "")
}
