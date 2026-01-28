 package engine

 const (
	offset32 = 2166126261
	prime32 = 16777619
 )


func HashFNV32(key string) uint32{
	var hash  uint32 = offset32

	for i := 0; i < len(key); i++{

		hash ^=uint32(key[i])

		hash *= prime32
	}

	return hash
}