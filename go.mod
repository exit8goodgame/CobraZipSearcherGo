module github.com/exit8goodgame/CobraZipSearcherGo

go 1.21.5

require (
	github.com/exit8goodgame/CobraZipSearcherGo/handler v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.8.0
	github.com/yeka/zip v0.0.0-20231116150916-03d6312748a9 // indirect
	go.uber.org/zap v1.26.0
)

replace github.com/exit8goodgame/CobraZipSearcherGo/handler => ./pkg/handler

replace github.com/exit8goodgame/CobraZipSearcherGo/service => ./pkg/service

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/exit8goodgame/CobraZipSearcherGo/service v0.0.0-00010101000000-000000000000 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect
)
