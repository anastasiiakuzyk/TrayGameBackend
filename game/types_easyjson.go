// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package game

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

func easyjson6601e8cdDecodeUntitledGameBackendGame(in *jlexer.Lexer, out *UpdatePosPackage) {
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
		case "uuid":
			out.UUID = string(in.String())
		case "left":
			out.Left = bool(in.Bool())
		case "right":
			out.Right = bool(in.Bool())
		case "top":
			out.Top = bool(in.Bool())
		case "bottom":
			out.Bottom = bool(in.Bool())
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
func easyjson6601e8cdEncodeUntitledGameBackendGame(out *jwriter.Writer, in UpdatePosPackage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	{
		const prefix string = ",\"left\":"
		out.RawString(prefix)
		out.Bool(bool(in.Left))
	}
	{
		const prefix string = ",\"right\":"
		out.RawString(prefix)
		out.Bool(bool(in.Right))
	}
	{
		const prefix string = ",\"top\":"
		out.RawString(prefix)
		out.Bool(bool(in.Top))
	}
	{
		const prefix string = ",\"bottom\":"
		out.RawString(prefix)
		out.Bool(bool(in.Bottom))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdatePosPackage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeUntitledGameBackendGame(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdatePosPackage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeUntitledGameBackendGame(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdatePosPackage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeUntitledGameBackendGame(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdatePosPackage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeUntitledGameBackendGame(l, v)
}
func easyjson6601e8cdDecodeUntitledGameBackendGame1(in *jlexer.Lexer, out *UpdateClientPackage) {
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
		case "alive":
			out.Alive = bool(in.Bool())
		case "map":
			easyjson6601e8cdDecodeUntitledGameBackendGame2(in, &out.Map)
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
func easyjson6601e8cdEncodeUntitledGameBackendGame1(out *jwriter.Writer, in UpdateClientPackage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"alive\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Alive))
	}
	{
		const prefix string = ",\"map\":"
		out.RawString(prefix)
		easyjson6601e8cdEncodeUntitledGameBackendGame2(out, in.Map)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateClientPackage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeUntitledGameBackendGame1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateClientPackage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeUntitledGameBackendGame1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateClientPackage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeUntitledGameBackendGame1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateClientPackage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeUntitledGameBackendGame1(l, v)
}
func easyjson6601e8cdDecodeUntitledGameBackendGame2(in *jlexer.Lexer, out *SyncMap) {
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
		case "players":
			if in.IsNull() {
				in.Skip()
				out.Players = nil
			} else {
				in.Delim('[')
				if out.Players == nil {
					if !in.IsDelim(']') {
						out.Players = make([]*Player, 0, 8)
					} else {
						out.Players = []*Player{}
					}
				} else {
					out.Players = (out.Players)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Player
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Player)
						}
						easyjson6601e8cdDecodeUntitledGameBackendGame3(in, v1)
					}
					out.Players = append(out.Players, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "width":
			out.Width = uint(in.Uint())
		case "height":
			out.Height = uint(in.Uint())
		case "helpers":
			if in.IsNull() {
				in.Skip()
				out.Helpers = nil
			} else {
				in.Delim('[')
				if out.Helpers == nil {
					if !in.IsDelim(']') {
						out.Helpers = make([]Helper, 0, 1)
					} else {
						out.Helpers = []Helper{}
					}
				} else {
					out.Helpers = (out.Helpers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Helper
					easyjson6601e8cdDecodeUntitledGameBackendGame4(in, &v2)
					out.Helpers = append(out.Helpers, v2)
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
func easyjson6601e8cdEncodeUntitledGameBackendGame2(out *jwriter.Writer, in SyncMap) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"players\":"
		out.RawString(prefix[1:])
		if in.Players == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Players {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					easyjson6601e8cdEncodeUntitledGameBackendGame3(out, *v4)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"width\":"
		out.RawString(prefix)
		out.Uint(uint(in.Width))
	}
	{
		const prefix string = ",\"height\":"
		out.RawString(prefix)
		out.Uint(uint(in.Height))
	}
	{
		const prefix string = ",\"helpers\":"
		out.RawString(prefix)
		if in.Helpers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Helpers {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson6601e8cdEncodeUntitledGameBackendGame4(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson6601e8cdDecodeUntitledGameBackendGame4(in *jlexer.Lexer, out *Helper) {
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
		case "position":
			easyjson6601e8cdDecodeUntitledGameBackendGame5(in, &out.Position)
		case "type":
			out.Type = string(in.String())
		case "uuid":
			out.UUID = string(in.String())
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
func easyjson6601e8cdEncodeUntitledGameBackendGame4(out *jwriter.Writer, in Helper) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix[1:])
		easyjson6601e8cdEncodeUntitledGameBackendGame5(out, in.Position)
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix)
		out.String(string(in.UUID))
	}
	out.RawByte('}')
}
func easyjson6601e8cdDecodeUntitledGameBackendGame5(in *jlexer.Lexer, out *Position) {
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
			out.X = uint(in.Uint())
		case "y":
			out.Y = uint(in.Uint())
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
func easyjson6601e8cdEncodeUntitledGameBackendGame5(out *jwriter.Writer, in Position) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.X))
	}
	{
		const prefix string = ",\"y\":"
		out.RawString(prefix)
		out.Uint(uint(in.Y))
	}
	out.RawByte('}')
}
func easyjson6601e8cdDecodeUntitledGameBackendGame3(in *jlexer.Lexer, out *Player) {
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
		case "uuid":
			out.UUID = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "alive":
			out.Alive = bool(in.Bool())
		case "effects":
			if in.IsNull() {
				in.Skip()
				out.Effects = nil
			} else {
				in.Delim('[')
				if out.Effects == nil {
					if !in.IsDelim(']') {
						out.Effects = make([]string, 0, 4)
					} else {
						out.Effects = []string{}
					}
				} else {
					out.Effects = (out.Effects)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Effects = append(out.Effects, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "kills":
			out.Kills = uint(in.Uint())
		case "position":
			if in.IsNull() {
				in.Skip()
				out.Position = nil
			} else {
				if out.Position == nil {
					out.Position = new(Position)
				}
				easyjson6601e8cdDecodeUntitledGameBackendGame5(in, out.Position)
			}
		case "trajectory":
			if in.IsNull() {
				in.Skip()
				out.Trajectory = nil
			} else {
				in.Delim('[')
				if out.Trajectory == nil {
					if !in.IsDelim(']') {
						out.Trajectory = make([]Position, 0, 4)
					} else {
						out.Trajectory = []Position{}
					}
				} else {
					out.Trajectory = (out.Trajectory)[:0]
				}
				for !in.IsDelim(']') {
					var v8 Position
					easyjson6601e8cdDecodeUntitledGameBackendGame5(in, &v8)
					out.Trajectory = append(out.Trajectory, v8)
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
func easyjson6601e8cdEncodeUntitledGameBackendGame3(out *jwriter.Writer, in Player) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"alive\":"
		out.RawString(prefix)
		out.Bool(bool(in.Alive))
	}
	{
		const prefix string = ",\"effects\":"
		out.RawString(prefix)
		if in.Effects == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Effects {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"kills\":"
		out.RawString(prefix)
		out.Uint(uint(in.Kills))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		if in.Position == nil {
			out.RawString("null")
		} else {
			easyjson6601e8cdEncodeUntitledGameBackendGame5(out, *in.Position)
		}
	}
	{
		const prefix string = ",\"trajectory\":"
		out.RawString(prefix)
		if in.Trajectory == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Trajectory {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjson6601e8cdEncodeUntitledGameBackendGame5(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}