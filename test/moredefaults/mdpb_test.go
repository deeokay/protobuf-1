// Code generated by protoc-gen-gogo.
// source: md.proto
// DO NOT EDIT!

/*
Package moredefaults is a generated protocol buffer package.

It is generated from these files:
	md.proto

It has these top-level messages:
	MoreDefaultsB
	MoreDefaultsA
*/
package moredefaults

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import encoding_json "encoding/json"
import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
// discarding unused import test "github.com/gogo/protobuf/test/example"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

func TestMoreDefaultsBProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsB(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &MoreDefaultsB{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMoreDefaultsAProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsA(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &MoreDefaultsA{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMoreDefaultsBJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsB(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		if _, ok := err.(*encoding_json.UnsupportedTypeError); ok {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	msg := &MoreDefaultsB{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestMoreDefaultsAJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsA(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		if _, ok := err.(*encoding_json.UnsupportedTypeError); ok {
			t.Skip(err)
		}
		t.Fatal(err)
	}
	msg := &MoreDefaultsA{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestMoreDefaultsBProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsB(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &MoreDefaultsB{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMoreDefaultsBProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsB(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &MoreDefaultsB{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMoreDefaultsAProtoText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsA(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &MoreDefaultsA{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestMoreDefaultsAProtoCompactText(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMoreDefaultsA(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &MoreDefaultsA{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
