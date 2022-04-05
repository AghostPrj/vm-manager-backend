/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 10:57
 * @Desc:
 */

package constData

import (
	"time"
)

const (
	DefaultUserName    = "admin"
	ApplicationName    = "vm-manager"
	AuthCodeHeaderKey  = "x-aghost-auth-code"
	UserInfoContextKey = "context_user_info"

	ConfGoProcNumKey = "app.runtime.proc.num"
	EnvGoProcNumKey  = "app_runtime_proc_num"
	DefaultGoProcNum = 2

	ConfHostNetTypeConfigKey = "app.host.net.type"
	EnvHostNetTypeConfigKey  = "app_host_net_type"

	HostNetTypeTap  = "tap"
	HostNetTypeDpdk = "dpdk"

	ConfServerListenPortKey = "app.server.listen.port"
	EnvServerListenPortKey  = "app_server_listen_port"
	DefaultServerListenPort = 33289

	ConfServerListenHostKey = "app.server.listen.host"
	EnvServerListenHostKey  = "app_server_listen_host"
	DefaultServerListenHost = ""

	ConfDbNameKey = "app.db.name"
	EnvDbNameKey  = "app_db_name"
	DefaultDbName = ApplicationName

	ConfDbHostKey = "app.db.host"
	EnvDbHostKey  = "app_db_host"
	DefaultDbHost = "127.0.0.1"

	ConfDbPortKey = "app.db.port"
	EnvDbPortKey  = "app_db_port"
	DefaultDbPort = "3306"

	ConfDbCharsetKey = "app.db.charset"
	EnvDbCharsetKey  = "app_db_charset"
	DefaultDbCharset = "utf8mb4"

	ConfDbCollationKey = "app.db.collation"
	EnvDbCollationKey  = "app_db_collation"
	DefaultDbCollation = "utf8mb4_general_ci"

	ConfDbLocKey = "app.db.loc"
	EnvDbLocKey  = "app_db_loc"
	DefaultDbLoc = "Asia%2FShanghai"

	ConfDbTimeoutKey = "app.db.timeout"
	EnvDbTimeoutKey  = "app_db_timeout"
	DefaultDbTimeout = "120s"

	ConfDbUserKey = "app.db.user"
	EnvDbUserKey  = "app_db_user"
	DefaultDbUser = "root"

	ConfDbPasswordKey = "app.db.password"
	EnvDbPasswordKey  = "app_db_password"
	DefaultDbPassword = "123456"

	ConfDbAutoMigrateKey = "app.db.auto_migrate"
	EnvDbAutoMigrateKey  = "app_db_auto_migrate"
	DefaultDbAutoMigrate = true

	ConfDbDsnKey = "app.db.dsn"
	EnvDbDsnKey  = "app_db_dsn"
	DefaultDbDsn = ""

	ConfDbMaxConnKey = "app.db.max_conn"
	EnvDbMaxConnKey  = "app_db_max_conn"
	DefaultDbMaxConn = 25

	ConfDbMaxIdleKey = "app.db.max_idle"
	EnvDbMaxIdleKey  = "app_db_max_idle"
	DefaultDbMaxIdle = 5

	ConfDbConnLifeKey = "app.db.conn_life"
	EnvDbConnLifeKey  = "app_db_conn_life"
	DefaultDbConnLife = time.Hour

	ConfDbConnMaxIdleKey = "app.db.conn_idle"
	EnvDbConnMaxIdleKey  = "app_db_conn_idle"
	DefaultDbConnMaxIdle = time.Minute * 15

	ConfDebugFlagKey = "app.debug"
	EnvDebugFlagKey  = "app_debug"
	DefaultDebugFlag = false

	ConfAuthExpireTimeKey = "app.auth.expire"
	EnvAuthExpireTimeKey  = "app_auth_expire"
	DefaultAuthExpireTime = time.Minute * 15
)
