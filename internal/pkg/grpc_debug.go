/**
 * @author:伯约
 * @date:2024/5/27
 * @note:
**/

package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	gmd "google.golang.org/grpc/metadata"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprint(*i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var headers arrayFlags

// JSONCodec is an implementation of encoding.Codec for JSON
type JSONCodec struct{}

func (j JSONCodec) Name() string {
	return "json"
}

func (j JSONCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j JSONCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func init() {
	// 注册 JSON 编解码器
	encoding.RegisterCodec(JSONCodec{})
}

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
	resp := &GrpcDebugResp{
		Code:    0,
		Message: "",
		Data:    nil,
	}
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.CallContentSubtype(JSONCodec{}.Name())),
		grpc.WithInsecure(),
	}
	if req.IPAddr == "" || req.GrpcPath == "" || req.GrpcParams == "" {
		resp.Code = 400
		resp.Message = "参数错误"
		return resp, nil
	}
	conn, err := grpc.Dial(req.IPAddr, opts...)
	if err != nil {
		return &GrpcDebugResp{
			Code:    500,
			Message: "grpc.Dial error",
			Data:    nil,
		}, err
	}
	defer conn.Close()

	var reply interface{}
	if headers != nil {
		md := gmd.MD{}
		for _, header := range headers {
			pairs := strings.Split(header, ":")
			if len(pairs) != 2 {
				return &GrpcDebugResp{
					Code:    500,
					Message: fmt.Sprintf("invalid header %s", header),
					Data:    nil,
				}, fmt.Errorf("invalid header")
			} else {
				md[strings.TrimSpace(pairs[0])] = append(md[strings.TrimSpace(pairs[0])], strings.TrimSpace(pairs[1]))
			}
		}
		ctx = gmd.NewOutgoingContext(ctx, md)
	}

	// Assuming req.GrpcParams is a JSON string that matches the request type
	var grpcReq interface{}
	err = json.Unmarshal([]byte(req.GrpcParams), &grpcReq)
	if err != nil {
		return &GrpcDebugResp{
			Code:    400,
			Message: "Invalid GrpcParams",
			Data:    nil,
		}, err
	}

	err = conn.Invoke(ctx, req.GrpcPath, grpcReq, &reply)
	if err != nil {
		return &GrpcDebugResp{
			Code:    500,
			Message: fmt.Sprintf("grpc.Invoke error %s", err.Error()),
			Data:    nil,
		}, err
	}

	return &GrpcDebugResp{
		Code:    0,
		Message: "",
		Data:    reply,
	}, nil
}
