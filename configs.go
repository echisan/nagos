package nagos

type RegistryConfig struct {
	Host string
	Port int
	ContextPath string
}

type Config struct {
	// 租户信息，对应 Nacos 的命名空间字段
	// optional
	Tenant string
	// 	配置 ID
	DataId string
	// 配置分组
	Group string
	// 配置内容
	Content string
}

func NewServerConfig(dataId, group, content string) *Config {
	return &Config{
		Tenant:  "",
		DataId:  dataId,
		Group:   group,
		Content: content,
	}
}

func (s *Config) SetTenant(tenant string) {
	s.Tenant = tenant
}


