package file

import (
	"os"
)

const (
	Read  = 4
	Write = 2
	Exec  = 1
)

// O_RDONLY：只读模式(read-only)
// O_WRONLY：只写模式(write-only)
// O_RDWR：读写模式(read-write)
// O_APPEND：追加模式(append)
// O_CREATE：文件不存在就创建(create a new file if none exists.)
// O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
// O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
// O_TRUNC：打开并清空文件******

type FileWriter struct {
	File     *os.File
	FilePath string
	IsAppend bool
}

func (f *FileWriter) WriteString(s string) error {
	_, err := f.File.WriteString(s)
	if err != nil {
		return err
	}
	f.File.Sync()
	return err
}

func (f *FileWriter) WriteByte(b []byte) error {
	_, err := f.File.Write(b)
	if err != nil {
		return nil
	}
	f.File.Sync()
	return err
}

func (f *FileWriter) Close() {
	f.File.Close()
}

// 0644 is means -rw-r--r--
// if you want to create a dir then use
// err := os.Mkdir("../src/p", os.ModePerm)
// where os.ModePerm means -rwxrwxrwx

//Write & Read
func NewFileWriter(filePath string, isAppend bool) (*FileWriter, error) {
	if isAppend {
		// If the file doesn't exist, create it, or open it and append to the file
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}

		return &FileWriter{
			File:     f,
			FilePath: filePath,
			IsAppend: isAppend,
		}, nil

	} else {
		//os.O_TRUNC will make is cover what you have, else can not cover
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}

		return &FileWriter{
			File:     f,
			FilePath: filePath,
			IsAppend: isAppend,
		}, nil
	}
}
