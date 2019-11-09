package common

import "testing"

func TestShabal256(t *testing.T) {

	digest := NewDegist()
	digest.Write([]byte("abc"))
	hash := digest.Sum(nil)
	t.Log("hash =",hash)
}
