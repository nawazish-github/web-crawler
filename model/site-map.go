package model

var urlReg = make(map[string][]string)

//GetURLReg returns the urlReg
func GetURLReg() map[string][]string {
	return urlReg
}

//addURLToURLRegistry maintains a registry of crawled URLs
func AddURLToURLRegistry(rawURL string) {
	if _, ok := urlReg[rawURL]; !ok {
		urlReg[rawURL] = []string{}
	}
}
