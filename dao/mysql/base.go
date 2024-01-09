package mysql

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Base struct {
	ID         int  `gorm:"column:id;" json:"id"`
	GmtCreated MyTime `gorm:"column:gmtCreated;DEFAULT:DEFAULT;" json:"gmtCreated"`
	GmtUpdated MyTime `gorm:"column:gmtUpdated;DEFAULT:DEFAULT;" json:"gmtUpdated"`
}

type MyTime struct {
	time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
