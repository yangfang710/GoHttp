package model

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type (
	CatProfile struct {
		ID        int64     `xorm:"'id' not null pk autoincr BIGINT(11)" json:"id"`
		HostID    string    `xorm:"'host_id' not null VARCHAR(32) unique(host)" json:"host_id"`
		Nickname  string    `xorm:"'nickname' not null VARCHAR(64)" json:"nickname"`
		Avatar    string    `xorm:"'avatar' not null VARCHAR(191)" json:"avatar"`
		MeowID    string    `xorm:"'meow_id' not null VARCHAR(32)" json:"meow_id"`
		Level     int32     `xorm:"'level' not null TINYINT(3)" json:"level"`
		Money     float64   `xorm:"'money' not null DOUBLE" json:"money"`
		IsWorked  bool      `xorm:"'is_worked' not null BOOL" json:"is_worked"`
		CreatedAt time.Time `xorm:"'created_at' not null created DATETIME" json:"created_at"`
		UpdatedAt time.Time `xorm:"'updated_at' not null updated DATETIME" json:"updated_at"`
	}

	catProfileStatic struct{}
)

var (
	CatProfileStatic = new(catProfileStatic)
)

func (CatProfile) TableName() string {
	return "cat_profile"
}

func (*catProfileStatic) Insert(ctx context.Context, catProfile *CatProfile) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Insert(catProfile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (*catProfileStatic) GetByHostID(ctx context.Context, hostID string) (*CatProfile, error) {

	session := GetSession(ctx)
	defer session.Close()

	var catProfile CatProfile
	found, err := session.Where("host_id = ?", hostID).Get(&catProfile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if !found {
		return nil, nil
	}

	return &catProfile, nil
}
