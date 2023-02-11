package oicq

import (
	"crypto/ecdh"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
)

// session is ecdh session in oicq.
type session struct {
	SvrPublicKeyVer uint16
	PublicKey       []byte
	ShareKey        []byte
}

const serverPublicKey = "04EBCA94D733E399B2DB96EACDD3F69A8BB0F74224E2B44E3357812211D2E62EFBC91BB553098E25E33A799ADC7F76FEB208DA7C6522CDB0719A305180CC54A82E"

func newSession() *session {
	e := &session{
		SvrPublicKeyVer: 1,
	}
	key, _ := hex.DecodeString(serverPublicKey)
	e.init(key)
	return e
}

type pubKeyResp struct {
	Meta struct {
		PubKeyVer uint16 `json:"KeyVer"`
		PubKey    string `json:"PubKey"`
	} `json:"PubKeyMeta"`
}

// fetchPubKey 从服务器获取PubKey
func (e *session) fetchPubKey(uin int64) {
	resp, err := http.Get("https://keyrotate.qq.com/rotate_key?cipher_suite_ver=305&uin=" + strconv.FormatInt(uin, 10))
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()
	pubKey := pubKeyResp{}
	err = json.NewDecoder(resp.Body).Decode(&pubKey)
	if err != nil {
		return
	}
	e.SvrPublicKeyVer = pubKey.Meta.PubKeyVer
	key, _ := hex.DecodeString(pubKey.Meta.PubKey)
	e.init(key) // todo check key sign
}

func (e *session) init(svrPubKey []byte) {
	p256 := ecdh.P256()
	local, _ := p256.GenerateKey(rand.Reader)
	remote, _ := p256.NewPublicKey(svrPubKey)
	share, _ := local.ECDH(remote)

	hash := md5.New()
	hash.Write(share[:16])
	e.ShareKey = hash.Sum(nil)
	e.PublicKey = local.PublicKey().Bytes()
}
