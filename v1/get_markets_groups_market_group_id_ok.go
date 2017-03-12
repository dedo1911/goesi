/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev7
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
type GetMarketsGroupsMarketGroupIdOk struct {

	/* description string */
	Description string `json:"description,omitempty"`

	/* market_group_id integer */
	MarketGroupId int32 `json:"market_group_id,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* parent_group_id integer */
	ParentGroupId int32 `json:"parent_group_id,omitempty"`

	/* types array */
	Types []int32 `json:"types,omitempty"`
}
