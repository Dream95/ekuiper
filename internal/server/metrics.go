// Copyright 2023-2025 EMQ Technologies Co., Ltd.
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

package server

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/process"

	"github.com/lf-edge/ekuiper/v2/internal/conf"
	"github.com/lf-edge/ekuiper/v2/pkg/infra"
)

type Metrics struct {
	kp *process.Process
}

func NewMetrics() *Metrics {
	var (
		kp *process.Process
		e  error
	)
	// Maybe panic in android, so add to safe run.
	err := infra.SafeRun(func() error {
		kp, e = process.NewProcess(int32(os.Getpid()))
		return e
	})
	if err != nil {
		conf.Log.Warnf("Can not initialize process for ekuiperd : %v", err)
	}
	return &Metrics{kp: kp}
}

func (m *Metrics) GetCpuUsage() string {
	if m.kp == nil {
		return ""
	}
	percent, _ := m.kp.CPUPercent()
	value := fmt.Sprintf("%.2f%%", percent)
	return value
}

func (m *Metrics) GetMemoryUsage() string {
	if m.kp == nil {
		return ""
	}
	mInfo, _ := m.kp.MemoryInfo()
	used := mInfo.RSS
	value := fmt.Sprintf("%d", used)
	return value
}
