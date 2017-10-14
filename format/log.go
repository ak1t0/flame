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

	l = CountOnionLogs(l)

	return l
}

func CountOnionLogs(logs OnionLogs) OnionLogs {
	var l OnionLogs
	counter := make(map[string]int)

	for _, v := range logs {
		if val, ok := counter[v.Address]; ok {
			counter[v.Address] = val + 1
		} else {
			counter[v.Address] = 1
		}
	}

	keys := make([]string, 0, len(counter))

	for k := range counter {
		keys = append(keys, k)
	}

	for _, k := range keys {
		l = append(l, OnionLog{Address: k, Access: counter[k]})
	}

	return l
}
