/**
 * @author:伯约
 * @date:2024/5/27
 * @note:
**/

package service

import (
	"context"
	"github.com/webxiaohua/GrpcDebug/internal/dao"
	"github.com/webxiaohua/GrpcDebug/internal/pkg"
)

// 获取服务列表
type GrpcDebugDiscoveryListReq struct {
}
type GrpcDebugDiscoveryListResp struct {
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func GrpcDebugDiscoveryList(ctx context.Context, req *GrpcDebugDiscoveryListReq) (resp *GrpcDebugDiscoveryListResp, err error) {
	resp = &GrpcDebugDiscoveryListResp{
		Code:    0,
		Message: "",
		Data:    []string{},
	}
	list, listErr := pkg.DbDiscoveryList(dao.DB)
	if listErr != nil {
		err = listErr
		return
	}
	for _, item := range list {
		resp.Data = append(resp.Data, item.DiscoveryID)
	}
	return
}

// 获取服务实例列表
type GrpcNodeListReq struct {
	DiscoveryId string `json:"discovery_id"`
}
type GrpcNodeListResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    *GrpcNodeListItem `json:"data"`
}
type GrpcNodeListItem struct {
	NodeList []*NodeItem `json:"node_list"`
	PathList []*PathItem `json:"path_list"'`
}

type NodeItem struct {
	Tag      string      `json:"tag"`       // 标识符
	Addr     string      `json:"addr"`      // 调用地址 ip:port
	PathList []*PathItem `json:"path_list"` // 调用路径列表
}
type PathItem struct {
	Path   string `json:"path"`   // 调用路径
	Params string `json:"params"` // 调用参数
}

func GrpcNodeList(ctx context.Context, req *GrpcNodeListReq) (resp *GrpcNodeListResp, err error) {
	resp = &GrpcNodeListResp{
		Code:    0,
		Message: "",
		Data: &GrpcNodeListItem{
			NodeList: []*NodeItem{}, // todo 接入服务发现
			PathList: []*PathItem{},
		},
	}
	list, listErr := pkg.DbPathListByDiscovery(dao.DB, req.DiscoveryId)
	if listErr != nil {
		err = listErr
		return
	}
	for _, item := range list {
		resp.Data.PathList = append(resp.Data.PathList, &PathItem{
			Path:   item.Path,
			Params: item.Params,
		})
	}
	return
}

// grpc 调试
type GrpcDebugReq struct {
	IPAddr     string `json:"ip_addr"`
	GrpcPath   string `json:"grpc_path"`
	GrpcParams string `json:"grpc_params"`
}
type GrpcDebugResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GrpcDebug(ctx context.Context, req *GrpcDebugReq) (*GrpcDebugResp, error) {
	pkgReq := &pkg.GrpcDebugReq{
		IPAddr:     req.IPAddr,
		GrpcPath:   req.GrpcPath,
		GrpcParams: req.GrpcParams,
	}
	resp := &GrpcDebugResp{}
	grpcResp, grpcRespErr := pkg.GrpcDebug(ctx, pkgReq)
	if grpcRespErr == nil {
		resp.Code = grpcResp.Code
		resp.Message = grpcResp.Message
		resp.Data = grpcResp.Data
	}
	return resp, grpcRespErr
}
