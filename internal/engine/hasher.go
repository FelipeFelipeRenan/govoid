 package engine

 const (
	offset32 = 2166126261
	prime32 = 16777619
 )


func HashFNV32(key []byte) uint32{
	var hash  uint32 = offset32

	for _, b := range key {
		hash ^= uint32(b)
		hash *= prime32


	}

	return hash
}