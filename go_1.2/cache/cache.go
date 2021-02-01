package cache

var Cache = cacheLastUsedId{
	lastInventoryId: 2,
	lastUserId:      2,
}

type cacheLastUsedId struct {
	lastUserId      int
	lastInventoryId int
}

func (s cacheLastUsedId) GetLastUserId() int {
	s.lastUserId++
	return s.lastUserId
}

func (s cacheLastUsedId) GetLastInventoryId() int {
	s.lastInventoryId++
	return s.lastInventoryId
}
