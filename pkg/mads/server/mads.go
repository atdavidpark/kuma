package server

import (
	"context"

	envoy_api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_cache "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	"github.com/envoyproxy/go-control-plane/pkg/server/sotw/v2"
	envoy_server "github.com/envoyproxy/go-control-plane/pkg/server/v2"

	observability_proto "github.com/kumahq/kuma/api/observability/v1alpha1"
)

type Server interface {
	observability_proto.MonitoringAssignmentDiscoveryServiceServer
}

func NewServer(config envoy_cache.Cache, callbacks envoy_server.Callbacks) Server {
	sotwServer := sotw.NewServer(context.Background(), config, callbacks)
	return &server{sotwServer}
}

var _ Server = &server{}

type server struct {
	sotw.Server
}

func (s *server) DeltaMonitoringAssignments(stream observability_proto.MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsServer) error {
	panic("implement me")
}

func (s *server) StreamMonitoringAssignments(stream observability_proto.MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsServer) error {
	return s.StreamHandler(stream, "")
}

func (s *server) FetchMonitoringAssignments(ctx context.Context, request *envoy_api.DiscoveryRequest) (*envoy_api.DiscoveryResponse, error) {
	panic("implement me")
}
