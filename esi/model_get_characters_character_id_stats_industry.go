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

/* A list of GetCharactersCharacterIdStatsIndustry. */
//easyjson:json
type GetCharactersCharacterIdStatsIndustryList []GetCharactersCharacterIdStatsIndustry

/* industry object */
//easyjson:json
type GetCharactersCharacterIdStatsIndustry struct {
	HackingSuccesses                           int64 `json:"hacking_successes,omitempty"`                              /* hacking_successes integer */
	JobsCancelled                              int64 `json:"jobs_cancelled,omitempty"`                                 /* jobs_cancelled integer */
	JobsCompletedCopyBlueprint                 int64 `json:"jobs_completed_copy_blueprint,omitempty"`                  /* jobs_completed_copy_blueprint integer */
	JobsCompletedInvention                     int64 `json:"jobs_completed_invention,omitempty"`                       /* jobs_completed_invention integer */
	JobsCompletedManufacture                   int64 `json:"jobs_completed_manufacture,omitempty"`                     /* jobs_completed_manufacture integer */
	JobsCompletedManufactureAsteroid           int64 `json:"jobs_completed_manufacture_asteroid,omitempty"`            /* jobs_completed_manufacture_asteroid integer */
	JobsCompletedManufactureAsteroidQuantity   int64 `json:"jobs_completed_manufacture_asteroid_quantity,omitempty"`   /* jobs_completed_manufacture_asteroid_quantity integer */
	JobsCompletedManufactureCharge             int64 `json:"jobs_completed_manufacture_charge,omitempty"`              /* jobs_completed_manufacture_charge integer */
	JobsCompletedManufactureChargeQuantity     int64 `json:"jobs_completed_manufacture_charge_quantity,omitempty"`     /* jobs_completed_manufacture_charge_quantity integer */
	JobsCompletedManufactureCommodity          int64 `json:"jobs_completed_manufacture_commodity,omitempty"`           /* jobs_completed_manufacture_commodity integer */
	JobsCompletedManufactureCommodityQuantity  int64 `json:"jobs_completed_manufacture_commodity_quantity,omitempty"`  /* jobs_completed_manufacture_commodity_quantity integer */
	JobsCompletedManufactureDeployable         int64 `json:"jobs_completed_manufacture_deployable,omitempty"`          /* jobs_completed_manufacture_deployable integer */
	JobsCompletedManufactureDeployableQuantity int64 `json:"jobs_completed_manufacture_deployable_quantity,omitempty"` /* jobs_completed_manufacture_deployable_quantity integer */
	JobsCompletedManufactureDrone              int64 `json:"jobs_completed_manufacture_drone,omitempty"`               /* jobs_completed_manufacture_drone integer */
	JobsCompletedManufactureDroneQuantity      int64 `json:"jobs_completed_manufacture_drone_quantity,omitempty"`      /* jobs_completed_manufacture_drone_quantity integer */
	JobsCompletedManufactureImplant            int64 `json:"jobs_completed_manufacture_implant,omitempty"`             /* jobs_completed_manufacture_implant integer */
	JobsCompletedManufactureImplantQuantity    int64 `json:"jobs_completed_manufacture_implant_quantity,omitempty"`    /* jobs_completed_manufacture_implant_quantity integer */
	JobsCompletedManufactureModule             int64 `json:"jobs_completed_manufacture_module,omitempty"`              /* jobs_completed_manufacture_module integer */
	JobsCompletedManufactureModuleQuantity     int64 `json:"jobs_completed_manufacture_module_quantity,omitempty"`     /* jobs_completed_manufacture_module_quantity integer */
	JobsCompletedManufactureOther              int64 `json:"jobs_completed_manufacture_other,omitempty"`               /* jobs_completed_manufacture_other integer */
	JobsCompletedManufactureOtherQuantity      int64 `json:"jobs_completed_manufacture_other_quantity,omitempty"`      /* jobs_completed_manufacture_other_quantity integer */
	JobsCompletedManufactureShip               int64 `json:"jobs_completed_manufacture_ship,omitempty"`                /* jobs_completed_manufacture_ship integer */
	JobsCompletedManufactureShipQuantity       int64 `json:"jobs_completed_manufacture_ship_quantity,omitempty"`       /* jobs_completed_manufacture_ship_quantity integer */
	JobsCompletedManufactureStructure          int64 `json:"jobs_completed_manufacture_structure,omitempty"`           /* jobs_completed_manufacture_structure integer */
	JobsCompletedManufactureStructureQuantity  int64 `json:"jobs_completed_manufacture_structure_quantity,omitempty"`  /* jobs_completed_manufacture_structure_quantity integer */
	JobsCompletedManufactureSubsystem          int64 `json:"jobs_completed_manufacture_subsystem,omitempty"`           /* jobs_completed_manufacture_subsystem integer */
	JobsCompletedManufactureSubsystemQuantity  int64 `json:"jobs_completed_manufacture_subsystem_quantity,omitempty"`  /* jobs_completed_manufacture_subsystem_quantity integer */
	JobsCompletedMaterialProductivity          int64 `json:"jobs_completed_material_productivity,omitempty"`           /* jobs_completed_material_productivity integer */
	JobsCompletedTimeProductivity              int64 `json:"jobs_completed_time_productivity,omitempty"`               /* jobs_completed_time_productivity integer */
	JobsStartedCopyBlueprint                   int64 `json:"jobs_started_copy_blueprint,omitempty"`                    /* jobs_started_copy_blueprint integer */
	JobsStartedInvention                       int64 `json:"jobs_started_invention,omitempty"`                         /* jobs_started_invention integer */
	JobsStartedManufacture                     int64 `json:"jobs_started_manufacture,omitempty"`                       /* jobs_started_manufacture integer */
	JobsStartedMaterialProductivity            int64 `json:"jobs_started_material_productivity,omitempty"`             /* jobs_started_material_productivity integer */
	JobsStartedTimeProductivity                int64 `json:"jobs_started_time_productivity,omitempty"`                 /* jobs_started_time_productivity integer */
	ReprocessItem                              int64 `json:"reprocess_item,omitempty"`                                 /* reprocess_item integer */
	ReprocessItemQuantity                      int64 `json:"reprocess_item_quantity,omitempty"`                        /* reprocess_item_quantity integer */
}
