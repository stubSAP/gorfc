package gorfc

type ConnectionParameters map[string]string

type Connection struct {
}

func ConnectionFromParams(connectionParams ConnectionParameters) (conn *Connection, err error) {
	return &Connection{}, nil
}
