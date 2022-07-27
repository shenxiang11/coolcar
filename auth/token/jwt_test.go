package token

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

//-----BEGIN PUBLIC KEY-----
//MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtY6fwIKXX2Y2GrS706yU
//p1tNlKPYpCfJJQIiKdAE4iD1gySWjnPHFJ5QfRPYyKDELY9jVUzD9MP/y9nSDlsm
//baXIDSFnR9V7FMatnhu4NDT6oqpX/UEGAta4xLRnE9X1RbVOBlaZSWxLctQ1U+4+
//p5OsMwcalYlUk+QVUf+pTFCXjCn4KAXYLk6S+9h6X4iWdhZ+x3wAFAIg4t+xtgtD
//gEAewCtviTazBJQsVWxr8lwLJz6vGkqK28y/azLXc5d3RVYruxvfEg6Aqfmnjl7h
//H6d4lpVRXPqwEANO72gfdcGIxkBZ8I4RZdT3ZW8xnreAetpaws+R3KBsR0aBHOeA
///JYYOASXwCHbA8SBlU9gZArU5nQkoX2NQ/N8F7BtDfnBMJJUC83O7/m2Rz+ohC/Q
//NGc4HKrcIpDckIh3l9dYl7mCBDM/Xu2f9IhFViLOkGtvuNejqzxHiZaRpG9PipqA
//mFvljwX8LF+MWfMTF80xuO8+iHZDL9ORTrcdjAOYFZhLpDrEFMx8WhdbAzmrTC80
//xSlSERHFzbTHnEJ6avNyR/+/NhvYBEIKTsikCBz1NTjhH8pV4K6/DmTRAOtbPPf6
//CggDsPfcUioNqtQYO10gxmu5g+476SLvu4+H1CMk7l1R5PecgSKgT3fhxUpKIRmO
//CHtk+NnjE86CHz6C7auuoqMCAwEAAQ==
//-----END PUBLIC KEY-----

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}

	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}

	tkn, err := g.GenerateToken("5f7c3168e2283aa722e351a3", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}

	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWEzIn0.soAAcyTcYU_yNmN0Lafzds2wcDqLyMx5nX9rpFW6JGIFoIovBSyZIlGCivngk4MaUvuO72vnDlSgnXxAz1eHdNnMwxosvgdSCD5RRduYXteHKj_yNEA1mVLuqmergAqJmN5Qw_JEsFCXaVb2SK5OyDb0F2k_cDza9RyuxkOwX9PzaZ8SFU1v-kUjDTgiE_mCV5nkgX9-GECbzcr_CI3LqMnZK86UGdSiXBeAxUacxNBWMz_IKmLvMLeCZasxli9LYDMzwwC6-0BV3ly3aLXxHie4c3pxw6nTDHsXj4_Oe4gvF74_yqRFLPfcFXoGTwum5YMbXkP7W3y_D5Slm2nVJlj-1sev8k5IGjHFYP8B4jL0leI7AxhlqMQp6rRgJtuPMwRZfkIs_D172i4goWIHP1H7PBv80ATFEN1SkzKpNqlxrpe4jEqdjQz2Bx55xCZgE8rvjnBS5WwTqHNxbBemTcr2eLWOLWs0w5--fAvxz8H-MZ7Jky_YCTvyna_OqwNqWqyv6ETjTSGpc1kM1uACcAfmezf9c1vFOU04JLo37gfpxoKgqdHD-0_AeEd0hPlf8BYDXxA97TP7m4zCrLV9BIEFXFRnAAaBCma32DIWXxQKczD5xz382L93mfyQ45uSKnPUMKsbHD_xdVWOGlcXH_-tEOXwpmIKebuPvTKUFG8"
	if tkn != want {
		t.Errorf("wrong token generated. want: %q, got: %q", want, tkn)
	}
}
