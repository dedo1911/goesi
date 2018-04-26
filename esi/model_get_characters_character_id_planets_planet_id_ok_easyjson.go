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

func easyjson29210505DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdOk
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
func easyjson29210505EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdOkList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29210505EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29210505EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29210505DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29210505DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson29210505DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdOk) {
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
		case "links":
			if in.IsNull() {
				in.Skip()
				out.Links = nil
			} else {
				in.Delim('[')
				if out.Links == nil {
					if !in.IsDelim(']') {
						out.Links = make([]GetCharactersCharacterIdPlanetsPlanetIdLink, 0, 2)
					} else {
						out.Links = []GetCharactersCharacterIdPlanetsPlanetIdLink{}
					}
				} else {
					out.Links = (out.Links)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdPlanetsPlanetIdLink
					(v4).UnmarshalEasyJSON(in)
					out.Links = append(out.Links, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pins":
			if in.IsNull() {
				in.Skip()
				out.Pins = nil
			} else {
				in.Delim('[')
				if out.Pins == nil {
					if !in.IsDelim(']') {
						out.Pins = make([]GetCharactersCharacterIdPlanetsPlanetIdPin, 0, 1)
					} else {
						out.Pins = []GetCharactersCharacterIdPlanetsPlanetIdPin{}
					}
				} else {
					out.Pins = (out.Pins)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetCharactersCharacterIdPlanetsPlanetIdPin
					easyjson29210505DecodeGithubComAntihaxGoesiEsi2(in, &v5)
					out.Pins = append(out.Pins, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "routes":
			if in.IsNull() {
				in.Skip()
				out.Routes = nil
			} else {
				in.Delim('[')
				if out.Routes == nil {
					if !in.IsDelim(']') {
						out.Routes = make([]GetCharactersCharacterIdPlanetsPlanetIdRoute, 0, 1)
					} else {
						out.Routes = []GetCharactersCharacterIdPlanetsPlanetIdRoute{}
					}
				} else {
					out.Routes = (out.Routes)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetCharactersCharacterIdPlanetsPlanetIdRoute
					(v6).UnmarshalEasyJSON(in)
					out.Routes = append(out.Routes, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson29210505EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Links) != 0 {
		const prefix string = ",\"links\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Links {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Pins) != 0 {
		const prefix string = ",\"pins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Pins {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjson29210505EncodeGithubComAntihaxGoesiEsi2(out, v10)
			}
			out.RawByte(']')
		}
	}
	if len(in.Routes) != 0 {
		const prefix string = ",\"routes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Routes {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29210505EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29210505EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29210505DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29210505DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson29210505DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdPin) {
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
					var v13 GetCharactersCharacterIdPlanetsPlanetIdContent
					(v13).UnmarshalEasyJSON(in)
					out.Contents = append(out.Contents, v13)
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
func easyjson29210505EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdPin) {
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
			for v14, v15 := range in.Contents {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
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
