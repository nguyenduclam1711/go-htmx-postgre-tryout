package models

import (
	"log"
	"strings"
)

type FieldConfig struct {
	NotNull    bool
	IsDefault  bool
	Unique     bool
	PrimaryKey bool
	DefaultVal string
	Serial     bool
}

func ParseFieldConfig(config string) FieldConfig {
	res := FieldConfig{}
	parts := strings.Split(config, ",")
	for _, part := range parts {
		trimedPart := strings.TrimSpace(part)
		switch {
		case trimedPart == "serial":
			res.Serial = true
		case trimedPart == "unique":
			res.Unique = true
		case strings.HasPrefix(trimedPart, "default"):
			res.IsDefault = true
			defaultVal, found := strings.CutPrefix(trimedPart, "default_")
			if found {
				res.DefaultVal = defaultVal
			}
		case trimedPart == "primarykey":
			res.PrimaryKey = true
		case trimedPart == "notnull":
			res.NotNull = true
		}
	}
	return res
}

func MapTypeToPostgres(goType string) string {
	switch goType {
	case "int", "int32", "int64":
		return "int"
	case "string":
		return "text"
	case "time.Time":
		return "timestamp"
	case "float32":
		return "real"
	case "float64":
		return "double precision"
	default:
		log.Fatalf("Don't know this type: %s", goType)
		return ""
	}
}
