import "maps"
type TimeMap struct {
	Store map[string]map[int]string
}

func Constructor() TimeMap {
	return TimeMap{Store: make(map[string]map[int]string)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	val, exists := this.Store[key]
	if !exists {
		this.Store[key] = make(map[int]string)
		val = this.Store[key]
	}
	val[timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	valsMap, exists := this.Store[key]
	if !exists {
		return ""
	}
	keys := maps.Keys(valsMap)

	//it would be more efficient to do a binary search
	closest := -1
	for key := range keys {
		if key == timestamp {
			return valsMap[key]
		}
		if key < timestamp && key > closest {
			closest = key
		}

	}
	return valsMap[closest]

}
