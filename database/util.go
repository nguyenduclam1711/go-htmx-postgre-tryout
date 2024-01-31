package database

import (
	"log"
	"strings"
)

type fieldConfig struct {
	notNull    bool
	isDefault  bool
	unique     bool
	primaryKey bool
	defaultVal string
	serial     bool
}

func parseFieldConfig(config string) fieldConfig {
	res := fieldConfig{}
	parts := strings.Split(config, ",")
	for _, part := range parts {
		trimedPart := strings.TrimSpace(part)
		switch {
		case trimedPart == "serial":
			res.serial = true
		case trimedPart == "unique":
			res.unique = true
		case strings.HasPrefix(trimedPart, "default"):
			res.isDefault = true
			defaultVal, found := strings.CutPrefix(trimedPart, "default_")
			if found {
				res.defaultVal = defaultVal
			}
		case trimedPart == "primarykey":
			res.primaryKey = true
		case trimedPart == "notnull":
			res.notNull = true
		}
	}
	return res
}

func mapTypeToPostgres(goType string) string {
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
