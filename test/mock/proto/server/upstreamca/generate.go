package mock_upstreamca

//go:generate sh -c "mockgen github.com/spiffe/spire/proto/server/upstreamca UpstreamCa > upstreamca.go"