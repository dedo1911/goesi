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

func easyjson93059dbfDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsLastWeek1List) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsLastWeek1List, 0, 8)
			} else {
				*out = GetFwLeaderboardsLastWeek1List{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsLastWeek1
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
func easyjson93059dbfEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsLastWeek1List) {
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
func (v GetFwLeaderboardsLastWeek1List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson93059dbfEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsLastWeek1List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson93059dbfEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeek1List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson93059dbfDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeek1List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson93059dbfDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson93059dbfDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsLastWeek1) {
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
		case "faction_id":
			out.FactionId = int32(in.Int32())
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
func easyjson93059dbfEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsLastWeek1) {
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
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsLastWeek1) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson93059dbfEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsLastWeek1) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson93059dbfEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeek1) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson93059dbfDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsLastWeek1) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson93059dbfDecodeGithubComAntihaxGoesiEsi1(l, v)
}
