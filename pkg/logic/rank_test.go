// Copyright 2016 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

// +build ignore

package logic_test

import (
	"github.com/guanggu-coder/ggcoder-server/pkg/logic"
	"github.com/guanggu-coder/ggcoder-server/pkg/model"
	"testing"
)

func TestGenRank(t *testing.T) {
	needRankTypes := []int{
		model.TypeTopic,
		model.TypeResource,
		model.TypeProject,
		model.TypeArticle,
		model.TypeBook,
	}

	for _, objtype := range needRankTypes {
		logic.DefaultRank.GenWeekRank(objtype)
		logic.DefaultRank.GenMonthRank(objtype)
	}
}
