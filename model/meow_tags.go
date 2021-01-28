package model

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

type (
	MeowTag struct {
		ID          int64     `xorm:"'id' not null BIGINT(22) pk autoincr"`
		Name        string    `xorm:"'name' not null VARCHAR(191) unique(name)"`
		HostID      string    `xorm:"'host_id' not null VARCHAR(32)"`
		Description string    `xorm:"'description' not null VARCHAR(255)"`
		Background  string    `xorm:"'background' not null VARCHAR(255)"`
		CreatedAt   time.Time `xorm:"'created_at' not null created DATETIME"`
		UpdatedAt   time.Time `xorm:"'updated_at' not null updated DATETIME"`
	}

	meowTagStatic struct{}
)

var (
	MeowTagStatic = new(meowTagStatic)
)

func (MeowTag) TableName() string {
	return "meow_tags"
}

func (mts *meowTagStatic) InsertOrUpdate(ctx context.Context, meowTag *MeowTag) error {

	session := GetSession(ctx)
	defer session.Close()

	preMeowTagRecord, err := mts.getByName(session, meowTag.Name)
	if err != nil {
		return errors.WithStack(err)
	}

	if preMeowTagRecord != nil {

		if _, err := session.ID(preMeowTagRecord.ID).Update(meowTag); err != nil {
			return errors.WithStack(err)
		}

		meowTag.Name = preMeowTagRecord.Name
		meowTag.HostID = preMeowTagRecord.HostID
		return nil
	}

	_, err = session.Insert(meowTag)
	return errors.WithStack(err)
}

func (mts *meowTagStatic) GetMeowTagByName(ctx context.Context, name string) (*MeowTag, error) {

	session := GetSession(ctx)
	defer session.Close()

	return mts.getByName(session, name)
}

func (*meowTagStatic) getByName(session *xorm.Session, name string) (*MeowTag, error) {

	var meowTag MeowTag
	found, err := session.Where("name = ?", name).Get(&meowTag)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !found {
		return nil, nil
	}

	return &meowTag, nil
}

func (mts *meowTagStatic) GetMeowTagByID(ctx context.Context, id int64) (*MeowTag, error) {

	session := GetSession(ctx)
	defer session.Close()

	var meowTag MeowTag
	found, err := session.Where("id = ?", id).Get(&meowTag)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !found {
		return nil, nil
	}

	return &meowTag, nil
}

func (*meowTagStatic) FindMeowTags(ctx context.Context) ([]MeowTag, error) {

	session := GetSession(ctx)
	defer session.Close()

	var meowTags []MeowTag
	err := session.Find(&meowTags)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return meowTags, nil
}
