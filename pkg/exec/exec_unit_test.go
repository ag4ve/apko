// Copyright 2022, 2023 Chainguard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func testLogger() *logrus.Entry {
	return logrus.NewEntry(&logrus.Logger{})
}

func TestNew(t *testing.T) {
	badOpt := func(e *Executor) error {
		return fmt.Errorf("synth error")
	}
	e, err := New(".", testLogger())
	require.NotNil(t, e)
	require.NoError(t, err)

	// Option fails
	e, err = New(".", testLogger(), badOpt)
	require.Nil(t, e)
	require.Error(t, err)
}
