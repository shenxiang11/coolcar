package poi

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"google.golang.org/protobuf/proto"
	"hash/fnv"
)

var poi = []string{
	"摩天大楼",
	"购物中心",
	"会议中心",
	"娱乐场所",
	"饭店",
	"体育馆",
	"剧院",
	"学校",
	"博物馆",
	"纪念碑",
	"广场",
	"钟楼",
	"市政厅",
	"教堂",
	"寺庙",
	"清真寺",
	"雕像",
	"车站",
	"机场",
	"发电厂",
	"天线",
	"烟囱",
	"水坝",
	"水塔",
	"灯塔",
	"桥梁",
}

type Manager struct {
}

func (m *Manager) Resolve(c context.Context, loc *rentalpb.Location) (string, error) {
	b, err := proto.Marshal(loc)
	if err != nil {
		return "", err
	}

	h := fnv.New32()
	h.Write(b)

	return poi[int(h.Sum32())%len(poi)], nil
}
