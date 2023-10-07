package model

const (

	//os 常量

	WIN    string = "windows"
	DARWIN string = "darwin"
	LINUX  string = "linux"
)

type Option struct {
	//系统标识
	Os string
	//打包后执行文件的名称
	RunName string
	//执行文件所在的目录
	AppPath string
	//执行文件同级目录生成应用目录
	DataPath string
	//应用目录下用来做文件缓存的目录
	TempPath string
	//应用目录下用来存放配置文件的目录
	ConfigPath string
	//应用目录下用来存放 log 的目录
	LogPath string
	//从配置文件中解析出的任务列表
	Tasks []Task
	//根据不同的操作系统定义的命令集合
	Command Command
}
