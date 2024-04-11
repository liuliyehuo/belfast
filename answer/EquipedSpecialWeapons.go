package answer

import (
	"github.com/bettercallmolly/belfast/connection"

	"github.com/bettercallmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

var validSC14001 protobuf.SC_14001

func EquipedSpecialWeapons(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_14001
	response.SpweaponBagSize = proto.Uint32(0)
	return client.SendMessage(14001, &response)
}

func init() {
	data := []byte{}
	panic("replayed packet: replace this with the actual data")
	proto.Unmarshal(data, &validSC14001)
}