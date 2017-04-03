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

package goesiv1

/* 200 ok object */
type GetDogmaEffectsEffectIdOk struct {

	/* description string */
	Description string `json:"description,omitempty"`

	/* disallow_auto_repeat boolean */
	DisallowAutoRepeat bool `json:"disallow_auto_repeat,omitempty"`

	/* discharge_attribute_id integer */
	DischargeAttributeId int32 `json:"discharge_attribute_id,omitempty"`

	/* display_name string */
	DisplayName string `json:"display_name,omitempty"`

	/* duration_attribute_id integer */
	DurationAttributeId int32 `json:"duration_attribute_id,omitempty"`

	/* effect_category integer */
	EffectCategory int32 `json:"effect_category,omitempty"`

	/* effect_id integer */
	EffectId int32 `json:"effect_id,omitempty"`

	/* electronic_chance boolean */
	ElectronicChance bool `json:"electronic_chance,omitempty"`

	/* falloff_attribute_id integer */
	FalloffAttributeId int32 `json:"falloff_attribute_id,omitempty"`

	/* icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`

	/* is_assistance boolean */
	IsAssistance bool `json:"is_assistance,omitempty"`

	/* is_offensive boolean */
	IsOffensive bool `json:"is_offensive,omitempty"`

	/* is_warp_safe boolean */
	IsWarpSafe bool `json:"is_warp_safe,omitempty"`

	/* modifiers array */
	Modifiers []GetDogmaEffectsEffectIdModifier `json:"modifiers,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* post_expression integer */
	PostExpression int32 `json:"post_expression,omitempty"`

	/* pre_expression integer */
	PreExpression int32 `json:"pre_expression,omitempty"`

	/* published boolean */
	Published bool `json:"published,omitempty"`

	/* range_attribute_id integer */
	RangeAttributeId int32 `json:"range_attribute_id,omitempty"`

	/* range_chance boolean */
	RangeChance bool `json:"range_chance,omitempty"`

	/* tracking_speed_attribute_id integer */
	TrackingSpeedAttributeId int32 `json:"tracking_speed_attribute_id,omitempty"`
}
