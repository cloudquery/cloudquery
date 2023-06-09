package tableoptions

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/xhit/go-str2duration/v2"
	"golang.org/x/exp/slices"
)

type changeCaseFunc func(string) string

// skipFields is a list of fields that should not be changed. This is useful for fields that are
// maps, where case needs to be preserved. Right now skipFields only supports top level fields,
// but recursive support could be added if needed later.
func changeCaseForObject(obj any, changeCase changeCaseFunc, skipFields ...string) {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() == reflect.String {
				nk := changeCase(k.String())
				v := iter.Value()
				if slices.Contains(skipFields, k.String()) {
					continue
				}
				changeCaseForObject(v.Interface(), changeCase)
				value.SetMapIndex(k, reflect.Value{})
				value.SetMapIndex(reflect.ValueOf(nk), v)
			}
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			changeCaseForObject(value.Index(i).Interface(), changeCase)
		}
	}
}

// processRelativeTimes processes relative times in the given object. Relative times are strings in the "now+<duration>" or "now-<duration>" format.
// Date truncation is supported by the use of the '%' symbol, as in "now+<duration>%<duration for truncation>". The first duration is the amount of time
// to add or subtract from the reference time, and the second duration is the amount of time to truncate to. The truncation duration can be omitted.
func processRelativeTimes(obj any, ref time.Time, fields []string) error {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() != reflect.String || !slices.Contains(fields, k.String()) {
				continue
			}

			v := iter.Value().Interface().(string)

			var (
				durDir int
				trunc  time.Duration
			)
			switch {
			case strings.HasPrefix(v, "now+"):
				durDir = 1
				v = v[4:]
			case strings.HasPrefix(v, "now-"):
				durDir = -1
				v = v[4:]
			case strings.HasPrefix(v, "now"):
				durDir = 0
				v = " " + v[3:]
			default:
				continue
			}
			if strings.Contains(v, "%") {
				parts := strings.SplitN(v, "%", 2)
				t, err := str2duration.ParseDuration(parts[1])
				if err != nil {
					return fmt.Errorf("field %s truncation: %w", k.String(), err)
				}
				trunc = t
				v = parts[0]
			}

			tt := ref
			if strings.TrimSpace(v) != "" {
				dur, err := str2duration.ParseDuration(v)
				if err != nil {
					return fmt.Errorf("field %s: %w", k.String(), err)
				}
				tt = tt.Add(time.Duration(durDir) * dur)
			}
			if trunc != 0 {
				tt = tt.Truncate(trunc)
			}

			value.SetMapIndex(k, reflect.Value{})
			value.SetMapIndex(k, reflect.ValueOf(tt))
		}
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			err := processRelativeTimes(value.Index(i).Interface(), ref, fields)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
