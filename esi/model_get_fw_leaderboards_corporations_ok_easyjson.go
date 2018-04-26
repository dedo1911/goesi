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

func easyjsonE4451086DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCorporationsOkList, 0, 1)
			} else {
				*out = GetFwLeaderboardsCorporationsOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCorporationsOk
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
func easyjsonE4451086EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCorporationsOkList) {
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
func (v GetFwLeaderboardsCorporationsOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4451086EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4451086EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4451086DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4451086DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonE4451086DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsOk) {
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
		case "kills":
			(out.Kills).UnmarshalEasyJSON(in)
		case "victory_points":
			easyjsonE4451086DecodeGithubComAntihaxGoesiEsi2(in, &out.VictoryPoints)
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
func easyjsonE4451086EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCorporationsOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"kills\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Kills).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"victory_points\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonE4451086EncodeGithubComAntihaxGoesiEsi2(out, in.VictoryPoints)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCorporationsOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4451086EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4451086EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4451086DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4451086DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjsonE4451086DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsVictoryPoints) {
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
		case "active_total":
			if in.IsNull() {
				in.Skip()
				out.ActiveTotal = nil
			} else {
				in.Delim('[')
				if out.ActiveTotal == nil {
					if !in.IsDelim(']') {
						out.ActiveTotal = make([]GetFwLeaderboardsCorporationsActiveTotal1, 0, 8)
					} else {
						out.ActiveTotal = []GetFwLeaderboardsCorporationsActiveTotal1{}
					}
				} else {
					out.ActiveTotal = (out.ActiveTotal)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetFwLeaderboardsCorporationsActiveTotal1
					(v4).UnmarshalEasyJSON(in)
					out.ActiveTotal = append(out.ActiveTotal, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "last_week":
			if in.IsNull() {
				in.Skip()
				out.LastWeek = nil
			} else {
				in.Delim('[')
				if out.LastWeek == nil {
					if !in.IsDelim(']') {
						out.LastWeek = make([]GetFwLeaderboardsCorporationsLastWeek1, 0, 8)
					} else {
						out.LastWeek = []GetFwLeaderboardsCorporationsLastWeek1{}
					}
				} else {
					out.LastWeek = (out.LastWeek)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetFwLeaderboardsCorporationsLastWeek1
					easyjsonE4451086DecodeGithubComAntihaxGoesiEsi3(in, &v5)
					out.LastWeek = append(out.LastWeek, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "yesterday":
			if in.IsNull() {
				in.Skip()
				out.Yesterday = nil
			} else {
				in.Delim('[')
				if out.Yesterday == nil {
					if !in.IsDelim(']') {
						out.Yesterday = make([]GetFwLeaderboardsCorporationsYesterday1, 0, 8)
					} else {
						out.Yesterday = []GetFwLeaderboardsCorporationsYesterday1{}
					}
				} else {
					out.Yesterday = (out.Yesterday)[:0]
				}
				for !in.IsDelim(']') {
					var v6 GetFwLeaderboardsCorporationsYesterday1
					easyjsonE4451086DecodeGithubComAntihaxGoesiEsi4(in, &v6)
					out.Yesterday = append(out.Yesterday, v6)
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
func easyjsonE4451086EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetFwLeaderboardsCorporationsVictoryPoints) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.ActiveTotal) != 0 {
		const prefix string = ",\"active_total\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.ActiveTotal {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.LastWeek) != 0 {
		const prefix string = ",\"last_week\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.LastWeek {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjsonE4451086EncodeGithubComAntihaxGoesiEsi3(out, v10)
			}
			out.RawByte(']')
		}
	}
	if len(in.Yesterday) != 0 {
		const prefix string = ",\"yesterday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Yesterday {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjsonE4451086EncodeGithubComAntihaxGoesiEsi4(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonE4451086DecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsYesterday1) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
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
func easyjsonE4451086EncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetFwLeaderboardsCorporationsYesterday1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Amount))
	}
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	out.RawByte('}')
}
func easyjsonE4451086DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsLastWeek1) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
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
func easyjsonE4451086EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetFwLeaderboardsCorporationsLastWeek1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Amount))
	}
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	out.RawByte('}')
}
