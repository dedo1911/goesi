/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev19
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

package goesiv2

/* 200 ok object */
type GetUniverseSystemsSystemIdOk struct {

	/* The constellation this solar system is in */
	ConstellationId int32 `json:"constellation_id,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* planets array */
	Planets []GetUniverseSystemsSystemIdPlanet `json:"planets,omitempty"`

	Position GetUniverseSystemsSystemIdPosition `json:"position,omitempty"`

	/* security_class string */
	SecurityClass string `json:"security_class,omitempty"`

	/* security_status number */
	SecurityStatus float32 `json:"security_status,omitempty"`

	/* stargates array */
	Stargates []int32 `json:"stargates,omitempty"`

	/* system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
}
