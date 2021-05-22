package node

type Node struct {
	LibName string `json:"libName"`
	CriticalFor []string `json:"criticalFor"`

}