package sync

import (
	"github.com/n42blockchain/N42/log"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
)

var defaultReadDuration = ttfbTimeout
var defaultWriteDuration = respTimeout // RESP_TIMEOUT

// SetRPCStreamDeadlines sets read and write deadlines for libp2p-based connection streams.
func SetRPCStreamDeadlines(stream network.Stream) {
	SetStreamReadDeadline(stream, defaultReadDuration)
	SetStreamWriteDeadline(stream, defaultWriteDuration)
}

// SetStreamReadDeadline for reading from libp2p connection streams, deciding when to close
// a connection based on a particular duration.
//
// NOTE: libp2p uses the system clock time for determining the deadline so we use
// time.Now() instead of the synchronized roughtime.Now(). If the system
// time is corrupted (i.e. time does not advance), the node will experience
// issues being able to properly close streams, leading to unexpected failures and possible
// memory leaks.
func SetStreamReadDeadline(stream network.Stream, duration time.Duration) {
	if err := stream.SetReadDeadline(time.Now().Add(duration)); err != nil &&
		!strings.Contains(err.Error(), "stream closed") {
		log.Debug("Could not set stream deadline", "peer", stream.Conn().RemotePeer(), "protocol", stream.Protocol(), "direction", stream.Stat().Direction, "err", err)
	}
}

// SetStreamWriteDeadline for writing to libp2p connection streams, deciding when to close
// a connection based on a particular duration.
//
// NOTE: libp2p uses the system clock time for determining the deadline so we use
// time.Now() instead of the synchronized roughtime.Now(). If the system
// time is corrupted (i.e. time does not advance), the node will experience
// issues being able to properly close streams, leading to unexpected failures and possible
// memory leaks.
func SetStreamWriteDeadline(stream network.Stream, duration time.Duration) {
	if err := stream.SetWriteDeadline(time.Now().Add(duration)); err != nil {
		log.Debug("Could not set stream deadline", "peer", stream.Conn().RemotePeer(), "protocol", stream.Protocol(), "direction", stream.Stat().Direction, "err", err)
	}
}
