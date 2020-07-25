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

/* A list of GetCharactersCharacterIdStatsMining. */
//easyjson:json
type GetCharactersCharacterIdStatsMiningList []GetCharactersCharacterIdStatsMining

/* mining object */
//easyjson:json
type GetCharactersCharacterIdStatsMining struct {
	DroneMine           int64 `json:"drone_mine,omitempty"`            /* drone_mine integer */
	OreArkonor          int64 `json:"ore_arkonor,omitempty"`           /* ore_arkonor integer */
	OreBistot           int64 `json:"ore_bistot,omitempty"`            /* ore_bistot integer */
	OreCrokite          int64 `json:"ore_crokite,omitempty"`           /* ore_crokite integer */
	OreDarkOchre        int64 `json:"ore_dark_ochre,omitempty"`        /* ore_dark_ochre integer */
	OreGneiss           int64 `json:"ore_gneiss,omitempty"`            /* ore_gneiss integer */
	OreHarvestableCloud int64 `json:"ore_harvestable_cloud,omitempty"` /* ore_harvestable_cloud integer */
	OreHedbergite       int64 `json:"ore_hedbergite,omitempty"`        /* ore_hedbergite integer */
	OreHemorphite       int64 `json:"ore_hemorphite,omitempty"`        /* ore_hemorphite integer */
	OreIce              int64 `json:"ore_ice,omitempty"`               /* ore_ice integer */
	OreJaspet           int64 `json:"ore_jaspet,omitempty"`            /* ore_jaspet integer */
	OreKernite          int64 `json:"ore_kernite,omitempty"`           /* ore_kernite integer */
	OreMercoxit         int64 `json:"ore_mercoxit,omitempty"`          /* ore_mercoxit integer */
	OreOmber            int64 `json:"ore_omber,omitempty"`             /* ore_omber integer */
	OrePlagioclase      int64 `json:"ore_plagioclase,omitempty"`       /* ore_plagioclase integer */
	OrePyroxeres        int64 `json:"ore_pyroxeres,omitempty"`         /* ore_pyroxeres integer */
	OreScordite         int64 `json:"ore_scordite,omitempty"`          /* ore_scordite integer */
	OreSpodumain        int64 `json:"ore_spodumain,omitempty"`         /* ore_spodumain integer */
	OreVeldspar         int64 `json:"ore_veldspar,omitempty"`          /* ore_veldspar integer */
}
