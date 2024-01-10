module github.com/thalesfu/ck2nebula

go 1.19

require github.com/thalesfu/nebulagolang v0.0.0-20240108145007-c093dd041ca6

require (
	github.com/samber/lo v1.39.0 // indirect
	github.com/vesoft-inc/fbthrift v0.0.0-20230214024353-fa2f34755b28 // indirect
	github.com/vesoft-inc/nebula-go/v3 v3.6.1 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/thalesfu/nebulagolang => ../nebulagolang
	github.com/thalesfu/paradoxtools => ../paradoxtools
)
