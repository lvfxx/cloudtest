// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build failed_limit

package sample

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestPass1(t *testing.T) {
	logrus.Infof("Passed test")
}

func TestPass2(t *testing.T) {
	logrus.Infof("Passed test")
}

func TestPass3(t *testing.T) {
	logrus.Infof("Passed test")
}

func TestFail1(t *testing.T) {
	logrus.Infof("Failed test")
	t.FailNow()
}

func TestFail2(t *testing.T) {
	logrus.Infof("Failed test")
	t.FailNow()
}

func TestFail3(t *testing.T) {
	logrus.Infof("Failed test")
	t.FailNow()
}