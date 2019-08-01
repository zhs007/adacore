package adacoredb

import (
	adacorebase "github.com/zhs007/adacore/base"
)

// AdaCoreDBKeyPrefix - This is the prefix of AdaCoreDBKey
const AdaCoreDBKeyPrefix = "h:"

// makeKey - Generate a database key via hashname
func makeKey(hashname string) string {
	return adacorebase.AppendString(AdaCoreDBKeyPrefix, hashname)
}