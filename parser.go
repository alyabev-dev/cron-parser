package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct{}

func (p *Parser) Parse(field string, min, max int, name string) *CronField {
	parsers := []func(string, int, int) ([]int, bool){
		p.parseWildcard,
		p.parseStep,
		p.parseRange,
		p.parseList,
		p.parseSingleValue,
	}

	for _, parser := range parsers {
		if values, ok := parser(field, min, max); ok {
			return &CronField{Name: name, Values: values}
		}
	}

	return &CronField{Name: name, Values: []int{}}
}

func (p *Parser) Print(field *CronField) {
	values := make([]string, len(field.Values))
	for i, val := range field.Values {
		values[i] = strconv.Itoa(val)
	}
	fmt.Printf("%-14s %s\n", field.Name, strings.Join(values, " "))
}

func (p *Parser) generateRange(start, end, step int) []int {
	var result []int
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}

func (p *Parser) parseWildcard(field string, min, max int) ([]int, bool) {
	if field == "*" {
		return p.generateRange(min, max, 1), true
	}
	return nil, false
}

func (p *Parser) parseStep(field string, min, max int) ([]int, bool) {
	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		if parts[0] == "*" {
			step, err := strconv.Atoi(parts[1])
			if err != nil || step <= 0 {
				return nil, false
			}
			return p.generateRange(min, max, step), true
		}
	}
	return nil, false
}

func (p *Parser) parseRange(field string, min, max int) ([]int, bool) {
	if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || start < min || end > max || start > end {
			return nil, false
		}
		return p.generateRange(start, end, 1), true
	}
	return nil, false
}

func (p *Parser) parseList(field string, min, max int) ([]int, bool) {
	if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		var values []int
		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil || val < min || val > max {
				continue
			}
			values = append(values, val)
		}
		if len(values) > 0 {
			return values, true
		}
	}
	return nil, false
}

func (p *Parser) parseSingleValue(field string, min, max int) ([]int, bool) {
	val, err := strconv.Atoi(field)
	if err == nil && val >= min && val <= max {
		return []int{val}, true
	}
	return nil, false
}
