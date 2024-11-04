package api

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type CommonRes struct {
	g.Meta  `mime:"application/json"`
	Code    int         `json:"code" des:"Status Code"`
	Message string      `json:"message" des:"Status Message"`
	Data    interface{} `json:"data"`
}

type ResponseStatusExample interface {
	GenerateJsonFile() (string, error)
}

func GenerateJsonFile(res ResponseStatusExample, errList map[int][]gcode.Code) (path string, err error) {
	for status, codes := range errList {
		examples := make([]CommonRes, 0)
		for _, code := range codes {
			example := CommonRes{
				Code:    code.Code(),
				Message: code.Message(),
				Data:    nil,
			}
			examples = append(examples, example)
		}
		jsonData, err := json.Marshal(examples)
		if err != nil {
			return "", err
		}
		// get original type name
		resName := GetResponseName(res)
		path := "./resource/openapi/" + resName
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, 0755)
		}
		os.WriteFile(fmt.Sprintf("%s/%d.json", path, status), jsonData, 0755)
	}
	return "", nil
}

func GetResponseName(res ResponseStatusExample) string {
	t := reflect.TypeOf(res)
	var realType reflect.Value
	if t.Kind() == reflect.Ptr {
		realType = reflect.New(t.Elem()).Elem()
	} else {
		realType = reflect.New(t).Elem()
	}

	var (
		pkgPath    string
		schemaName = gstr.TrimLeft(realType.Type().String(), "*")
	)
	// Pointer type has no PkgPath.
	for realType.Type().Kind() == reflect.Ptr {
		realType = realType.Elem()
	}
	if pkgPath = realType.Type().PkgPath(); pkgPath != "" && pkgPath != "." {
		schemaName = gstr.Replace(pkgPath, `/`, `.`) + gstr.SubStrFrom(schemaName, ".")
	}
	schemaName = gstr.ReplaceByMap(schemaName, map[string]string{
		` `: ``,
		`{`: ``,
		`}`: ``,
	})
	return schemaName
}
