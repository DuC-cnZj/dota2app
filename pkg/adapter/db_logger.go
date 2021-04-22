package adapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DuC-cnZj/dota2app/pkg/dlog"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLoggerAdapter struct {
	level logger.LogLevel
}

func (g *GormLoggerAdapter) LogMode(level logger.LogLevel) logger.Interface {
	g.level = level

	return g
}

func (g *GormLoggerAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Info {
		dlog.Infof(s, i...)
	}
}

func (g *GormLoggerAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Warn {
		dlog.Warningf(s, i...)
	}
}

func (g *GormLoggerAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	if g.level >= logger.Error {
		dlog.Errorf(s, i...)
	}
}

func (g *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	const (
		traceStr      = "%s [%.3fms] [rows:%v] %s"
		traceWarnStr  = "%s %s [%.3fms] [rows:%v] %s"
		traceErrStr   = "%s %s [%.3fms] [rows:%v] %s"
		slowThreshold = 200 * time.Millisecond
	)
	if g.level > logger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && g.level >= logger.Error:
			if errors.Is(err, gorm.ErrRecordNotFound) {
				sql, rows := fc()
				if rows == -1 {
					dlog.Debugf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
				} else {
					dlog.Debugf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
				}
				return
			}
			sql, rows := fc()
			if rows == -1 {
				dlog.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				dlog.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > slowThreshold && g.level >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
			if rows == -1 {
				dlog.Warningf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				dlog.Warningf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case g.level == logger.Info:
			sql, rows := fc()
			if rows == -1 {
				dlog.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				dlog.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
