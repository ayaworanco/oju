package track

import (
	"encoding/json"
	"errors"
	"flag"
	"oju/internal/utils"
)

const (
	NO_TARGET_POINTED    = "no_target_pointed"
	ESSENTIAL_DATA_EMPTY = "essential_data_empty"
)

type Track struct {
	id         string
	Resource   TrackResource     `json:"resource"`
	Target     TrackResource     `json:"target"`
	Attributes map[string]string `json:"attributes"`
}

type TrackResource struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

func (track *Track) set_id() {
	track.id = utils.GenerateId()
}

func (track Track) GetID() string {
	return track.id
}

func (track Track) Print() {
	var service string
	if flag.Lookup("test.v") == nil {
		service = get_track_target(track)
		print_track(track, service)
	}
}

func Parse(packet string) (Track, error) {
	var track Track
	unmarshal_error := json.Unmarshal([]byte(packet), &track)

	if unmarshal_error != nil {
		return Track{}, unmarshal_error
	}

	if track.Resource.Action == "" && track.Target.Action == "" && track.Resource.Name == "" {
		return Track{}, errors.New(ESSENTIAL_DATA_EMPTY)
	}

	track.set_id()
	return track, nil
}
