/*
 * @date: 2021/12/15
 * @desc: ...
 */

package config

// MysqlConfig mysql 配置
type MysqlConfig struct {
	DBName   string `mapstructure:"dbname" json:"dbname"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host                string `mapstructure:"host" json:"host"`
	Port                int    `mapstructure:"port" json:"port"`
	DB                  int    `mapstructure:"db" json:"db"`
	Username            string `mapstructure:"username" json:"username"`
	Password            string `mapstructure:"password" json:"password"`
	ConnectTimeout      int    `mapstructure:"connectTimeout" json:"connectTimeout"`
	PoolMaxIdleConns    int    `mapstructure:"poolMaxIdleConns" json:"poolMaxIdleConns"`
	PoolMaxOpenConns    int    `mapstructure:"poolMaxOpenConns" json:"poolMaxOpenConns"`
	PoolConnMaxLifetime int    `mapstructure:"poolConnMaxLifetime" json:"poolConnMaxLifetime"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
}

// OrmDatabasePoolConfig 配置
type OrmDatabasePoolConfig struct {
	Status          string `mapstructure:"status" json:"status"` // enable 开启数据库连接池 disable 不启用数据库连接池
	MaxIdleConns    int    `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" json:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogLevel     string `mapstructure:"logLevel" json:"logLevel"`
	LogPath      string `mapstructure:"logPath" json:"logPath"`
	LogInConsole bool   `mapstructure:"logInConsole" json:"logInConsole"`
	MaxSize      int    `mapstructure:"maxSize" json:"maxSize"`
	MaxBackups   int    `mapstructure:"maxBackups" json:"maxBackups"`
	MaxAge       int    `mapstructure:"maxAge" json:"maxAge"`
	Compress     bool   `mapstructure:"compress" json:"compress"`
}

// JWTConfig 配置信息
type JWTConfig struct {
	SigningKey    string `mapstructure:"signingKey" json:"signingKey"`
	TokenKey      string `mapstructure:"tokenKey" json:"tokenKey"`
	EffectiveTime int    `mapstructure:"effectiveTime" json:"effectiveTime"`
}

// LDAPConfig ldap 设置
type LDAPConfig struct {
	LdapHost       string `mapstructure:"ldapHost" json:"ldapHost"`
	LdapPort       int    `mapstructure:"ldapPort" json:"ldapPort"`
	BaseDN         string `mapstructure:"baseDN" json:"baseDN"`
	SearchProperty string `mapstructure:"searchProperty" json:"searchProperty"`
	BindDN         string `mapstructure:"bindDN" json:"bindDN"`
	BindPassword   string `mapstructure:"bindPassword" json:"bindPassword"`
}

// AuthConfig Auth配置
type AuthConfig struct {
	JWTInfo  JWTConfig  `mapstructure:"jwt" json:"jwt"`
	LADPInfo LDAPConfig `mapstructure:"ldap" json:"ldap"`
}

// OrmConfig 配置
type OrmConfig struct {
	TablePrefix string `mapstructure:"tablePrefix" json:"tablePrefix"`
}

//
// SmConfig
// @Description: 网易云信
//
type SmConfig struct {
	SendSmBaseUrl   string `mapstructure:"sendSmBaseUrl" json:"sendSmBaseUrl"`
	AppSecret       string `mapstructure:"appSecret" json:"appSecret"`
	AppKey          string `mapstructure:"appKey" json:"appKey"`
	SMTemplateCode  int    `mapstructure:"SMTemplateCode" json:"SMTemplateCode"`
	CodeLen         int    `mapstructure:"codeLen" json:"codeLen"`
	VerifySmBaseUrl string `mapstructure:"verifySmBaseUrl" json:"verifySmBaseUrl"`
}

// ServerConfig 全局配置
type ServerConfig struct {
	Host                string                `mapstructure:"host" json:"host"`
	Port                int                   `mapstructure:"port" json:"port"`
	LogInfo             LogConfig             `mapstructure:"log" json:"log"`
	SmInfo              SmConfig              `mapstructure:"sm" json:"sm"`
	OrmInfo             OrmConfig             `mapstructure:"orm" json:"orm"`
	AuthInfo            AuthConfig            `mapstructure:"auth" json:"auth"`
	DatabaseInfo        DatabaseConfig        `mapstructure:"database" json:"database"`
	OrmDatabasePoolInfo OrmDatabasePoolConfig `mapstructure:"ormDatabasePool" json:"ormDatabasePool"`
}
