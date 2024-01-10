package ck2nebula

type Culture struct {
	VID                   string `nebulakey:"tagvid" nebulatagname:"culture" nebulatagcomment:"culture"`
	Name                  string `nebulaproperty:"name" description:"culture name" nebulaindexes:"name"`
	Code                  string `nebulaproperty:"code" description:"culture code" nebulaindexes:"code"`
	FromDynastyPrefix     string `nebulaproperty:"from_dynasty_prefix" description:"from dynasty prefix" nebulaindexes:"from_dynasty_prefix"`
	MaleParonym           string `nebulaproperty:"male_paron_ym" description:"male paron ym" nebulaindexes:"male_paron_ym"`
	FemaleParonym         string `nebulaproperty:"female_paron_ym" description:"female paron ym" nebulaindexes:"female_paron_ym"`
	DynastyTitleNames     bool   `nebulaproperty:"dynasty_title_names" description:"dynasty title names" nebulaindexes:"dynasty_title_names"`
	FounderNamedDynasties bool   `nebulaproperty:"founder_named_dynasties" description:"founder named dynasties" nebulaindexes:"founder_named_dynasties"`
	Caste                 bool   `nebulaproperty:"caste" description:"caste" nebulaindexes:"caste"`
	DynastyNameFirst      bool   `nebulaproperty:"dynasty_name_first" description:"dynasty name first" nebulaindexes:"dynasty_name_first"`
	DukesCalledKings      bool   `nebulaproperty:"dukes_called_kings" description:"dukes called kings" nebulaindexes:"dukes_called_kings"`
	CountTitlesHidden     bool   `nebulaproperty:"count_titles_hidden" description:"count titles hidden" nebulaindexes:"count_titles_hidden"`
	BaronTitlesHidden     bool   `nebulaproperty:"baron_titles_hidden" description:"baron titles hidden" nebulaindexes:"baron_titles_hidden"`
	AllowLooting          bool   `nebulaproperty:"allow_looting" description:"allow looting" nebulaindexes:"allow_looting"`
}
