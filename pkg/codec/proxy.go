// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"github.com/fagongzi/goetty"
	"github.com/fagongzi/util/protoc"
)

const (
	// RedisBegin tag for redis command
	RedisBegin = 0x01
	// WatcherNotifyBegin tag for notify
	WatcherNotifyBegin = 0x02
	// WatcherNotifySyncBegin tag for notify sync
	WatcherNotifySyncBegin = 0x03
	// WatcherNotifyRspBegin tag for notify rsp
	WatcherNotifyRspBegin = 0x04
)

// ProxyDecoder proxy decoder base on goetty
type ProxyDecoder struct {
}

// ProxyEncoder proxy encoder base on goetty
type ProxyEncoder struct {
}

// Decode return a decoded msg or wait for next tcp packet
func (decoder *ProxyDecoder) Decode(in *goetty.ByteBuf) (bool, interface{}, error) {
	_ = "STUB: not implemented"

	// remember the begin read index,
	// if we found has no enough data, we will resume this read index,
	// and waiting for next.
	return false, nil, nil
}

// Encode encode proxy message
func (encoder *ProxyEncoder) Encode(data interface{}, out *goetty.ByteBuf) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteProxyMessage write a proxy to the buf
func WriteProxyMessage(tag byte, req protoc.PB, out *goetty.ByteBuf) error {
	_ = "STUB: not implemented"
	return nil
}

func hasEnoughData(in *goetty.ByteBuf, backupReaderIndex int) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}

func readRedis(in *goetty.ByteBuf, size int) (bool, interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func readPB(in *goetty.ByteBuf, size int, pb protoc.PB) (bool, interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}
