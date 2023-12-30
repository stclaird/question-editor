package questions

import (
	"sort"
	"strings"
)

func splitCatSubcat( catSubCat string) ( string, string ) {
    catSubCatSplit := strings.Split(catSubCat, "-")
    return catSubCatSplit[0], catSubCatSplit[1]
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

