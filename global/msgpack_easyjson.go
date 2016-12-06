// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package global

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonF2feebfaDecodeGithubComAiyunGomqttGlobal(in *jlexer.Lexer, out *JsonMsg) {
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
		case "facc":
			out.FAcc = string(in.String())
		case "ftopic":
			out.FTopic = string(in.String())
		case "type":
			out.Type = int(in.Int())
		case "qos":
			out.Qos = int(in.Int())
		case "time":
			out.Time = int(in.Int())
		case "nick":
			out.Nick = string(in.String())
		case "msgid":
			out.MsgID = string(in.String())
		case "msg":
			if in.IsNull() {
				in.Skip()
				out.Msg = nil
			} else {
				out.Msg = in.Bytes()
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
func easyjsonF2feebfaEncodeGithubComAiyunGomqttGlobal(out *jwriter.Writer, in JsonMsg) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"facc\":")
	out.String(string(in.FAcc))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ftopic\":")
	out.String(string(in.FTopic))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	out.Int(int(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"qos\":")
	out.Int(int(in.Qos))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"time\":")
	out.Int(int(in.Time))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nick\":")
	out.String(string(in.Nick))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"msgid\":")
	out.String(string(in.MsgID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"msg\":")
	out.Base64Bytes(in.Msg)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JsonMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF2feebfaEncodeGithubComAiyunGomqttGlobal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JsonMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF2feebfaEncodeGithubComAiyunGomqttGlobal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JsonMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF2feebfaDecodeGithubComAiyunGomqttGlobal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JsonMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF2feebfaDecodeGithubComAiyunGomqttGlobal(l, v)
}
func easyjsonF2feebfaDecodeGithubComAiyunGomqttGlobal1(in *jlexer.Lexer, out *JsonData) {
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
		case "msgs":
			if in.IsNull() {
				in.Skip()
				out.Msgs = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Msgs = make([]*JsonMsg, 0, 8)
				} else {
					out.Msgs = []*JsonMsg{}
				}
				for !in.IsDelim(']') {
					var v4 *JsonMsg
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(JsonMsg)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Msgs = append(out.Msgs, v4)
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
func easyjsonF2feebfaEncodeGithubComAiyunGomqttGlobal1(out *jwriter.Writer, in JsonData) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"msgs\":")
	if in.Msgs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Msgs {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JsonData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF2feebfaEncodeGithubComAiyunGomqttGlobal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JsonData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF2feebfaEncodeGithubComAiyunGomqttGlobal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JsonData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF2feebfaDecodeGithubComAiyunGomqttGlobal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JsonData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF2feebfaDecodeGithubComAiyunGomqttGlobal1(l, v)
}