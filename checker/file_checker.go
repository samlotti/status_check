package checker

///
/// Checks to see if a file is present, optional is max age of file
///

import (
	"fmt"
	"github.com/sosodev/duration"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileCheck struct {
	BaseCheck
	FilePath string
	MaxAge   time.Duration
	MinSize  int64
}

func (p *FileCheck) CheckRule() *RuleResult {

	st, err := os.Stat(p.FilePath)
	if err != nil {
		return &RuleResult{
			Passed: false,
			Name:   p.Name,
			Extra:  fmt.Sprintf("error stat of file: %s", err),
		}
	}

	if st.Size() < p.MinSize {
		return &RuleResult{
			Passed: false,
			Name:   p.Name,
			Extra:  fmt.Sprintf("file size too small: %d", st.Size()),
		}
	}

	if p.MaxAge > 0 {
		if time.Now().Sub(st.ModTime()) > p.MaxAge {
			return &RuleResult{
				Passed: false,
				Name:   p.Name,
				Extra:  fmt.Sprintf("file is too old: %s", st.ModTime().String()),
			}
		}
	}

	return &RuleResult{
		Passed: true,
		Name:   p.Name,
		Extra:  "",
	}
}

var fc IRule = &FileCheck{}

func FileParser(sections []string) IRule {
	var fc = &FileCheck{
		BaseCheck: BaseCheck{sections[1]},
		FilePath:  sections[2],
	}
	for k, v := range ToMap(sections[3]) {
		switch k {
		case "minSize":
			i, e := strconv.Atoi(v)
			if e != nil {
				panic(fmt.Sprintf("invalid number for minSize in file command: %s", strings.Join(sections, ":")))
			}
			fc.MinSize = int64(i)
		case "maxAge":
			d, e := duration.Parse(v)
			if e != nil {
				panic(fmt.Sprintf("invalid duration (%s) for maxAge in file command: %s : %s, see: https://www.digi.com/resources/documentation/digidocs//90001488-13/reference/r_iso_8601_duration_format.htm", v, strings.Join(sections, ":"), e))
			}
			fc.MaxAge = d.ToTimeDuration()
		default:
			panic(fmt.Sprintf("unknown entry in file command: %s", strings.Join(sections, ":")))
		}
	}
	return fc
}

func ToMap(str string) map[string]string {
	r := make(map[string]string)

	kvs := strings.Split(str, ",")
	for _, kventry := range kvs {
		kvpair := strings.Split(kventry, "=")
		if len(kvpair) != 2 {
			panic(fmt.Sprintf("invalid key value pair in file checker: %s", kventry))
		}
		r[kvpair[0]] = kvpair[1]
	}

	return r

}
