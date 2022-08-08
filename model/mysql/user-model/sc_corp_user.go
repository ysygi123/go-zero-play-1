package user_model

import (
	"context"
	"go-zero-play-1/common/symysql"
	"time"
)

var ScCorpUserModel *ScCorpUser

// ScCorpUser 企业用户表
type ScCorpUser struct {
	Id                  uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	BindUserId          uint      `gorm:"column:bind_user_id;type:int(11) unsigned;default:0;comment:绑定蝉客user_id" json:"bind_user_id"`
	Cid                 uint      `gorm:"column:cid;type:int(11) unsigned;comment:系统内部公司ID" json:"cid"`
	CorpId              string    `gorm:"column:corp_id;type:varchar(64);comment:授权方企业id" json:"corp_id"`
	UserId              string    `gorm:"column:user_id;type:varchar(64);comment:企业微信用户的userid" json:"user_id"`
	SuitUserId          string    `gorm:"column:suit_user_id;type:varchar(64);comment:企业微信用户的userid(第三方)" json:"suit_user_id"`
	OpenUserid          string    `gorm:"column:open_userid;type:varchar(64);comment:全局唯一" json:"open_userid"`
	Avatar              string    `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	UserType            uint      `gorm:"column:user_type;type:tinyint(4) unsigned;comment:登录用户的类型：1.创建者 2.内部系统管理员 3.外部系统管理员 4.分级管理员 5.成员" json:"user_type"`
	CreateTime          int64     `gorm:"column:create_time;type:bigint(11)" json:"create_time"`
	UpdateTime          int64     `gorm:"column:update_time;type:bigint(11)" json:"update_time"`
	RoleId              int       `gorm:"column:role_id;type:smallint(6);default:4;comment:4" json:"role_id"`
	BackUserId          string    `gorm:"column:back_user_id;type:varchar(64);comment:员工移除后暂存的原userID" json:"back_user_id"`
	Name                string    `gorm:"column:name;type:varchar(64);comment:员工名称" json:"name"`
	Mobile              string    `gorm:"column:mobile;type:char(11);comment:手机号;NOT NULL" json:"mobile"`
	Department          string    `gorm:"column:department;type:json;comment:部门ID" json:"department"`
	Position            string    `gorm:"column:position;type:varchar(64);comment:职位" json:"position"`
	Gender              int       `gorm:"column:gender;type:tinyint(1);default:0;comment:性别，0表示未定义，1表示男性，2表示女性;NOT NULL" json:"gender"`
	Email               string    `gorm:"column:email;type:varchar(64);comment:email" json:"email"`
	IsLeaderInDept      string    `gorm:"column:is_leader_in_dept;type:json;comment:对应DEPARTMENT，是否领导标识" json:"is_leader_in_dept"`
	ThumbAvatar         string    `gorm:"column:thumb_avatar;type:varchar(256);comment:缩略头像" json:"thumb_avatar"`
	Telephone           string    `gorm:"column:telephone;type:varchar(16);comment:电话" json:"telephone"`
	Alias               string    `gorm:"column:alias;type:varchar(64);comment:别名" json:"alias"`
	Status              int       `gorm:"column:status;type:tinyint(1);default:0;comment:激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。已激活代表已激活企业微信或已关注微工作台（原企业号）。未激活代表既未激活企业微信又未关注微工作台（原企业号）。;NOT NULL" json:"status"`
	Address             string    `gorm:"column:address;type:varchar(128);comment:地址" json:"address"`
	HideMobile          int       `gorm:"column:hide_mobile;type:tinyint(1);default:0;comment:隐藏手机号;NOT NULL" json:"hide_mobile"`
	EnglishName         string    `gorm:"column:english_name;type:varchar(16);comment:英文名" json:"english_name"`
	MainDepartment      int       `gorm:"column:main_department;type:int(11);default:0;comment:主部门;NOT NULL" json:"main_department"`
	QrCode              string    `gorm:"column:qr_code;type:varchar(256);comment:员工个人二维码" json:"qr_code"`
	Order               string    `gorm:"column:order;type:json;comment:在部门中的排序" json:"order"`
	ExternalPosition    string    `gorm:"column:external_position;type:varchar(32);comment:对外职务" json:"external_position"`
	Extattr             string    `gorm:"column:extattr;type:text;comment:扩展属性" json:"extattr"`
	ExternalProfile     string    `gorm:"column:external_profile;type:text;comment:成员对外属性" json:"external_profile"`
	DelStatus           int       `gorm:"column:del_status;type:tinyint(4);default:0;comment:删除状态 0否 1是" json:"del_status"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime;comment:首次同步时间" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
	DeletedAt           time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	IsArchives          int       `gorm:"column:is_archives;type:tinyint(4);default:2;comment:会话是否开通 1开 2关;NOT NULL" json:"is_archives"`
	IsOpen              int       `gorm:"column:is_open;type:tinyint(4);default:2;comment:是否开启 1：开启 2:关闭;NOT NULL" json:"is_open"`
	SysUserId           int64     `gorm:"column:sys_user_id;type:bigint(20);comment:后台系统账号id" json:"sys_user_id"`
	CommutingStatus     int       `gorm:"column:commuting_status;type:tinyint(1);default:0;comment:上下班状态，默认0，1上班，2下班;NOT NULL" json:"commuting_status"`
	AuthorizationStatus int       `gorm:"column:authorization_status;type:tinyint(3);default:1;comment:企业微信官网 授权状态 1授权 2未授权;NOT NULL" json:"authorization_status"`
	IsLinkMission       int       `gorm:"column:is_link_mission;type:tinyint(3);default:0;comment:0没关联 1关联 只有在authorization_status为2的时候才会生效;NOT NULL" json:"is_link_mission"`
}

func (s *ScCorpUser) TableName() string {
	return "sc_corp_user"
}

func (s *ScCorpUser) GetUser(ctx context.Context, id int) (data *ScCorpUser, err error) {
	data = &ScCorpUser{}
	err = symysql.GetDbSession(ctx).Where("id=?", id).First(data).Error
	return
}
