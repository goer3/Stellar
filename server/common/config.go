package common

// 配置
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" yaml:"log" json:"log"`
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
