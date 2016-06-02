package modelJWT

//JWT Model
type JWT struct {
	Token    string `json:"token"`
	Issuer   string `json:"iss"`
	Audience string `json:"aud"`
	IssuedAt int64  `json:"iat"`
	Expires  int64  `json:"exp:"`
	JTI      string `json:"jti"`
}
