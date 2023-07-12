package transport

import "kit.golaxy.org/plugins/transport/binaryutil"

// MsgPayload 数据传输
type MsgPayload struct {
	Data []byte // 数据
}

func (m *MsgPayload) Read(p []byte) (int, error) {
	bs := binaryutil.NewByteStream(p)
	if err := bs.WriteBytes(m.Data); err != nil {
		return 0, err
	}
	return bs.BytesWritten(), nil
}

func (m *MsgPayload) Write(p []byte) (int, error) {
	bs := binaryutil.NewByteStream(p)
	data, err := bs.ReadBytesRef()
	if err != nil {
		return 0, err
	}
	m.Data = data
	return bs.BytesRead(), nil
}

func (m *MsgPayload) Size() int {
	return binaryutil.SizeofBytes(m.Data)
}

func (MsgPayload) MsgId() MsgId {
	return MsgId_Payload
}
