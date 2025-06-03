package genconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// simpleYamlParser is a minimalistic YAML generator that parses Go structs
// and produces a YAML sample configuration file with optional defaults.
type simpleYamlParser struct {
	indent        int
	originalValue any
	currVal       any
	defaultOnly   bool
}

// Parse recursively walks through a struct and generates a YAML-formatted string.
func (p *simpleYamlParser) Parse(v any) string {
	p.originalValue = v

	var sb strings.Builder

	for _, v := range reflect.VisibleFields(reflect.TypeOf(v)) {
		sb.WriteString(p.vToYaml(v))
	}

	return strings.TrimSpace(sb.String())
}

// _getValue safely retrieves the reflect.Value for a given field name.
func (c *simpleYamlParser) _getValue(field string) reflect.Value {
	defer func() {
		recover()
	}()

	if c.currVal != nil {
		return reflect.ValueOf(c.currVal).FieldByName(field)
	}
	return reflect.ValueOf(c.originalValue).FieldByName(field)
}

// vToYaml converts a single struct field into its YAML representation.
func (c *simpleYamlParser) vToYaml(v reflect.StructField) string {
	var str string

	if v.Type == nil {
		panic("nil type")
	}

	if v.Type.Kind() == reflect.Ptr {
		panic("pointer type not supported")
	}

	switch v.Type.Kind() {
	case reflect.Struct:
		name := v.Tag.Get("yaml")
		str += strings.Repeat(" ", c.indent*2) + name + ":\n"
		c.indent++
		for i := 0; i < v.Type.NumField(); i++ {
			currVal := c.currVal
			c.currVal = c._getValue(v.Name).Interface()
			str += c.vToYaml(v.Type.Field(i))
			c.currVal = currVal
		}
		c.indent--
		if c.indent == 0 {
			str += "\n"
		}
	case reflect.Map:
		name := v.Tag.Get("yaml")
		str += strings.Repeat(" ", c.indent*2) + name + ":\n"

		mapValue := c._getValue(v.Name).Interface()
		mapKeys := reflect.ValueOf(mapValue).MapKeys()

		for _, key := range mapKeys {
			c.indent++
			str += strings.Repeat(" ", c.indent*2) + key.String() + ":\n"
			c.indent++

			val := reflect.ValueOf(mapValue).MapIndex(key)
			currVal := c.currVal
			c.currVal = val.Interface()

			for j := 0; j < val.Type().NumField(); j++ {
				str += c.vToYaml(val.Type().Field(j))
			}

			c.currVal = currVal
			c.indent--
			c.indent--
		}
		if c.indent == 0 {
			str += "\n"
		}
	case reflect.Slice:
		name := v.Tag.Get("yaml")
		comment := v.Tag.Get("comment")
		sliceValue := c._getValue(v.Name).Interface()
		var output []string

		var values []string
		b, err := json.Marshal(sliceValue)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(b, &values); err != nil {
			panic(err)
		}

		if len(values) == 0 || c.defaultOnly {
			output = strings.Split(v.Tag.Get("default"), ",")
		} else {
			for _, val := range values {
				if comment == "" {
					output = append(output, fmt.Sprintf("%v", val))
				} else {
					output = append(output, fmt.Sprintf("%v # %v", val, comment))
				}
			}
		}

		str += strings.Repeat(" ", c.indent*2) + name + ":\n"
		c.indent++
		for _, item := range output {
			str += strings.Repeat(" ", c.indent*2) + "- " + strings.TrimSpace(item) + "\n"
		}
		c.indent--
	case reflect.String, reflect.Int, reflect.Bool, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		defaultVal := v.Tag.Get("default")
		comment := v.Tag.Get("comment")
		required := v.Tag.Get("required")
		if required != "false" {
			required = "true"
		}

		valStr := fmt.Sprintf("%v", c._getValue(v.Name))
		if !c.defaultOnly && !strings.Contains(valStr, "<invalid reflect.Value>") {
			defaultVal = valStr
		}

		yamlName := v.Tag.Get("yaml")
		str += strings.Repeat(" ", c.indent*2) + yamlName + ": " + defaultVal
		if comment != "" {
			str += " # " + comment
		}
		if required == "false" {
			if comment != "" {
				str += " (optional)"
			} else {
				str += " # (optional)"
			}
		}
		str += "\n"
	}

	return str
}

// GenConfig generates a default YAML sample file named "config.yaml.sample".
func GenConfig(cfg any) {
	GenConfigTo(cfg, "config.yaml.sample")
}

// GenConfigTo generates a YAML sample config from a struct and writes it to the specified file.
func GenConfigTo(cfg any, filename string) {
	parser := simpleYamlParser{defaultOnly: true}

	if _, err := os.Stat(filename); err == nil {
		if err := os.Remove(filename); err != nil {
			panic(err)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(parser.Parse(cfg)); err != nil {
		panic(err)
	}
}
