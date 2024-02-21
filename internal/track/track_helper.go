package track

import (
	"fmt"
)

func get_track_target(track Track) string {
	var service string
	if track.Target.Name == "" {
		service = NO_TARGET_POINTED
	} else {
		service = track.Target.Name
	}
	return service
}

func print_track(track Track, service string) {
	fmt.Println("=> TRACK from ", track.Resource)
	fmt.Println("[id]: ", track.GetID())
	fmt.Println("[resource.name]: ", get_resource_name(track.Resource.Name))
	fmt.Println("[resource.action]: ", get_resource_name(track.Resource.Action))
	fmt.Println("[target.name]: ", get_resource_name(track.Target.Name))
	fmt.Println("[target.action]: ", get_resource_name(track.Target.Action))

	fmt.Println("[attributes]:")
	for key, value := range track.Attributes {
		fmt.Printf("\t[%s]: %s\n", key, value)
	}
}

func get_resource_name(name string) string {
	if name == "" {
		return NO_TARGET_POINTED
	}
	return name
}
