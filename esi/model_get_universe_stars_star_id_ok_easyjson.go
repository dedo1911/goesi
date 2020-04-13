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

func easyjsonFddeaa82DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseStarsStarIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseStarsStarIdOkList, 0, 0)
			} else {
				*out = GetUniverseStarsStarIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseStarsStarIdOk
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
func easyjsonFddeaa82EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseStarsStarIdOkList) {
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
func (v GetUniverseStarsStarIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFddeaa82EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStarsStarIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFddeaa82EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStarsStarIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFddeaa82DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStarsStarIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFddeaa82DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFddeaa82DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseStarsStarIdOk) {
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
		case "age":
			out.Age = int64(in.Int64())
		case "luminosity":
			out.Luminosity = float32(in.Float32())
		case "name":
			out.Name = string(in.String())
		case "radius":
			out.Radius = int64(in.Int64())
		case "solar_system_id":
			out.SolarSystemId = int32(in.Int32())
		case "spectral_class":
			out.SpectralClass = string(in.String())
		case "temperature":
			out.Temperature = int32(in.Int32())
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
func easyjsonFddeaa82EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseStarsStarIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Age != 0 {
		const prefix string = ",\"age\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Age))
	}
	if in.Luminosity != 0 {
		const prefix string = ",\"luminosity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Luminosity))
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
	if in.Radius != 0 {
		const prefix string = ",\"radius\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Radius))
	}
	if in.SolarSystemId != 0 {
		const prefix string = ",\"solar_system_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SolarSystemId))
	}
	if in.SpectralClass != "" {
		const prefix string = ",\"spectral_class\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SpectralClass))
	}
	if in.Temperature != 0 {
		const prefix string = ",\"temperature\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Temperature))
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
func (v GetUniverseStarsStarIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFddeaa82EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseStarsStarIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFddeaa82EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseStarsStarIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFddeaa82DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseStarsStarIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFddeaa82DecodeGithubComAntihaxGoesiEsi1(l, v)
}
