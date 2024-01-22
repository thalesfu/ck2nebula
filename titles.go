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
	[]*Title_Dynasty) {
	rts := make([]*Title, len(titles))
	rtlts := make([]*Title_LiegeTitle, 0)
	rtbts := make([]*Title_BaseTitle, 0)
	rtdlts := make([]*Title_DejureLiegeTitle, 0)
	rtalts := make([]*Title_AssimilatingLiegeTitle, 0)
	rtds := make([]*Title_Dynasty, 0)

	i := 0
	for _, title := range titles {
		rts[i] = NewTitleByData(title)

		if title.Liege != nil {
			if title.Liege.ID == "" && title.Liege.Title != "" {
				title.Liege.ID = title.Liege.Title
			}
			tlt := NewTitle_LiegeTitle(rts[i], NewTitleByData(title.Liege))
			tlt.PlayID = title.PlayID
			rtlts = append(rtlts, tlt)
		}

		if title.BaseTitle != nil {
			if title.BaseTitle.ID == "" && title.BaseTitle.Title != "" {
				title.BaseTitle.ID = title.BaseTitle.Title
			}
			tbt := NewTitle_BaseTitle(rts[i], NewTitleByData(title.BaseTitle))
			tbt.PlayID = title.PlayID
			rtbts = append(rtbts, tbt)
		}

		if title.DeJureLiege != nil {
			if title.DeJureLiege.ID == "" && title.DeJureLiege.Title != "" {
				title.DeJureLiege.ID = title.DeJureLiege.Title
			}
			tdlt := NewTitle_DejureLiegeTitle(rts[i], NewTitleByData(title.DeJureLiege))
			tdlt.PlayID = title.PlayID
			rtdlts = append(rtdlts, tdlt)
		}

		if title.AssimilatingLiege != nil {
			if title.AssimilatingLiege.ID == "" && title.AssimilatingLiege.Title != "" {
				title.AssimilatingLiege.ID = title.AssimilatingLiege.Title
			}
			talt := NewTitle_AssimilatingLiegeTitle(rts[i], NewTitleByData(title.AssimilatingLiege))
			talt.DeJureAssYears = title.DeJureAssYears
			talt.PlayID = title.PlayID
			rtalts = append(rtalts, talt)
		}

		if title.Dynasty != 0 {
			td := NewTitle_Dynasty(rts[i], NewDynasty(title.PlayID, title.Dynasty))
			td.PlayID = title.PlayID
			rtds = append(rtds, td)
		}

		i++
	}

	return rts, rtbts, rtlts, rtdlts, rtalts, rtds
}
