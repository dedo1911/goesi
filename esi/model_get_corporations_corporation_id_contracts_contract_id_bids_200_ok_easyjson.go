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

func easyjsonFece3e7dDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCorporationsCorporationIdContractsContractIdBids200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCorporationsCorporationIdContractsContractIdBids200OkList, 0, 1)
			} else {
				*out = GetCorporationsCorporationIdContractsContractIdBids200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCorporationsCorporationIdContractsContractIdBids200Ok
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
func easyjsonFece3e7dEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCorporationsCorporationIdContractsContractIdBids200OkList) {
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
func (v GetCorporationsCorporationIdContractsContractIdBids200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFece3e7dEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdContractsContractIdBids200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFece3e7dEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdContractsContractIdBids200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFece3e7dDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdContractsContractIdBids200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFece3e7dDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonFece3e7dDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCorporationsCorporationIdContractsContractIdBids200Ok) {
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
		case "amount":
			out.Amount = float32(in.Float32())
		case "bid_id":
			out.BidId = int32(in.Int32())
		case "bidder_id":
			out.BidderId = int32(in.Int32())
		case "date_bid":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateBid).UnmarshalJSON(data))
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
func easyjsonFece3e7dEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCorporationsCorporationIdContractsContractIdBids200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Float32(float32(in.Amount))
	}
	if in.BidId != 0 {
		const prefix string = ",\"bid_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.BidId))
	}
	if in.BidderId != 0 {
		const prefix string = ",\"bidder_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.BidderId))
	}
	if true {
		const prefix string = ",\"date_bid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DateBid).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCorporationsCorporationIdContractsContractIdBids200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFece3e7dEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCorporationsCorporationIdContractsContractIdBids200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFece3e7dEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCorporationsCorporationIdContractsContractIdBids200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFece3e7dDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCorporationsCorporationIdContractsContractIdBids200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFece3e7dDecodeGithubComAntihaxGoesiEsi1(l, v)
}
