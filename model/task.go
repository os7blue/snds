package model

const (
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
	Name *string
	//加密 key，如果没填则不加密
	Key string
	//网盘类型
	Type string
	//cron 表达式
	Cron string
}

func (Task) NameExist(tasks []Task, name string) bool {

	for _, task := range tasks {

		if *task.Name == name {
			return true
		}

	}

	return false
}
