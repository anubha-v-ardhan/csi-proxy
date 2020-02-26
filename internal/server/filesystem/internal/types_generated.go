// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package internal

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

type VersionedAPI interface {
	Register(grpcServer *grpc.Server)
}

// All the functions this group's server needs to define.
type ServerInterface interface {
	IsMountPoint(context.Context, *IsMountPointRequest, apiversion.Version) (*IsMountPointResponse, error)
	LinkPath(context.Context, *LinkPathRequest, apiversion.Version) (*LinkPathResponse, error)
	Mkdir(context.Context, *MkdirRequest, apiversion.Version) (*MkdirResponse, error)
	PathExists(context.Context, *PathExistsRequest, apiversion.Version) (*PathExistsResponse, error)
	Rmdir(context.Context, *RmdirRequest, apiversion.Version) (*RmdirResponse, error)
}