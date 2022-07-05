package utils

import (
	"Oenone/common/base"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const PATH = "static/"

func SaveUploadImages(ctx *gin.Context, files []*multipart.FileHeader, path string, suffix string) bool {
	downloadPath := filepath.Join(base.GLOBAL_RESOURCE[base.WorkDir].(string), PATH, path)
	if err := os.MkdirAll(downloadPath, os.ModePerm); err != nil {
		log.Println("[SaveUploadImages] 创建目录错误: " + err.Error())
		return false
	}
	// 多线程写入，高并发需要异步消息队列
	var wg sync.WaitGroup
	savedResult := make([]bool, len(files))
	for index, file := range files {
		wg.Add(1)
		go func(index int, file *multipart.FileHeader) {
			defer wg.Done()
			err := ctx.SaveUploadedFile(file, filepath.Join(downloadPath, strconv.Itoa(index))+suffix)
			if err != nil {
				log.Println(err)
				return
			}
			savedResult[index] = true
		}(index, file)
	}
	wg.Wait()
	return funk.All(savedResult)
}
