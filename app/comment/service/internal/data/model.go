package data

import (
	"gorm.io/gorm"
	"time"
)

type CommentDraft struct {
	gorm.Model
	Uuid   string `gorm:"index;size:36"`
	Status int32  `gorm:"default:1"`
}

type Comment struct {
	CommentId    int32 `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CreationId   int32          `gorm:"index:creation"`
	CreationType int32          `gorm:"index:creation"`
	Uuid         string         `gorm:"index;size:36"`
	Agree        int32          `gorm:"index;type:int unsigned;default:0"`
	Comment      int32          `gorm:"type:int unsigned;default:0"`
}

type CommentAgree struct {
	gorm.Model
	CommentId int32  `gorm:"uniqueIndex:idx_unique"`
	Uuid      string `gorm:"uniqueIndex:idx_unique;size:36"`
	Status    int32  `gorm:"default:1"`
}

type SubComment struct {
	CommentId int32 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	RootId    int32          `gorm:"index"`
	ParentId  int32          `gorm:"index"`
	Uuid      string         `gorm:"index;size:36"`
	Reply     string         `gorm:"index;size:36"`
	Agree     int32          `gorm:"type:int unsigned;default:0"`
}

type Record struct {
	gorm.Model
	Uuid     string `gorm:"size:36"`
	CommonId int32
	Ip       string
}
