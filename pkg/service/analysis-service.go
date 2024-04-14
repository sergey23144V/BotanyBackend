package service

import (
	context "context"
	"errors"
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
		analysis, err = a.CreateExcelTypeTrialSiteAnalysis(ctx, nil, transect, ecomorph)
		if err != nil {
			return nil, err
		}
		analysis.UserId = userId
		analysis, err = a.repository.AnalysisRepository.CreatAnalysis(ctx, analysis)
		if err != nil {
			return nil, err
		}
	} else if input.TypeAnalysis == api.TypeAnalysis_TypeAnalysisPlant {
		transect, err := a.repository.TransectRepository.GetTransectByIdForAnalysis(ctx, input.Transect)
		if err != nil {
			return nil, err
		}
		filter := a.repository.EcomorphRepository.GetWhereList(input.Ecomorph)
		ecomorph, err := a.repository.EcomorphRepository.GetListEcomorphById(ctx, filter)
		if err != nil {
			return nil, err
		}
		analysis, err = a.CreateExcelTypeAnalysisPlantAnalysis(ctx, nil, transect, ecomorph)
		if err != nil {
			return nil, err
		}
		if analysis == nil {
			return nil, errors.New("Failed Create Analysis")
		}
		analysis.UserId = userId
		analysis, err = a.repository.AnalysisRepository.CreatAnalysis(ctx, analysis)
		if err != nil {
			return nil, err
		}
	}
	return analysis, nil
}

func (t AnalysisServiceImpl) CreateExcelTypeAnalysisPlantAnalysis(ctx context.Context, idInput *resource.Identifier, transect *api.Transect, ecomorph []*api.Ecomorph) (*api.Analysis, error) {

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
	err = t.BasicFormTypePlant(sheetName, indexPlant, ecomorph, f)
	if err != nil {
		return nil, err
	}
	number := 0
	typePlants := map[string]int{}
	typePlantPoints := map[string]int{}
	ecomorphsAnalis := make(map[string]map[string]int)
	for _, item := range transect.TrialSite {
		for _, plant := range item.Plant {
			if typePlants[plant.TypePlant.Id.ResourceId] == 0 {
				indexPlant += 2
				number++
				typePlants[plant.TypePlant.Id.ResourceId] = indexPlant
				err = f.SetCellValue(sheetName, "A"+strconv.Itoa(indexPlant), number)
				if err != nil {
					return nil, err
				}

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

				err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), plant.TypePlant.Subtitle)
				if err != nil {
					return nil, err
				}

				err = f.MergeCell(sheetName, "C"+strconv.Itoa(indexPlant), "C"+strconv.Itoa(indexPlant+1))
				if err != nil {
					fmt.Println("Ошибка при объединении ячеек:", err)
					return nil, err
				}
				ecomorphsColumb := 'D'
				for _, ecomorphItem := range ecomorph {
					ecomorphsEntity := GetEcomorphsEntityFromTypePlant(plant.TypePlant, ecomorphItem)

					if ecomorphsEntity != nil {
						err = f.SetCellValue(sheetName, string(ecomorphsColumb)+strconv.Itoa(indexPlant), ecomorphsEntity.DisplayTable)
						if err != nil {
							return nil, err
						}
						err = f.SetCellValue(sheetName, string(ecomorphsColumb)+strconv.Itoa(indexPlant+1), ecomorphsEntity.Score)
						if err != nil {
							return nil, err
						}
						typePlantPoints[plant.TypePlant.Id.ResourceId] += int(ecomorphsEntity.Score)
					} else {
						continue
					}
					if ecomorphsAnalis[ecomorphItem.Id.ResourceId] == nil {
						ecomorphsAnalis[ecomorphItem.Id.ResourceId] = make(map[string]int)
					}
					ecomorphsAnalis[ecomorphItem.Id.ResourceId][ecomorphsEntity.DisplayTable]++
					ecomorphsColumb++
				}
			}
		}

	}

	indexPlant += 2
	ecomorphsColumb := 'D'

	for _, ecomorphsEntity := range ecomorphsAnalis {
		indexEcomorphs := indexPlant
		for key, value := range ecomorphsEntity {
			err = f.SetCellValue(sheetName, string(ecomorphsColumb)+strconv.Itoa(indexEcomorphs), key+" = "+strconv.Itoa(int(float64(value)/float64(number)*100))+"%")
			if err != nil {
				return nil, err
			}
			indexEcomorphs++
		}
		ecomorphsColumb++
	}

	for plant, index := range typePlants {

		err = f.SetCellValue(sheetName, string(ecomorphsColumb)+strconv.Itoa(index), "Итог = "+strconv.Itoa(typePlantPoints[plant]))
		if err != nil {
			return nil, err
		}

	}

	var id *resource.Identifier
	if idInput != nil {
		id = idInput
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	path := "./analysis/" + id.ResourceId + ".xlsx"

	analysis := &api.Analysis{
		Id:           id,
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

func (t AnalysisServiceImpl) BasicFormTypePlant(sheetName string, indexPlant int, ecomorph []*api.Ecomorph, f *excelize.File) error {

	err := f.SetCellValue(sheetName, "A"+strconv.Itoa(indexPlant), "№")
	if err != nil {
		return err
	}
	f.SetColWidth(sheetName, "A", "A", 3)

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
	err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), "Название на латыни")
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "C2", "C4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}
	column := 'D'
	for _, item := range ecomorph {
		err = f.SetCellValue(sheetName, string(column)+strconv.Itoa(indexPlant), item.Title)
		if err != nil {
			return err
		}

		column++
	}

	return err
}

func (t AnalysisServiceImpl) CreateExcelTypeTrialSiteAnalysis(ctx context.Context, idInput *resource.Identifier, transect *api.Transect, ecomorph *api.Ecomorph) (*api.Analysis, error) {

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
	err = t.BasicFormTrialSite(sheetName, indexPlant, ecomorph.Title, f)
	if err != nil {
		return nil, err
	}
	indexPlant += 3
	number := 0
	columnTrialSite := 'D'
	typePlants := map[string]int{}
	ecomorphUnique := map[string]bool{}
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

			if typePlants[plant.TypePlant.Id.ResourceId] == 0 {
				number++
				err = f.SetCellValue(sheetName, "A"+strconv.Itoa(indexPlant), number)
				if err != nil {
					return nil, err
				}

				err = f.MergeCell(sheetName, "A"+strconv.Itoa(indexPlant), "A"+strconv.Itoa(indexPlant+1))
				if err != nil {
					fmt.Println("Ошибка при объединении ячеек:", err)
					return nil, err
				}
				typePlants[plant.TypePlant.Id.ResourceId] = indexPlant
				err = f.SetCellValue(sheetName, "B"+strconv.Itoa(indexPlant), plant.TypePlant.Title)
				if err != nil {
					return nil, err
				}

				err = f.MergeCell(sheetName, "B"+strconv.Itoa(indexPlant), "B"+strconv.Itoa(indexPlant+1))
				if err != nil {
					fmt.Println("Ошибка при объединении ячеек:", err)
					return nil, err
				}
				ecomorphsEntity := GetEcomorphsEntityFromTypePlant(plant.TypePlant, ecomorph)
				err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), ecomorphsEntity.Title)
				if err != nil {
					return nil, err
				}
				if !ecomorphUnique[ecomorphsEntity.Id.ResourceId] {
					ecomorphUnique[ecomorphsEntity.Id.ResourceId] = true
				}
				err = f.MergeCell(sheetName, "C"+strconv.Itoa(indexPlant), "C"+strconv.Itoa(indexPlant+1))
				if err != nil {
					fmt.Println("Ошибка при объединении ячеек:", err)
					return nil, err
				}
				indexPlant += 2
				indexTrialSite += 2
			}

			err = f.SetCellValue(sheetName, string(columnTrialSite)+strconv.Itoa(typePlants[plant.TypePlant.Id.ResourceId]), plant.Coverage)
			if err != nil {
				return nil, err
			}

			err = f.SetCellValue(sheetName, string(columnTrialSite)+strconv.Itoa(typePlants[plant.TypePlant.Id.ResourceId]+1), RevealBall(int(plant.Coverage)))
			if err != nil {
				return nil, err
			}
			f.SetColWidth(sheetName, string(columnTrialSite), string(columnTrialSite), 5)
		}
		columnTrialSite++
	}

	err = f.SetCellValue(sheetName, "B"+strconv.Itoa(indexPlant), "всего видов: "+strconv.Itoa(number))
	if err != nil {
		return nil, err
	}
	err = f.MergeCell(sheetName, "B"+strconv.Itoa(indexPlant), "B"+strconv.Itoa(indexPlant+1))
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return nil, err
	}

	err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), "всего "+ecomorph.Title+" : "+strconv.Itoa(len(ecomorphUnique)))
	if err != nil {
		return nil, err
	}
	err = f.MergeCell(sheetName, "C"+strconv.Itoa(indexPlant), "C"+strconv.Itoa(indexPlant+1))
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return nil, err
	}
	columnTrialSite = 'D'

	for _, trialSite := range transect.TrialSite {
		err = f.SetCellValue(sheetName, string(columnTrialSite)+strconv.Itoa(indexPlant), strconv.Itoa(len(trialSite.Plant))+" вида")
		if err != nil {
			return nil, err
		}
		err = f.MergeCell(sheetName, string(columnTrialSite)+strconv.Itoa(indexPlant), string(columnTrialSite)+strconv.Itoa(indexPlant+1))
		if err != nil {
			fmt.Println("Ошибка при объединении ячеек:", err)
			return nil, err
		}
		columnTrialSite++
	}

	var id *resource.Identifier
	if idInput != nil {
		id = idInput
	} else {
		id = &resource.Identifier{ResourceId: pkg.GenerateUUID()}
	}

	path := "./analysis/" + id.ResourceId + ".xlsx"

	analysis := &api.Analysis{
		Id:           id,
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

func RevealBall(coverage int) string {
	switch {
	case coverage < 1:
		return "+"
	case coverage < 5:
		return "1"
	case coverage < 25:
		return "2"
	case coverage < 50:
		return "3"
	case coverage < 75:
		return "4"
	default:
		return "5"
	}
}

func GetEcomorphsEntityFromTypePlant(typePlant *api.TypePlant, ecomorph *api.Ecomorph) *api.EcomorphsEntity {
	for _, ecomorphEmtiti := range typePlant.EcomorphsEntity {
		if ecomorphEmtiti.Ecomorphs.Id.ResourceId == ecomorph.Id.ResourceId {
			return ecomorphEmtiti
		}
	}
	return nil
}

func (t AnalysisServiceImpl) BasicFormTrialSite(sheetName string, indexPlant int, ecomorphTitle string, f *excelize.File) error {

	err := f.SetCellValue(sheetName, "A"+strconv.Itoa(indexPlant), "№")
	if err != nil {
		return err
	}
	err = f.MergeCell(sheetName, "A2", "A4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}
	f.SetColWidth(sheetName, "A", "A", 3)

	err = f.SetCellValue(sheetName, "B"+strconv.Itoa(indexPlant), "Вид")
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "B2", "B4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}

	f.SetColWidth(sheetName, "B", "B", 25)

	err = f.SetCellValue(sheetName, "C"+strconv.Itoa(indexPlant), ecomorphTitle)
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "C2", "C4")
	if err != nil {
		fmt.Println("Ошибка при объединении ячеек:", err)
		return err
	}

	f.SetColWidth(sheetName, "C", "C", 25)

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

func (a AnalysisServiceImpl) RepeatedAnalysis(ctx context.Context, input *api.InputUpdateAnalysis) (*api.Analysis, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	analysis, err := a.GetAnalysis(ctx, &api.IdRequest{Id: input.Id})
	if err != nil {
		return nil, err
	}
	if analysis.TypeAnalysis == api.TypeAnalysis_TypeAnalysisTransect {
		transect, err := a.repository.TransectRepository.GetTransectByIdForAnalysis(ctx, analysis.Transect)
		if err != nil {
			return nil, err
		}
		ecomorph, err := a.repository.EcomorphRepository.GetEcomorphById(ctx, input.Ecomorph[0])
		if err != nil {
			return nil, err
		}
		analysis, err = a.CreateExcelTypeTrialSiteAnalysis(ctx, analysis.Id, transect, ecomorph)
		if err != nil {
			return nil, err
		}
		analysis.UserId = userId
		analysis, err = a.repository.AnalysisRepository.CreatAnalysis(ctx, analysis)
		if err != nil {
			return nil, err
		}
	} else if analysis.TypeAnalysis == api.TypeAnalysis_TypeAnalysisPlant {
		transect, err := a.repository.TransectRepository.GetTransectByIdForAnalysis(ctx, analysis.Transect)
		if err != nil {
			return nil, err
		}
		filter := a.repository.EcomorphRepository.GetWhereList(input.Ecomorph)
		ecomorph, err := a.repository.EcomorphRepository.GetListEcomorphById(ctx, filter)
		if err != nil {
			return nil, err
		}
		analysis, err = a.CreateExcelTypeAnalysisPlantAnalysis(ctx, analysis.Id, transect, ecomorph)
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

func (a AnalysisServiceImpl) GetAnalysis(ctx context.Context, request *api.IdRequest) (*api.Analysis, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	return a.repository.GetAnalysis(ctx, &api.Analysis{Id: request.Id, UserId: userId})
}

func (a AnalysisServiceImpl) GetListAnalysis(ctx context.Context, request *api.PagesRequest) (*api.AnalysisList, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	var page *api.PagesResponse
	list, err := a.repository.GetListAnalysis(ctx, &api.Analysis{UserId: userId}, request)
	if err != nil {
		return nil, err
	}
	if request != nil {
		page = &api.PagesResponse{Page: request.Page, Limit: request.Limit, Total: int32(len(list))}
	}
	return &api.AnalysisList{List: list, Page: page}, nil
}

func (a AnalysisServiceImpl) DeleteAnalysis(ctx context.Context, request *api.IdRequest) (*api.BoolResponse, error) {
	userId := middlewares.GetUserIdFromContext(ctx)
	result := &api.BoolResponse{Result: true}
	err := a.repository.DeleteAnalysis(ctx, &api.Analysis{Id: request.Id, UserId: userId})
	if err != nil {
		result.Result = false
		return result, err
	}
	return result, nil
}
