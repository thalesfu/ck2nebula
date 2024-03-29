package main

import (
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang/build"
)

func main() {

	//build.RebuildTagWithIndexes[ck2nebula.CultureGroup](ck2nebula.SPACE)
	//build.RebuildTagWithIndexes[ck2nebula.Culture](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.CultureGroup_Culture](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.ReligionGroup](ck2nebula.SPACE)
	//build.RebuildTagWithIndexes[ck2nebula.Religion](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.ReligionGroup_Religion](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Title](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Title_LiegeTitle](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Title_BaseTitle](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Title_DejureLiegeTitle](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Title_AssimilatingLiegeTitle](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Title_Dynasty](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Title_People](ck2nebula.SPACE)
	//
	//build.RebuildTagWithIndexes[ck2nebula.Story](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Story_Title](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Modifier](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Province](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Story_Province](ck2nebula.SPACE)
	//
	//build.RebuildEdgeWithIndexes[ck2nebula.Province_Modifier](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Province_Culture](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Province_Religion](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Province_Title](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Building](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Baron](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Baron_Title](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Baron_Building](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Province_Baron](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Story_Baron](ck2nebula.SPACE)
	//
	//build.RebuildTagWithIndexes[ck2nebula.Dynasty](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Dynasty_Culture](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Dynasty_Religion](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Story_Dynasty](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Trait](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Trait_OppositeTrait](ck2nebula.SPACE)

	//build.RebuildTagWithIndexes[ck2nebula.Objective](ck2nebula.SPACE)

	build.RebuildTagWithIndexes[ck2nebula.People](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Culture](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_GFXCulture](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Religion](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_SecretReligion](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.Story_People](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Trait](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Modifier](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_ClaimTitle](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_ClaimTitle](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Dynasty](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_FamilyPeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_HostPeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_EmpirePeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_KillPeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_RelatePeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_LoverPeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_GuardianPeople](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Ambition](ck2nebula.SPACE)
	//build.RebuildEdgeWithIndexes[ck2nebula.People_Focus](ck2nebula.SPACE)

	//build.RebuildEdgeWithIndexes[ck2nebula.Story_Player](ck2nebula.SPACE)
}
