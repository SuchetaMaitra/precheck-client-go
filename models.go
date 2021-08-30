package precheck

// Order -
type Order struct {
	ID      int         `json:"id,omitempty"`
	fcpaths []OrderItem `json:"fcpaths,omitempty"`
}

// OrderItem -
type OrderItem struct {
	portpos  portpos `json:"portpos,omitempty"`
	wwn      string  `json:"wwn"`
	hostpeed int     `json:"hostspeed"`
}

// portpos -
type portpos struct {
	node     int `json:"node,omitempty"`
	slot     int `json:"slot,omitempty"`
	cardport int `json:"cardport,omitempty"`
}
