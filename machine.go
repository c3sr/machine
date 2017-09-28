package machine

import (
	"encoding/json"

	proto "github.com/gogo/protobuf/proto"
)

// JSONMarshal marshals the Machine info as JSON and returns the result.
func (s *Machine) JSONMarshal() (p []byte, err error) {
	return json.Marshal(s)
}

// JSONUnmarshal unmarshals the received bytes into Machine.
func JSONUnmarshal(p []byte) (s *Machine, err error) {
	s = &Machine{}
	err = json.Unmarshal(p, &s)
	return s, err
}

// ProtoMarshal marshals the Machine info as Protobuf and returns the result.
func (s *Machine) ProtoMarshal() (p []byte, err error) {
	return proto.Marshal(s)
}

// ProtoUnmarshal unmarshals the received bytes into Machine.
func ProtoUnmarshal(p []byte) (s *Machine, err error) {
	s = &Machine{}
	err = proto.Unmarshal(p, s)
	return s, err
}
