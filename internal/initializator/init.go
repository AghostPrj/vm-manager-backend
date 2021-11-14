/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 10:58
 * @Desc:
 */

package initializator

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/model/userModel"
	"github.com/AghostPrj/vm-manager-backend/internal/model/vmDiskModel"
	"github.com/AghostPrj/vm-manager-backend/internal/model/vmListModel"
	"github.com/AghostPrj/vm-manager-backend/internal/model/vmMacModel"
	"github.com/AghostPrj/vm-manager-backend/internal/model/vmPciModel"
	"github.com/AghostPrj/vm-manager-backend/internal/model/vmPortModel"
	"github.com/AghostPrj/vm-manager-backend/internal/service/userService"
	"github.com/ggg17226/aghost-go-base/pkg/utils/configUtils"
	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitApp() {
	configUtils.SetConfigFileName(constData.ApplicationName)
	bindAppConfigKey()
	bindAppConfigDefaultValue()
	configUtils.InitConfigAndLog()
	checkAppConfig()
	initDbClient()
}

func bindAppConfigKey() {
	configUtils.ConfigKeyList = append(configUtils.ConfigKeyList,
		[]string{constData.ConfGoProcNumKey, constData.EnvGoProcNumKey},
		[]string{constData.ConfHostNetTypeConfigKey, constData.EnvHostNetTypeConfigKey},
		[]string{constData.ConfServerListenPortKey, constData.EnvServerListenPortKey},
		[]string{constData.ConfServerListenHostKey, constData.EnvServerListenHostKey},
		[]string{constData.ConfDbNameKey, constData.EnvDbNameKey},
		[]string{constData.ConfDbHostKey, constData.EnvDbHostKey},
		[]string{constData.ConfDbPortKey, constData.EnvDbPortKey},
		[]string{constData.ConfDbCharsetKey, constData.EnvDbCharsetKey},
		[]string{constData.ConfDbCollationKey, constData.EnvDbCollationKey},
		[]string{constData.ConfDbLocKey, constData.EnvDbLocKey},
		[]string{constData.ConfDbTimeoutKey, constData.EnvDbTimeoutKey},
		[]string{constData.ConfDbUserKey, constData.EnvDbUserKey},
		[]string{constData.ConfDbPasswordKey, constData.EnvDbPasswordKey},
		[]string{constData.ConfDbAutoMigrateKey, constData.EnvDbAutoMigrateKey},
		[]string{constData.ConfDbDsnKey, constData.EnvDbDsnKey},
		[]string{constData.ConfDbMaxConnKey, constData.EnvDbMaxConnKey},
		[]string{constData.ConfDbMaxIdleKey, constData.EnvDbMaxIdleKey},
		[]string{constData.ConfDbConnLifeKey, constData.EnvDbConnLifeKey},
		[]string{constData.ConfDbConnMaxIdleKey, constData.EnvDbConnMaxIdleKey},
		[]string{constData.ConfDebugFlagKey, constData.EnvDebugFlagKey},
		[]string{constData.ConfAuthExpireTimeKey, constData.EnvAuthExpireTimeKey},
	)
}

func bindAppConfigDefaultValue() {
	viper.SetDefault(constData.ConfGoProcNumKey, constData.DefaultGoProcNum)
	viper.SetDefault(constData.ConfServerListenPortKey, constData.DefaultServerListenPort)
	viper.SetDefault(constData.ConfServerListenHostKey, constData.DefaultServerListenHost)
	viper.SetDefault(constData.ConfDbNameKey, constData.DefaultDbName)
	viper.SetDefault(constData.ConfDbHostKey, constData.DefaultDbHost)
	viper.SetDefault(constData.ConfDbPortKey, constData.DefaultDbPort)
	viper.SetDefault(constData.ConfDbCharsetKey, constData.DefaultDbCharset)
	viper.SetDefault(constData.ConfDbCollationKey, constData.DefaultDbCollation)
	viper.SetDefault(constData.ConfDbLocKey, constData.DefaultDbLoc)
	viper.SetDefault(constData.ConfDbTimeoutKey, constData.DefaultDbTimeout)
	viper.SetDefault(constData.ConfDbUserKey, constData.DefaultDbUser)
	viper.SetDefault(constData.ConfDbPasswordKey, constData.DefaultDbPassword)
	viper.SetDefault(constData.ConfDbAutoMigrateKey, constData.DefaultDbAutoMigrate)
	viper.SetDefault(constData.ConfDbDsnKey, constData.DefaultDbDsn)
	viper.SetDefault(constData.ConfDbMaxConnKey, constData.DefaultDbMaxConn)
	viper.SetDefault(constData.ConfDbMaxIdleKey, constData.DefaultDbMaxIdle)
	viper.SetDefault(constData.ConfDbConnLifeKey, constData.DefaultDbConnLife)
	viper.SetDefault(constData.ConfDbConnMaxIdleKey, constData.DefaultDbConnMaxIdle)
	viper.SetDefault(constData.ConfDebugFlagKey, constData.DefaultDebugFlag)
	viper.SetDefault(constData.ConfAuthExpireTimeKey, constData.DefaultAuthExpireTime)

}

func checkAppConfig() {
	checkHostNetType()
	checkHttpListenConfig()
}

func initDbClient() {
	dsn := viper.GetString(constData.ConfDbDsnKey)
	if len(dsn) <= 1 {
		dsn = viper.GetString(constData.ConfDbUserKey) + ":" + viper.GetString(constData.ConfDbPasswordKey) +
			"@(" + viper.GetString(constData.ConfDbHostKey) + ":" + viper.GetString(constData.ConfDbPortKey) + ")/" +
			viper.GetString(constData.ConfDbNameKey) + "?charset=" + viper.GetString(constData.ConfDbCharsetKey) +
			"&collation=" + viper.GetString(constData.ConfDbCollationKey) + "&loc=" + viper.GetString(constData.ConfDbLocKey) +
			"&readTimeout=" + viper.GetString(constData.ConfDbTimeoutKey) + "&writeTimeout=" +
			viper.GetString(constData.ConfDbTimeoutKey) + "&tls=false&parseTime=True"
	}
	log.WithFields(log.Fields{
		"op":            "startup",
		"dsn":           dsn,
		"max_conn":      viper.GetInt(constData.ConfDbMaxConnKey),
		"max_idle":      viper.GetInt(constData.ConfDbMaxIdleKey),
		"max_conn_life": viper.GetDuration(constData.ConfDbConnLifeKey),
		"max_conn_idle": viper.GetDuration(constData.ConfDbConnMaxIdleKey),
	}).Trace("db dsn")
	DBClient, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                 gorm_logrus.New(),
		PrepareStmt:            false,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"op":  "startup",
			"err": err,
		}).Panic("init db client error")
	}
	Db, err := DBClient.DB()
	if err != nil {
		log.WithFields(log.Fields{
			"op":  "startup",
			"err": err,
		}).Panic("get db client error")
	}

	Db.SetMaxOpenConns(viper.GetInt(constData.ConfDbMaxConnKey))
	Db.SetMaxIdleConns(viper.GetInt(constData.ConfDbMaxIdleKey))
	Db.SetConnMaxLifetime(viper.GetDuration(constData.ConfDbConnLifeKey))
	Db.SetConnMaxIdleTime(viper.GetDuration(constData.ConfDbConnMaxIdleKey))

	global.DBClient = DBClient

	if viper.GetBool(constData.ConfDbAutoMigrateKey) {
		global.DBClient.AutoMigrate(&vmDiskModel.VmDisk{})
		global.DBClient.AutoMigrate(&vmPciModel.VmPci{})
		global.DBClient.AutoMigrate(&vmMacModel.VmMac{})
		global.DBClient.AutoMigrate(&vmPortModel.VmPort{})
		global.DBClient.AutoMigrate(&vmListModel.VmList{})
		global.DBClient.AutoMigrate(&userModel.User{})
		if !userService.CheckDefaultUserExist() {
			user, password, err := userService.CreateUserWithoutPassword(constData.DefaultUserName)
			if err != nil {
				log.WithField("err", err).
					WithField("op", "init").
					Panic("create default admin user error")
			}
			log.WithFields(log.Fields{
				"user":     user.Name,
				"password": password,
				"otp":      user.Otp,
			}).Info("create default admin user")
		}
	}

	log.WithFields(log.Fields{
		"op": "startup",
	}).Trace("db inited")

}
