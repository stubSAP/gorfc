package gorfc

type ConnectionParameters map[string]string

func ConnectionFromParams(connectionParams ConnectionParameters) (conn *Connection, err error) {
	return &Connection{}, nil
}

type Connection struct {
}

func (conn *Connection) Close() (err error) {
	return nil
}

func (conn *Connection) Call(goFuncName string, params interface{}) (result map[string]interface{}, err error) {
	return nil, nil
}
