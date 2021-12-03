package jsonBenchmark

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func deleteUnnecessaryFieldsFastJson(entitiesMap map[int][]byte, skipFields []string) map[int][]byte {
	newEntitiesMap := make(map[int][]byte, len(entitiesMap))
	for idx, entityJson := range entitiesMap {
		parser := fastJsonParserPool.Get()
		obj, _ := parser.ParseBytes(entityJson)
		for _, f := range skipFields {
			obj.Del(f)
		}
		newEntitiesMap[idx] = obj.MarshalTo(nil)
		fastJsonParserPool.Put(parser)
	}
	return newEntitiesMap
}

func deleteUnnecessaryFieldsSJson(entitiesMap map[int][]byte, skipFields []string) map[int][]byte {
	newEntitiesMap := make(map[int][]byte, len(entitiesMap))
	for idx, entityJson := range entitiesMap {
		newJs := make([]byte, len(entityJson))
		copy(newJs, entityJson)
		for _, f := range skipFields {
			newJs, _ = sjson.DeleteBytes(newJs, f)
		}
		newEntitiesMap[idx] = newJs
	}
	return newEntitiesMap
}

var opts = &sjson.Options{
	Optimistic:     false,
	ReplaceInPlace: true,
}

func deleteUnnecessaryFieldsSJsonBySet(entitiesMap map[int][]byte, fields []string) map[int][]byte {
	newEntitiesMap := make(map[int][]byte, len(entitiesMap))
	for idx, entityJson := range entitiesMap {
		newJs := make([]byte, len(entityJson))
		for _, f := range fields {
			val := gjson.GetBytes(entityJson, f).Raw
			newJs, _ = sjson.SetRawBytesOptions(
				newJs,
				f,
				[]byte(val),
				opts,
			)
		}
		newEntitiesMap[idx] = newJs
	}
	return newEntitiesMap
}

func deleteUnnecessaryFieldsGJsonMap(entitiesMap map[int][]byte, skipFields []string) map[int][]byte {
	newEntitiesMap := make(map[int][]byte, len(entitiesMap))
	for idx, entityJson := range entitiesMap {
		newJs := gjson.Parse(string(entityJson)).Value().(map[string]interface{})
		for _, f := range skipFields {
			delete(newJs, f)
		}
		j, _ := jsoniter.Marshal(newJs)
		newEntitiesMap[idx] = j
	}
	return newEntitiesMap
}
