// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
)

type AnalysisMutation struct {
	CreatAnalysis    *api.Analysis     `json:"creatAnalysis"`
	RepeatedAnalysis *api.Analysis     `json:"repeatedAnalysis"`
	DeleteAnalysis   *api.BoolResponse `json:"deleteAnalysis"`
}

type AnalysisQuery struct {
	GetAnalysis     *api.Analysis     `json:"getAnalysis"`
	GetListAnalysis *api.AnalysisList `json:"getListAnalysis"`
}

type AuthMutation struct {
	SignUpUser *api.SignInUserResponse `json:"signUpUser,omitempty"`
	SignInUser *api.SignInUserResponse `json:"signInUser,omitempty"`
}

type EcomorphMutation struct {
	InsertEcomorph     *api.Ecomorph     `json:"insertEcomorph"`
	UpdateEcomorph     *api.Ecomorph     `json:"updateEcomorph"`
	DeleteEcomorphByID *api.BoolResponse `json:"deleteEcomorphById"`
}

type EcomorphQuery struct {
	GetEcomorphByID *api.Ecomorph      `json:"getEcomorphById"`
	GetListEcomorph *api.EcomorphsList `json:"getListEcomorph"`
}

type EcomorphsEntityMutation struct {
	InsertEcomorphEntity     *api.EcomorphsEntity `json:"insertEcomorphEntity,omitempty"`
	UpdateEcomorphEntity     *api.EcomorphsEntity `json:"updateEcomorphEntity,omitempty"`
	DeleteEcomorphEntityByID *api.BoolResponse    `json:"deleteEcomorphEntityByID,omitempty"`
}

type EcomorphsEntityQuery struct {
	GetEcomorphEntityByID *api.EcomorphsEntity     `json:"getEcomorphEntityByID,omitempty"`
	GetAllEcomorphEntity  *api.EcomorphsEntityList `json:"getAllEcomorphEntity,omitempty"`
}

type IDRequest struct {
	ID string `json:"id"`
}

type ImgQuery struct {
	GetImgByID *api.Img     `json:"getImgByID,omitempty"`
	GetListImg *api.ImgList `json:"getListImg,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type TransectMutation struct {
	CreateTransect         *api.Transect     `json:"createTransect,omitempty"`
	UpTransect             *api.Transect     `json:"upTransect,omitempty"`
	AddTrialSiteToTransect *api.Transect     `json:"addTrialSiteToTransect,omitempty"`
	DeleteTransect         *api.BoolResponse `json:"deleteTransect,omitempty"`
}

type TransectQuery struct {
	GetTransect    *api.Transect     `json:"getTransect,omitempty"`
	GetAllTransect *api.TransectList `json:"getAllTransect,omitempty"`
}

type TrialSiteMutation struct {
	CreateTrialSite      *api.TrialSite    `json:"createTrialSite,omitempty"`
	UpTrialSite          *api.TrialSite    `json:"upTrialSite,omitempty"`
	AddPlantsToTrialSite *api.TrialSite    `json:"addPlantsToTrialSite,omitempty"`
	DeleteTrialSite      *api.BoolResponse `json:"deleteTrialSite,omitempty"`
	CreatePlant          *api.Plant        `json:"createPlant,omitempty"`
	UpdatePlant          *api.Plant        `json:"updatePlant,omitempty"`
	DeletePlant          *api.BoolResponse `json:"deletePlant,omitempty"`
}

type TrialSiteQuery struct {
	GetTrialSite    *api.TrialSite     `json:"getTrialSite,omitempty"`
	GetAllTrialSite *api.TrialSiteList `json:"getAllTrialSite,omitempty"`
	GetPlant        *api.Plant         `json:"getPlant,omitempty"`
	GetAllPlant     *api.PlantList     `json:"getAllPlant,omitempty"`
}

type TypePlantMutation struct {
	CreateTypePlant               *api.TypePlant    `json:"createTypePlant,omitempty"`
	UpdateTypePlant               *api.TypePlant    `json:"updateTypePlant,omitempty"`
	AddEcomorphsEntityToTypePlant *api.TypePlant    `json:"addEcomorphsEntityToTypePlant,omitempty"`
	DeleteTypePlant               *api.BoolResponse `json:"deleteTypePlant,omitempty"`
}

type TypePlantQuery struct {
	GetTypePlant    *api.TypePlant     `json:"getTypePlant,omitempty"`
	GetAllTypePlant *api.TypePlantList `json:"getAllTypePlant,omitempty"`
}

type TypeAnalysis string

const (
	TypeAnalysisTypeAnalysisPlant    TypeAnalysis = "TypeAnalysisPlant"
	TypeAnalysisTypeAnalysisTransect TypeAnalysis = "TypeAnalysisTransect"
)

var AllTypeAnalysis = []TypeAnalysis{
	TypeAnalysisTypeAnalysisPlant,
	TypeAnalysisTypeAnalysisTransect,
}

func (e TypeAnalysis) IsValid() bool {
	switch e {
	case TypeAnalysisTypeAnalysisPlant, TypeAnalysisTypeAnalysisTransect:
		return true
	}
	return false
}

func (e TypeAnalysis) String() string {
	return string(e)
}

func (e *TypeAnalysis) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TypeAnalysis(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TypeAnalysis", str)
	}
	return nil
}

func (e TypeAnalysis) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
