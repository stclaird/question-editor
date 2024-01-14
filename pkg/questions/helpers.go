package questions

import (
	b64 "encoding/base64"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/stclaird/question-editor/pkg/common/models"
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

func generateQuestionFileName(q models.Question) (name string ){

    //Generate File name for Question

    // b := make([]byte, 4)
    // rand.Read(b) // Doesn’t actually fail
    // s := hex.EncodeToString(b)

    slug := generateQuestionID(q)[:7]

    cat := strings.ToLower(q.Category)
    catsafe := safeFileName(cat)

    subcat := strings.ToLower(q.Subcategory)
    subcatsafe := safeFileName(subcat)

    name = fmt.Sprintf("%s-%s-%s.json", catsafe, subcatsafe, slug )

    return name
}

func generateQuestionID(q models.Question) string {
    s := q.Text
    se := b64.StdEncoding.EncodeToString([]byte(s))
    return se
}

func safeFileName(s string) string {
    //Remove any undesirable characters from a file name
    re := regexp.MustCompile(`[\\/:*?"<>|]`)
    cleanString := re.ReplaceAllString(s, "")

    return cleanString
}
