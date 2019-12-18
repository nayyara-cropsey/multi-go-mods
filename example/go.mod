module github.com/nayyara-samuel/multi-go-mods/example

go 1.12

require (
	github.com/nayyara-samuel/multi-go-mods/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
)

replace github.com/nayyara-samuel/multi-go-mods/common => ../common
