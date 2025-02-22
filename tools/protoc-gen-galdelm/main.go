package main

import (
	"io"
	"log"
	"os"

	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"
)

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var req plugin.CodeGeneratorRequest
	if err = proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func processReq(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}

	var resp plugin.CodeGeneratorResponse
	for _, fname := range req.FileToGenerate {
		out := fname + ".dump"
		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name: proto.String(out),
		})
	}
	return &resp
}

func emitResp(resp *plugin.CodeGeneratorResponse) error {
	buf, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(buf)
	return err
}

func run() error {
	req, err := parseReq(os.Stdin)
	if err != nil {
		return err
	}

	resp := processReq(req)

	return emitResp(resp)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
