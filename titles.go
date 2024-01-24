package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
)

func GenerateTitles(titles map[string]*save.Title) (
	[]*Title,
	[]*Title_BaseTitle,
	[]*Title_LiegeTitle,
	[]*Title_DejureLiegeTitle,
	[]*Title_AssimilatingLiegeTitle,
	[]*Title_Dynasty,
	[]*Title_People) {
	rts := make([]*Title, len(titles))
	rtlts := make([]*Title_LiegeTitle, 0)
	rtbts := make([]*Title_BaseTitle, 0)
	rtdlts := make([]*Title_DejureLiegeTitle, 0)
	rtalts := make([]*Title_AssimilatingLiegeTitle, 0)
	rtds := make([]*Title_Dynasty, 0)
	rtp := make([]*Title_People, 0)

	i := 0
	for _, title := range titles {
		rts[i] = NewTitleByData(title)

		if title.Liege != nil {
			if title.Liege.ID == "" && title.Liege.Title != "" {
				title.Liege.ID = title.Liege.Title
			}
			tlt := NewTitle_LiegeTitle(rts[i], NewTitleByData(title.Liege))
			rtlts = append(rtlts, tlt)
		}

		if title.BaseTitle != nil {
			if title.BaseTitle.ID == "" && title.BaseTitle.Title != "" {
				title.BaseTitle.ID = title.BaseTitle.Title
			}
			tbt := NewTitle_BaseTitle(rts[i], NewTitleByData(title.BaseTitle))
			rtbts = append(rtbts, tbt)
		}

		if title.DeJureLiege != nil {
			if title.DeJureLiege.ID == "" && title.DeJureLiege.Title != "" {
				title.DeJureLiege.ID = title.DeJureLiege.Title
			}
			tdlt := NewTitle_DejureLiegeTitle(rts[i], NewTitleByData(title.DeJureLiege))
			rtdlts = append(rtdlts, tdlt)
		}

		if title.AssimilatingLiege != nil {
			if title.AssimilatingLiege.ID == "" && title.AssimilatingLiege.Title != "" {
				title.AssimilatingLiege.ID = title.AssimilatingLiege.Title
			}
			talt := NewTitle_AssimilatingLiegeTitle(rts[i], NewTitleByData(title.AssimilatingLiege))
			talt.DeJureAssYears = title.DeJureAssYears
			rtalts = append(rtalts, talt)
		}

		if title.Dynasty != 0 {
			td := NewTitle_Dynasty(rts[i], NewDynasty(title.PlayID, title.Dynasty))
			rtds = append(rtds, td)
		}

		if title.Holder != 0 {
			tpe := NewTitle_People(rts[i], NewPeople(title.PlayID, title.Holder))
			rtp = append(rtp, tpe)
		}

		i++
	}

	return rts, rtbts, rtlts, rtdlts, rtalts, rtds, rtp
}
