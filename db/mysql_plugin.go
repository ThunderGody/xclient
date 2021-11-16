package db

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

const (
	callbackBeforeName = "core:before"
	callBackAfterName  = "core:after"
	startTime          = "_start_time"
)

type TracePlugin struct {
}

func (tp *TracePlugin) Name() string {
	return "tracePlugin"
}

func (tp *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// before
	_ = db.Callback().Create().Before("gorm:before_create").Register(callbackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callbackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callbackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:before_update").Register(callbackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callbackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callbackBeforeName, before)

	// after
	_ = db.Callback().Create().After("gorm:after_create").Register(callbackBeforeName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callbackBeforeName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callbackBeforeName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callbackBeforeName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callbackBeforeName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callbackBeforeName, after)

	return
}

var _ gorm.Plugin = &TracePlugin{}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
}

func after(db *gorm.DB) {
	_ts, isExists := db.InstanceGet(startTime)
	if !isExists {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	fmt.Println(fmt.Sprintf("sql: %s, cost time: %fs",
		db.Statement.SQL.String()),
		time.Since(ts).Seconds())
}
