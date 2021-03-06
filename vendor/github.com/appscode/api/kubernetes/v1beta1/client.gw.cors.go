// Code generated by protoc-gen-grpc-gateway-cors
// source: client.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportClientsCorsPatterns returns an array of grpc gatway mux patterns for
// Clients service to enable CORS.
func ExportClientsCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Clients_Copy_0,
		pattern_Clients_Nodes_0,
		pattern_Clients_Apps_0,
		pattern_Clients_App_0,
		pattern_Clients_AppPods_0,
		pattern_Clients_Namespaces_0,
		pattern_Clients_Secrets_0,
		pattern_Clients_Secret_0,
		pattern_Clients_Jobs_0,
		pattern_Clients_Pods_0,
		pattern_Clients_Services_0,
		pattern_Clients_ReplicationControllers_0,
		pattern_Clients_ConfigMaps_0,
		pattern_Clients_ConfigMap_0,
	}
}
