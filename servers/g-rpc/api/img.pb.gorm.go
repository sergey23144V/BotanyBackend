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

type ImgORM struct {
	gorm.Model
	CreatedAt *time.Time
	DeletedAt *time.Time `gorm:"index:soft_delete"`
	Id        string     `gorm:"type:uuid;primaryKey"`
	Name      string
	Path      string
	UpdatedAt *time.Time
	UserId    *string `gorm:"type:uuid;foreignKey:auth.User"`
}

// TableName overrides the default tablename generated by GORM
func (ImgORM) TableName() string {
	return "imgs"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Img) ToORM(ctx context.Context) (ImgORM, error) {
	to := ImgORM{}
	var err error
	if prehook, ok := interface{}(m).(ImgWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Decode(&Img{}, m.Id); err != nil {
		return to, err
	} else if v != nil {
		to.Id = v.(string)
	}
	to.Name = m.Name
	to.Path = m.Path
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
	if posthook, ok := interface{}(m).(ImgWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ImgORM) ToPB(ctx context.Context) (Img, error) {
	to := Img{}
	var err error
	if prehook, ok := interface{}(m).(ImgWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource.Encode(&Img{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Name = m.Name
	to.Path = m.Path
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
	if posthook, ok := interface{}(m).(ImgWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Img the arg will be the target, the caller the one being converted from

// ImgBeforeToORM called before default ToORM code
type ImgWithBeforeToORM interface {
	BeforeToORM(context.Context, *ImgORM) error
}

// ImgAfterToORM called after default ToORM code
type ImgWithAfterToORM interface {
	AfterToORM(context.Context, *ImgORM) error
}

// ImgBeforeToPB called before default ToPB code
type ImgWithBeforeToPB interface {
	BeforeToPB(context.Context, *Img) error
}

// ImgAfterToPB called after default ToPB code
type ImgWithAfterToPB interface {
	AfterToPB(context.Context, *Img) error
}

// DefaultCreateImg executes a basic gorm create call
func DefaultCreateImg(ctx context.Context, in *Img, db *gorm.DB) (*Img, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ImgORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadImg(ctx context.Context, in *Img, db *gorm.DB) (*Img, error) {
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
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ImgORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ImgORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ImgORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteImg(ctx context.Context, in *Img, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ImgORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ImgORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteImgSet(ctx context.Context, in []*Img, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&ImgORM{})).(ImgORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ImgORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ImgORM{})).(ImgORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ImgORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Img, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Img, *gorm.DB) error
}

// DefaultStrictUpdateImg clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateImg(ctx context.Context, in *Img, db *gorm.DB) (*Img, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateImg")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &ImgORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithAfterStrictUpdateSave); ok {
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

type ImgORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchImg executes a basic gorm update call with patch behavior
func DefaultPatchImg(ctx context.Context, in *Img, updateMask *field_mask.FieldMask, db *gorm.DB) (*Img, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Img
	var err error
	if hook, ok := interface{}(&pbObj).(ImgWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadImg(ctx, &Img{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ImgWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskImg(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ImgWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateImg(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ImgWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ImgWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Img, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ImgWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Img, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ImgWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Img, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type ImgWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Img, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetImg executes a bulk gorm update call with patch behavior
func DefaultPatchSetImg(ctx context.Context, objects []*Img, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Img, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Img, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchImg(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskImg patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskImg(ctx context.Context, patchee *Img, patcher *Img, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Img, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Path" {
			patchee.Path = patcher.Path
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

// DefaultListImg executes a gorm list call
func DefaultListImg(ctx context.Context, db *gorm.DB) ([]*Img, error) {
	in := Img{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ImgORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ImgORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Img{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ImgORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type ImgORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]ImgORM) error
}
