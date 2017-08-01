package format

// OnionLog is data format for the whole
type OnionLog struct {
	Address string
	Access  int
	Result  string
}

// OnionLog slice
type OnionLogs []OnionLog

// OnionLogs constructor
func NewOnionLogs(log []string) OnionLogs {
	var l OnionLogs
	for _, v := range log {
		l = append(l, OnionLog{Address: v})
	}

	return l
}
