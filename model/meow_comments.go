package model

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type (
	MeowComment struct {
		ID        int64     `xorm:"'id' not null BIGINT(22) pk autoincr"`
		Content   string    `xorm:"'content' not null VARCHAR(191)"`
		HostID    string    `xorm:"'host_id' not null VARCHAR(32)"`
		TieziID   int64     `xorm:"'tiezi_id' not null int(11)"`
		CreatedAt time.Time `xorm:"'created_at' not null created DATETIME"`
		UpdatedAt time.Time `xorm:"'updated_at' not null updated DATETIME"`
	}

	meowCommentStatic struct{}
)

var (
	MeowCommentStatic = new(meowCommentStatic)
)

func (MeowComment) TableName() string {
	return "meow_comments"
}

func (mtzs *meowCommentStatic) InsertOrUpdate(ctx context.Context, meowComment *MeowComment) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Insert(meowComment)
	return errors.WithStack(err)
}

func (*meowCommentStatic) FindMeowCommentByTieziID(ctx context.Context, tieziID int64) ([]MeowComment, int64, error) {

	session := GetSession(ctx)
	defer session.Close()

	var meowComments []MeowComment
	err := session.Where("tiezi_id = ?", tieziID).Desc("id").Find(&meowComments)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	totalCnt, err := session.Where("tiezi_id = ?", tieziID).Count(&MeowComment{})
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return meowComments, totalCnt, nil
}
