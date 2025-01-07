package common

// 配置
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" yaml:"system" json:"system"`
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
