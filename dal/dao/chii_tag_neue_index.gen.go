// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameTagIndex = "chii_tag_neue_index"

// TagIndex mapped from table <chii_tag_neue_index>
type TagIndex struct {
	ID          uint32 `gorm:"column:tag_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:""`
	Name        string `gorm:"column:tag_name;type:varchar(30);not null" json:""`
	Cat         int8   `gorm:"column:tag_cat;type:tinyint(3);not null;comment:0=条目 1=日志 2=天窗" json:""` // 0=条目 1=日志 2=天窗
	Type        uint8  `gorm:"column:tag_type;type:tinyint(3);not null" json:""`
	Results     uint32 `gorm:"column:tag_results;type:mediumint(8) unsigned;not null" json:""`
	CreatedTime uint32 `gorm:"column:tag_dateline;type:int(10) unsigned;not null" json:""`
	UpdatedTime uint32 `gorm:"column:tag_lasttouch;type:int(10) unsigned;not null" json:""`
}

// TableName TagIndex's table name
func (*TagIndex) TableName() string {
	return TableNameTagIndex
}
