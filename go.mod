module github.com/vigneshv59/ion-sfu

go 1.13

require (
	github.com/bep/debounce v1.2.0
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/envoyproxy/go-control-plane v0.9.9-0.20201210154907-fd9021fe5dad // indirect
	github.com/gammazero/deque v0.0.0-20201010052221-3932da5530cc
	github.com/gammazero/workerpool v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/improbable-eng/grpc-web v0.13.0
	github.com/lucsky/cuid v1.0.2
	github.com/pion/dtls/v2 v2.0.4
	github.com/pion/ice/v2 v2.0.15
	github.com/pion/ion-log v1.0.0
	github.com/pion/ion-sfu v0.0.0-00010101000000-000000000000
	github.com/pion/logging v0.2.2
	github.com/pion/rtcp v1.2.6
	github.com/pion/rtp v1.6.2
	github.com/pion/sdp/v3 v3.0.4
	github.com/pion/transport v0.12.2
	github.com/pion/turn/v2 v2.0.5
	github.com/pion/webrtc/v3 v3.0.4
	github.com/prometheus/client_golang v1.9.0
	github.com/rs/cors v1.7.0 // indirect
	github.com/soheilhy/cmux v0.1.4
	github.com/sourcegraph/jsonrpc2 v0.0.0-20200429184054-15c2290dcb37
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/ini.v1 v1.51.1 // indirect
)

replace github.com/pion/ion-sfu => github.com/vigneshv59/ion-sfu v1.8.1mod
