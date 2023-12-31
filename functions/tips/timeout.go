// Copyright 2019 Google LLC
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

// [START functions_concepts_after_timeout]

// Package tips contains tips for writing Cloud Functions in Go.
package tips

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("Timeout", Timeout)
}

// Timeout sleeps for 2 minutes and may time out before finishing.
func Timeout(w http.ResponseWriter, r *http.Request) {
	log.Println("Function execution started...")
	time.Sleep(2 * time.Minute)
	log.Println("Function completed!")
	fmt.Fprintln(w, "Function completed!")
}

// [END functions_concepts_after_timeout]
