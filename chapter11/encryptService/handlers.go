package main

import (
	"context"

	proto "github.com/alfonsovgs/hands_web_service/chapter11/encryptService/proto"
)

// Encrypter holds the information about methods
type Encrypter struct{}

// Encript converts a message into cipher and returns response
func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = EncryptString(req.Key, req.Message)
	return nil
}

// Decript converts a cipher into message and returns response
func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = DecryptString(req.Key, req.Message)
	return nil
}
