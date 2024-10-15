// Code generated by 'yaegi extract git.golaxy.org/framework/net/gtp'. DO NOT EDIT.

package fwlib

import (
	"git.golaxy.org/framework/net/gtp"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["git.golaxy.org/framework/net/gtp/gtp"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AsymmetricEncryption_ECDSA_P256":        reflect.ValueOf(gtp.AsymmetricEncryption_ECDSA_P256),
		"AsymmetricEncryption_None":              reflect.ValueOf(gtp.AsymmetricEncryption_None),
		"AsymmetricEncryption_RSA256":            reflect.ValueOf(gtp.AsymmetricEncryption_RSA256),
		"BlockCipherMode_CBC":                    reflect.ValueOf(gtp.BlockCipherMode_CBC),
		"BlockCipherMode_CFB":                    reflect.ValueOf(gtp.BlockCipherMode_CFB),
		"BlockCipherMode_CTR":                    reflect.ValueOf(gtp.BlockCipherMode_CTR),
		"BlockCipherMode_GCM":                    reflect.ValueOf(gtp.BlockCipherMode_GCM),
		"BlockCipherMode_None":                   reflect.ValueOf(gtp.BlockCipherMode_None),
		"BlockCipherMode_OFB":                    reflect.ValueOf(gtp.BlockCipherMode_OFB),
		"Code_AuthFailed":                        reflect.ValueOf(gtp.Code_AuthFailed),
		"Code_ContinueFailed":                    reflect.ValueOf(gtp.Code_ContinueFailed),
		"Code_Customize":                         reflect.ValueOf(constant.MakeFromLiteral("32", token.INT, 0)),
		"Code_EncryptFailed":                     reflect.ValueOf(gtp.Code_EncryptFailed),
		"Code_Reject":                            reflect.ValueOf(gtp.Code_Reject),
		"Code_SessionDeath":                      reflect.ValueOf(gtp.Code_SessionDeath),
		"Code_SessionNotFound":                   reflect.ValueOf(gtp.Code_SessionNotFound),
		"Code_Shutdown":                          reflect.ValueOf(gtp.Code_Shutdown),
		"Code_VersionError":                      reflect.ValueOf(gtp.Code_VersionError),
		"Compression_Brotli":                     reflect.ValueOf(gtp.Compression_Brotli),
		"Compression_Deflate":                    reflect.ValueOf(gtp.Compression_Deflate),
		"Compression_Gzip":                       reflect.ValueOf(gtp.Compression_Gzip),
		"Compression_LZ4":                        reflect.ValueOf(gtp.Compression_LZ4),
		"Compression_None":                       reflect.ValueOf(gtp.Compression_None),
		"Compression_Snappy":                     reflect.ValueOf(gtp.Compression_Snappy),
		"DefaultMsgCreator":                      reflect.ValueOf(gtp.DefaultMsgCreator),
		"ErrNotDeclared":                         reflect.ValueOf(&gtp.ErrNotDeclared).Elem(),
		"Flag_Auth":                              reflect.ValueOf(gtp.Flag_Auth),
		"Flag_AuthOK":                            reflect.ValueOf(gtp.Flag_AuthOK),
		"Flag_Compressed":                        reflect.ValueOf(gtp.Flag_Compressed),
		"Flag_Continue":                          reflect.ValueOf(gtp.Flag_Continue),
		"Flag_ContinueOK":                        reflect.ValueOf(gtp.Flag_ContinueOK),
		"Flag_Customize":                         reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"Flag_EncryptOK":                         reflect.ValueOf(gtp.Flag_EncryptOK),
		"Flag_Encrypted":                         reflect.ValueOf(gtp.Flag_Encrypted),
		"Flag_Encryption":                        reflect.ValueOf(gtp.Flag_Encryption),
		"Flag_HelloDone":                         reflect.ValueOf(gtp.Flag_HelloDone),
		"Flag_MAC":                               reflect.ValueOf(gtp.Flag_MAC),
		"Flag_Ping":                              reflect.ValueOf(gtp.Flag_Ping),
		"Flag_Pong":                              reflect.ValueOf(gtp.Flag_Pong),
		"Flag_ReqTime":                           reflect.ValueOf(gtp.Flag_ReqTime),
		"Flag_RespTime":                          reflect.ValueOf(gtp.Flag_RespTime),
		"Flag_Signature":                         reflect.ValueOf(gtp.Flag_Signature),
		"Flag_VerifyEncryption":                  reflect.ValueOf(gtp.Flag_VerifyEncryption),
		"Flags_None":                             reflect.ValueOf(gtp.Flags_None),
		"Hash_Fnv1a128":                          reflect.ValueOf(gtp.Hash_Fnv1a128),
		"Hash_Fnv1a32":                           reflect.ValueOf(gtp.Hash_Fnv1a32),
		"Hash_Fnv1a64":                           reflect.ValueOf(gtp.Hash_Fnv1a64),
		"Hash_None":                              reflect.ValueOf(gtp.Hash_None),
		"Hash_SHA256":                            reflect.ValueOf(gtp.Hash_SHA256),
		"LoadPrivateKeyFile":                     reflect.ValueOf(gtp.LoadPrivateKeyFile),
		"LoadPublicKeyFile":                      reflect.ValueOf(gtp.LoadPublicKeyFile),
		"MsgId_Auth":                             reflect.ValueOf(gtp.MsgId_Auth),
		"MsgId_ChangeCipherSpec":                 reflect.ValueOf(gtp.MsgId_ChangeCipherSpec),
		"MsgId_Continue":                         reflect.ValueOf(gtp.MsgId_Continue),
		"MsgId_Customize":                        reflect.ValueOf(constant.MakeFromLiteral("16", token.INT, 0)),
		"MsgId_ECDHESecretKeyExchange":           reflect.ValueOf(gtp.MsgId_ECDHESecretKeyExchange),
		"MsgId_Finished":                         reflect.ValueOf(gtp.MsgId_Finished),
		"MsgId_Heartbeat":                        reflect.ValueOf(gtp.MsgId_Heartbeat),
		"MsgId_Hello":                            reflect.ValueOf(gtp.MsgId_Hello),
		"MsgId_None":                             reflect.ValueOf(gtp.MsgId_None),
		"MsgId_Payload":                          reflect.ValueOf(gtp.MsgId_Payload),
		"MsgId_Rst":                              reflect.ValueOf(gtp.MsgId_Rst),
		"MsgId_SyncTime":                         reflect.ValueOf(gtp.MsgId_SyncTime),
		"NamedCurve_None":                        reflect.ValueOf(gtp.NamedCurve_None),
		"NamedCurve_P256":                        reflect.ValueOf(gtp.NamedCurve_P256),
		"NamedCurve_X25519":                      reflect.ValueOf(gtp.NamedCurve_X25519),
		"NewMsgCreator":                          reflect.ValueOf(gtp.NewMsgCreator),
		"PaddingMode_None":                       reflect.ValueOf(gtp.PaddingMode_None),
		"PaddingMode_PSS":                        reflect.ValueOf(gtp.PaddingMode_PSS),
		"PaddingMode_Pkcs1v15":                   reflect.ValueOf(gtp.PaddingMode_Pkcs1v15),
		"PaddingMode_Pkcs7":                      reflect.ValueOf(gtp.PaddingMode_Pkcs7),
		"PaddingMode_X923":                       reflect.ValueOf(gtp.PaddingMode_X923),
		"ParseAsymmetricEncryption":              reflect.ValueOf(gtp.ParseAsymmetricEncryption),
		"ParseBlockCipherMode":                   reflect.ValueOf(gtp.ParseBlockCipherMode),
		"ParseCipherSuite":                       reflect.ValueOf(gtp.ParseCipherSuite),
		"ParseCompression":                       reflect.ValueOf(gtp.ParseCompression),
		"ParseHash":                              reflect.ValueOf(gtp.ParseHash),
		"ParseNamedCurve":                        reflect.ValueOf(gtp.ParseNamedCurve),
		"ParsePaddingMode":                       reflect.ValueOf(gtp.ParsePaddingMode),
		"ParseSecretKeyExchange":                 reflect.ValueOf(gtp.ParseSecretKeyExchange),
		"ParseSignatureAlgorithm":                reflect.ValueOf(gtp.ParseSignatureAlgorithm),
		"ParseSymmetricEncryption":               reflect.ValueOf(gtp.ParseSymmetricEncryption),
		"ReadPrivateKey":                         reflect.ValueOf(gtp.ReadPrivateKey),
		"ReadPublicKey":                          reflect.ValueOf(gtp.ReadPublicKey),
		"SecretKeyExchange_ECDHE":                reflect.ValueOf(gtp.SecretKeyExchange_ECDHE),
		"SecretKeyExchange_None":                 reflect.ValueOf(gtp.SecretKeyExchange_None),
		"SymmetricEncryption_AES":                reflect.ValueOf(gtp.SymmetricEncryption_AES),
		"SymmetricEncryption_ChaCha20":           reflect.ValueOf(gtp.SymmetricEncryption_ChaCha20),
		"SymmetricEncryption_ChaCha20_Poly1305":  reflect.ValueOf(gtp.SymmetricEncryption_ChaCha20_Poly1305),
		"SymmetricEncryption_None":               reflect.ValueOf(gtp.SymmetricEncryption_None),
		"SymmetricEncryption_XChaCha20":          reflect.ValueOf(gtp.SymmetricEncryption_XChaCha20),
		"SymmetricEncryption_XChaCha20_Poly1305": reflect.ValueOf(gtp.SymmetricEncryption_XChaCha20_Poly1305),
		"Version_V1_0":                           reflect.ValueOf(gtp.Version_V1_0),

		// type definitions
		"AsymmetricEncryption":      reflect.ValueOf((*gtp.AsymmetricEncryption)(nil)),
		"BlockCipherMode":           reflect.ValueOf((*gtp.BlockCipherMode)(nil)),
		"CipherSuite":               reflect.ValueOf((*gtp.CipherSuite)(nil)),
		"Code":                      reflect.ValueOf((*gtp.Code)(nil)),
		"Compression":               reflect.ValueOf((*gtp.Compression)(nil)),
		"Flag":                      reflect.ValueOf((*gtp.Flag)(nil)),
		"Flags":                     reflect.ValueOf((*gtp.Flags)(nil)),
		"Hash":                      reflect.ValueOf((*gtp.Hash)(nil)),
		"IMsgCreator":               reflect.ValueOf((*gtp.IMsgCreator)(nil)),
		"Msg":                       reflect.ValueOf((*gtp.Msg)(nil)),
		"MsgAuth":                   reflect.ValueOf((*gtp.MsgAuth)(nil)),
		"MsgChangeCipherSpec":       reflect.ValueOf((*gtp.MsgChangeCipherSpec)(nil)),
		"MsgCompressed":             reflect.ValueOf((*gtp.MsgCompressed)(nil)),
		"MsgContinue":               reflect.ValueOf((*gtp.MsgContinue)(nil)),
		"MsgECDHESecretKeyExchange": reflect.ValueOf((*gtp.MsgECDHESecretKeyExchange)(nil)),
		"MsgFinished":               reflect.ValueOf((*gtp.MsgFinished)(nil)),
		"MsgHead":                   reflect.ValueOf((*gtp.MsgHead)(nil)),
		"MsgHeartbeat":              reflect.ValueOf((*gtp.MsgHeartbeat)(nil)),
		"MsgHello":                  reflect.ValueOf((*gtp.MsgHello)(nil)),
		"MsgId":                     reflect.ValueOf((*gtp.MsgId)(nil)),
		"MsgMAC":                    reflect.ValueOf((*gtp.MsgMAC)(nil)),
		"MsgMAC32":                  reflect.ValueOf((*gtp.MsgMAC32)(nil)),
		"MsgMAC64":                  reflect.ValueOf((*gtp.MsgMAC64)(nil)),
		"MsgPacket":                 reflect.ValueOf((*gtp.MsgPacket)(nil)),
		"MsgPacketLen":              reflect.ValueOf((*gtp.MsgPacketLen)(nil)),
		"MsgPayload":                reflect.ValueOf((*gtp.MsgPayload)(nil)),
		"MsgReader":                 reflect.ValueOf((*gtp.MsgReader)(nil)),
		"MsgRst":                    reflect.ValueOf((*gtp.MsgRst)(nil)),
		"MsgSyncTime":               reflect.ValueOf((*gtp.MsgSyncTime)(nil)),
		"MsgWriter":                 reflect.ValueOf((*gtp.MsgWriter)(nil)),
		"NamedCurve":                reflect.ValueOf((*gtp.NamedCurve)(nil)),
		"PaddingMode":               reflect.ValueOf((*gtp.PaddingMode)(nil)),
		"SecretKeyExchange":         reflect.ValueOf((*gtp.SecretKeyExchange)(nil)),
		"SignatureAlgorithm":        reflect.ValueOf((*gtp.SignatureAlgorithm)(nil)),
		"SymmetricEncryption":       reflect.ValueOf((*gtp.SymmetricEncryption)(nil)),
		"Version":                   reflect.ValueOf((*gtp.Version)(nil)),

		// interface wrapper definitions
		"_IMsgCreator": reflect.ValueOf((*_git_golaxy_org_framework_net_gtp_IMsgCreator)(nil)),
		"_Msg":         reflect.ValueOf((*_git_golaxy_org_framework_net_gtp_Msg)(nil)),
		"_MsgReader":   reflect.ValueOf((*_git_golaxy_org_framework_net_gtp_MsgReader)(nil)),
		"_MsgWriter":   reflect.ValueOf((*_git_golaxy_org_framework_net_gtp_MsgWriter)(nil)),
	}
}

// _git_golaxy_org_framework_net_gtp_IMsgCreator is an interface wrapper for IMsgCreator type
type _git_golaxy_org_framework_net_gtp_IMsgCreator struct {
	IValue     interface{}
	WDeclare   func(msg gtp.Msg)
	WNew       func(msgId uint8) (gtp.Msg, error)
	WUndeclare func(msgId uint8)
}

func (W _git_golaxy_org_framework_net_gtp_IMsgCreator) Declare(msg gtp.Msg) {
	W.WDeclare(msg)
}
func (W _git_golaxy_org_framework_net_gtp_IMsgCreator) New(msgId uint8) (gtp.Msg, error) {
	return W.WNew(msgId)
}
func (W _git_golaxy_org_framework_net_gtp_IMsgCreator) Undeclare(msgId uint8) {
	W.WUndeclare(msgId)
}

// _git_golaxy_org_framework_net_gtp_Msg is an interface wrapper for Msg type
type _git_golaxy_org_framework_net_gtp_Msg struct {
	IValue interface{}
	WClone func() gtp.MsgReader
	WMsgId func() uint8
	WRead  func(p []byte) (n int, err error)
	WSize  func() int
	WWrite func(p []byte) (n int, err error)
}

func (W _git_golaxy_org_framework_net_gtp_Msg) Clone() gtp.MsgReader {
	return W.WClone()
}
func (W _git_golaxy_org_framework_net_gtp_Msg) MsgId() uint8 {
	return W.WMsgId()
}
func (W _git_golaxy_org_framework_net_gtp_Msg) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _git_golaxy_org_framework_net_gtp_Msg) Size() int {
	return W.WSize()
}
func (W _git_golaxy_org_framework_net_gtp_Msg) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}

// _git_golaxy_org_framework_net_gtp_MsgReader is an interface wrapper for MsgReader type
type _git_golaxy_org_framework_net_gtp_MsgReader struct {
	IValue interface{}
	WClone func() gtp.MsgReader
	WMsgId func() uint8
	WRead  func(p []byte) (n int, err error)
	WSize  func() int
}

func (W _git_golaxy_org_framework_net_gtp_MsgReader) Clone() gtp.MsgReader {
	return W.WClone()
}
func (W _git_golaxy_org_framework_net_gtp_MsgReader) MsgId() uint8 {
	return W.WMsgId()
}
func (W _git_golaxy_org_framework_net_gtp_MsgReader) Read(p []byte) (n int, err error) {
	return W.WRead(p)
}
func (W _git_golaxy_org_framework_net_gtp_MsgReader) Size() int {
	return W.WSize()
}

// _git_golaxy_org_framework_net_gtp_MsgWriter is an interface wrapper for MsgWriter type
type _git_golaxy_org_framework_net_gtp_MsgWriter struct {
	IValue interface{}
	WMsgId func() uint8
	WSize  func() int
	WWrite func(p []byte) (n int, err error)
}

func (W _git_golaxy_org_framework_net_gtp_MsgWriter) MsgId() uint8 {
	return W.WMsgId()
}
func (W _git_golaxy_org_framework_net_gtp_MsgWriter) Size() int {
	return W.WSize()
}
func (W _git_golaxy_org_framework_net_gtp_MsgWriter) Write(p []byte) (n int, err error) {
	return W.WWrite(p)
}
