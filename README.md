# StudyGolang
go语言相关体现
go 测试与性能
go test -bench .
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out


go 包安装
// go get -v -u github.com/gpmgo/gopm
go build github.com/gpmgo/gopm
go install github.com/gpmgo/gopm

//go get golang.org/x/text 国内安装不了
//gopm get -g -v golang.org/x/text
//git clone https://github.com/golang/text.git

// go get github.com/golang/tools/godoc 国内安装不了
//gopm get -g -v golang.org/x/tools
go build golang.org/x/tools/cmd/godoc

//gopm get -g -v golang.org/x/tools/cmd/goimports
go build golang.org/x/tools/cmd/goimports
go install golang.org/x/tools/cmd/goimports

//gopm get -g -v golang.org/x/net/html

阿里云学习地址
https://edu.aliyun.com/course/explore/golang?spm=5176.8764728.aliyun-edu-course-header.3.1e96a0beUC2uRc

配置文件
go get github.com/spf13/viper