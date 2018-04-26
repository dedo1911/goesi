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

func easyjson1bb669caDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdPinList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdPinList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdPinList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdPin
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
func easyjson1bb669caEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdPinList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdPinList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1bb669caEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdPinList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1bb669caEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPinList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1bb669caDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPinList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1bb669caDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson1bb669caDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdPin) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "contents":
			if in.IsNull() {
				in.Skip()
				out.Contents = nil
			} else {
				in.Delim('[')
				if out.Contents == nil {
					if !in.IsDelim(']') {
						out.Contents = make([]GetCharactersCharacterIdPlanetsPlanetIdContent, 0, 4)
					} else {
						out.Contents = []GetCharactersCharacterIdPlanetsPlanetIdContent{}
					}
				} else {
					out.Contents = (out.Contents)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdPlanetsPlanetIdContent
					(v4).UnmarshalEasyJSON(in)
					out.Contents = append(out.Contents, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "expiry_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiryTime).UnmarshalJSON(data))
			}
		case "extractor_details":
			(out.ExtractorDetails).UnmarshalEasyJSON(in)
		case "factory_details":
			(out.FactoryDetails).UnmarshalEasyJSON(in)
		case "install_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.InstallTime).UnmarshalJSON(data))
			}
		case "last_cycle_start":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastCycleStart).UnmarshalJSON(data))
			}
		case "latitude":
			out.Latitude = float32(in.Float32())
		case "longitude":
			out.Longitude = float32(in.Float32())
		case "pin_id":
			out.PinId = int64(in.Int64())
		case "schematic_id":
			out.SchematicId = int32(in.Int32())
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjson1bb669caEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdPin) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Contents) != 0 {
		const prefix string = ",\"contents\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Contents {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"expiry_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.ExpiryTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"extractor_details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ExtractorDetails).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"factory_details\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.FactoryDetails).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"install_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.InstallTime).MarshalJSON())
	}
	if true {
		const prefix string = ",\"last_cycle_start\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastCycleStart).MarshalJSON())
	}
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Longitude))
	}
	if in.PinId != 0 {
		const prefix string = ",\"pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PinId))
	}
	if in.SchematicId != 0 {
		const prefix string = ",\"schematic_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SchematicId))
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdPin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1bb669caEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdPin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1bb669caEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1bb669caDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdPin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1bb669caDecodeGithubComAntihaxGoesiEsi1(l, v)
}
