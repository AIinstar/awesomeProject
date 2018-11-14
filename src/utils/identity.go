package utils

import (
	"configfile"
	"utils/log"
)

func CreateEntityCall(entityId, name, entityType, parentId, jwtString string) (code int, err error) {
	args := make(map[string]interface{}, 0)
	args["entity_id"] = entityId
	args["name"] = name
	args["type"] = entityType
	args["parent_id"] = parentId
	_, code, err = RestPost(configfile.IdentityUrl+"/auth/entity", args, "", jwtString)
	if err != nil {
		log.Error(err.Error())
		return
	}
	return
}

func DeleteEntityCall(entityId, jwtString string) (code int, err error) {
	_, code, err = RestDelete(configfile.IdentityUrl+"/auth/entity/"+entityId, nil, "", jwtString)
	if err != nil {
		log.Error(err.Error())
		return
	}
	return
}

func GetEntityChildrenCall(entityId, relationType, jwtString string) (children []string, code int, err error) {
	data := make(map[string]interface{}, 0)
	children = make([]string, 0)
	data["type"] = relationType
	resp, code, err := RestPost(configfile.IdentityUrl+"/auth/entity/"+entityId+"/children", data, "", jwtString)
	if err != nil {
		return
	}
	//log.Debug(resp)
	for _, v := range resp["data"].(map[string]interface{})["children"].([]interface{}) {
		children = append(children, v.(string))
	}
	return
}
