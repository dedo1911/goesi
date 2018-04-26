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

func easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseSystemsSystemIdOkList, 0, 1)
			} else {
				*out = GetUniverseSystemsSystemIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseSystemsSystemIdOk
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
func easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseSystemsSystemIdOkList) {
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
func (v GetUniverseSystemsSystemIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdOk) {
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
		case "constellation_id":
			out.ConstellationId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "planets":
			if in.IsNull() {
				in.Skip()
				out.Planets = nil
			} else {
				in.Delim('[')
				if out.Planets == nil {
					if !in.IsDelim(']') {
						out.Planets = make([]GetUniverseSystemsSystemIdPlanet, 0, 1)
					} else {
						out.Planets = []GetUniverseSystemsSystemIdPlanet{}
					}
				} else {
					out.Planets = (out.Planets)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetUniverseSystemsSystemIdPlanet
					easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Planets = append(out.Planets, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "position":
			easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi3(in, &out.Position)
		case "security_class":
			out.SecurityClass = string(in.String())
		case "security_status":
			out.SecurityStatus = float32(in.Float32())
		case "star_id":
			out.StarId = int32(in.Int32())
		case "stargates":
			if in.IsNull() {
				in.Skip()
				out.Stargates = nil
			} else {
				in.Delim('[')
				if out.Stargates == nil {
					if !in.IsDelim(']') {
						out.Stargates = make([]int32, 0, 16)
					} else {
						out.Stargates = []int32{}
					}
				} else {
					out.Stargates = (out.Stargates)[:0]
				}
				for !in.IsDelim(']') {
					var v5 int32
					v5 = int32(in.Int32())
					out.Stargates = append(out.Stargates, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "stations":
			if in.IsNull() {
				in.Skip()
				out.Stations = nil
			} else {
				in.Delim('[')
				if out.Stations == nil {
					if !in.IsDelim(']') {
						out.Stations = make([]int32, 0, 16)
					} else {
						out.Stations = []int32{}
					}
				} else {
					out.Stations = (out.Stations)[:0]
				}
				for !in.IsDelim(']') {
					var v6 int32
					v6 = int32(in.Int32())
					out.Stations = append(out.Stations, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "system_id":
			out.SystemId = int32(in.Int32())
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
func easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseSystemsSystemIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConstellationId != 0 {
		const prefix string = ",\"constellation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ConstellationId))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.Planets) != 0 {
		const prefix string = ",\"planets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Planets {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi2(out, v8)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi3(out, in.Position)
	}
	if in.SecurityClass != "" {
		const prefix string = ",\"security_class\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SecurityClass))
	}
	if in.SecurityStatus != 0 {
		const prefix string = ",\"security_status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.SecurityStatus))
	}
	if in.StarId != 0 {
		const prefix string = ",\"star_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.StarId))
	}
	if len(in.Stargates) != 0 {
		const prefix string = ",\"stargates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Stargates {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v10))
			}
			out.RawByte(']')
		}
	}
	if len(in.Stations) != 0 {
		const prefix string = ",\"stations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Stations {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v12))
			}
			out.RawByte(']')
		}
	}
	if in.SystemId != 0 {
		const prefix string = ",\"system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SystemId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseSystemsSystemIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseSystemsSystemIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseSystemsSystemIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPosition) {
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
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
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
func easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetUniverseSystemsSystemIdPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		const prefix string = ",\"x\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		const prefix string = ",\"z\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}
func easyjsonF4e2485aDecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetUniverseSystemsSystemIdPlanet) {
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
		case "asteroid_belts":
			if in.IsNull() {
				in.Skip()
				out.AsteroidBelts = nil
			} else {
				in.Delim('[')
				if out.AsteroidBelts == nil {
					if !in.IsDelim(']') {
						out.AsteroidBelts = make([]int32, 0, 16)
					} else {
						out.AsteroidBelts = []int32{}
					}
				} else {
					out.AsteroidBelts = (out.AsteroidBelts)[:0]
				}
				for !in.IsDelim(']') {
					var v13 int32
					v13 = int32(in.Int32())
					out.AsteroidBelts = append(out.AsteroidBelts, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "moons":
			if in.IsNull() {
				in.Skip()
				out.Moons = nil
			} else {
				in.Delim('[')
				if out.Moons == nil {
					if !in.IsDelim(']') {
						out.Moons = make([]int32, 0, 16)
					} else {
						out.Moons = []int32{}
					}
				} else {
					out.Moons = (out.Moons)[:0]
				}
				for !in.IsDelim(']') {
					var v14 int32
					v14 = int32(in.Int32())
					out.Moons = append(out.Moons, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "planet_id":
			out.PlanetId = int32(in.Int32())
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
func easyjsonF4e2485aEncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetUniverseSystemsSystemIdPlanet) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.AsteroidBelts) != 0 {
		const prefix string = ",\"asteroid_belts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v15, v16 := range in.AsteroidBelts {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v16))
			}
			out.RawByte(']')
		}
	}
	if len(in.Moons) != 0 {
		const prefix string = ",\"moons\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.Moons {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v18))
			}
			out.RawByte(']')
		}
	}
	if in.PlanetId != 0 {
		const prefix string = ",\"planet_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PlanetId))
	}
	out.RawByte('}')
}
