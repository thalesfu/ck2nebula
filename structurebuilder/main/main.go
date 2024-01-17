package main

import (
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang/build"
)

func main() {

	//build.RebuildTagWithIndexes[ck2nebula.CultureGroup](ck2nebula.SPACE)
	//build.RebuildTagWithIndexes[ck2nebula.Culture](ck2nebula.SPACE)
	build.RebuildEdgeWithIndexes[ck2nebula.CultureGroup_Culture](ck2nebula.SPACE)
}
