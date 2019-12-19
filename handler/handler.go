package handler

import (
	"fmt"
	"io"
	"log"

	"github.com/Xuanwo/storage/types/pairs"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"

	"github.com/Xuanwo/bard/contexts"
	"github.com/Xuanwo/bard/model"
	"github.com/Xuanwo/bard/utils"
)

// Create will handle poem create.
func Create(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		log.Printf("Parse form file: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	defer file.Close()

	contextType := info.Header.Get(context.ContentTypeHeaderKey)
	p := model.NewPoem(info.Filename, contextType)

	ef, err := utils.Encrypt(contexts.Server.Key, p.ID, file)
	if err != nil {
		log.Printf("utils Encrypt: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	err = contexts.Storage.Write(p.ID, ef, pairs.WithSize(info.Size))
	if err != nil {
		log.Printf("Storage write failed: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	err = contexts.DB.Create(p).Error
	if err != nil {
		log.Printf("DB create failed: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	_, _ = ctx.JSON(iris.Map{
		"url": fmt.Sprintf("%s/%s", contexts.Server.PublicURL, p.ShortID),
	})
}

// Get will handle poem content get.
func Get(ctx iris.Context) {
	id := ctx.Params().Get("short_id")

	p := &model.Poem{}
	err := contexts.DB.Where("short_id = ?", id).First(p).Error
	if err != nil {
		log.Printf("DB query failed: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	r, err := contexts.Storage.Read(p.ID)
	if err != nil {
		log.Printf("Storage read failed: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	df, err := utils.Decrypt(contexts.Server.Key, p.ID, r)
	if err != nil {
		log.Printf("utils Decrypt: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.ContentType(p.ContentType)
	_, err = io.Copy(ctx, df)
	if err != nil {
		log.Printf("HTTP write failed: %v", err)
		return
	}
}
