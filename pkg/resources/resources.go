// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"regexp"
	"time"

	"github.com/googleinterns/cloudai-gcp-test-resource-reaper/reaperconfig"
)

// A Resource represents a single GCP resource instance of any
// type supported by the Reaper.
type Resource struct {
	Name        string
	Zone        string
	TimeCreated time.Time
	Type        reaperconfig.ResourceType
}

// NewResource constructs a Resource struct.
func NewResource(name, zone string, timeCreated time.Time, resourceType reaperconfig.ResourceType) Resource {
	return Resource{name, zone, timeCreated, resourceType}
}

// TimeAlive returns how long a resource has been running.
func (resource Resource) TimeAlive() float64 {
	timeAlive := time.Since(resource.TimeCreated)
	numSeconds := timeAlive.Seconds()
	return numSeconds
}

// ShouldAddResourceToWatchlist determines whether a Resource should be watched
// by checking if its name matches the skip filter or name filter regex from the
// ResourceConfig and ReaperConfig. If a resource matches both the skip filter
// and name filter, then the skip filter wins and the resource will NOT be watched.
// An empty string for the skip filter will be interpreted as unset, and therefore
// will not match any resources.
func ShouldAddResourceToWatchlist(resource Resource, nameFilter, skipFilter string) bool {
	if len(nameFilter) == 0 {
		return false
	}
	resourceName := resource.Name
	if len(skipFilter) > 0 {
		skipMatch, err := regexp.MatchString(skipFilter, resourceName)
		if err != nil || skipMatch {
			return false
		}
	}
	nameMatch, err := regexp.MatchString(nameFilter, resourceName)
	if err != nil {
		return false
	}
	return nameMatch
}