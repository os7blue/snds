package model

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

const (

	//os 常量

	WIN    string = "windows"
	DARWIN string = "darwin"
	LINUX  string = "linux"

	//网盘类型
	BAIDU_PAN string = "BAIDU"
	ALI_PAN   string = "ALI"
)

type Task struct {
	//网盘服务商 token
	Token string
	//网盘根路径
	Path string
	//需备份路径 文件或者文件夹路径
	LocalPath []string
	//任务名称（会用来命名备份文件）
	Name string
	//加密 key，和备份时日期结合加密（yyyyMMddhhmmss 格式），如果没填则不加密
	Key string
	//网盘类型
	Type string
	//cron 表达式
	Cron string
}
