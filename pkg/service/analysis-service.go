package service

import (
	context "context"
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/pkg"
	"github.com/sergey23144V/BotanyBackend/pkg/middlewares"
	"github.com/sergey23144V/BotanyBackend/pkg/repository"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type AnalysisService interface {
	CreatAnalysis(context.Context, *api.InputCreateAnalysis) (*api.Analysis, error)
	RepeatedAnalysis(context.Context, *api.InputUpdateAnalysis) (*api.Analysis, error)
	GetAnalysis(context.Context, *api.IdRequest) (*api.Analysis, error)
	GetListAnalysis(context.Context, *api.PagesRequest) (*api.AnalysisList, error)
	DeleteAnalysis(context.Context, *api.IdRequest) (*api.BoolResponse, error)
}

type AnalysisServiceImpl struct {
	repository *repository.Repository
}

func NewAnalysisServiceImpl(repository *repository.Repository) AnalysisService {
	return AnalysisServiceImpl{repository}
}

func (a AnalysisServiceImpl) CreatAnalysis(ctx context.Context, input *api.InputCreateAnalysis) (*api.Analysis, error) {
	var analysis *api.Analysis
	userId := middlewares.GetUserIdFromContext(ctx)
	if input.TypeAnalysis == api.TypeAnalysis_TypeAnalysisTransect {
		transect, err := a.repository.TransectRepository.GetTransectByIdForAnalysis(ctx, input.Transect)
		if err != nil {
			return nil, err
		}
		ecomorph, err := a.repository.EcomorphRepository.GetEcomorphById(ctx, input.Ecomorph[0])
		if err != nil {
			return nil, err
		}
		analysis, err = a.CreateExcelTypeTrialSiteAnalysis(ctx, transect, ecomorph)
		if err != nil {
			return nil, err
		}
		analysis.UserId = userId
		analysis, err = a.repository.AnalysisRepository.CreatAnalysis(ctx, analysis)
		if err != nil {
			return nil, err
		}
	}
	return analysis, nil
}

func (t AnalysisServiceImpl) CreateExcelTypeTrialSiteAnalysis(ctx context.Context, transect *api.Transect, ecomorph *api.Ecomorph) (*api.Analysis, error) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "Лист 1"
	err := f.SetSheetName("Sheet1", sheetName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	f.SetActiveSheet(1)
	indexPlant := 2
	err = t.BasicForm(sheetName, indexPlant, ecomorph.Title, f)
	if err != nil {
		return nil, err
	}
	indexPlant += 3
	number := 1
	columnTrialSite := 'D'
	indexTrialSite := indexPlant
	for i, item := range transect.TrialSite {
		err = f.SetCellValue(sheetName, string(columnTrialSite)+"3", strconv.Itoa(int(item.Covered))+" %")
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue(sheetName, string(columnTrialSite)+"4", i+1)
		if err != nil {
			return nil, err
		}
		for _, plant := range item.Plant {

			err = f.SetCellValue(sheetName, "A"+strconv.Itoa(indexPlant), number)
			if err != nil {
				return nil, err
			}
			number++
			err = f.MergeCell(sheetName, "A"+strconv.Itoa(indexPlant), "A"+strconv.Itoa(indexPlant+1))
			if err != nil {
				fmt.Println("Ошибка при объединении ячеек:", err)
				return nil, err
			}

			err = f.SetCellValue(sheetName, "B"+strconv.Itoa(indexPlant), plant.TypePlant.Title)
			if err != nil {
				return nil, err
			}

			err = f.MergeCell(sheetName, "B"+strconv.Itoa(indexPlant), "B"+strconv.Itoa(indexPlant+1))
			if err != nil {
				fmt.Println("Ошибка при объединении ячеек:", err)
				return nil, err
			}

			err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), GetEcomorphsEntityNameFromTypePlant(plant.TypePlant, ecomorph))
			if err != nil {
				return nil, err
			}

			err = f.MergeCell(sheetName, "C"+strconv.Itoa(indexPlant), "C"+strconv.Itoa(indexPlant+1))
			if err != nil {
				fmt.Println("Ошибка при объединении ячеек:", err)
				return nil, err
			}

			err = f.SetCellValue(sheetName, string(columnTrialSite)+strconv.Itoa(indexTrialSite), plant.Coverage)
			if err != nil {
				return nil, err
			}

			err = f.SetCellValue(sheetName, string(columnTrialSite)+strconv.Itoa(indexTrialSite+1), plant.Count)
			if err != nil {
				return nil, err
			}
			indexPlant += 2
			indexTrialSite += 2
		}
		columnTrialSite++
	}

	id := pkg.GenerateUUID()
	path := "../analysis/" + id + ".xlsx"
	analysis := &api.Analysis{
		Id:           &resource.Identifier{ResourceId: id},
		TypeAnalysis: api.TypeAnalysis_TypeAnalysisTransect,
		Title:        "",
		Transect:     transect,
		Path:         path,
	}

	if err := f.SaveAs(path); err != nil {
		fmt.Println(err)
	}

	return analysis, nil
}

func GetEcomorphsEntityNameFromTypePlant(typePlant *api.TypePlant, ecomorph *api.Ecomorph) string {
	for _, ecomorphEmtiti := range typePlant.EcomorphsEntity {
		if ecomorphEmtiti.Ecomorphs.Id.ResourceId == ecomorph.Id.ResourceId {
			return ecomorphEmtiti.Title
		}
	}
	return ""
}

func (t AnalysisServiceImpl) BasicForm(sheetName string, indexPlant int, ecomorphTitle string, f *excelize.File) error {

	err := f.SetCellValue(sheetName, "A"+strconv.Itoa(indexPlant), "№")
	if err != nil {
		return err
	}
	err = f.MergeCell(sheetName, "A2", "A4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}

	err = f.SetCellValue(sheetName, "B"+strconv.Itoa(indexPlant), "Вид")
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "B2", "B4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}

	err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), ecomorphTitle)
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "C2", "C4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}

	err = f.SetCellValue(sheetName, "D"+strconv.Itoa(indexPlant), "Проективное покрытие каждой площадки номер пробной площадки")
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "D"+strconv.Itoa(indexPlant), "I"+strconv.Itoa(indexPlant))
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}
	return err
}

func (a AnalysisServiceImpl) RepeatedAnalysis(ctx context.Context, analysis *api.InputUpdateAnalysis) (*api.Analysis, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServiceImpl) GetAnalysis(ctx context.Context, request *api.IdRequest) (*api.Analysis, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServiceImpl) GetListAnalysis(ctx context.Context, request *api.PagesRequest) (*api.AnalysisList, error) {
	//TODO implement me
	panic("implement me")
}

func (a AnalysisServiceImpl) DeleteAnalysis(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	//TODO implement me
	panic("implement me")
}
