package model

// 注册实例的请求参数
type RegisterInstance struct {
	Ip          string  `json:"ip"`
	Port        int64   `json:"port"`
	NamespaceId string  `json:"namespaceId"` // optional
	Weight      float64 `json:"weight"`      // optional
	Enable      bool    `json:"enable"`      // optional 是否上线
	Healthy     bool    `json:"healthy"`     // optional
	Metadata    string  `json:"metadata"`    // optional
	ClusterName string  `json:"clusterName"` // optional
	ServiceName string  `json:"serviceName"`
	GroupName   string  `json:"groupName"` // optional
	Ephemeral   bool    `json:"ephemeral"` // optional 是否临时实例
}


