package node

import "github.com/kryloffgregory/totoro/server/depends"

type Manager struct {
	dao *dao
}

func NewManager() *Manager {
	return &Manager{
		dao: NewDAO(),
	}
}

func (m *Manager) AddNode(libName string) error {
	return m.dao.UpsertNode(&Node{
		LibName:     libName,
		CriticalFor: []string{},
	})
}

func (m *Manager) AddAffected(libName string, uid string) error {
	node, err := m.dao.GetNode(libName)
	if err != nil {
		return err
	}

	node.CriticalFor = append(node.CriticalFor, uid)

	return m.dao.UpsertNode(node)
}

func (m *Manager) GetAffectedForNodeUpdate(libName string) ([]string, error) {
	node, err := m.dao.GetNode(libName)
	if err != nil {
		return nil, err
	}

	return node.CriticalFor, nil
}

func (m *Manager) GetAffectedForNodeDelete(libName string) ([]string, error) {
	node, err := m.dao.GetNode(libName)
	if err != nil {
		return nil, err
	}

	var result []string

	rdeps, err := depends.GetRDepends(libName)
	if err != nil {
		return nil, err
	}

	for _, rdep := range rdeps {
		node, err := m.dao.GetNode(rdep)
		if err != nil {
			return nil, err
		}
		if node == nil {
			continue
		}
		result = append(result, node.CriticalFor...)
	}

	return node.CriticalFor, nil
}
