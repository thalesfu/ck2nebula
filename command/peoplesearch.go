package command

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
)

func FindPeopleByTrait(space *nebulagolang.Space, playId int, trait string) *nebulagolang.ResultT[[]*ck2nebula.People] {
	return nebulagolang.QueryVertexesByQueryToSlice[*ck2nebula.People](space, getFindPeopleByTraitCommand(playId, trait))
}

func getFindPeopleByTraitCommand(playId int, trait string) string {
	commands := []string{
		fmt.Sprintf("lookup on trait where trait.name==\"%s\" yield vertex as v", trait),
		"yield id($-.v) as vid",
		fmt.Sprintf("go from $-.vid over people_trait reversely "+
			"where properties($$).isdead==false and properties($$).story_id==%d "+
			"yield $$ as v, properties($$).age as age, properties($$).isfemale as isfemale", playId),
		"order by $-.isfemale asc, $-.age asc",
	}
	return nebulagolang.CommandPipelineCombine(commands...)
}
