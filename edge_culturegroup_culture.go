package ck2nebula

type CultureGroup_Culture struct {
	CultureGroup *CultureGroup `nebulaedgename:"culturegroup_culture" nebulaedgecomment:"culture group -> culture" nebulakey:"edgefrom"`
	Culture      *Culture      `nebulakey:"edgeto"`
}
