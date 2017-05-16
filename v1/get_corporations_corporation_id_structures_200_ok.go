/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.6.dev11
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

package goesiv1

/* 200 ok object */
type GetCorporationsCorporationIdStructures200Ok struct {
	CorporationId   int32                                              `json:"corporation_id,omitempty"`    /* ID of the corporation that owns the structure */
	CurrentVul      []GetCorporationsCorporationIdStructuresCurrentVul `json:"current_vul,omitempty"`       /* This week's vulnerability windows, Monday is day 0 */
	FuelExpires     string                                             `json:"fuel_expires,omitempty"`      /* Date on which the structure will run out of fuel */
	NextVul         []GetCorporationsCorporationIdStructuresNextVul    `json:"next_vul,omitempty"`          /* Next week's vulnerability windows, Monday is day 0 */
	ProfileId       int32                                              `json:"profile_id,omitempty"`        /* The id of the ACL profile for this citadel */
	Services        []GetCorporationsCorporationIdStructuresService    `json:"services,omitempty"`          /* Contains a list of service upgrades, and their state */
	StateTimerEnd   string                                             `json:"state_timer_end,omitempty"`   /* Date at which the structure will move to it's next state */
	StateTimerStart string                                             `json:"state_timer_start,omitempty"` /* Date at which the structure entered it's current state */
	StructureId     int64                                              `json:"structure_id,omitempty"`      /* The Item ID of the structure */
	SystemId        int32                                              `json:"system_id,omitempty"`         /* The solar system the structure is in */
	TypeId          int32                                              `json:"type_id,omitempty"`           /* The type id of the structure */
	UnanchorsAt     string                                             `json:"unanchors_at,omitempty"`      /* Date at which the structure will unanchor */
}
