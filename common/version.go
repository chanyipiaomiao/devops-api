package common

var (
	// GoVersion Go版本
	GoVersion = "go1.10.3"

	// AppName 程序名称
	AppName = "devops-api"

	// AppVersion 程序版本号
	AppVersion = "1.0.0"

	// AppDescription 程序描述
	AppDescription = "happy with devops-api"

	// CommitHash git commit id
	CommitHash = ""

	// BuildDate 编译日期
	BuildDate = "2018-08-29"

	// Author 作者
	Author = "helei"

	// GitHub 地址
	GitHub = "https://github.com/chanyipiaomiao"
)

// GetVersion 获取版本信息
func GetVersion() map[string]string {
	return map[string]string{
		"goVersion":      GoVersion,
		"appName":        AppName,
		"appVersion":     AppVersion,
		"commitHash":     CommitHash,
		"buildDate":      BuildDate,
		"author":         Author,
		"gitHub":         GitHub,
		"appDescription": AppDescription,
	}
}
