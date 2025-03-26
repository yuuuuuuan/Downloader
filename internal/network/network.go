package network

type Network struct {
	URL string
	Req []byte
	Res []byte
}

func (n *Network) Post() error {
	return nil
}
