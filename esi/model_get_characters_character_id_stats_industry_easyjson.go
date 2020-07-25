// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson996343d4DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsIndustryList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdStatsIndustryList, 0, 0)
			} else {
				*out = GetCharactersCharacterIdStatsIndustryList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdStatsIndustry
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson996343d4EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdStatsIndustryList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStatsIndustryList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson996343d4EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsIndustryList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson996343d4EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsIndustryList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson996343d4DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsIndustryList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson996343d4DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson996343d4DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsIndustry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "hacking_successes":
			out.HackingSuccesses = int64(in.Int64())
		case "jobs_cancelled":
			out.JobsCancelled = int64(in.Int64())
		case "jobs_completed_copy_blueprint":
			out.JobsCompletedCopyBlueprint = int64(in.Int64())
		case "jobs_completed_invention":
			out.JobsCompletedInvention = int64(in.Int64())
		case "jobs_completed_manufacture":
			out.JobsCompletedManufacture = int64(in.Int64())
		case "jobs_completed_manufacture_asteroid":
			out.JobsCompletedManufactureAsteroid = int64(in.Int64())
		case "jobs_completed_manufacture_asteroid_quantity":
			out.JobsCompletedManufactureAsteroidQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_charge":
			out.JobsCompletedManufactureCharge = int64(in.Int64())
		case "jobs_completed_manufacture_charge_quantity":
			out.JobsCompletedManufactureChargeQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_commodity":
			out.JobsCompletedManufactureCommodity = int64(in.Int64())
		case "jobs_completed_manufacture_commodity_quantity":
			out.JobsCompletedManufactureCommodityQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_deployable":
			out.JobsCompletedManufactureDeployable = int64(in.Int64())
		case "jobs_completed_manufacture_deployable_quantity":
			out.JobsCompletedManufactureDeployableQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_drone":
			out.JobsCompletedManufactureDrone = int64(in.Int64())
		case "jobs_completed_manufacture_drone_quantity":
			out.JobsCompletedManufactureDroneQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_implant":
			out.JobsCompletedManufactureImplant = int64(in.Int64())
		case "jobs_completed_manufacture_implant_quantity":
			out.JobsCompletedManufactureImplantQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_module":
			out.JobsCompletedManufactureModule = int64(in.Int64())
		case "jobs_completed_manufacture_module_quantity":
			out.JobsCompletedManufactureModuleQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_other":
			out.JobsCompletedManufactureOther = int64(in.Int64())
		case "jobs_completed_manufacture_other_quantity":
			out.JobsCompletedManufactureOtherQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_ship":
			out.JobsCompletedManufactureShip = int64(in.Int64())
		case "jobs_completed_manufacture_ship_quantity":
			out.JobsCompletedManufactureShipQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_structure":
			out.JobsCompletedManufactureStructure = int64(in.Int64())
		case "jobs_completed_manufacture_structure_quantity":
			out.JobsCompletedManufactureStructureQuantity = int64(in.Int64())
		case "jobs_completed_manufacture_subsystem":
			out.JobsCompletedManufactureSubsystem = int64(in.Int64())
		case "jobs_completed_manufacture_subsystem_quantity":
			out.JobsCompletedManufactureSubsystemQuantity = int64(in.Int64())
		case "jobs_completed_material_productivity":
			out.JobsCompletedMaterialProductivity = int64(in.Int64())
		case "jobs_completed_time_productivity":
			out.JobsCompletedTimeProductivity = int64(in.Int64())
		case "jobs_started_copy_blueprint":
			out.JobsStartedCopyBlueprint = int64(in.Int64())
		case "jobs_started_invention":
			out.JobsStartedInvention = int64(in.Int64())
		case "jobs_started_manufacture":
			out.JobsStartedManufacture = int64(in.Int64())
		case "jobs_started_material_productivity":
			out.JobsStartedMaterialProductivity = int64(in.Int64())
		case "jobs_started_time_productivity":
			out.JobsStartedTimeProductivity = int64(in.Int64())
		case "reprocess_item":
			out.ReprocessItem = int64(in.Int64())
		case "reprocess_item_quantity":
			out.ReprocessItemQuantity = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson996343d4EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdStatsIndustry) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HackingSuccesses != 0 {
		const prefix string = ",\"hacking_successes\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.HackingSuccesses))
	}
	if in.JobsCancelled != 0 {
		const prefix string = ",\"jobs_cancelled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCancelled))
	}
	if in.JobsCompletedCopyBlueprint != 0 {
		const prefix string = ",\"jobs_completed_copy_blueprint\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedCopyBlueprint))
	}
	if in.JobsCompletedInvention != 0 {
		const prefix string = ",\"jobs_completed_invention\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedInvention))
	}
	if in.JobsCompletedManufacture != 0 {
		const prefix string = ",\"jobs_completed_manufacture\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufacture))
	}
	if in.JobsCompletedManufactureAsteroid != 0 {
		const prefix string = ",\"jobs_completed_manufacture_asteroid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureAsteroid))
	}
	if in.JobsCompletedManufactureAsteroidQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_asteroid_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureAsteroidQuantity))
	}
	if in.JobsCompletedManufactureCharge != 0 {
		const prefix string = ",\"jobs_completed_manufacture_charge\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureCharge))
	}
	if in.JobsCompletedManufactureChargeQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_charge_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureChargeQuantity))
	}
	if in.JobsCompletedManufactureCommodity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_commodity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureCommodity))
	}
	if in.JobsCompletedManufactureCommodityQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_commodity_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureCommodityQuantity))
	}
	if in.JobsCompletedManufactureDeployable != 0 {
		const prefix string = ",\"jobs_completed_manufacture_deployable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureDeployable))
	}
	if in.JobsCompletedManufactureDeployableQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_deployable_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureDeployableQuantity))
	}
	if in.JobsCompletedManufactureDrone != 0 {
		const prefix string = ",\"jobs_completed_manufacture_drone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureDrone))
	}
	if in.JobsCompletedManufactureDroneQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_drone_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureDroneQuantity))
	}
	if in.JobsCompletedManufactureImplant != 0 {
		const prefix string = ",\"jobs_completed_manufacture_implant\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureImplant))
	}
	if in.JobsCompletedManufactureImplantQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_implant_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureImplantQuantity))
	}
	if in.JobsCompletedManufactureModule != 0 {
		const prefix string = ",\"jobs_completed_manufacture_module\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureModule))
	}
	if in.JobsCompletedManufactureModuleQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_module_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureModuleQuantity))
	}
	if in.JobsCompletedManufactureOther != 0 {
		const prefix string = ",\"jobs_completed_manufacture_other\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureOther))
	}
	if in.JobsCompletedManufactureOtherQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_other_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureOtherQuantity))
	}
	if in.JobsCompletedManufactureShip != 0 {
		const prefix string = ",\"jobs_completed_manufacture_ship\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureShip))
	}
	if in.JobsCompletedManufactureShipQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_ship_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureShipQuantity))
	}
	if in.JobsCompletedManufactureStructure != 0 {
		const prefix string = ",\"jobs_completed_manufacture_structure\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureStructure))
	}
	if in.JobsCompletedManufactureStructureQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_structure_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureStructureQuantity))
	}
	if in.JobsCompletedManufactureSubsystem != 0 {
		const prefix string = ",\"jobs_completed_manufacture_subsystem\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureSubsystem))
	}
	if in.JobsCompletedManufactureSubsystemQuantity != 0 {
		const prefix string = ",\"jobs_completed_manufacture_subsystem_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedManufactureSubsystemQuantity))
	}
	if in.JobsCompletedMaterialProductivity != 0 {
		const prefix string = ",\"jobs_completed_material_productivity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedMaterialProductivity))
	}
	if in.JobsCompletedTimeProductivity != 0 {
		const prefix string = ",\"jobs_completed_time_productivity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsCompletedTimeProductivity))
	}
	if in.JobsStartedCopyBlueprint != 0 {
		const prefix string = ",\"jobs_started_copy_blueprint\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsStartedCopyBlueprint))
	}
	if in.JobsStartedInvention != 0 {
		const prefix string = ",\"jobs_started_invention\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsStartedInvention))
	}
	if in.JobsStartedManufacture != 0 {
		const prefix string = ",\"jobs_started_manufacture\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsStartedManufacture))
	}
	if in.JobsStartedMaterialProductivity != 0 {
		const prefix string = ",\"jobs_started_material_productivity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsStartedMaterialProductivity))
	}
	if in.JobsStartedTimeProductivity != 0 {
		const prefix string = ",\"jobs_started_time_productivity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobsStartedTimeProductivity))
	}
	if in.ReprocessItem != 0 {
		const prefix string = ",\"reprocess_item\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ReprocessItem))
	}
	if in.ReprocessItemQuantity != 0 {
		const prefix string = ",\"reprocess_item_quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ReprocessItemQuantity))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStatsIndustry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson996343d4EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsIndustry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson996343d4EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsIndustry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson996343d4DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsIndustry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson996343d4DecodeGithubComAntihaxGoesiEsi1(l, v)
}
