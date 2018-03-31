package common

import (
	ip2region "github.com/chanyipiaomiao/ip2region/binding/golang"
)

// QueryIP IP查询
type QueryIP struct {
	DBPath string
}

// NewQueryIP IP对象
func NewQueryIP(dbPath string) *QueryIP {
	return &QueryIP{DBPath: dbPath}
}

// Query 根据IP查询
func (q *QueryIP) Query(ip string) (*ip2region.IpInfo, error) {
	region, err := ip2region.New(q.DBPath)
	defer region.Close()
	if err != nil {
		return nil, err
	}

	r, err := region.MemorySearch(ip)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
