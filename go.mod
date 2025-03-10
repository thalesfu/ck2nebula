module github.com/thalesfu/ck2nebula

go 1.23.6

require (
	github.com/thalesfu/golangutils v0.0.0-20250310030459-a6ea23977f07
	github.com/thalesfu/nebulagolang v0.0.0-20240108145007-c093dd041ca6
	github.com/thalesfu/paradoxtools v0.0.0-00010101000000-000000000000
)

require (
	github.com/ftrvxmtrx/tga v0.0.0-20150524081124-bd8e8d5be13a // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/vesoft-inc/fbthrift v0.0.0-20230214024353-fa2f34755b28 // indirect
	github.com/vesoft-inc/nebula-go/v3 v3.7.0 // indirect
	golang.org/x/image v0.23.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/thalesfu/nebulagolang => ../nebulagolang
	github.com/thalesfu/paradoxtools => ../paradoxtools
)
