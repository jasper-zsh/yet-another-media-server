package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"strings"
	"yet-another-media-server/media_library/media_library"
)

type File struct {
	gorm.Model
	LibraryID uint `gorm:"uniqueIndex:idx_file"`
	Type      string
	MediaID   uint
	Media     Media
	ParentID  uint
	Parent    *File
	Children  []File `gorm:"foreignkey:ParentID"`
	FilePath  string `gorm:"uniqueIndex:idx_file"`
	Ext       JSONStringMap
}

func (f *File) ToProto() *media_library.File {
	result := &media_library.File{
		Id:        int32(f.ID),
		CreatedAt: f.CreatedAt.Unix(),
		UpdatedAt: f.UpdatedAt.Unix(),
		Type:      f.Type,
		FilePath:  f.FilePath,
		Children:  make([]*media_library.File, 0),
		Ext:       f.Ext,
	}
	for _, t := range f.Children {
		result.Children = append(result.Children, t.ToProto())
	}
	return result
}

type JSONStringMap map[string]string

func (m JSONStringMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	ba, err := m.MarshalJSON()
	return string(ba), err
}

func (m *JSONStringMap) Scan(val interface{}) error {
	if val == nil {
		*m = make(JSONStringMap)
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := map[string]string{}
	err := json.Unmarshal(ba, &t)
	*m = t
	return err
}

func (m JSONStringMap) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	t := (map[string]string)(m)
	return json.Marshal(t)
}

func (m *JSONStringMap) UnmarshalJSON(b []byte) error {
	t := map[string]string{}
	err := json.Unmarshal(b, &t)
	*m = JSONStringMap(t)
	return err
}

func (m JSONStringMap) GormDataType() string {
	return "jsonstringmap"
}

func (JSONStringMap) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	}
	return ""
}

func (jm JSONStringMap) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	data, _ := jm.MarshalJSON()
	switch db.Dialector.Name() {
	case "mysql":
		if v, ok := db.Dialector.(*mysql.Dialector); ok && !strings.Contains(v.ServerVersion, "MariaDB") {
			return gorm.Expr("CAST(? AS JSON)", string(data))
		}
	}
	return gorm.Expr("?", string(data))
}
