package nagos

import (
	"fmt"
	"strings"
)

type RegistryConfig struct {
	Host        string
	Port        int
	ContextPath string
}

// 根据注册中心配置提供的一个快捷获取url的方法
func (rc *RegistryConfig) Url() string {
	cp := rc.ContextPath
	if !strings.HasPrefix(rc.ContextPath, "/") {
		cp = "/" + rc.ContextPath
	}
	if strings.HasSuffix(cp, "/") {
		cp = string(cp[:len(cp)-1])
	}
	return fmt.Sprintf("http://%s:%d%s", rc.Host, rc.Port, cp)
}

func (rc *RegistryConfig) UrlWithPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return rc.Url() + path
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
