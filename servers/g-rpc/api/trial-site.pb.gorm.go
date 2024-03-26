package api

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/v2/gorm"
	resource "github.com/infobloxopen/atlas-app-toolkit/v2/gorm/resource"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	strings "strings"
	time "time"
)

type AddPlantTrialSiteRequestORM struct {
	Count       int32
	Covered     int32
	IdPlant     interface{}
	IdTrialSite interface{}
}

// TableName overrides the default tablename generated by GORM
func (AddPlantTrialSiteRequestORM) TableName() string {
	return "add_plant_trial_site_requests"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *AddPlantTrialSiteRequest) ToORM(ctx context.Context) (AddPlantTrialSiteRequestORM, error) {
	to := AddPlantTrialSiteRequestORM{}
	var err error
	if prehook, ok := interface{}(m).(AddPlantTrialSiteRequestWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Decode(nil, m.IdTrialSite); err != nil {
		return to, err
	} else if v != nil {
		to.IdTrialSite = v
	}
	if v, err := resource.Decode(nil, m.IdPlant); err != nil {
		return to, err
	} else if v != nil {
		to.IdPlant = v
	}
	to.Covered = m.Covered
	to.Count = m.Count
	if posthook, ok := interface{}(m).(AddPlantTrialSiteRequestWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddPlantTrialSiteRequestORM) ToPB(ctx context.Context) (AddPlantTrialSiteRequest, error) {
	to := AddPlantTrialSiteRequest{}
	var err error
	if prehook, ok := interface{}(m).(AddPlantTrialSiteRequestWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Encode(nil, m.IdTrialSite); err != nil {
		return to, err
	} else {
		to.IdTrialSite = v
	}
	if v, err := resource.Encode(nil, m.IdPlant); err != nil {
		return to, err
	} else {
		to.IdPlant = v
	}
	to.Covered = m.Covered
	to.Count = m.Count
	if posthook, ok := interface{}(m).(AddPlantTrialSiteRequestWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type AddPlantTrialSiteRequest the arg will be the target, the caller the one being converted from

// AddPlantTrialSiteRequestBeforeToORM called before default ToORM code
type AddPlantTrialSiteRequestWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddPlantTrialSiteRequestORM) error
}

// AddPlantTrialSiteRequestAfterToORM called after default ToORM code
type AddPlantTrialSiteRequestWithAfterToORM interface {
	AfterToORM(context.Context, *AddPlantTrialSiteRequestORM) error
}

// AddPlantTrialSiteRequestBeforeToPB called before default ToPB code
type AddPlantTrialSiteRequestWithBeforeToPB interface {
	BeforeToPB(context.Context, *AddPlantTrialSiteRequest) error
}

// AddPlantTrialSiteRequestAfterToPB called after default ToPB code
type AddPlantTrialSiteRequestWithAfterToPB interface {
	AfterToPB(context.Context, *AddPlantTrialSiteRequest) error
}

type TrialSiteORM struct {
	gorm.Model
	CountTypes    int32
	Covered       int32
	CreatedAt     *time.Time
	DeletedAt     *time.Time
	Dominant      *TypePlantORM `gorm:"foreignKey:DominantId;references:Id"`
	DominantId    *string
	Id            string `gorm:"type:uuid;primaryKey"`
	Rating        int32
	SubDominant   *TypePlantORM `gorm:"foreignKey:SubDominantId;references:Id"`
	SubDominantId *string
	Title         string
	TransectId    *string
	UpdatedAt     *time.Time
	UserId        *string `gorm:"type:uuid;foreignKey:auth.User"`
}

// TableName overrides the default tablename generated by GORM
func (TrialSiteORM) TableName() string {
	return "trial_sites"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TrialSite) ToORM(ctx context.Context) (TrialSiteORM, error) {
	to := TrialSiteORM{}
	var err error
	if prehook, ok := interface{}(m).(TrialSiteWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Decode(&TrialSite{}, m.Id); err != nil {
		return to, err
	} else if v != nil {
		to.Id = v.(string)
	}
	to.Title = m.Title
	to.Covered = m.Covered
	to.Rating = m.Rating
	to.CountTypes = m.CountTypes
	if m.Dominant != nil {
		tempDominant, err := m.Dominant.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Dominant = &tempDominant
	}
	if m.SubDominant != nil {
		tempSubDominant, err := m.SubDominant.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.SubDominant = &tempSubDominant
	}
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		t := m.DeletedAt.AsTime()
		to.DeletedAt = &t
	}
	if m.UserId != nil {
		if v, err := resource.Decode(nil, m.UserId); err != nil {
			return to, err
		} else if v != nil {
			vv := v.(string)
			to.UserId = &vv
		}
	}
	if posthook, ok := interface{}(m).(TrialSiteWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TrialSiteORM) ToPB(ctx context.Context) (TrialSite, error) {
	to := TrialSite{}
	var err error
	if prehook, ok := interface{}(m).(TrialSiteWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Encode(&TrialSite{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Title = m.Title
	to.Covered = m.Covered
	to.Rating = m.Rating
	to.CountTypes = m.CountTypes
	if m.Dominant != nil {
		tempDominant, err := m.Dominant.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Dominant = &tempDominant
	}
	if m.SubDominant != nil {
		tempSubDominant, err := m.SubDominant.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.SubDominant = &tempSubDominant
	}
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if m.DeletedAt != nil {
		to.DeletedAt = timestamppb.New(*m.DeletedAt)
	}
	if m.UserId != nil {
		if v, err := resource.Encode(nil, *m.UserId); err != nil {
			return to, err
		} else {
			to.UserId = v
		}
	}
	if posthook, ok := interface{}(m).(TrialSiteWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TrialSite the arg will be the target, the caller the one being converted from

// TrialSiteBeforeToORM called before default ToORM code
type TrialSiteWithBeforeToORM interface {
	BeforeToORM(context.Context, *TrialSiteORM) error
}

// TrialSiteAfterToORM called after default ToORM code
type TrialSiteWithAfterToORM interface {
	AfterToORM(context.Context, *TrialSiteORM) error
}

// TrialSiteBeforeToPB called before default ToPB code
type TrialSiteWithBeforeToPB interface {
	BeforeToPB(context.Context, *TrialSite) error
}

// TrialSiteAfterToPB called after default ToPB code
type TrialSiteWithAfterToPB interface {
	AfterToPB(context.Context, *TrialSite) error
}

// DefaultCreateAddPlantTrialSiteRequest executes a basic gorm create call
func DefaultCreateAddPlantTrialSiteRequest(ctx context.Context, in *AddPlantTrialSiteRequest, db *gorm.DB) (*AddPlantTrialSiteRequest, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddPlantTrialSiteRequestORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddPlantTrialSiteRequestORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AddPlantTrialSiteRequestORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddPlantTrialSiteRequestORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

// DefaultApplyFieldMaskAddPlantTrialSiteRequest patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAddPlantTrialSiteRequest(ctx context.Context, patchee *AddPlantTrialSiteRequest, patcher *AddPlantTrialSiteRequest, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*AddPlantTrialSiteRequest, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"IdTrialSite" {
			patchee.IdTrialSite = patcher.IdTrialSite
			continue
		}
		if f == prefix+"IdPlant" {
			patchee.IdPlant = patcher.IdPlant
			continue
		}
		if f == prefix+"Covered" {
			patchee.Covered = patcher.Covered
			continue
		}
		if f == prefix+"Count" {
			patchee.Count = patcher.Count
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAddPlantTrialSiteRequest executes a gorm list call
func DefaultListAddPlantTrialSiteRequest(ctx context.Context, db *gorm.DB) ([]*AddPlantTrialSiteRequest, error) {
	in := AddPlantTrialSiteRequest{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddPlantTrialSiteRequestORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(AddPlantTrialSiteRequestORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	ormResponse := []AddPlantTrialSiteRequestORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddPlantTrialSiteRequestORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*AddPlantTrialSiteRequest{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AddPlantTrialSiteRequestORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddPlantTrialSiteRequestORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type AddPlantTrialSiteRequestORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]AddPlantTrialSiteRequestORM) error
}

// DefaultCreateTrialSite executes a basic gorm create call
func DefaultCreateTrialSite(ctx context.Context, in *TrialSite, db *gorm.DB) (*TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Preload("SubDominant").Preload("Dominant").Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TrialSiteORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadTrialSite(ctx context.Context, in *TrialSite, db *gorm.DB) (*TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := TrialSiteORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(TrialSiteORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type TrialSiteORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteTrialSite(ctx context.Context, in *TrialSite, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&TrialSiteORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type TrialSiteORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteTrialSiteSet(ctx context.Context, in []*TrialSite, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&TrialSiteORM{})).(TrialSiteORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&TrialSiteORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&TrialSiteORM{})).(TrialSiteORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type TrialSiteORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*TrialSite, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*TrialSite, *gorm.DB) error
}

// DefaultStrictUpdateTrialSite clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateTrialSite(ctx context.Context, in *TrialSite, db *gorm.DB) (*TrialSite, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateTrialSite")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &TrialSiteORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Preload("SubDominant").Preload("Dominant").Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type TrialSiteORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchTrialSite executes a basic gorm update call with patch behavior
func DefaultPatchTrialSite(ctx context.Context, in *TrialSite, updateMask *field_mask.FieldMask, db *gorm.DB) (*TrialSite, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj TrialSite
	var err error
	if hook, ok := interface{}(&pbObj).(TrialSiteWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadTrialSite(ctx, &TrialSite{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(TrialSiteWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskTrialSite(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(TrialSiteWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateTrialSite(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(TrialSiteWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type TrialSiteWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *TrialSite, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *TrialSite, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *TrialSite, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *TrialSite, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetTrialSite executes a bulk gorm update call with patch behavior
func DefaultPatchSetTrialSite(ctx context.Context, objects []*TrialSite, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*TrialSite, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*TrialSite, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchTrialSite(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskTrialSite patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTrialSite(ctx context.Context, patchee *TrialSite, patcher *TrialSite, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*TrialSite, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedDominant bool
	var updatedSubDominant bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Title" {
			patchee.Title = patcher.Title
			continue
		}
		if f == prefix+"Covered" {
			patchee.Covered = patcher.Covered
			continue
		}
		if f == prefix+"Rating" {
			patchee.Rating = patcher.Rating
			continue
		}
		if f == prefix+"CountTypes" {
			patchee.CountTypes = patcher.CountTypes
			continue
		}
		if !updatedDominant && strings.HasPrefix(f, prefix+"Dominant.") {
			updatedDominant = true
			if patcher.Dominant == nil {
				patchee.Dominant = nil
				continue
			}
			if patchee.Dominant == nil {
				patchee.Dominant = &TypePlant{}
			}
			if o, err := DefaultApplyFieldMaskTypePlant(ctx, patchee.Dominant, patcher.Dominant, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"Dominant.", db); err != nil {
				return nil, err
			} else {
				patchee.Dominant = o
			}
			continue
		}
		if f == prefix+"Dominant" {
			updatedDominant = true
			patchee.Dominant = patcher.Dominant
			continue
		}
		if !updatedSubDominant && strings.HasPrefix(f, prefix+"SubDominant.") {
			updatedSubDominant = true
			if patcher.SubDominant == nil {
				patchee.SubDominant = nil
				continue
			}
			if patchee.SubDominant == nil {
				patchee.SubDominant = &TypePlant{}
			}
			if o, err := DefaultApplyFieldMaskTypePlant(ctx, patchee.SubDominant, patcher.SubDominant, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"SubDominant.", db); err != nil {
				return nil, err
			} else {
				patchee.SubDominant = o
			}
			continue
		}
		if f == prefix+"SubDominant" {
			updatedSubDominant = true
			patchee.SubDominant = patcher.SubDominant
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedDeletedAt && strings.HasPrefix(f, prefix+"DeletedAt.") {
			if patcher.DeletedAt == nil {
				patchee.DeletedAt = nil
				continue
			}
			if patchee.DeletedAt == nil {
				patchee.DeletedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"DeletedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.DeletedAt, patchee.DeletedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"DeletedAt" {
			updatedDeletedAt = true
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
		if f == prefix+"UserId" {
			patchee.UserId = patcher.UserId
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTrialSite executes a gorm list call
func DefaultListTrialSite(ctx context.Context, db *gorm.DB) ([]*TrialSite, error) {
	in := TrialSite{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []TrialSiteORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TrialSiteORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TrialSite{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TrialSiteORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type TrialSiteORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]TrialSiteORM) error
}
