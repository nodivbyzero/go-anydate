package goanydate

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidDateFormat = errors.New("invalid date format")

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func isShortMonth(month string) bool {
	for _, n := range shortMonthNames {
		if strings.EqualFold(month, n) {
			return true
		}
	}
	return false
}

func isLongMonth(month string) bool {
	for _, n := range longMonthNames {
		if strings.EqualFold(month, n) {
			return true
		}
	}

	return false
}

func isShortWeekDay(day string) bool {
	for _, n := range shortDayNames {
		if strings.EqualFold(day, n) {
			return true
		}
	}

	return false
}

func isLongWeekDay(day string) bool {
	for _, n := range longDayNames {
		if strings.EqualFold(day, n) {
			return true
		}
	}

	return false
}

func isTzAbbr(v string) bool {
	abbrs := []string{"EET", "EEST", "SAST", "CAT", "WAT", "EAT", "GMT", "HST", "HDT", "AKST", "AKDT", "EST", "CST", "CDT", "MST", "MDT", "EDT", "AST", "ADT", "NST", "NDT", "PST", "PDT", "IST", "IDT", "PKT", "KST", "JST", "ACST", "ACDT", "AEST", "AEDT", "AWST", "UTC", "CET", "CEST", "BST", "MSK", "NZST", "NZDT"}
	for _, a := range abbrs {
		if strings.EqualFold(v, a) {
			return true
		}
	}

	return false
}

func isAmPm(v string) bool {
	if strings.EqualFold(v, "am") || strings.EqualFold(v, "pm") {
		return true
	}
	return false
}

func isPlusMinus(v string) bool {
	return v == "+" || v == "-" || v == "Z"
}

func isNanoSep(v string) bool {
	return v == "." || v == ","
}

func isNumber(v string) bool {
	for _, r := range v {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

type adChunk struct {
	Value string
	Type  string
}

type adDetector struct {
}

// Parse attempts to extract date components from the input string
func (d *adDetector) parse(input string) []adChunk {
	var chunks []adChunk
	var cur strings.Builder
	var chunkType string

	for _, r := range input {
		switch {
		case unicode.IsDigit(r):
			if chunkType == "" || chunkType == "digit" {
				chunkType = "digit"
				cur.WriteRune(r)
			} else {
				if cur.Len() > 0 {
					chunks = append(chunks, adChunk{
						Value: cur.String(),
						Type:  chunkType,
					})
					cur.Reset()
				}
				chunkType = "digit"
				cur.WriteRune(r)
			}
		case unicode.IsLetter(r):
			if chunkType == "" || chunkType == "letter" {
				chunkType = "letter"
				cur.WriteRune(r)
			} else {
				if cur.Len() > 0 {
					chunks = append(chunks, adChunk{
						Value: cur.String(),
						Type:  chunkType,
					})
					cur.Reset()
				}
				chunkType = "letter"
				cur.WriteRune(r)
			}
		default:
			// separator
			if cur.Len() > 0 {
				chunks = append(chunks, adChunk{
					Value: cur.String(),
					Type:  chunkType,
				})
				cur.Reset()
				chunkType = "sep"
				cur.WriteRune(r)
			}
		}
	}

	// Add the last chunk if exists
	if cur.Len() > 0 {
		chunks = append(chunks, adChunk{
			Value: cur.String(),
			Type:  chunkType,
		})
	}

	return chunks
}

type componentType uint8

const (
	ctSep componentType = iota
	ctYear
	ctMonth
	ctMonthNum
	ctDay
	ctWeekday
	ctAmPm
	ctHour
	ctMin
	ctSec
	ctNano
	ctTzAbbr
	ctTzSign
	ctTzHour
	ctTzMin
)

type adComponent struct {
	Value string
	Type  componentType
}

func (c *adComponent) GoFmt() string {
	switch c.Type {
	case ctSep:
		return c.Value
	case ctYear:
		if len(c.Value) == 2 {
			return "06"
		}
		return "2006"
	case ctMonth:
		if isShortMonth(c.Value) {
			return "Jan"
		}
		return "January"
	case ctMonthNum:
		if len(c.Value) == 1 {
			return "1"
		}
		return "01"
	case ctDay:
		if len(c.Value) == 1 {
			return "2"
		}
		return "02"
	case ctWeekday:
		if isShortWeekDay(c.Value) {
			return "Mon"
		}
		return "Monday"
	case ctAmPm:
		if c.Value == "am" || c.Value == "pm" {
			return "pm"
		}
		return "PM"
	case ctHour:
		if len(c.Value) == 1 {
			return "3"
		}
		return "15"
	case ctMin:
		if len(c.Value) == 1 {
			return "4"
		}
		return "04"
	case ctSec:
		if len(c.Value) == 1 {
			return "5"
		}
		return "05"
	case ctNano:
		if strings.HasSuffix(c.Value, "0") {
			return strings.Repeat("0", len(c.Value))
		}
		return strings.Repeat("9", len(c.Value))
	case ctTzAbbr:
		return "MST"
	case ctTzSign:
		if c.Value == "Z" {
			return "Z"
		}
		return "-"
	case ctTzHour:
		return "07"
	case ctTzMin:
		return "00"
	}
	return ""
}

func (d *adDetector) goFmt(components []adComponent) string {
	s := strings.Builder{}
	for _, c := range components {
		_, _ = s.WriteString(c.GoFmt())
	}

	return s.String()
}

func (d *adDetector) extractPattern(input string) (string, error) {
	components := d.parse(input)
	result := []adComponent{}
	prev := adComponent{}
	plusminus := false
	componentsMap := map[componentType]int{}
	rl := 0

	add := func(v string, vt componentType) {
		result = append(result, adComponent{Value: v, Type: vt})
		componentsMap[vt] = len(result) - 1
	}
	replaceType := func(old componentType, new componentType) bool {
		for i := len(result) - 1; i >= 0; i-- {
			if result[i].Type == old {
				result[i].Type = new
				componentsMap[new] = componentsMap[old]
				delete(componentsMap, old)
				return true
			}
		}
		return false
	}
	added := func(key componentType) bool {
		_, exists := componentsMap[key]
		return exists
	}

	for _, c := range components {
		switch c.Type {
		case "letter":
			if isAmPm(c.Value) {
				add(c.Value, ctAmPm)
			} else if !added(ctMonth) && (isShortMonth(c.Value) || isLongMonth(c.Value)) {
				add(c.Value, ctMonth)
				if added(ctMonthNum) && len(result) >= 2 {
					replaceType(ctMonthNum, ctDay)
				}
				componentsMap[ctMonthNum] = len(result) - 1
			} else if !added(ctWeekday) && (isShortWeekDay(c.Value) || isLongWeekDay(c.Value)) {
				add(c.Value, ctWeekday)
			} else if added(ctHour) && added(ctMin) && isTzAbbr(c.Value) {
				add(c.Value, ctTzAbbr)
			} else if c.Value == "Z" {
				add(c.Value, ctTzSign)
			} else {
				add(c.Value, ctSep)
			}
		case "digit":
			switch len(c.Value) {
			case 1:
				if !added(ctMonthNum) {
					add(c.Value, ctMonthNum)
				} else if !added(ctDay) {
					add(c.Value, ctDay)
				} else if !added(ctHour) {
					add(c.Value, ctHour)
				} else if !added(ctMin) {
					add(c.Value, ctMin)
				} else if !added(ctSec) {
					add(c.Value, ctSec)
				} else if isNanoSep(prev.Value) {
					add(c.Value, ctNano)
				}
			case 2:
				if prev.Value == ":" {
					if !added(ctMin) {
						// has leading sign?
						if rl >= 3 && isPlusMinus(result[rl-3].Value) {
							result[rl-3].Type = ctTzSign
							result[rl-2].Type = ctTzHour
							delete(componentsMap, ctHour)
							add(c.Value, ctTzMin)
						} else {
							if added(ctYear) && !added(ctMonthNum) && !added(ctDay) { // YYYY:MM:DD
								add(c.Value, ctMonthNum)
							} else if added(ctYear) && added(ctMonthNum) && !added(ctDay) && !added(ctHour) {
								add(c.Value, ctDay)
							} else {
								type2add := ctMin
								if rl >= 2 {
									prevComp := result[rl-2]
									if prevComp.Type != ctHour && isNumber(prevComp.Value) {
										if len(prevComp.Value) == 4 && !added(ctHour) {
											type2add = ctHour // MM-DD-YYYY:HH:MM:SS
										}
										if len(prevComp.Value) <= 2 {
											delete(componentsMap, result[rl-2].Type)
											componentsMap[ctHour] = rl - 2
											result[rl-2].Type = ctHour
										}
									}
								}
								add(c.Value, type2add)
							}
						}
					} else if !added(ctSec) && added(ctMin) {
						add(c.Value, ctSec)
					} else if !added(ctTzMin) && added(ctTzHour) {
						add(c.Value, ctTzMin)
					}
				} else if !added(ctMonthNum) {
					add(c.Value, ctMonthNum)
				} else if !added(ctDay) {
					add(c.Value, ctDay)
				} else if !added(ctYear) {
					add(c.Value, ctYear)
				} else if added(ctTzHour) && !added(ctTzMin) && prev.Value == ":" {
					add(c.Value, ctTzMin)
				} else if !added(ctHour) && !plusminus {
					add(c.Value, ctHour)
				} else if !added(ctMin) && added(ctHour) {
					add(c.Value, ctMin)
				} else if plusminus {
					result[rl-1].Type = ctTzSign
					add(c.Value, ctTzHour)
				} else if isNanoSep(prev.Value) {
					add(c.Value, ctNano)
				}
			case 4:
				if added(ctHour) && !added(ctTzHour) && plusminus {
					result[rl-1].Type = ctTzSign
					componentsMap[ctTzSign] = rl - 1
					add(c.Value[0:2], ctTzHour)
					add(c.Value[2:4], ctTzMin)
				} else if !added(ctYear) {
					add(c.Value, ctYear)
					if prev.Value == ":" && rl >= 4 && result[rl-2].Type == ctMin && result[rl-4].Type == ctHour { //MM:DD:YYYY
						result[rl-2].Type = ctDay
						componentsMap[ctDay] = rl - 2
						result[rl-4].Type = ctMonthNum
						componentsMap[ctMonthNum] = rl - 4
						delete(componentsMap, ctMin)
						delete(componentsMap, ctHour)
					}
				} else if isNanoSep(prev.Value) {
					add(c.Value, ctNano)
				}
			case 8:
				if !added(ctYear) && !added(ctMonthNum) && !added(ctDay) {
					add(c.Value[0:4], ctYear)
					add(c.Value[4:6], ctMonthNum)
					add(c.Value[6:8], ctDay)
				} else if rl > 0 && isNanoSep(prev.Value) && !added(ctNano) {
					add(c.Value, ctNano)
				}

			case 14:
				if !added(ctYear) && !added(ctMonthNum) && !added(ctDay) {
					add(c.Value[0:4], ctYear)
					add(c.Value[4:6], ctMonthNum)
					add(c.Value[6:8], ctDay)
					add(c.Value[8:10], ctHour)
					add(c.Value[10:12], ctMin)
					add(c.Value[12:14], ctSec)
				}

			default:
				if rl > 0 && isNanoSep(prev.Value) && !added(ctNano) {
					add(c.Value, ctNano)
				}
			}

		case "sep":
			add(c.Value, ctSep)
		}

		rl = len(result)
		prev = result[rl-1]
		plusminus = (prev.Value == "+" || prev.Value == "-")
	}

	// validate
	month := 0
	indexMonthNum := -1
	_, monthAdded := componentsMap[ctMonth]
	if !monthAdded {
		indexMonthNum, monthAdded = componentsMap[ctMonthNum]
		if monthAdded {
			v, err := strconv.Atoi(result[indexMonthNum].Value)
			if err != nil {
				return "", ErrInvalidDateFormat
			}
			month = v
		}
	}

	day := 0
	indexDay, dayAdded := componentsMap[ctDay]
	if dayAdded {
		v, err := strconv.Atoi(result[indexDay].Value)
		if err != nil {
			return "", ErrInvalidDateFormat
		}
		if v < 1 || v > 31 {
			return "", ErrInvalidDateFormat
		}
		day = v
	}

	if month != 0 && day != 0 {
		if month > 12 && day <= 12 {
			result[indexMonthNum].Type = ctDay
			result[indexDay].Type = ctMonthNum
			month, day = day, month
		}

		if day < 1 || day > 31 {
			return "", ErrInvalidDateFormat
		}
		if month < 1 || month > 12 {
			return "", ErrInvalidDateFormat
		}
	}

	return d.goFmt(result), nil
}

// Attempts to detect the correct Go time layout format for parsing a given time string.
// Parameters:
//   - input: A string representing a date and/or time in various possible formats
//
// Returns:
//   - A string representing the Go time layout that matches the input format
//   - An error if the input format cannot be recognized or parsed
func DetectFormat(input string) (string, error) {
	input = strings.TrimSpace(input)

	d := adDetector{}
	return d.extractPattern(input)
}
