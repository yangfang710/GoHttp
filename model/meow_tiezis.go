package model

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

type (
	MeowTiezi struct {
		ID        int64     `xorm:"'id' not null BIGINT(22) pk autoincr"`
		Title     string    `xorm:"'title' not null VARCHAR(191) unique(title)"`
		Content   string    `xorm:"'content' not null VARCHAR(191)"`
		HostID    string    `xorm:"'host_id' not null VARCHAR(32)"`
		Weight    int32     `xorm:"'weight' not null int(11)"`
		ReadCount int32     `xorm:"'read_count' not null int(11)"`
		TagID     int64     `xorm:"'tag_id' not null int(11)"`
		CreatedAt time.Time `xorm:"'created_at' not null created DATETIME"`
		UpdatedAt time.Time `xorm:"'updated_at' not null updated DATETIME"`
	}

	meowTieziStatic struct{}
)

var (
	MeowTieziStatic = new(meowTieziStatic)
)

func (MeowTiezi) TableName() string {
	return "meow_tiezis"
}

func (mtzs *meowTieziStatic) InsertOrUpdate(ctx context.Context, meowTiezi *MeowTiezi) error {

	session := GetSession(ctx)
	defer session.Close()

	preMeowTieziRecord, err := mtzs.getByID(session, meowTiezi.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if preMeowTieziRecord != nil {

		if _, err := session.ID(preMeowTieziRecord.ID).Update(meowTiezi); err != nil {
			return errors.WithStack(err)
		}

		meowTiezi.Title = preMeowTieziRecord.Title
		meowTiezi.Content = preMeowTieziRecord.Content
		meowTiezi.HostID = preMeowTieziRecord.HostID
		meowTiezi.Weight = preMeowTieziRecord.Weight
		meowTiezi.TagID = preMeowTieziRecord.TagID
		return nil
	}

	_, err = session.Insert(meowTiezi)
	return errors.WithStack(err)
}

func (mtzs *meowTieziStatic) GetMeowTieziByID(ctx context.Context, id int64) (*MeowTiezi, error) {

	session := GetSession(ctx)
	defer session.Close()

	return mtzs.getByID(session, id)
}

func (*meowTieziStatic) getByID(session *xorm.Session, id int64) (*MeowTiezi, error) {

	var meowTiezi MeowTiezi
	found, err := session.Where("id = ?", id).Get(&meowTiezi)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !found {
		return nil, nil
	}

	return &meowTiezi, nil
}

func (*meowTieziStatic) FindMeowTiezis(ctx context.Context, tagID int64, page, pageSize int) ([]MeowTiezi, int64, error) {

	session := GetSession(ctx)
	defer session.Close()

	var meowTiezis []MeowTiezi
	err := session.Where("tag_id = ?", tagID).Desc("weight").Limit(pageSize, pageSize*(page-1)).Find(&meowTiezis)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	totalCnt, err := session.Where("tag_id = ?", tagID).Count(&MeowTiezi{})
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return meowTiezis, totalCnt, nil
}

func (*meowTieziStatic) CountByTagID(ctx context.Context, tagID int64) (int64, error) {

	session := GetSession(ctx)
	defer session.Close()

	totalCnt, err := session.Where("tag_id = ?", tagID).Count(&MeowTiezi{})
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return totalCnt, nil
}

func (mtzs *meowTieziStatic) Increase(ctx context.Context, tieziID int64) (bool, error) {

	session := GetSession(ctx)
	defer session.Close()

	affectedCnt, err := session.Where("id = ? ", tieziID).Incr("read_count").Update(&MeowTiezi{})
	if err != nil {
		return false, errors.WithStack(err)
	}

	return affectedCnt > 0, nil
}
