package api

import (
	"encoding/json"

	"github.com/thash/asana/utils"
)

type Me_t struct {
	Id         string `json:"gid"`
	Name       string
	Email      string
	Workspaces []Base
	Photo      map[string]string
}

func Me() Me_t {
	var me map[string]Me_t
	err := json.Unmarshal(Get("/api/1.0/users/me", nil), &me)
	utils.Check(err)
	return me["data"]
}
