package ck2nebula

import (
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/CK2/save"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GenerateRelations(file *save.SaveFile, translations map[string]string) []*People_RelatePeople {
	if len(file.Relations) == 0 {
		return []*People_RelatePeople{}
	}

	pairs := make(map[*save.Character]map[*save.Character][]*save.PersonRelation)

	relations := make([]*People_RelatePeople, 0)

	for k, rd := range file.Relations {
		var srcString string
		if strings.HasPrefix(k, "rel_") {
			srcString = strings.TrimPrefix(k, "rel_")
		} else if strings.HasPrefix(k, "diplo_") {
			srcString = strings.TrimPrefix(k, "diplo_")
		}
		srcId, err := strconv.Atoi(srcString)
		if err != nil {
			continue
		}

		src := file.Characters[srcId]

		srcpart, ok := pairs[src]

		if !ok {
			srcpart = make(map[*save.Character][]*save.PersonRelation)
			pairs[src] = srcpart
		}

		for dstString, rel := range rd {
			dstId, err := strconv.Atoi(dstString)
			if err != nil {
				continue
			}

			dst := file.Characters[dstId]

			dstpart, ok := srcpart[dst]

			if !ok {
				dstpart = make([]*save.PersonRelation, 0)
				srcpart[dst] = dstpart
			}

			srcpart[dst] = append(srcpart[dst], rel)
		}
	}

	for src, srcpart := range pairs {
		for dst, rels := range srcpart {
			relation := generateRelations(src, dst, rels, translations, file)
			relations = append(relations, relation...)
		}
	}

	return relations
}

func generateRelations(src *save.Character, dst *save.Character, rels []*save.PersonRelation, translations map[string]string, file *save.SaveFile) []*People_RelatePeople {
	relations := make([]*People_RelatePeople, 0)
	rank := 0

	for _, rel := range rels {

		relv := golangutils.IndirectValue(reflect.ValueOf(rel))
		relt := relv.Type()

		if rel.RY != 0 {
			relation := generateReignedYearsRelation(src, dst, rel.RY)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.LooterHostilityDays > 0 {
			relation := generateLooterHostilityDaysRelation(src, dst, rel.LooterHostilityDays)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.Alliance {
			relation := generateAllianceRelation(src, dst)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.CanCallToWar {
			relation := generateCanCallToWarRelation(src, dst)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.DisableNonAggressionPacts {
			relation := generateDisableNonAggressionPactsRelation(src, dst)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.IsLooter {
			relation := generateIsLooterRelation(src, dst)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.Succession {
			relation := generateSuccessionRelation(src, dst)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++
		}

		if rel.NonAggressionPact {
			relation := generateNonAggressionPactRelation(src, dst)
			relation.Rank = rank
			relations = append(relations, relation)
			rank++

		}

		for i := 0; i < relt.NumField(); i++ {
			ft := relt.Field(i)
			fv := relv.Field(i)
			ftn := golangutils.IndirectValue(reflect.New(ft.Type)).Type().Name()
			var relation *People_RelatePeople

			switch ftn {
			case "CommonRelation":
				if !fv.IsNil() {
					cr := golangutils.IndirectValue(fv).Interface().(save.CommonRelation)
					relation = generateCommonRelation(src, dst, &cr, ft.Tag.Get("paradox_field"), translations)
				}
			case "CasusBelli":
				if !fv.IsNil() {
					cb := golangutils.IndirectValue(fv).Interface().(save.CasusBelli)
					relation = generateCasusBelliRelation(src, dst, &cb, ft.Tag.Get("paradox_field"), file)
				}
			case "Truce":
				if !fv.IsNil() {
					t := golangutils.IndirectValue(fv).Interface().(save.Truce)
					relation = generateTruceRelation(src, dst, &t, ft.Tag.Get("paradox_field"), file)
				}
			case "Tributary":
				if !fv.IsNil() {
					t := golangutils.IndirectValue(fv).Interface().(save.Tributary)
					relation = generateTributaryRelation(src, dst, &t, ft.Tag.Get("paradox_field"), file)
				}
			case "MarriageTie":
				if !fv.IsNil() {
					mt := golangutils.IndirectValue(fv).Interface().(save.MarriageTie)
					relation = generateMarriageTieRelation(src, dst, &mt, ft.Tag.Get("paradox_field"), file)
				}
			}

			if relation != nil {
				relation.Rank = rank
				relations = append(relations, relation)
				rank++
			}

		}
	}

	return relations
}

func generateMarriageTieRelation(src *save.Character, dst *save.Character, mt *save.MarriageTie, get string, file *save.SaveFile) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = get
	relation.Name = "婚姻关系"
	relation.Detail = relation.Name

	if mt.First > 0 {
		f, ok := file.Characters[mt.First]
		if ok {
			fd, ok := file.Dynasties[f.Dynasty]
			if ok {
				s, ok := file.Characters[mt.Second]
				if ok {
					sd, ok := file.Dynasties[s.Dynasty]
					if ok {
						relation.Detail += fmt.Sprintf(" %s家族的%s和%s家族的%s结婚", fd.Name, f.Name, sd.Name, s.Name)
					}
				}
			}
		}
	}

	return relation
}

func generateTributaryRelation(src *save.Character, dst *save.Character, t *save.Tributary, code string, file *save.SaveFile) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = code
	relation.Name = "朝贡"
	relation.Detail = relation.Name

	if t.TributaryType != "" {
		relation.Detail += fmt.Sprintf(" 朝贡类型:%s", t.TributaryType)
	}

	if t.Tributary != 0 {
		c, ok := file.Characters[t.Tributary]
		if ok {
			cd, ok := file.Dynasties[c.Dynasty]
			if ok {
				relation.Detail += fmt.Sprintf(" 朝贡者:%s家族的%s", cd.Name, c.Name)
			}
		}
	}

	if t.IncomeTributePercentage > 0 {
		relation.Detail += fmt.Sprintf(" 朝贡收入:%f%%", t.IncomeTributePercentage)
	}

	if t.ReinforceTributePercentage > 0 {
		relation.Detail += fmt.Sprintf(" 朝贡补给:%f%%", t.ReinforceTributePercentage)
	}

	return relation

}

func generateTruceRelation(src *save.Character, dst *save.Character, t *save.Truce, code string, file *save.SaveFile) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = code
	relation.Name = "休战协议"
	relation.Due = time.Time(t.Truce)
	relation.Detail = relation.Name

	if t.Name != "" {
		relation.Detail += fmt.Sprintf(" %s", t.Name)
	}

	if t.Attacker > 0 {
		c, ok := file.Characters[t.Attacker]
		if ok {
			cd, ok := file.Dynasties[c.Dynasty]
			if ok {
				relation.Detail += fmt.Sprintf(" 攻击者:%s家族的%s", cd.Name, c.Name)
			}
		}
	}

	if t.Defender > 0 {
		c, ok := file.Characters[t.Defender]
		if ok {
			cd, ok := file.Dynasties[c.Dynasty]
			if ok {
				relation.Detail += fmt.Sprintf(" 防御者:%s家族的%s", cd.Name, c.Name)
			}
		}
	}

	if t.PeaceOffer > 0 {
		relation.Detail += fmt.Sprintf(" peace offer:%d", t.PeaceOffer)
	}

	if !relation.Due.IsZero() {
		relation.Detail += fmt.Sprintf(" 截止日期:%s", relation.Due.Format("2006-01-02"))
	}

	return relation
}

func generateCasusBelliRelation(src *save.Character, dst *save.Character, cb *save.CasusBelli, code string, file *save.SaveFile) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = code
	relation.Name = "开战理由"
	relation.Due = time.Time(cb.Date)
	relation.Detail = relation.Name

	if cb.CasusBelli != "" {
		relation.Detail += fmt.Sprintf(" %s", cb.CasusBelli)
	}

	if cb.Actor > 0 {
		c, ok := file.Characters[cb.Actor]
		if ok {
			cd, ok := file.Dynasties[c.Dynasty]
			if ok {
				relation.Detail += fmt.Sprintf(" 发起者:%s家族的%s", cd.Name, c.Name)
			}
		}
	}

	if cb.Recipient > 0 {
		c, ok := file.Characters[cb.Recipient]
		if ok {
			cd, ok := file.Dynasties[c.Dynasty]
			if ok {
				relation.Detail += fmt.Sprintf(" 目标:%s家族的%s", cd.Name, c.Name)
			}
		}
	}

	if cb.ThirdParty > 0 {
		c, ok := file.Characters[cb.ThirdParty]
		if ok {
			cd, ok := file.Dynasties[c.Dynasty]
			if ok {
				relation.Detail += fmt.Sprintf(" 第三方:%s家族的%s", cd.Name, c.Name)
			}
		}
	}

	if cb.LandedTitle != nil {
		title, ok := file.Titles[cb.LandedTitle.ID]
		if ok {
			relation.Detail += fmt.Sprintf(" 领地:%s", title.Name)
		}
	}

	if !relation.Due.IsZero() {
		relation.Detail += fmt.Sprintf(" 截止日期:%s", relation.Due.Format("2006-01-02"))
	}

	return relation
}

func generateLooterHostilityDaysRelation(src *save.Character, dst *save.Character, days int) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "looter_hostility_days"
	relation.Name = "劫掠敌对"
	relation.Detail = fmt.Sprintf("劫掠敌对:%d天", days)

	return relation
}

func generateNonAggressionPactRelation(src *save.Character, dst *save.Character) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "non_aggression_pact"
	relation.Name = "互不侵犯条约"
	relation.Detail = "互不侵犯条约"

	return relation
}

func generateSuccessionRelation(src *save.Character, dst *save.Character) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "succession"
	relation.Name = "入赘"
	relation.Detail = "入赘"

	return relation
}

func generateIsLooterRelation(src *save.Character, dst *save.Character) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "is_looter"
	relation.Name = "劫掠"
	relation.Detail = "劫掠"

	return relation
}

func generateDisableNonAggressionPactsRelation(src *save.Character, dst *save.Character) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "disable_non_aggression_pacts"
	relation.Name = "不允许缔结互不侵犯条约"
	relation.Detail = "不允许缔结互不侵犯条约"

	return relation
}

func generateCanCallToWarRelation(src *save.Character, dst *save.Character) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "can_call_to_war"
	relation.Name = "可召集战争"
	relation.Detail = "可召集战争"

	return relation
}

func generateAllianceRelation(src *save.Character, dst *save.Character) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "alliance"
	relation.Name = "同盟"
	relation.Detail = "同盟"

	return relation
}

func generateCommonRelation(src *save.Character, dst *save.Character, cr *save.CommonRelation, code string, translations map[string]string) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = code
	relation.Name = translations[code]
	relation.Due = time.Time(cr.Date)
	relation.Detail = relation.Name

	if cr.OriginDescription != "" {
		relation.Detail += fmt.Sprintf(" %s", cr.OriginDescription)
	}

	if !relation.Due.IsZero() {
		relation.Detail += fmt.Sprintf(" 截止日期:%s", relation.Due.Format("2006-01-02"))
	}

	return relation
}

func generateReignedYearsRelation(src *save.Character, dst *save.Character, ry int) *People_RelatePeople {
	relation := NewPeople_RelatePeople(NewPeopleByData(src), NewPeopleByData(dst))
	relation.Code = "reigned_years"
	relation.Name = "统治时长"
	relation.Detail = fmt.Sprintf("统治时长:%d年", ry)

	return relation
}
