// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

//go:build !gogit

package git

import (
	"path/filepath"
	"testing"

	"code.gitea.io/gitea/modules/setting"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetLanguageStats(t *testing.T) {
	setting.AppDataPath = t.TempDir()
	repoPath := filepath.Join(testReposDir, "language_stats_repo")
	gitRepo, err := openRepositoryWithDefaultContext(repoPath)
	require.NoError(t, err)

	defer gitRepo.Close()

	stats, err := gitRepo.GetLanguageStats("8fee858da5796dfb37704761701bb8e800ad9ef3")
	require.NoError(t, err)

	assert.Equal(t, map[string]int64{
		"Python": 134,
		"Java":   112,
	}, stats)
}

func TestMergeLanguageStats(t *testing.T) {
	assert.Equal(t, map[string]int64{
		"PHP":    1,
		"python": 10,
		"JAVA":   700,
	}, mergeLanguageStats(map[string]int64{
		"PHP":    1,
		"python": 10,
		"Java":   100,
		"java":   200,
		"JAVA":   400,
	}))
}
