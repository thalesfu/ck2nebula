package ck2nebula

import (
	"github.com/thalesfu/paradoxtools/CK2/save"
	"sort"
	"strings"
)

func GenerateTitles(titles map[string]*save.Title) (
	[]*Title,
	[]*Title_BaseTitle,
	[]*Title_LiegeTitle,
	[]*Title_DejureLiegeTitle,
	[]*Title_AssimilatingLiegeTitle,
	[]*Title_Dynasty,
	[]*Title_People,
	map[int]int,
	map[int]map[int]bool) {
	rts := make([]*Title, len(titles))
	rtlts := make([]*Title_LiegeTitle, 0)
	rtbts := make([]*Title_BaseTitle, 0)
	rtdlts := make([]*Title_DejureLiegeTitle, 0)
	rtalts := make([]*Title_AssimilatingLiegeTitle, 0)
	rtds := make([]*Title_Dynasty, 0)
	rtp := make([]*Title_People, 0)
	supremeRuler := make(map[int]int)
	rulerChain := make(map[int]map[int]bool)

	// Sort title IDs by level (barony→county→duchy→kingdom→empire) so that
	// higher-level titles are processed last and win the supremeRuler overwrite.
	// This prevents map-iteration non-determinism from causing per-run oscillation.
	keys := make([]string, 0, len(titles))
	for k := range titles {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(a, b int) bool {
		return titleLevelSort(keys[a]) < titleLevelSort(keys[b])
	})

	i := 0
	for _, k := range keys {
		title := titles[k]
		rts[i] = NewTitleByData(title)

		if title.Liege != nil {
			if title.Liege.ID == "" && title.Liege.Title != "" {
				title.Liege.ID = title.Liege.Title
			}
			tlt := NewTitle_LiegeTitle(rts[i], NewTitleByData(titles[title.Liege.ID]))
			rtlts = append(rtlts, tlt)
		}

		if title.BaseTitle != nil {
			if title.BaseTitle.ID == "" && title.BaseTitle.Title != "" {
				title.BaseTitle.ID = title.BaseTitle.Title
			}
			tbt := NewTitle_BaseTitle(rts[i], NewTitleByData(titles[title.BaseTitle.ID]))
			rtbts = append(rtbts, tbt)
		}

		if title.DeJureLiege != nil {
			if title.DeJureLiege.ID == "" && title.DeJureLiege.Title != "" {
				title.DeJureLiege.ID = title.DeJureLiege.Title
			}
			tdlt := NewTitle_DejureLiegeTitle(rts[i], NewTitleByData(titles[title.DeJureLiege.ID]))
			rtdlts = append(rtdlts, tdlt)
		}

		if title.AssimilatingLiege != nil {
			if title.AssimilatingLiege.ID == "" && title.AssimilatingLiege.Title != "" {
				title.AssimilatingLiege.ID = title.AssimilatingLiege.Title
			}
			talt := NewTitle_AssimilatingLiegeTitle(rts[i], NewTitleByData(titles[title.AssimilatingLiege.ID]))
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

		supremeRuler[title.Holder], rulerChain[title.Holder] = getSupremeRuler(title, titles, nil)

		i++
	}

	return rts, rtbts, rtlts, rtdlts, rtalts, rtds, rtp, supremeRuler, rulerChain
}

func getSupremeRuler(title *save.Title, titles map[string]*save.Title, rulerChain map[int]bool) (int, map[int]bool) {
	if rulerChain == nil {
		rulerChain = make(map[int]bool)
	}

	rulerChain[title.Holder] = true

	if title.Liege == nil {
		return title.Holder, rulerChain
	}

	if title.Liege.ID == "" {
		return title.Holder, rulerChain
	}

	if title.Liege.ID == title.ID {
		return title.Holder, rulerChain
	}

	if l, ok := titles[title.Liege.ID]; ok {
		return getSupremeRuler(l, titles, rulerChain)
	}

	return title.Holder, rulerChain
}

// titleLevelSort returns a sort key so that lower-level titles sort before higher ones.
// Processing order: b_ < c_ < d_ < k_ < e_, so empire titles are processed last
// and their supremeRuler determination wins over lower-level titles for the same holder.
func titleLevelSort(id string) int {
	switch {
	case strings.HasPrefix(id, "e_"):
		return 4
	case strings.HasPrefix(id, "k_"):
		return 3
	case strings.HasPrefix(id, "d_"):
		return 2
	case strings.HasPrefix(id, "c_"):
		return 1
	default:
		return 0
	}
}
