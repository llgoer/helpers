package helpers

import (
	"testing"
)

func TestHash(t *testing.T) {
	tmd5 := Md5("123456")
	equal(t, "e10adc3949ba59abbe56e057f20f883e", tmd5)

	md5FileHash, _ := Md5File("./test/llgoer.txt")
	equal(t, "4a7e30dbc0fbdd3e9529625a9382f46f", md5FileHash)

	tsha1 := Sha1("123456")
	equal(t, "7c4a8d09ca3762af61e59520943dc26494f8941b", tsha1)

	sha1FileHash, _ := Sha1File("./test/llgoer.txt")
	equal(t, "6dc9167367e459549c43ae056c9f53428c8b6ab7", sha1FileHash)

	tcrc32 := Crc32("123456")
	equal(t, uint32(158520161), tcrc32)
}
