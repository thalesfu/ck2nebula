package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
)

func GenerateTitles(path string, savePath string) ([]*Title, []*Title_LiegeTitle, []*Title_BaseTitle, []*Title_DejureLiegeTitle, []*Title_AssimilatingLiegeTitle) {
	titles, ok := save.LoadTitles(path, savePath)

	if !ok {
		return nil, nil, nil, nil, nil
	}

	rts := make([]*Title, len(titles))
	rtlts := make([]*Title_LiegeTitle, 0)
	rtbts := make([]*Title_BaseTitle, 0)
	rtdlts := make([]*Title_DejureLiegeTitle, 0)
	rtalts := make([]*Title_AssimilatingLiegeTitle, 0)

	i := 0
	for _, title := range titles {
		rts[i] = NewTitleByData(title)
		if title.IsDynamic {
			rts[i].IsDynamic = true
		}

		if title.Liege != nil {
			if title.Liege.ID == "" && title.Liege.Title != "" {
				title.Liege.ID = title.Liege.Title
			}
			tlt := NewTitle_LiegeTitle(rts[i], NewTitleByData(title.Liege))
			tlt.PlayID = title.PlayID
			tlt.PlayDate = title.PlayDate
			rtlts = append(rtlts, tlt)
		}

		if title.BaseTitle != nil {
			if title.BaseTitle.ID == "" && title.BaseTitle.Title != "" {
				title.BaseTitle.ID = title.BaseTitle.Title
			}
			tbt := NewTitle_BaseTitle(rts[i], NewTitleByData(title.BaseTitle))
			tbt.PlayID = title.PlayID
			tbt.PlayDate = title.PlayDate
			rtbts = append(rtbts, tbt)
		}

		if title.DeJureLiege != nil {
			if title.DeJureLiege.ID == "" && title.DeJureLiege.Title != "" {
				title.DeJureLiege.ID = title.DeJureLiege.Title
			}
			tdlt := NewTitle_DejureLiegeTitle(rts[i], NewTitleByData(title.DeJureLiege))
			tdlt.PlayID = title.PlayID
			tdlt.PlayDate = title.PlayDate
			rtdlts = append(rtdlts, tdlt)
		}

		if title.AssimilatingLiege != nil {
			if title.AssimilatingLiege.ID == "" && title.AssimilatingLiege.Title != "" {
				title.AssimilatingLiege.ID = title.AssimilatingLiege.Title
			}
			talt := NewTitle_AssimilatingLiegeTitle(rts[i], NewTitleByData(title.AssimilatingLiege))
			talt.DeJureAssYears = title.DeJureAssYears
			talt.PlayID = title.PlayID
			talt.PlayDate = title.PlayDate
			rtalts = append(rtalts, talt)
		}

		i++
	}

	return rts, rtlts, rtbts, rtdlts, rtalts
}
