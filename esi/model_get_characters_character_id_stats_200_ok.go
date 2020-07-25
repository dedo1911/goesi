/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.7.2
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetCharactersCharacterIdStats200Ok. */
//easyjson:json
type GetCharactersCharacterIdStats200OkList []GetCharactersCharacterIdStats200Ok

/* Aggregate stats for a year */
//easyjson:json
type GetCharactersCharacterIdStats200Ok struct {
	Character GetCharactersCharacterIdStatsCharacter `json:"character,omitempty"`
	Combat    GetCharactersCharacterIdStatsCombat    `json:"combat,omitempty"`
	Industry  GetCharactersCharacterIdStatsIndustry  `json:"industry,omitempty"`
	Inventory GetCharactersCharacterIdStatsInventory `json:"inventory,omitempty"`
	Isk       GetCharactersCharacterIdStatsIsk       `json:"isk,omitempty"`
	Market    GetCharactersCharacterIdStatsMarket    `json:"market,omitempty"`
	Mining    GetCharactersCharacterIdStatsMining    `json:"mining,omitempty"`
	Module    GetCharactersCharacterIdStatsModule    `json:"module,omitempty"`
	Orbital   GetCharactersCharacterIdStatsOrbital   `json:"orbital,omitempty"`
	Pve       GetCharactersCharacterIdStatsPve       `json:"pve,omitempty"`
	Social    GetCharactersCharacterIdStatsSocial    `json:"social,omitempty"`
	Travel    GetCharactersCharacterIdStatsTravel    `json:"travel,omitempty"`
	Year      int32                                  `json:"year,omitempty"` /* Gregorian year for this set of aggregates */
}
