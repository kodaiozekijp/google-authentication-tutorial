package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// 変数idTokenに、入手したIDトークンを代入する
	idToken := `eyJhbGciOiJSUzI1NiIsImtpZCI6ImI1MDljNTEzODc2OGY3Y2YyZTgyN2UwNGIyN2U3ZTRjYmM3YmI5MTkiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI2Mjk3ODY4NzY4NDItdTk1ODgydGZuNWM2M3A0bmc1ZHRpa2g3Y3Y2NzdhbTEuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI2Mjk3ODY4NzY4NDItdTk1ODgydGZuNWM2M3A0bmc1ZHRpa2g3Y3Y2NzdhbTEuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDEzODU4MDE4MzI4MjAxMTk3MzMiLCJub25jZSI6ImFiY2RlZmciLCJuYmYiOjE3NTMyNzgxMDMsIm5hbWUiOiLlsI_plqLluoPlpKciLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jSVhDZHhmZ3ZaaDJBVDVtNjBuTDRlQ1lFeVJLMUJPZmNUbUl5cWdBaWFPQWRXdE5nPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IuW6g-WkpyIsImZhbWlseV9uYW1lIjoi5bCP6ZaiIiwiaWF0IjoxNzUzMjc4NDAzLCJleHAiOjE3NTMyODIwMDMsImp0aSI6ImI5NWFiMTk0ZDIwMWNjNDA3NTY3NDY5Y2IzMzU3Mjg3ZDVhYmIwNmQifQ.eSdYfhZ2QBslsGpM7dJF6EPDb1g5kMgIOHwD6Uks16kCGbzccOaNruBn5Q3QnJfbvMfODa4_J7udT42CJjJXH-NUQ_8-ZQ3NX9YDLQgxhVhBoZwhKGXpHydpZlRC5t-U3S-BvUVedP13Bg8Di_sfEof7EMnq-94hyCA2VQIZy24Ijh0lS9RnDmOGwBvM8FnAq-_WiMzjV76u0Ycm9dMBfblwxl2H73HdzKf7iCet_RkhLtDj7DrE58ohKMv97mJvWrakjZ4bOs-iUps0j-A1hEXGdpWkGm2V5ael2buwrFBYi-gXVvV6X6l_jV8A-RiWmkUin23SdToO2QQO15xMxg`

	// IDトークンから、ヘッダー・ペイロード, 署名を入手する
	dataArray := strings.Split(idToken, ".")
	header, payload, signature := dataArray[0], dataArray[1], dataArray[2]

	// headerをbase64 decodeする
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// payloadをbase64 decodeする
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// 公開鍵構造体を作る
	N := "rHz-FQE9gjFJR_FhnzhBMPpa8NJ2nCfnXLr5LWDJOOaiGqI__Nrm6HHUCpMi52_pLqqVkCihR9xbscZ6UKr9wjp-7YTDN6A9i7QqQAJyNRIMCkJR1z6D95_pam_mIkBVnYjJ_LskOyOHI65Yvuaw6oA9iFlSyucn4B-jZRmp7JyGyU8UMohaOvJB7_boaIoEx_QY8YdoANKrp0WGawEkW6RgopgiHB7D0CXU-c_GDp0TjWCZegQzoV_fDD5eH5mc2Ai3dBylZxgQ-ZxMakYS01nmVr1atkpHT1L9W7PiCP60C8WG1aLIzZTLcABK3BWCmZ3-wBZtHZ0y9kSP35aowQ"
	E := "AQAB"

	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	// 検証するデータ
	message := sha256.Sum256([]byte(header + "." + payload))

	// 署名をbase64 decodeする
	sigData, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// sigDataを公開鍵で復号し、messageと一致するか検証する
	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}
