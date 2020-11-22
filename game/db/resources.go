package db

import (
	"github.com/ghtalpo/egb/common/fileutil"
	"github.com/ghtalpo/egb/common/stringutil"
)

const languageFallback = "en"

var resources Resources

// GetResources is
func GetResources() *Resources {
	return getResources()
}

func getResources() *Resources {
	return &resources
}

// Resources represents of json, sprite, text cache
type Resources struct {
	jsonScenarioSetup map[string]interface{} // scenarioSetupData
	jsonMessage       map[string]interface{} //

	language string
}

// Initialize is
func (r *Resources) Initialize() {
	r.setLanguage(getGameLanguage())
}

func (r *Resources) setLanguage(lang string) {
	r.language = lang

	r.jsonScenarioSetup = fileutil.LoadJSON("_resources/json/scenario_setup.json", r.language)
	r.jsonMessage = fileutil.LoadJSON("_resources/json/message.json", "")
}

func readIntVec(json map[string]interface{}, key string) []byte {
	d := json[key].(map[string]interface{})

	len := len(d)
	v := make([]byte, len)

	for i := range v {
		vv := d[stringutil.Itoa(i)]
		b, _ := vv.(float64)
		v[i] = byte(b)
	}
	return v
}

////////////////////////////////////////////////////////////////////////////////////////////
func (r *Resources) getMessageEx(json map[string]interface{}, key string) string {
	j, bOk := json[key]
	if !bOk {
		return key
	}
	d := j.(map[string]interface{})
	v, ok1 := d[r.language]
	if ok1 {
		return v.(string)
	}
	vv, ok2 := d[languageFallback]
	if ok2 {
		return vv.(string)
	}
	return key
}

// GetMessage return localized string from message.json
func (r *Resources) GetMessage(key string) string {
	return r.getMessageEx(r.jsonMessage, key)
}

func (r *Resources) getDictionaryElement(json map[string]interface{}, key string, keys []string, idx int) interface{} {
	k := keys[idx]
	//
	v := json[key].(map[string]interface{})
	return v[k]
}

func (r *Resources) getArrayElement(json map[string]interface{}, key string, idx int) interface{} {
	v := json[key].(map[string]interface{})
	vv := v[stringutil.Itoa(idx)]
	return vv
}

// ////////////////////////////////////////////////////////////////////////////////////////////
func (r *Resources) getDictionaryElementFromScenarioSetup(key string, keys []string, idx int) interface{} {
	return r.getDictionaryElement(r.jsonScenarioSetup, key, keys, idx)
}

func (r *Resources) getArraryElementFromScenarioSetup(key string, idx int) interface{} {
	return r.getArrayElement(r.jsonScenarioSetup, key, idx)
}

// GetTopMenu gets
func (r *Resources) GetTopMenu(abilityIdx int) string {
	return r.getDictionaryElementFromScenarioSetup("top_menu", []string{"new_game", "continue", "options"}, abilityIdx).(string)
}
