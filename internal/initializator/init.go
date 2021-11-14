/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 10:58
 * @Desc:
 */

package initializator

import (
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/ggg17226/aghost-go-base/pkg/utils/configUtils"
	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitApp() {
	configUtils.SetConfigFileName(ApplicationName)
	bindAppConfigKey()
	bindAppConfigDefaultValue()
	configUtils.InitConfigAndLog()
	checkAppConfig()
	initDbClient()
}

func bindAppConfigKey() {
	configUtils.ConfigKeyList = append(configUtils.ConfigKeyList,
		[]string{ConfGoProcNumKey, EnvGoProcNumKey},
		[]string{ConfHostNetTypeConfigKey, EnvHostNetTypeConfigKey},
		[]string{ConfServerListenPortKey, EnvServerListenPortKey},
		[]string{ConfServerListenHostKey, EnvServerListenHostKey},
		[]string{ConfDbNameKey, EnvDbNameKey},
		[]string{ConfDbHostKey, EnvDbHostKey},
		[]string{ConfDbPortKey, EnvDbPortKey},
		[]string{ConfDbCharsetKey, EnvDbCharsetKey},
		[]string{ConfDbCollationKey, EnvDbCollationKey},
		[]string{ConfDbLocKey, EnvDbLocKey},
		[]string{ConfDbTimeoutKey, EnvDbTimeoutKey},
		[]string{ConfDbUserKey, EnvDbUserKey},
		[]string{ConfDbPasswordKey, EnvDbPasswordKey},
		[]string{ConfDbAutoMigrateKey, EnvDbAutoMigrateKey},
		[]string{ConfDbDsnKey, EnvDbDsnKey},
		[]string{ConfDbMaxConnKey, EnvDbMaxConnKey},
		[]string{ConfDbMaxIdleKey, EnvDbMaxIdleKey},
		[]string{ConfDbConnLifeKey, EnvDbConnLifeKey},
		[]string{ConfDbConnMaxIdleKey, EnvDbConnMaxIdleKey},
	)
}

func bindAppConfigDefaultValue() {
	viper.SetDefault(ConfGoProcNumKey, DefaultGoProcNum)
	viper.SetDefault(ConfServerListenPortKey, DefaultServerListenPort)
	viper.SetDefault(ConfServerListenHostKey, DefaultServerListenHost)
	viper.SetDefault(ConfDbNameKey, DefaultDbName)
	viper.SetDefault(ConfDbHostKey, DefaultDbHost)
	viper.SetDefault(ConfDbPortKey, DefaultDbPort)
	viper.SetDefault(ConfDbCharsetKey, DefaultDbCharset)
	viper.SetDefault(ConfDbCollationKey, DefaultDbCollation)
	viper.SetDefault(ConfDbLocKey, DefaultDbLoc)
	viper.SetDefault(ConfDbTimeoutKey, DefaultDbTimeout)
	viper.SetDefault(ConfDbUserKey, DefaultDbUser)
	viper.SetDefault(ConfDbPasswordKey, DefaultDbPassword)
	viper.SetDefault(ConfDbAutoMigrateKey, DefaultDbAutoMigrate)
	viper.SetDefault(ConfDbDsnKey, DefaultDbDsn)
	viper.SetDefault(ConfDbMaxConnKey, DefaultDbMaxConn)
	viper.SetDefault(ConfDbMaxIdleKey, DefaultDbMaxIdle)
	viper.SetDefault(ConfDbConnLifeKey, DefaultDbConnLife)
	viper.SetDefault(ConfDbConnMaxIdleKey, DefaultDbConnMaxIdle)

}

func checkAppConfig() {
	checkHostNetType()
	checkHttpListenConfig()
}

func initDbClient() {
	dsn := viper.GetString(ConfDbDsnKey)
	if len(dsn) <= 1 {
		dsn = viper.GetString(ConfDbUserKey) + ":" + viper.GetString(ConfDbPasswordKey) +
			"@(" + viper.GetString(ConfDbHostKey) + ":" + viper.GetString(ConfDbPortKey) + ")/" +
			viper.GetString(ConfDbNameKey) + "?charset=" + viper.GetString(ConfDbCharsetKey) +
			"&collation=" + viper.GetString(ConfDbCollationKey) + "&loc=" + viper.GetString(ConfDbLocKey) +
			"&readTimeout=" + viper.GetString(ConfDbTimeoutKey) + "&writeTimeout=" +
			viper.GetString(ConfDbTimeoutKey) + "&tls=false&parseTime=True"
	}
	log.WithFields(log.Fields{
		"op":            "startup",
		"dsn":           dsn,
		"max_conn":      viper.GetInt(ConfDbMaxConnKey),
		"max_idle":      viper.GetInt(ConfDbMaxIdleKey),
		"max_conn_life": viper.GetDuration(ConfDbConnLifeKey),
		"max_conn_idle": viper.GetDuration(ConfDbConnMaxIdleKey),
	}).Trace("db dsn")
	DBClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	Db.SetMaxOpenConns(viper.GetInt(ConfDbMaxConnKey))
	Db.SetMaxIdleConns(viper.GetInt(ConfDbMaxIdleKey))
	Db.SetConnMaxLifetime(viper.GetDuration(ConfDbConnLifeKey))
	Db.SetConnMaxIdleTime(viper.GetDuration(ConfDbConnMaxIdleKey))

	global.DBClient = DBClient

	if viper.GetBool(ConfDbAutoMigrateKey) {
		// global.DBClient.AutoMigrate(&QuestionModel.Question{})
	}

	log.WithFields(log.Fields{
		"op": "startup",
	}).Trace("db inited")

}
