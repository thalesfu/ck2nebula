package ck2nebula

type CultureGroup struct {
	VID  string `nebulakey:"tagvid" nebulatagname:"culturegroup" nebulatagcomment:"culture group"`
	Name string `nebulaproperty:"name" description:"culture group name" nebulaindexes:"name"`
	Code string `nebulaproperty:"code" description:"culture group code" nebulaindexes:"code"`
}
