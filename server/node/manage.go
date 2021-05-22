package node


func AddNode(libName string) error{
	return UpsertNode(&Node{
		LibName:     libName,
		CriticalFor: []string{},
	})
}

func AddAffected(libName string, uid string) error {
	node, err:=GetNode(libName)
	if err!=nil {
		return err
	}

	node.CriticalFor = append(node.CriticalFor, uid)

	return UpsertNode(node)
}

func GetAffectedForNodeUpdate(libName string) ([]string,error) {
	node, err:=GetNode(libName)
	if err!=nil {
		return nil, err
	}

	return node.CriticalFor, nil
}

func GetAffectedForNodeDelete(libName string) ([]string, error) {

}

