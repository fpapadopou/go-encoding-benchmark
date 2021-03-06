// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package encoding

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

func easyjsonE453ad8fDecodeGithubComFpapadopouGoEncodingBenchmarkEncoding(in *jlexer.Lexer, out *object) {
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
		case "string_prop_1":
			out.StringProp1 = string(in.String())
		case "string_prop_2":
			out.StringProp2 = string(in.String())
		case "list":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]embedded, 0, 2)
					} else {
						out.List = []embedded{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v1 embedded
					(v1).UnmarshalEasyJSON(in)
					out.List = append(out.List, v1)
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
func easyjsonE453ad8fEncodeGithubComFpapadopouGoEncodingBenchmarkEncoding(out *jwriter.Writer, in object) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"string_prop_1\":"
		out.RawString(prefix[1:])
		out.String(string(in.StringProp1))
	}
	{
		const prefix string = ",\"string_prop_2\":"
		out.RawString(prefix)
		out.String(string(in.StringProp2))
	}
	{
		const prefix string = ",\"list\":"
		out.RawString(prefix)
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.List {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v object) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComFpapadopouGoEncodingBenchmarkEncoding(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v object) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComFpapadopouGoEncodingBenchmarkEncoding(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *object) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComFpapadopouGoEncodingBenchmarkEncoding(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *object) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComFpapadopouGoEncodingBenchmarkEncoding(l, v)
}
func easyjsonE453ad8fDecodeGithubComFpapadopouGoEncodingBenchmarkEncoding1(in *jlexer.Lexer, out *embedded) {
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
		case "string_prop":
			out.StringProp = string(in.String())
		case "int_prop":
			out.IntProp = int64(in.Int64())
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
func easyjsonE453ad8fEncodeGithubComFpapadopouGoEncodingBenchmarkEncoding1(out *jwriter.Writer, in embedded) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"string_prop\":"
		out.RawString(prefix[1:])
		out.String(string(in.StringProp))
	}
	{
		const prefix string = ",\"int_prop\":"
		out.RawString(prefix)
		out.Int64(int64(in.IntProp))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v embedded) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComFpapadopouGoEncodingBenchmarkEncoding1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v embedded) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComFpapadopouGoEncodingBenchmarkEncoding1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *embedded) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComFpapadopouGoEncodingBenchmarkEncoding1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *embedded) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComFpapadopouGoEncodingBenchmarkEncoding1(l, v)
}
