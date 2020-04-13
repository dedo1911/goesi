// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package meta

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

func easyjson8356b13dDecodeGithubComAntihaxGoesiMeta(in *jlexer.Lexer, out *GetVerifyOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetVerifyOkList, 0, 0)
			} else {
				*out = GetVerifyOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetVerifyOk
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
func easyjson8356b13dEncodeGithubComAntihaxGoesiMeta(out *jwriter.Writer, in GetVerifyOkList) {
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
func (v GetVerifyOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8356b13dEncodeGithubComAntihaxGoesiMeta(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetVerifyOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8356b13dEncodeGithubComAntihaxGoesiMeta(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetVerifyOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8356b13dDecodeGithubComAntihaxGoesiMeta(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetVerifyOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8356b13dDecodeGithubComAntihaxGoesiMeta(l, v)
}
func easyjson8356b13dDecodeGithubComAntihaxGoesiMeta1(in *jlexer.Lexer, out *GetVerifyOk) {
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
		case "CharacterID":
			out.CharacterID = int32(in.Int32())
		case "CharacterName":
			out.CharacterName = string(in.String())
		case "CharacterOwnerHash":
			out.CharacterOwnerHash = string(in.String())
		case "ExpiresOn":
			out.ExpiresOn = string(in.String())
		case "IntellectualProperty":
			out.IntellectualProperty = string(in.String())
		case "Scopes":
			out.Scopes = string(in.String())
		case "TokenType":
			out.TokenType = string(in.String())
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
func easyjson8356b13dEncodeGithubComAntihaxGoesiMeta1(out *jwriter.Writer, in GetVerifyOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CharacterID != 0 {
		const prefix string = ",\"CharacterID\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CharacterID))
	}
	if in.CharacterName != "" {
		const prefix string = ",\"CharacterName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CharacterName))
	}
	if in.CharacterOwnerHash != "" {
		const prefix string = ",\"CharacterOwnerHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CharacterOwnerHash))
	}
	if in.ExpiresOn != "" {
		const prefix string = ",\"ExpiresOn\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExpiresOn))
	}
	if in.IntellectualProperty != "" {
		const prefix string = ",\"IntellectualProperty\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IntellectualProperty))
	}
	if in.Scopes != "" {
		const prefix string = ",\"Scopes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Scopes))
	}
	if in.TokenType != "" {
		const prefix string = ",\"TokenType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TokenType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetVerifyOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8356b13dEncodeGithubComAntihaxGoesiMeta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetVerifyOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8356b13dEncodeGithubComAntihaxGoesiMeta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetVerifyOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8356b13dDecodeGithubComAntihaxGoesiMeta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetVerifyOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8356b13dDecodeGithubComAntihaxGoesiMeta1(l, v)
}
