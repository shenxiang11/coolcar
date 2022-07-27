package token

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtY6fwIKXX2Y2GrS706yU
p1tNlKPYpCfJJQIiKdAE4iD1gySWjnPHFJ5QfRPYyKDELY9jVUzD9MP/y9nSDlsm
baXIDSFnR9V7FMatnhu4NDT6oqpX/UEGAta4xLRnE9X1RbVOBlaZSWxLctQ1U+4+
p5OsMwcalYlUk+QVUf+pTFCXjCn4KAXYLk6S+9h6X4iWdhZ+x3wAFAIg4t+xtgtD
gEAewCtviTazBJQsVWxr8lwLJz6vGkqK28y/azLXc5d3RVYruxvfEg6Aqfmnjl7h
H6d4lpVRXPqwEANO72gfdcGIxkBZ8I4RZdT3ZW8xnreAetpaws+R3KBsR0aBHOeA
/JYYOASXwCHbA8SBlU9gZArU5nQkoX2NQ/N8F7BtDfnBMJJUC83O7/m2Rz+ohC/Q
NGc4HKrcIpDckIh3l9dYl7mCBDM/Xu2f9IhFViLOkGtvuNejqzxHiZaRpG9PipqA
mFvljwX8LF+MWfMTF80xuO8+iHZDL9ORTrcdjAOYFZhLpDrEFMx8WhdbAzmrTC80
xSlSERHFzbTHnEJ6avNyR/+/NhvYBEIKTsikCBz1NTjhH8pV4K6/DmTRAOtbPPf6
CggDsPfcUioNqtQYO10gxmu5g+476SLvu4+H1CMk7l1R5PecgSKgT3fhxUpKIRmO
CHtk+NnjE86CHz6C7auuoqMCAwEAAQ==
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{PublicKey: pubKey}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			"valid_token",
			"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWEzIn0.soAAcyTcYU_yNmN0Lafzds2wcDqLyMx5nX9rpFW6JGIFoIovBSyZIlGCivngk4MaUvuO72vnDlSgnXxAz1eHdNnMwxosvgdSCD5RRduYXteHKj_yNEA1mVLuqmergAqJmN5Qw_JEsFCXaVb2SK5OyDb0F2k_cDza9RyuxkOwX9PzaZ8SFU1v-kUjDTgiE_mCV5nkgX9-GECbzcr_CI3LqMnZK86UGdSiXBeAxUacxNBWMz_IKmLvMLeCZasxli9LYDMzwwC6-0BV3ly3aLXxHie4c3pxw6nTDHsXj4_Oe4gvF74_yqRFLPfcFXoGTwum5YMbXkP7W3y_D5Slm2nVJlj-1sev8k5IGjHFYP8B4jL0leI7AxhlqMQp6rRgJtuPMwRZfkIs_D172i4goWIHP1H7PBv80ATFEN1SkzKpNqlxrpe4jEqdjQz2Bx55xCZgE8rvjnBS5WwTqHNxbBemTcr2eLWOLWs0w5--fAvxz8H-MZ7Jky_YCTvyna_OqwNqWqyv6ETjTSGpc1kM1uACcAfmezf9c1vFOU04JLo37gfpxoKgqdHD-0_AeEd0hPlf8BYDXxA97TP7m4zCrLV9BIEFXFRnAAaBCma32DIWXxQKczD5xz382L93mfyQ45uSKnPUMKsbHD_xdVWOGlcXH_-tEOXwpmIKebuPvTKUFG8",
			time.Unix(1516239122, 0),
			"5f7c3168e2283aa722e351a3",
			false,
		},
		{
			"token_expired",
			"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWEzIn0.soAAcyTcYU_yNmN0Lafzds2wcDqLyMx5nX9rpFW6JGIFoIovBSyZIlGCivngk4MaUvuO72vnDlSgnXxAz1eHdNnMwxosvgdSCD5RRduYXteHKj_yNEA1mVLuqmergAqJmN5Qw_JEsFCXaVb2SK5OyDb0F2k_cDza9RyuxkOwX9PzaZ8SFU1v-kUjDTgiE_mCV5nkgX9-GECbzcr_CI3LqMnZK86UGdSiXBeAxUacxNBWMz_IKmLvMLeCZasxli9LYDMzwwC6-0BV3ly3aLXxHie4c3pxw6nTDHsXj4_Oe4gvF74_yqRFLPfcFXoGTwum5YMbXkP7W3y_D5Slm2nVJlj-1sev8k5IGjHFYP8B4jL0leI7AxhlqMQp6rRgJtuPMwRZfkIs_D172i4goWIHP1H7PBv80ATFEN1SkzKpNqlxrpe4jEqdjQz2Bx55xCZgE8rvjnBS5WwTqHNxbBemTcr2eLWOLWs0w5--fAvxz8H-MZ7Jky_YCTvyna_OqwNqWqyv6ETjTSGpc1kM1uACcAfmezf9c1vFOU04JLo37gfpxoKgqdHD-0_AeEd0hPlf8BYDXxA97TP7m4zCrLV9BIEFXFRnAAaBCma32DIWXxQKczD5xz382L93mfyQ45uSKnPUMKsbHD_xdVWOGlcXH_-tEOXwpmIKebuPvTKUFG8",
			time.Unix(1517239122, 0),
			"",
			true,
		},
		{
			"bad_token",
			"bad_token",
			time.Unix(1516239122, 0),
			"",
			true,
		},
		{
			"wrong_signature",
			"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzIyZTM1MWEzIn0.soAAcyTcYU_yNmN0Lafzds2wcDqLyMx5nX9rpFW6JGIFoIovBSyZIlGCivngk4MaUvuO71vnDlSgnXxAz1eHdNnMwxosvgdSCD5RRduYXteHKj_yNEA1mVLuqmergAqJmN5Qw_JEsFCXaVb2SK5OyDb0F2k_cDza9RyuxkOwX9PzaZ8SFU1v-kUjDTgiE_mCV5nkgX9-GECbzcr_CI3LqMnZK86UGdSiXBeAxUacxNBWMz_IKmLvMLeCZasxli9LYDMzwwC6-0BV3ly3aLXxHie4c3pxw6nTDHsXj4_Oe4gvF74_yqRFLPfcFXoGTwum5YMbXkP7W3y_D5Slm2nVJlj-1sev8k5IGjHFYP8B4jL0leI7AxhlqMQp6rRgJtuPMwRZfkIs_D172i4goWIHP1H7PBv80ATFEN1SkzKpNqlxrpe4jEqdjQz2Bx55xCZgE8rvjnBS5WwTqHNxbBemTcr2eLWOLWs0w5--fAvxz8H-MZ7Jky_YCTvyna_OqwNqWqyv6ETjTSGpc1kM1uACcAfmezf9c1vFOU04JLo37gfpxoKgqdHD-0_AeEd0hPlf8BYDXxA97TP7m4zCrLV9BIEFXFRnAAaBCma32DIWXxQKczD5xz382L93mfyQ45uSKnPUMKsbHD_xdVWOGlcXH_-tEOXwpmIKebuPvTKUFG8",
			time.Unix(1516239122, 0),
			"",
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}

			accountID, err := v.Verify(c.tkn)

			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}

			if c.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}

			if accountID != c.want {
				t.Errorf("wrong account id. want: %q, got %q", c.want, accountID)
			}
		})
	}
}
