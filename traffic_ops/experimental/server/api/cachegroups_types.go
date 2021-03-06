// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type CachegroupsTypes struct {
	Name        string                `db:"name" json:"name"`
	Description string                `db:"description" json:"description"`
	CreatedAt   time.Time             `db:"created_at" json:"createdAt"`
	Links       CachegroupsTypesLinks `json:"_links" db:-`
}

type CachegroupsTypesLinks struct {
	Self string `db:"self" json:"_self"`
}

type CachegroupsTypesLink struct {
	ID  string `db:"cachegroups_type" json:"name"`
	Ref string `db:"cachegroups_types_name_ref" json:"_ref"`
}

// @Title getCachegroupsTypesById
// @Description retrieves the cachegroups_types information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    CachegroupsTypes
// @Resource /api/2.0
// @Router /api/2.0/cachegroups_types/{id} [get]
func getCachegroupsType(name string, db *sqlx.DB) (interface{}, error) {
	ret := []CachegroupsTypes{}
	arg := CachegroupsTypes{}
	arg.Name = name
	queryStr := "select *, concat('" + API_PATH + "cachegroups_types/', name) as self"
	queryStr += " from cachegroups_types WHERE name=:name"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getCachegroupsTypess
// @Description retrieves the cachegroups_types
// @Accept  application/json
// @Success 200 {array}    CachegroupsTypes
// @Resource /api/2.0
// @Router /api/2.0/cachegroups_types [get]
func getCachegroupsTypes(db *sqlx.DB) (interface{}, error) {
	ret := []CachegroupsTypes{}
	queryStr := "select *, concat('" + API_PATH + "cachegroups_types/', name) as self"
	queryStr += " from cachegroups_types"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postCachegroupsTypes
// @Description enter a new cachegroups_types
// @Accept  application/json
// @Param                 Body body     CachegroupsTypes   true "CachegroupsTypes object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cachegroups_types [post]
func postCachegroupsType(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v CachegroupsTypes
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO cachegroups_types("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ",created_at"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ",:created_at"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putCachegroupsTypes
// @Description modify an existing cachegroups_typesentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     CachegroupsTypes   true "CachegroupsTypes object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/cachegroups_types/{id}  [put]
func putCachegroupsType(name string, payload []byte, db *sqlx.DB) (interface{}, error) {
	var arg CachegroupsTypes
	err := json.Unmarshal(payload, &arg)
	arg.Name = name
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE cachegroups_types SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",created_at = :created_at"
	sqlString += " WHERE name=:name"
	result, err := db.NamedExec(sqlString, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delCachegroupsTypesById
// @Description deletes cachegroups_types information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    CachegroupsTypes
// @Resource /api/2.0
// @Router /api/2.0/cachegroups_types/{id} [delete]
func delCachegroupsType(name string, db *sqlx.DB) (interface{}, error) {
	arg := CachegroupsTypes{}
	arg.Name = name
	result, err := db.NamedExec("DELETE FROM cachegroups_types WHERE name=:name", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
