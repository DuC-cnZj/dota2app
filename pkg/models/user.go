package models

import (
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/scopes"

	"github.com/DuC-cnZj/dota2app/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	ID int `json:"id" gorm:"primaryKey;"`

	Name     string `json:"name" gorm:"type:varchar(80);not null;default:'';"`
	Email    string `json:"email" gorm:"type:varchar(80);uniqueIndex:uniq_email;not null;"`
	Password string `json:"-" gorm:"type:varchar(255);not null;default:'';"`
	Mobile   string `json:"mobile" gorm:"type:varchar(40);"`
	Note     string `json:"note" gorm:"VARCHAR(255);"`
	Intro    string `json:"intro" gorm:"type:text;"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Avatar          File `gorm:"polymorphic:Fileable;polymorphicValue:avatar;"`
	BackgroundImage File `gorm:"polymorphic:Fileable;polymorphicValue:background_image;"`
}

func (user *User) HistoryAvatars() []*File {
	var avatars = make([]*File, 0)
	utils.DB().Where("`user_id` = ? AND `fileable_type` = ? AND `fileable_id` IS NULL", user.ID, TypeAvatar).Find(&avatars)

	return avatars
}

func (user *User) HistoryAvatarsWithPaginate(page *int, size *int) ([]*File, int64) {
	var (
		total   int64
		avatars = make([]*File, 0)
		db      = utils.DB()
	)

	db.Scopes(scopes.Paginate(page, size), scopes.OrderByIdDesc()).Where("`user_id` = ? AND `fileable_type` = ? AND `fileable_id` IS NULL", user.ID, TypeAvatar).Find(&avatars)
	db.Model(&File{}).Where("`user_id` = ? AND `fileable_type` = ? AND `fileable_id` IS NULL", user.ID, TypeAvatar).Count(&total)

	return avatars, total
}

func (user *User) HistoryBackgrounds() []*File {
	var backgrounds = make([]*File, 0)
	utils.DB().Where("`user_id` = ? AND `fileable_type` = ? AND `fileable_id` IS NULL", user.ID, TypeBackgroundImage).Find(&backgrounds)

	return backgrounds
}

func (user *User) HistoryBackgroundsWithPaginate(page *int, size *int) ([]*File, int64) {
	var (
		total       int64
		backgrounds = make([]*File, 0)
		db          = utils.DB()
	)

	db.Model(&File{}).Where("`user_id` = ? AND `fileable_type` = ? AND `fileable_id` IS NULL", user.ID, TypeBackgroundImage).Count(&total)
	db.Scopes(scopes.Paginate(page, size), scopes.OrderByIdDesc()).Where("`user_id` = ? AND `fileable_type` = ? AND `fileable_id` IS NULL", user.ID, TypeBackgroundImage).Find(&backgrounds)

	return backgrounds, total
}
