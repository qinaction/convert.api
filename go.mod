module convert.api

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/aliyun/aliyun-log-go-sdk v0.1.12
	github.com/coreos/etcd v3.3.27+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.7.0
	github.com/go-redis/redis/v7 v7.0.0-beta.4
	github.com/json-iterator/go v1.1.12
	github.com/kr/text v0.2.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.12.1 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/outreach-golang/etcd v1.4.1
	github.com/outreach-golang/logger v1.2.10
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/ugorji/go v1.2.0 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20211029224645-99673261e6eb // indirect
	golang.org/x/sys v0.0.0-20211031064116-611d5d643895 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211029142109-e255c875f7c7 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)
