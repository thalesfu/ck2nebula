package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
	"time"
)

func GeneratePeople(file *save.SaveFile, cultureMap map[string]*Culture, religionMap map[string]*Religion, historyPeople map[int]map[int]*People, historyDynasty map[int]*Dynasty, supremeRuler map[int]int, rulerChain map[int]map[int]bool) ([]*People, []*People_Culture, []*People_GFXCulture, []*People_Religion, []*People_SecretReligion, []*People_Trait, []*People_Modifier, []*People_ClaimTitle, []*People_Dynasty, []*People_FamilyPeople, []*People_HostPeople, []*People_EmpirePeople, []*People_KillPeople, []*People_LoverPeople, []*People_GuardianPeople, []*People_Ambition, []*People_Focus) {
	rp := make([]*People, len(file.Characters))
	rpcs := make([]*People_Culture, 0)
	rpgcs := make([]*People_GFXCulture, 0)
	rprs := make([]*People_Religion, 0)
	rpsrs := make([]*People_SecretReligion, 0)
	rpts := make([]*People_Trait, 0)
	rpms := make([]*People_Modifier, 0)
	rpcts := make([]*People_ClaimTitle, 0)
	rpds := make([]*People_Dynasty, 0)
	rpfps := make([]*People_FamilyPeople, 0)
	rphps := make([]*People_HostPeople, 0)
	rpeps := make([]*People_EmpirePeople, 0)
	rpkps := make([]*People_KillPeople, 0)
	rplps := make([]*People_LoverPeople, 0)
	rpgps := make([]*People_GuardianPeople, 0)
	rpas := make([]*People_Ambition, 0)
	rpfs := make([]*People_Focus, 0)
	pm := make(map[int]*People)

	modifiers := GetAllModifier(SPACE).Data
	traitsMap := GetAllTraitsMap(SPACE).Data
	objectives := GetAllObjective(SPACE).Data

	i := 0
	for _, c := range file.Characters {
		rp[i] = NewPeopleByData(c)
		if d, ok := file.Dynasties[c.Dynasty]; ok {
			rp[i].DynastyName = d.Name
		}

		rp[i].Culture = getMyCulture(c, file.Characters, historyPeople, file.Dynasties, historyDynasty)
		rp[i].Religion = getMyReligion(c, file.Characters, historyPeople, file.Dynasties, historyDynasty)

		if r, ok := religionMap[rp[i].Religion]; ok {
			rp[i].ReligionName = r.Name
		}
		if sr, ok := religionMap[rp[i].SecretReligion]; ok {
			rp[i].SecretReligionName = sr.Name
		}
		if cul, ok := cultureMap[rp[i].Culture]; ok {
			rp[i].CultureName = cul.Name
		}
		if gcul, ok := cultureMap[rp[i].GFXCulture]; ok {
			rp[i].GFXCultureName = gcul.Name
		}

		rp[i].Father, rp[i].Mother, rp[i].GrandFather, rp[i].GrandMother, rp[i].MaternalGrandFather, rp[i].MaternalGrandMother = getParents(c, file.Characters)

		if sr, ok := supremeRuler[c.Emp]; ok {
			rp[i].SupreMeRuler = sr
		}

		if rc, ok := rulerChain[c.Emp]; ok {
			rp[i].IsUnderMyRule = rc[file.Player.ID]
		}

		for _, spouse := range c.Spouse {
			if ps, ok := file.Characters[spouse]; ok {
				if time.Time(ps.DeathDay).IsZero() {
					rp[i].Married = true
				}
			}
		}

		for _, consort := range c.Consort {
			if pc, ok := file.Characters[consort]; ok {
				if time.Time(pc.DeathDay).IsZero() {
					rp[i].ConsortCount = rp[i].ConsortCount + 1
				}
			}
		}

		if c.ConsortOf > 0 {
			if m, ok := file.Characters[c.ConsortOf]; ok {
				if time.Time(m.DeathDay).IsZero() {
					rp[i].IsConsort = true
				}
			}
		}

		for _, modifier := range c.Modifier {
			if m, ok := modifiers[getModifierVid(modifier.Modifier)]; ok {
				processPeopleModify(rp[i], m)
			}
		}

		for _, trait := range c.Traits {
			if t, ok := traitsMap[trait]; ok {
				processPeopleTrait(rp[i], t)
			}
		}

		if rp[i].Religion == "buddhist" {
			rp[i].ModifiedLearning += 4
		} else if rp[i].Religion == "taoist" {
			rp[i].ModifiedStewardship += 2
		}

		if rp[i].Culture != "" {
			rpcs = append(rpcs, NewPeople_Culture(rp[i], NewCulture(rp[i].Culture)))
		}

		if rp[i].GFXCulture != "" {
			rpgcs = append(rpgcs, NewPeople_GFXCulture(rp[i], NewCulture(rp[i].GFXCulture)))
		}

		if rp[i].Religion != "" {
			rprs = append(rprs, NewPeople_Religion(rp[i], NewReligion(rp[i].Religion)))
		}

		if rp[i].SecretReligion != "" {
			rpsrs = append(rpsrs, NewPeople_SecretReligion(rp[i], NewReligion(rp[i].SecretReligion)))
		}

		if len(c.Traits) > 0 {
			for _, trait := range c.Traits {
				rpts = append(rpts, NewPeople_Trait(rp[i], NewTrait(trait)))
			}
		}

		if len(c.Modifier) > 0 {
			for _, modifier := range c.Modifier {
				pm := NewPeople_Modifier(rp[i], NewModifier(modifier.Modifier))
				pm.Abate = time.Time(modifier.Date)
				pm.Hidden = bool(modifier.Hidden)
				rpms = append(rpms, pm)
			}
		}

		if len(c.Claim) > 0 {
			for _, claim := range c.Claim {
				pct := NewPeople_ClaimTitle(rp[i], NewTitle(rp[i].StoryID, claim.Title.ID))
				pct.Pressed = bool(claim.Pressed)
				pct.Weak = bool(claim.Weak)
				rpcts = append(rpcts, pct)
			}
		}

		if c.Dynasty != 0 {
			rpds = append(rpds, NewPeople_Dynasty(rp[i], NewDynasty(rp[i].StoryID, c.Dynasty)))
		}

		var sn string

		if rp[i].IsFemale {
			sn = "daughter"
		} else {
			sn = "son"
		}

		if c.Father != 0 {
			father := NewPeople(rp[i].StoryID, c.Father)
			fatherRelation := NewPeople_FamilyPeople(rp[i], father)
			fatherRelation.Relation = "father"
			rpfps = append(rpfps, fatherRelation)

			sonRelation := NewPeople_FamilyPeople(father, rp[i])
			sonRelation.Relation = sn
			rpfps = append(rpfps, sonRelation)
		}

		if c.RealFather != 0 {
			realFather := NewPeople(rp[i].StoryID, c.RealFather)
			realFatherRelation := NewPeople_FamilyPeople(rp[i], realFather)
			realFatherRelation.Relation = "real_father"
			rpfps = append(rpfps, realFatherRelation)

			sonRelation := NewPeople_FamilyPeople(realFather, rp[i])
			sonRelation.Relation = "real_" + sn
			rpfps = append(rpfps, sonRelation)
		}

		if c.Mother != 0 {
			mother := NewPeople(rp[i].StoryID, c.Mother)
			motherRelation := NewPeople_FamilyPeople(rp[i], mother)
			motherRelation.Relation = "mother"
			rpfps = append(rpfps, motherRelation)

			sonRelation := NewPeople_FamilyPeople(mother, rp[i])
			sonRelation.Relation = sn
			rpfps = append(rpfps, sonRelation)
		}

		if len(c.Spouse) > 0 {
			for _, spouse := range c.Spouse {
				spousePeople := NewPeople(rp[i].StoryID, spouse)
				spouseRelation := NewPeople_FamilyPeople(rp[i], spousePeople)
				spouseRelation.Relation = "spouse"
				rpfps = append(rpfps, spouseRelation)
			}
		}

		if len(c.Consort) > 0 {
			for _, consort := range c.Consort {
				consortPeople := NewPeople(rp[i].StoryID, consort)
				consortRelation := NewPeople_FamilyPeople(rp[i], consortPeople)
				consortRelation.Relation = "consort"
				rpfps = append(rpfps, consortRelation)
			}
		}

		if c.ConsortOf != 0 {
			consortOfPeople := NewPeople(rp[i].StoryID, c.ConsortOf)
			consortOfRelation := NewPeople_FamilyPeople(rp[i], consortOfPeople)
			consortOfRelation.Relation = "consort_of"
			rpfps = append(rpfps, consortOfRelation)
		}

		if len(c.Lover) > 0 {
			for _, lover := range c.Lover {
				loverPeople := NewPeople(rp[i].StoryID, lover)
				loverRelation := NewPeople_LoverPeople(rp[i], loverPeople)
				rplps = append(rplps, loverRelation)
			}
		}

		if c.Guardian != 0 {
			loverPeople := NewPeople(rp[i].StoryID, c.Guardian)
			loverRelation := NewPeople_GuardianPeople(rp[i], loverPeople)
			rpgps = append(rpgps, loverRelation)
		}

		if c.Host != 0 {
			hostPeople := NewPeople(rp[i].StoryID, c.Host)
			hostRelation := NewPeople_HostPeople(rp[i], hostPeople)
			rphps = append(rphps, hostRelation)
		}

		if c.Emp != 0 {
			empirePeople := NewPeople(rp[i].StoryID, c.Emp)
			empireRelation := NewPeople_EmpirePeople(rp[i], empirePeople)
			rpeps = append(rpeps, empireRelation)
		}

		if c.Killer != 0 {
			killerPeople := NewPeople(rp[i].StoryID, c.Killer)
			killerRelation := NewPeople_KillPeople(rp[i], killerPeople)
			rpkps = append(rpkps, killerRelation)
		}

		pm[rp[i].ID] = rp[i]

		i++
	}

	if len(file.ActiveFocuses) > 0 {
		for _, f := range file.ActiveFocuses {
			if f.Scope != nil && f.Scope.Char > 0 {
				p, ok := pm[f.Scope.Char]
				if ok {
					rpfs = append(rpfs, NewPeople_Focus(p, NewObjective(f.Type)))

					o, ok := objectives[getObjectiveVid(f.Type)]
					if ok {
						processPeopleObject(p, o)
					}
				}
			}
		}
	}

	if len(file.ActiveAmbitions) > 0 {
		for _, a := range file.ActiveAmbitions {
			if a.Scope != nil && a.Scope.Char > 0 {
				p, ok := pm[a.Scope.Char]
				if ok {
					rpas = append(rpas, NewPeople_Ambition(p, NewObjective(a.Type)))

					o, ok := objectives[getObjectiveVid(a.Type)]
					if ok {
						processPeopleObject(p, o)
					}
				}
			}
		}
	}

	return rp, rpcs, rpgcs, rprs, rpsrs, rpts, rpms, rpcts, rpds, rpfps, rphps, rpeps, rpkps, rplps, rpgps, rpas, rpfs
}

func getParents(c *save.Character, characters map[int]*save.Character) (int, int, int, int, int, int) {
	var grandFather, grandMother, maternalGrandFather, maternalGrandMother int

	if c.Father != 0 {
		if f, ok := characters[c.Father]; ok {
			grandFather = f.Father
			grandMother = f.Mother
		}
	}

	if c.Mother != 0 {
		if m, ok := characters[c.Mother]; ok {
			maternalGrandFather = m.Father
			maternalGrandMother = m.Mother
		}
	}

	return c.Father, c.Mother, grandFather, grandMother, maternalGrandFather, maternalGrandMother
}

func processPeopleObject(p *People, o *Objective) {
	p.ModifiedDiplomacy += o.Diplomacy
	p.ModifiedMartial += o.Martial
	p.ModifiedStewardship += o.Stewardship
	p.ModifiedIntrigue += o.Intrigue
	p.ModifiedLearning += o.Learning
	p.ModifiedHealth += o.Health
	p.ModifiedFertility += o.Fertility
	p.ModifiedSexAppeal += o.SexAppealOpinion
	p.ModifiedCombatRating += o.CombatRating
}

func processPeopleModify(p *People, m *Modifier) {
	p.ModifiedDiplomacy += m.Diplomacy
	p.ModifiedMartial += m.DiplomacyPenalty
	p.ModifiedMartial += m.Martial
	p.ModifiedMartial += m.MartialPenalty
	p.ModifiedStewardship += m.Stewardship
	p.ModifiedStewardship += m.StewardshipPenalty
	p.ModifiedIntrigue += m.Intrigue
	p.ModifiedIntrigue += m.IntriguePenalty
	p.ModifiedLearning += m.Learning
	p.ModifiedLearning += m.LearningPenalty
	p.ModifiedHealth += m.Health
	p.ModifiedHealth += m.HealthPenalty
	p.ModifiedFertility += m.Fertility
	p.ModifiedFertility += m.FertilityPenalty
	p.ModifiedSexAppeal += m.SexAppealOpinion
	p.ModifiedCombatRating += m.CombatRating
}

func processPeopleTrait(p *People, t *Trait) {
	p.ModifiedDiplomacy += t.Diplomacy
	p.ModifiedMartial += t.DiplomacyPenalty
	p.ModifiedMartial += t.Martial
	p.ModifiedMartial += t.MartialPenalty
	p.ModifiedStewardship += t.Stewardship
	p.ModifiedStewardship += t.StewardshipPenalty
	p.ModifiedIntrigue += t.Intrigue
	p.ModifiedIntrigue += t.IntriguePenalty
	p.ModifiedLearning += t.Learning
	p.ModifiedLearning += t.LearningPenalty
	p.ModifiedHealth += t.Health
	p.ModifiedHealth += t.HealthPenalty
	p.ModifiedFertility += t.Fertility
	p.ModifiedFertility += t.FertilityPenalty
	p.ModifiedSexAppeal += t.SexAppealOpinion
	p.ModifiedCombatRating += t.CombatRating
}

func getMyCulture(me *save.Character, characterMap map[int]*save.Character, historyPeople map[int]map[int]*People, dynasties map[int]*save.Dynasty, historyDynasty map[int]*Dynasty) string {
	if me.Culture != "" {
		return me.Culture
	}

	if me.Father != 0 {
		if father, ok := characterMap[me.Father]; ok {
			if father.Culture != "" {
				return father.Culture
			} else {
				if father.IsHistory {
					if history, ok := getHistoryPeople(father.ID, historyPeople); ok {
						if history.Culture != "" {
							return history.Culture
						}
					}
				}

				culture := getMyCulture(father, characterMap, historyPeople, nil, nil)

				if culture != "" {
					return culture
				}
			}
		}
	}

	if me.Dynasty != 0 {
		if dynasty, ok := dynasties[me.Dynasty]; ok {
			if dynasty.Culture != "" {
				return dynasty.Culture
			}
		}

		if history, ok := historyDynasty[me.Dynasty]; ok {
			if history.Culture != "" {
				return history.Culture
			}
		}
	}

	return ""
}

func getMyReligion(me *save.Character, characterMap map[int]*save.Character, historyPeople map[int]map[int]*People, dynasties map[int]*save.Dynasty, historyDynasty map[int]*Dynasty) string {
	if me.Religion != "" {
		return me.Religion
	}

	if me.Father != 0 {
		if father, ok := characterMap[me.Father]; ok {
			if father.Culture != "" {
				return father.Religion
			} else {
				if father.IsHistory {
					if history, ok := getHistoryPeople(father.ID, historyPeople); ok {
						if history.Religion != "" {
							return history.Religion
						}
					}
				}

				religion := getMyReligion(father, characterMap, historyPeople, nil, nil)

				if religion != "" {
					return religion
				}
			}
		}
	}

	if me.Dynasty != 0 {
		if dynasty, ok := dynasties[me.Dynasty]; ok {
			if dynasty.Religion != "" {
				return dynasty.Religion
			}
		}

		if history, ok := historyDynasty[me.Dynasty]; ok {
			if history.Religion != "" {
				return history.Religion
			}
		}
	}

	return ""
}

func getHistoryPeople(peopleId int, historyPeople map[int]map[int]*People) (*People, bool) {
	key := peopleId % 10000
	if _, ok := historyPeople[key]; !ok {
		return nil, false
	}

	if p, ok := historyPeople[key][peopleId]; ok {
		return p, true
	}

	return nil, false
}
