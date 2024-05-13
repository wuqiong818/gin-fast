package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户个人信息表
type User struct {
	Id               int            `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username         string         `json:"username" gorm:"column:username;comment:'用户虚拟名称，限制长度位10位'"`
	Password         string         `json:"password" gorm:"column:password"`
	Age              int            `json:"age" gorm:"column:age"`
	Sex              int            `json:"sex" gorm:"column:sex"`
	Avatar           string         `json:"avatar" gorm:"column:avatar;comment:'头像，存储头像的路径'"`
	RealName         string         `json:"realName" gorm:"column:real_name;comment:'真实名称'"`
	Telephone        string         `json:"telephone" gorm:"column:telephone;comment:'电话号码'"`
	Signature        string         `json:"signature" gorm:"column:signature;comment:'用户个性签名'"`
	Status           int            `json:"status" gorm:"column:status;comment:'1正常2禁用'"`
	Province         string         `json:"province" gorm:"column:province"`
	City             string         `json:"city" gorm:"column:city"`
	FansNum          int            `json:"fansNum" gorm:"column:fans_num;NOT NULL"`
	FollowersNum     int            `json:"followersNum" gorm:"column:followers_num"`
	PostsNum         int            `json:"postsNum" gorm:"column:posts_num;comment:'帖子数'"`
	CollectionsNum   int            `json:"collectionsNum" gorm:"column:collections_num;NOT NULL;comment:'收藏数'"`
	NotificationsNum int            `json:"notificationsNum" gorm:"column:notifications_num;comment:'系统系统通知数量'"`
	StewardId        int            `json:"stewardId" gorm:"column:steward_id;default:NULL"` //管家ID
	CreatedAt        time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
	CreateBy         int            `json:"createBy" gorm:"column:create_by"`
	UpdateBy         int            `json:"updateBy" gorm:"column:update_by"`
	WeChatNum        string         `json:"wechat_num" gorm:"column:wechat_num;comment:'微信号'"`
	// 定义与 Attention 表中 UserId 的一对多关系
	Attentions []Attention `gorm:"foreignkey:UserId"`
	// 定义与 Attention 表中 FanId 的一对多关系
	Fans []Attention `gorm:"foreignkey:FanId"`
}

func (u *User) TableName() string {
	return "user"
}

// Attention 关注表
type Attention struct {
	Id        int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId    int    `json:"userId" gorm:"column:user_id"`
	FanId     int    `json:"fanId" gorm:"column:fan_id"`
	CreateBy  int    `json:"createBy" gorm:"column:create_by"`
	UpdateBy  int    `json:"updateBy" gorm:"column:update_by"`
	CreatedAt string `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt string `json:"deletedAt" gorm:"column:deleted_at"`
}

func (a *Attention) TableName() string {
	return "attention"
}
