package operation

import (
	"archive/zip"
	"github.com/robfig/cron/v3"
	"github.com/tjfoc/gmsm/sm4"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"snds/model"
	"snds/option"
	"time"
)

type start struct {
}

func (s *start) StartTask() error {

	// 创建一个Cron调度器
	c := cron.New()
	//
	//for _, t := range option.Option.Tasks {
	//
	//	// 添加定时任务
	//	_, err := c.AddFunc(t.Cron, func() {
	//
	//		option.Logger.Infof("创建定时任务：%s", t.Name)
	//		s.RunTask(t)
	//
	//	})
	//	if err != nil {
	//		option.Logger.Errorf("创建定时任务[%s]失败:%s", t.Name, err.Error())
	//		return err
	//	}
	//
	//}
	s.RunTask(option.Option.Tasks[0])

	// 启动Cron调度器
	option.Logger.Infof("所有定时任务创建完成并启动")
	c.Start()

	select {}

}

func (s *start) RunTask(t model.Task) error {

	currentTime := time.Now()
	now := currentTime.Format("200601021504")

	option.Logger.Infof("任务：%s,%s次任务开始执行", t.Name, now)

	tempFileName := t.Name + now
	tempDir := filepath.Join(option.Option.TempPath, tempFileName)

	//创建临时目录
	err := os.Mkdir(tempDir, 0755) // 0755 是目录权限
	if err != nil {
		option.Logger.Errorf("任务：%s,%s次任务创建临时目录失败：%s", t.Name, now, err.Error())
		return err
	}

	//将任务本地文件列表中的文件拷贝到临时目录
	err = s.copyFilesAndDirs(t, tempDir)
	if err != nil {
		option.Logger.Errorf("任务：%s,%s次任务复制文件至临时目录失败：%s", t.Name, now, err.Error())
		return err
	}
	option.Logger.Infof("任务：%s,%s次任务复制文件至临时目录完成", t.Name, now)

	//打包临时目录
	zipPath := filepath.Join(option.Option.TempPath, tempFileName) + ".zip"
	option.Logger.Infof("临时目录%s", zipPath)
	err = s.zipFolder(tempDir, zipPath)
	if err != nil {
		option.Logger.Errorf("任务：%s,%s次任务打包 zip 失败：%s", t.Name, now, err.Error())
		return err
	}
	option.Logger.Infof("任务：%s,%s次任务打包完成", t.Name, now)

	//加密
	finallyFilePath := zipPath
	if t.Key != "" {
		encryptPath := tempDir + ".snds"
		key := []byte("1234567890123456")
		err := s.encryptFile(zipPath, encryptPath, key)
		if err != nil {
			option.Logger.Errorf("任务：%s,%s次任务加密失败：%s", t.Name, now, err.Error())

			return err
		}
		finallyFilePath = encryptPath
		option.Logger.Info("密码" + string(key))

		err = s.decryptFile(finallyFilePath, option.Option.TempPath+"/123.zip", key)
		if err != nil {
			return err
		}

		option.Logger.Infof("任务：%s,%s次任务打包加密完成，最终打包文件：%s,密钥：%s", t.Name, now, finallyFilePath, string(key))

	}

	return nil
}

func (s *start) copyFilesAndDirs(task model.Task, toPath string) error {
	for _, sourcePath := range task.LocalPath {
		info, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		destinationPath := filepath.Join(toPath, filepath.Base(sourcePath))

		if info.IsDir() {
			if err := copyDir(sourcePath, destinationPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destinationPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyDir(sourceDir, destinationDir string) error {
	if err := os.MkdirAll(destinationDir, 0755); err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		destinationPath := filepath.Join(destinationDir, entry.Name())

		if entry.IsDir() {
			if err := copyDir(sourcePath, destinationPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destinationPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(sourceFile, destinationFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

func (s *start) zipFolder(sourceDir, targetZipPath string) error {
	// 创建目标zip文件
	targetZipFile, err := os.Create(targetZipPath)
	if err != nil {
		return err
	}
	defer targetZipFile.Close()

	// 创建zip写入器
	zipWriter := zip.NewWriter(targetZipFile)
	defer zipWriter.Close()

	// 遍历文件夹并将文件和子文件夹添加到zip文件中
	err = filepath.Walk(sourceDir, func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建zip文件中的文件或目录条目
		zipFileHeader, err := zip.FileInfoHeader(file)
		if err != nil {
			return err
		}

		// 将文件或目录的路径修改为相对于源文件夹的路径
		zipFileHeader.Name, err = filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		// 写入文件或目录条目到zip文件
		if file.Mode().IsDir() {
			zipFileHeader.Name += "/"
		}

		zipEntry, err := zipWriter.CreateHeader(zipFileHeader)
		if err != nil {
			return err
		}

		if !file.Mode().IsDir() {
			fileToZip, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer fileToZip.Close()

			_, err = io.Copy(zipEntry, fileToZip)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (s *start) encryptFile(inputFile, outputFile string, key []byte) error {

	option.Logger.Infof("开始加密，in:%s out:%s", inputFile, outputFile)

	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := sm4.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, len(plaintext))
	block.Encrypt(ciphertext, plaintext)

	err = ioutil.WriteFile(outputFile, ciphertext, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *start) decryptFile(inputFile, outputFile string, key []byte) error {
	option.Logger.Infof("开始解密，in:%s out:%s", inputFile, outputFile)
	ciphertext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := sm4.NewCipher(key)
	if err != nil {
		return err
	}

	plaintext := make([]byte, len(ciphertext))
	block.Decrypt(plaintext, ciphertext)

	err = ioutil.WriteFile(outputFile, plaintext, 0644)
	if err != nil {
		return err
	}

	return nil
}
