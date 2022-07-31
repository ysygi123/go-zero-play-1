package symysql

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

var db *gorm.DB

func InitSyMysql(master string, slave []string) (err error) {
	db, err = gorm.Open(mysql.Open(master), &gorm.Config{
		Logger: &syMysqlLogger{},
	})
	if err != nil {
		return
	}

	if len(slave) != 0 {
		rp := make([]gorm.Dialector, 0, len(slave))
		for _, v := range slave {
			rp = append(rp, mysql.Open(v))
		}

		err = db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: rp,
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			return
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return
}

func GetDbSession(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

type syMysqlLogger struct {
	//LogLevel logger.LogLevel
}

func (s *syMysqlLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *s
	//newlogger.LogLevel = level
	return &newlogger
}
func (s *syMysqlLogger) Info(ctx context.Context, msg string, params ...interface{}) {
	logx.WithContext(ctx).Infow(msg)
}
func (s *syMysqlLogger) Warn(ctx context.Context, msg string, params ...interface{}) {
	logx.WithContext(ctx).Infow(msg)
}
func (s *syMysqlLogger) Error(ctx context.Context, msg string, params ...interface{}) {
	logx.WithContext(ctx).Errorw(msg)
}
func (s *syMysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

}
