package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ImageAdd(r *http.Request, folder string) interface{} {
	var fileNames []string
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return err
	}
	files := r.MultipartForm.File["file"]
	for _, fileHeader := range files {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			return nil
		}
		file, err := fileHeader.Open()
		errorManager(err)

		defer file.Close()
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		errorManager(err)

		filetype := http.DetectContentType(buff)
		checkType(filetype)

		_, err = file.Seek(0, io.SeekStart)
		errorManager(err)

		var fileName string = fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
		f, err := os.Create(folder + fileName)
		errorManager(err)

		defer f.Close()
		pr := &Progress{
			TotalSize: fileHeader.Size,
		}
		_, err = io.Copy(f, io.TeeReader(file, pr))
		errorManager(err)
		fileNames = append(fileNames, fileName)
	}
	return strings.Join(fileNames, ",")
}

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
type Progress struct {
	TotalSize int64
	BytesRead int64
}

func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		return
	}
}

func errorManager(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func checkType(filetype string) {
	if filetype != "image/jpeg" && filetype != "image/png" {
		panic("The provided file format is not allowed. Please upload a JPEG or PNG image")
	}
}
