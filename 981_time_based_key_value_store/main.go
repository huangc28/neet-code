package main

type Data struct {
	Val       string
	Timestamp int
}

type TimeMap struct {
	datamap map[string][]Data
}

func Constructor() TimeMap {
	return TimeMap{datamap: make(map[string][]Data)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.datamap[key] = append(
		this.datamap[key],
		Data{
			Timestamp: timestamp,
			Val:       value,
		},
	)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	vals, exists := this.datamap[key]

	if !exists {
		return ""
	}

	left, right := 0, len(vals)-1
	sol := Data{Timestamp: -1, Val: ""}

	for left <= right {
		mid := (left + right) / 2
		if timestamp == vals[mid].Timestamp {
			return vals[mid].Val
		}

		if timestamp > vals[mid].Timestamp {
			sol = max(sol, vals[mid])
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return sol.Val
}

func max(i, j Data) Data {
	if i.Timestamp > j.Timestamp {
		return i
	}
	return j
}
