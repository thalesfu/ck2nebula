package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
	"time"
)

func GeneratePeople(characters map[int]*save.Character, cm map[string]string, rm map[string]string) (
	[]*People,
	[]*People_Culture,
	[]*People_GFXCulture,
	[]*People_Religion,
	[]*People_SecretReligion,
	[]*People_Trait,
	[]*People_Modifier,
	[]*People_ClaimTitle,
	[]*People_Dynasty,
	[]*People_FamilyPeople,
	[]*People_HostPeople,
	[]*People_EmpirePeople,
	[]*People_KillPeople,
	[]*People_LoverPeople,
	[]*People_GuardianPeople) {
	rp := make([]*People, len(characters))
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

	i := 0
	for _, c := range characters {
		rp[i] = NewPeopleByData(c)
		rp[i].ReligionName = rm[rp[i].Religion]
		rp[i].SecretReligionName = rm[rp[i].SecretReligion]
		rp[i].CultureName = cm[rp[i].Culture]
		rp[i].GFXCultureName = cm[rp[i].GFXCulture]

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
				pct := NewPeople_ClaimTitle(rp[i], NewTitle(rp[i].PlayID, claim.Title.ID))
				pct.Pressed = bool(claim.Pressed)
				pct.Weak = bool(claim.Weak)
				rpcts = append(rpcts, pct)
			}
		}

		if c.Dynasty != 0 {
			rpds = append(rpds, NewPeople_Dynasty(rp[i], NewDynasty(rp[i].PlayID, c.Dynasty)))
		}

		var sn string

		if rp[i].IsFemale {
			sn = "daughter"
		} else {
			sn = "son"
		}

		if c.Father != 0 {
			father := NewPeople(rp[i].PlayID, c.Father)
			fatherRelation := NewPeople_FamilyPeople(rp[i], father)
			fatherRelation.Relation = "father"
			rpfps = append(rpfps, fatherRelation)

			sonRelation := NewPeople_FamilyPeople(father, rp[i])
			sonRelation.Relation = sn
			rpfps = append(rpfps, sonRelation)
		}

		if c.RealFather != 0 {
			realFather := NewPeople(rp[i].PlayID, c.RealFather)
			realFatherRelation := NewPeople_FamilyPeople(rp[i], realFather)
			realFatherRelation.Relation = "real_father"
			rpfps = append(rpfps, realFatherRelation)

			sonRelation := NewPeople_FamilyPeople(realFather, rp[i])
			sonRelation.Relation = "real_son"
			rpfps = append(rpfps, sonRelation)
		}

		if c.Mother != 0 {
			mother := NewPeople(rp[i].PlayID, c.Mother)
			motherRelation := NewPeople_FamilyPeople(rp[i], mother)
			motherRelation.Relation = "mother"
			rpfps = append(rpfps, motherRelation)

			sonRelation := NewPeople_FamilyPeople(mother, rp[i])
			sonRelation.Relation = sn
			rpfps = append(rpfps, sonRelation)
		}

		if len(c.Spouse) > 0 {
			for _, spouse := range c.Spouse {
				spousePeople := NewPeople(rp[i].PlayID, spouse)
				spouseRelation := NewPeople_FamilyPeople(rp[i], spousePeople)
				spouseRelation.Relation = "spouse"
				rpfps = append(rpfps, spouseRelation)
			}
		}

		if len(c.Consort) > 0 {
			for _, consort := range c.Consort {
				consortPeople := NewPeople(rp[i].PlayID, consort)
				consortRelation := NewPeople_FamilyPeople(rp[i], consortPeople)
				consortRelation.Relation = "consort"
				rpfps = append(rpfps, consortRelation)
			}
		}

		if c.ConsortOf != 0 {
			consortOfPeople := NewPeople(rp[i].PlayID, c.ConsortOf)
			consortOfRelation := NewPeople_FamilyPeople(rp[i], consortOfPeople)
			consortOfRelation.Relation = "consort_of"
			rpfps = append(rpfps, consortOfRelation)
		}

		if len(c.Lover) > 0 {
			for _, lover := range c.Lover {
				loverPeople := NewPeople(rp[i].PlayID, lover)
				loverRelation := NewPeople_LoverPeople(rp[i], loverPeople)
				rplps = append(rplps, loverRelation)
			}
		}

		if c.Guardian != 0 {
			loverPeople := NewPeople(rp[i].PlayID, c.Guardian)
			loverRelation := NewPeople_GuardianPeople(rp[i], loverPeople)
			rpgps = append(rpgps, loverRelation)
		}

		if c.Host != 0 {
			hostPeople := NewPeople(rp[i].PlayID, c.Host)
			hostRelation := NewPeople_HostPeople(rp[i], hostPeople)
			rphps = append(rphps, hostRelation)
		}

		if c.Emp != 0 {
			empirePeople := NewPeople(rp[i].PlayID, c.Emp)
			empireRelation := NewPeople_EmpirePeople(rp[i], empirePeople)
			rpeps = append(rpeps, empireRelation)
		}

		if c.Killer != 0 {
			killerPeople := NewPeople(rp[i].PlayID, c.Killer)
			killerRelation := NewPeople_KillPeople(rp[i], killerPeople)
			rpkps = append(rpkps, killerRelation)
		}

		i++
	}

	return rp, rpcs, rpgcs, rprs, rpsrs, rpts, rpms, rpcts, rpds, rpfps, rphps, rpeps, rpkps, rplps, rpgps
}
