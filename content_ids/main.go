package main

import (
	"crypto/md5"
	"fmt"

	"github.com/google/uuid"
)

var namespace = md5.Sum([]byte("cr"))

func main() {
	prefixes := map[string]string{
		"0": "01",
		"1": "01",
		"2": "23",
		"3": "23",
		"4": "45",
		"5": "45",
		"6": "67",
		"7": "67",
		"8": "89",
		"9": "89",
		"a": "ab",
		"b": "ab",
		"c": "cd",
		"d": "cd",
		"e": "ef",
		"f": "ef",
	}
	buckets := map[string]int{
		"01": 0,
		"23": 0,
		"45": 0,
		"67": 0,
		"89": 0,
		"ab": 0,
		"cd": 0,
		"ef": 0,
	}

	ls := contentIDs()
	fmt.Println("length: ", len(ls))

	for _, contentID := range ls {
		prefix := toUUID(contentID)[:1]
		bucketKey := prefixes[prefix]

		buckets[bucketKey]++
	}

	fmt.Println(buckets)
}

func toUUID(val string) string {
	uuidV5 := uuid.NewSHA1(namespace, []byte(val))
	return uuidV5.String()
}

func contentIDs() []string {
	return []string{
		"GPSGR3VWXP96",
		"GPSGY2P59DJY",
		"GPSGYZX4ZW86",
		"GPSGRQ49Z2GY",
		"GPSG6WEM8QE6",
		"GPSGY097E27Y",
	}
}
