package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ImgService interface {
	GetImgByID(context.Context, *api.IdRequest) (*api.Img, error)
	GetListImg(context.Context, *api.PagesRequest) (*api.ImgList, error)
	SaveImg(w http.ResponseWriter, r *http.Request)
	GetImageFromRequest(r *http.Request) (*api.Img, error)
}

type ImgServiceImpl struct {
	repository *repository.Repository
}

func (a ImgServiceImpl) GetImageFromRequest(r *http.Request) (*api.Img, error) {
	maxFileSize := 10 * 1024 * 1024 // Например, 10 МБ

	r.ParseMultipartForm(int64(maxFileSize))
	file, handler, err := r.FormFile("image")
	if err != nil {
		return nil, errors.New("Не удается получить файл")
	}
	defer file.Close()

	// Получаем расширение файла
	fileExt := filepath.Ext(handler.Filename)
	id := pkg.GenerateUUID()
	// Создаем новый файл для сохранения изображения на сервере
	path := "./img/"
	newFileName := path + id + fileExt
	newFileName2 := "/img/" + id + fileExt
	newFile, err := os.Create(newFileName)
	if err != nil {
		return nil, errors.New("Не удается создать файл для сохранения изображения")
	}
	defer newFile.Close()

	// Копируем содержимое файла в созданный файл
	_, err = io.Copy(newFile, file)
	if err != nil {
		return nil, errors.New("Ошибка при копировании файла")
	}

	userId := middlewares.GetUserIdFromContext(r.Context())
	localfile := &api.Img{
		Id:     &resource.Identifier{ResourceId: pkg.GenerateUUID()},
		Name:   handler.Filename,
		Path:   newFileName2,
		UserId: userId,
	}
	// В этом моменте, изображение сохранено на сервере с именем newFileName
	log.Info(context.Background(), "Изображение успешно загружено и сохранено как :", newFileName)
	return localfile, nil
}
func (i ImgServiceImpl) SaveImg(w http.ResponseWriter, r *http.Request) {

	err := middlewares.ValidToken(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	file, err := i.GetImageFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	img, err := i.repository.ImgRepositoryImpl.CreateImg(r.Context(), file)
	if err != nil {
		return
	}
	data := pkg.StructToMap(img)
	jsonData, err := json.Marshal(data)

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (i ImgServiceImpl) GetImgByID(ctx context.Context, request *api.IdRequest) (*api.Img, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return i.repository.ImgRepositoryImpl.GetImgByID(ctx, &api.Img{Id: request.Id, UserId: userId})
}

func (i ImgServiceImpl) GetListImg(ctx context.Context, response *api.PagesRequest) (*api.ImgList, error) {
	var page *api.PagesResponse
	userId := middlewares.GetUserIdFromContext(ctx)
	getList, err := i.repository.ImgRepositoryImpl.GetListImg(ctx, &api.ImgORM{UserId: &userId.ResourceId}, response)
	if response != nil {
		page = &api.PagesResponse{Page: response.Page, Limit: response.Limit, Total: int32(len(getList))}
	}
	if err != nil {
		return nil, err
	}
	result := &api.ImgList{List: getList, Page: page}
	return result, nil
}

func NewImgServiceImpl(repository *repository.Repository) ImgService {
	return ImgServiceImpl{repository}
}
