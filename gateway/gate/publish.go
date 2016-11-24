package gate

import (
	"errors"

	"github.com/aiyun/gomqtt/global"
	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
	"github.com/aiyun/gomqtt/uuid"
	"github.com/corego/tools"

	"bytes"

	"fmt"

	rpc "github.com/aiyun/gomqtt/proto"
)

//@ToDo
//从客户端过来的只能是私聊或者单播
func publish(ci *connInfo, p *proto.PublishPacket) error {
	switch ci.payloadProtoType {
	case global.PayloadText:
		// text格式，需要生成MsgID
		mid := tools.String2Bytes(uuid.GenStr())
		tps := bytes.Split(p.Topic(), topicSep)
		if len(tps) != 2 {
			return errors.New("invalid publish topic, text topic need to be topic--acc")
		}

		qos := qosTrans(p.QoS())

		err := ci.rpc.pubText(&rpc.PubTextMsg{
			Cid:   ci.id,
			ToAcc: tps[1],
			Ttp:   tps[0],
			Qos:   int32(qos),
			Mid:   mid,
			Msg:   p.Payload(),
		})
		if err != nil {
			return fmt.Errorf("pubText rpc error: %v", err)
		}

	case global.PayloadProtobuf:

	}

	// need give back the ack
	if p.QoS() >= 1 {
		pb := proto.NewPubackPacket()
		pb.SetPacketID(p.PacketID())
		service.WritePacket(ci.c, pb)
	}

	return nil
}

func puback(ci *connInfo, p *proto.PubackPacket) error {
	return nil
}
