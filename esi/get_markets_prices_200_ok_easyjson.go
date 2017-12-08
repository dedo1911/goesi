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

func easyjson88693ff7DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetMarketsPrices200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetMarketsPrices200OkList, 0, 2)
			} else {
				*out = GetMarketsPrices200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetMarketsPrices200Ok
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
func easyjson88693ff7EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetMarketsPrices200OkList) {
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
func (v GetMarketsPrices200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson88693ff7EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsPrices200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson88693ff7EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsPrices200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson88693ff7DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsPrices200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson88693ff7DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson88693ff7DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetMarketsPrices200Ok) {
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
		case "type_id":
			out.TypeId = int32(in.Int32())
		case "average_price":
			out.AveragePrice = float64(in.Float64())
		case "adjusted_price":
			out.AdjustedPrice = float64(in.Float64())
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
func easyjson88693ff7EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetMarketsPrices200Ok) {
	out.RawByte('{')
	first := true
	_ = first
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
	if in.AveragePrice != 0 {
		const prefix string = ",\"average_price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.AveragePrice))
	}
	if in.AdjustedPrice != 0 {
		const prefix string = ",\"adjusted_price\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.AdjustedPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetMarketsPrices200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson88693ff7EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetMarketsPrices200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson88693ff7EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetMarketsPrices200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson88693ff7DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetMarketsPrices200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson88693ff7DecodeGithubComAntihaxGoesiEsi1(l, v)
}
