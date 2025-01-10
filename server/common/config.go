package common

// 配置解析入口
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" yaml:"log" json:"log"`
	MySQL  MySQLConfiguration  `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
	Redis  RedisConfiguration  `mapstructure:"redis" yaml:"redis" json:"redis"`
	Auth   AuthConfiguration   `mapstructure:"auth" yaml:"auth" json:"auth"`
}

// 系统配置
type SystemConfiguration struct {
	Listen ListenConfiguration `mapstructure:"listen" yaml:"listen" json:"listen"`
	Role   RoleConfiguration   `mapstructure:"role" yaml:"role" json:"role"`
}

// 监听配置
type ListenConfiguration struct {
	Host string `mapstructure:"host" yaml:"host" json:"host"`
	Port string `mapstructure:"port" yaml:"port" json:"port"`
}

// 角色配置
type RoleConfiguration struct {
	WebServer      bool `mapstructure:"web-server" yaml:"web-server" json:"webServer"`
	LeaderElection bool `mapstructure:"leader-election" yaml:"leader-election" json:"leaderElection"`
	Worker         bool `mapstructure:"worker" yaml:"worker" json:"worker"`
}

// 日志配置
type LogConfiguration struct {
	System LoggerConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Access LoggerConfiguration `mapstructure:"access" yaml:"access" json:"access"`
}

// 日志基础配置
type LoggerConfiguration struct {
	Enabled    bool                       `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	Level      string                     `mapstructure:"level" yaml:"level" json:"level"`
	Format     string                     `mapstructure:"format" yaml:"format" json:"format"`
	PathPrefix string                     `mapstructure:"path-prefix" yaml:"path-prefix" json:"pathPrefix"`
	Rolling    LoggerRollingConfiguration `mapstructure:"rolling" yaml:"rolling" json:"rolling"`
}

// 日志切割配置
type LoggerRollingConfiguration struct {
	Enabled    bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	MaxSize    int  `mapstructure:"max-size" yaml:"max-size" json:"maxSize"`
	MaxBackups int  `mapstructure:"max-backups" yaml:"max-backups" json:"maxBackups"`
	MaxAge     int  `mapstructure:"max-age" yaml:"max-age" json:"maxAge"`
	Compress   bool `mapstructure:"compress" yaml:"compress" json:"compress"`
}

// MySQL 配置
type MySQLConfiguration struct {
	Host         string `mapstructure:"host" yaml:"host" json:"host"`
	Port         int    `mapstructure:"port" yaml:"port" json:"port"`
	Database     string `mapstructure:"database" yaml:"database" json:"database"`
	Username     string `mapstructure:"username" yaml:"username" json:"username"`
	Password     string `mapstructure:"password" yaml:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
	Params       string `mapstructure:"params" yaml:"params" json:"params"`
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns" json:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns" json:"maxIdleConns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" yaml:"max-idle-time" json:"maxIdleTime"`
}

// Redis 配置
type RedisConfiguration struct {
	Host         string `mapstructure:"host" yaml:"host" json:"host"`
	Port         int    `mapstructure:"port" yaml:"port" json:"port"`
	Database     int    `mapstructure:"database" yaml:"database" json:"database"`
	Password     string `mapstructure:"password" yaml:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns" json:"maxOpenConns"`
	MinIdleConns int    `mapstructure:"min-idle-conns" yaml:"min-idle-conns" json:"minIdleConns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns" json:"maxIdleConns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" yaml:"max-idle-time" json:"maxIdleTime"`
}

// 用户登录相关配置
type AuthConfiguration struct {
	JWT   JWTConfiguration   `mapstructure:"jwt" yaml:"jwt" json:"jwt"`
	Login LoginConfiguration `mapstructure:"login" yaml:"login" json:"login"`
}

// JWT 配置
type JWTConfiguration struct {
	Realm   string `mapstructure:"realm" yaml:"realm" json:"realm"`
	Key     string `mapstructure:"key" yaml:"key" json:"key"`
	Timeout int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
}

// 登录配置
type LoginConfiguration struct {
	ErrorLimit    int  `mapstructure:"error-limit" yaml:"error-limit" json:"errorLimit"`
	ErrorLockTime int  `mapstructure:"error-lock-time" yaml:"error-lock-time" json:"errorLockTime"`
	MultiDevice   bool `mapstructure:"multi-device" yaml:"multi-device" json:"multiDevice"`
}
