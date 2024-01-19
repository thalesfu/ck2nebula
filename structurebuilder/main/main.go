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
	build.RebuildEdgeWithIndexes[ck2nebula.Title_DejureLiegeTitle](ck2nebula.SPACE)
	build.RebuildEdgeWithIndexes[ck2nebula.Title_AssimilatingLiegeTitle](ck2nebula.SPACE)
}