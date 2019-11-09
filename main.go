package main

import (
	"crypto/rand"
	"encoding/hex"
    "./curve25519sign"
	"encoding/binary"
	"fmt"
	"io"
)


func randBytes(data []byte) {
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		panic(err)
	}
}
func Int64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) uint64 {
	return (binary.BigEndian.Uint64(buf))
}
// 助记词：mirror size agree lust control journey keep goodbye survive forever secret crack
// 公钥：08700f8795ee9088f0de1a13081edbf144ca0c6528975a8240dead5d0f575228
// 私钥：57960452661bf6cbd791d8acacb31a8e6b64a7fc8140ba01b521a5eebeba970b
// accountID：9082845425700467299
// 签名数据：accc1f26d687011abc6aea6e345d4f3c3841e869ba68486c59465524567e5909e0e1bedc0bf57706df7d4cdcb09dd3320b222702424f6e0bc9a0aeb499f7dbda
// 签名的原始数据:7e0cbfcaaca8726333d89ca79b019c1446f0c3addf6d9a00b0c96c8c
// address:   33d89ca79b019c1446f0c3addf6d9a00b0c96c8c十六进制accountID:      7e0cbfcaaca87263
func main()  {
	fmt.Println("abc")

	// var kp curve25519KeyPair
	//
	//
	// rand := bytes.NewReader(make([]byte,32))

	// hash := sha256.New()
	//
	// if err := kp.generate(rand); err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// fmt.Println(kp.pub)
	//
	// fmt.Println(kp.priv)
	// msg := make([]byte, 200)
	//
	// var priv, pub [32]byte
	// var random [64]byte

	// Test for random values of the keys, nonce and message
	// for i := 0; i < 100; i++ {
	// 	randBytes(priv[:])
	// 	priv[0] &= 248
	// 	priv[31] &= 63
	// 	priv[31] |= 64
	//
	// 	fmt.Println(priv)
	// 	curve25519.ScalarBaseMult(&pub, &priv)
	// 	randBytes(random[:])
	// 	randBytes(msg)
	// 	fmt.Println("random=",random)
	// 	sig := curve25519sign.Sign(&priv, msg, random)
	// 	v := curve25519sign.Verify(pub, msg, sig)
	// 	// assert.True(t, v, "Verify must work")
	//
	// 	fmt.Println(v)
	//
	// }


	var AccountId uint64 = 4413664253157728500
	AccountIdlen:= hex.EncodeToString(Int64ToBytes(AccountId))
	println(AccountIdlen)

	var AccountIdBytes [8]byte
	AccountIdByteslen,err := hex.Decode(AccountIdBytes[:],[]byte("3d407c40ef4e60f4"))
	println(AccountIdByteslen,err)



	var pub [32]byte
	rlen,err := hex.Decode(pub[:],[]byte("08700f8795ee9088f0de1a13081edbf144ca0c6528975a8240dead5d0f575228"))
	println(rlen,err)

	message := make([]byte, 28)
	rlen,err = hex.Decode(message[:],[]byte("7e0cbfcaaca8726333d89ca79b019c1446f0c3addf6d9a00b0c96c8c"))
	println(rlen,err)

	// sig := make([]byte, 64)

	var sig [64]byte

	rlen,err = hex.Decode(sig[:],[]byte("accc1f26d687011abc6aea6e345d4f3c3841e869ba68486c59465524567e5909e0e1bedc0bf57706df7d4cdcb09dd3320b222702424f6e0bc9a0aeb499f7dbda"))
	println(rlen,err)


	v := curve25519sign.VerifyMsg(message,sig[:],pub[:],  true)
	// assert.True(t, v, "Verify must work")

	fmt.Println("Verify=",v)
}
