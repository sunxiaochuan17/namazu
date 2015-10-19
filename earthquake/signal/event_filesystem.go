// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signal

import "github.com/satori/go.uuid"

// implements Event
type FilesystemEvent struct {
	BasicEvent
}

type FilesystemOp int

const (
	NilFilesystemOp FilesystemOp = iota
	Read
)

func NewFilesystemEvent(entityID string, op FilesystemOp, path string, m map[string]interface{}) (Event, error) {
	action := &FilesystemEvent{}
	action.InitSignal()
	action.SetID(uuid.NewV4().String())
	action.SetEntityID(entityID)
	action.SetType("event")
	action.SetClass("FilesystemEvent")
	action.SetDeferred(true)
	opt := map[string]interface{}{
		"op":   op,
		"path": path,
	}
	for k, v := range m {
		opt[k] = v
	}
	action.SetOption(opt)
	return action, nil
}

// implements Event
func (this *FilesystemEvent) DefaultFaultAction() (Action, error) {
	return NewFilesystemFaultAction(this)
}
