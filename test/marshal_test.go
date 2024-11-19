package test

import (
	"encoding/json"
	"fmt"
	"github.com/fire1220/marshaljson"
	"testing"
	"time"
)

type GoodInfo struct {
	Title    string    `json:"title" default:"ABC"`
	Like     string    `json:"like"`
	PlayTime time.Time `json:"play_time" datetime:"2006-01-02 15:04:05"`
}

func (t GoodInfo) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

type Good struct {
	ID          int32       `json:"id" default:"456"`
	ValFloat    float64     `json:"val_float" default:"111"`
	Val         uint        `json:"val" default:"-111"`
	ValBool     bool        `json:"val_bool" default:"true"`
	ValSlice    []bool      `json:"val_slice" default:"[]"`
	ValMap      map[int]int `json:"val_map" default:"{}"`
	ValStruct   GoodInfo    `json:"val_struct"`
	ValStruct2  GoodInfo    `json:"val_struct2" default:"{}"`
	Name        string      `json:"name" default:"123"`
	PlayTime    time.Time   `json:"play_time" datetime:"2006-01-02 15:04:05"`
	ExecuteTime time.Time   `json:"execute_time" datetime:"2006-01-02" default:"-"`
	CreatedAt   time.Time   `json:"created_at" datetime:"2006-01-02 15:04:05" default:"0000-00-00"`
	UpdatedAt   time.Time   `json:"updated_at" default:""`
}

func (t *Good) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(*t)
}

func TestMarshal(t *testing.T) {
	good := Good{ID: 0, Name: "", PlayTime: time.Now(), ExecuteTime: time.Now()}
	goodList := make([]Good, 0)
	goodList = append(goodList, good)
	bytes, err := json.Marshal(goodList)
	fmt.Printf("%s\n", bytes)
	fmt.Println(err)
}
