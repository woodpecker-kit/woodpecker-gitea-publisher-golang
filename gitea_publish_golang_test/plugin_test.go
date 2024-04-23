package gitea_publish_golang_test

import (
	"fmt"
	"github.com/woodpecker-kit/woodpecker-gitea-publisher-golang/gitea_publish_golang"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock GiteaPublishGolang")

	testDataPathRoot, errTestDataPathRoot := testGoldenKit.GetOrCreateTestDataFullPath("args_plugin_test")
	if errTestDataPathRoot != nil {
		t.Fatal(errTestDataPathRoot)
	}

	// successArgs
	successArgsWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "successArgs")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	successArgsSettings := mockPluginSettings()
	successArgsSettings.GiteaApiKey = "foo key"

	// notSupport
	notSupportWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "notSupport")),
		wd_mock.FastCurrentStatus("not_support"),
	)
	notSupportSettings := mockPluginSettings()

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       gitea_publish_golang.Settings
		workRoot       string

		isDryRun          bool
		wantArgFlagNotErr bool
	}{
		{
			name:              "successArgs",
			woodpeckerInfo:    successArgsWoodpeckerInfo,
			settings:          successArgsSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:           "notSupport",
			woodpeckerInfo: notSupportWoodpeckerInfo,
			settings:       notSupportSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.OnlyArgsCheck()
			errPluginRun := p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(p.GetWoodPeckerInfo())
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(p.Settings, "print Settings")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
				infoShot := p.ShortInfo()
				wd_log.VerboseJsonf(infoShot, "print WoodpeckerInfoShort")
			} else {
				if errPluginRun == nil {
					t.Fatalf("test case [ %s ], wantArgFlagNotErr %v, but p.Exec() not error", tc.name, tc.wantArgFlagNotErr)
				}
				t.Logf("check args error: %v", errPluginRun)
			}
		})
	}
}

func TestPlugin(t *testing.T) {
	t.Log("do GiteaPublishGolang")
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}
	t.Log("mock GiteaPublishGolang args")

	testDataFolderFullPath := testGoldenKit.GetTestDataFolderFullPath()
	projectRootPath := filepath.Dir(filepath.Dir(testDataFolderFullPath))

	// tagPipeline
	tagPipelineWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(projectRootPath),
		wd_mock.FastTag("v1.0.0", "new tag"),
		wd_mock.WithCiForgeInfo(
			wd_mock.WithCiForgeType(valCiForgeType),
			wd_mock.WithCiForgeUrl(valCiForgeUrl),
		),
		wd_mock.WithCiSystemInfo(
			wd_mock.WithCiSystemHost(valCiSystemHost),
			wd_mock.WithCiSystemUrl(valCiSystemUrl),
		),
		wd_mock.WithRepositoryInfo(
			wd_mock.WithCIRepoName(valCiRepoName),
			wd_mock.WithCIRepoOwner(valCiRepoOwner),
			wd_mock.WithCIRepo(fmt.Sprintf("%s/%s", valCiRepoOwner, valCiRepoName)),
		),
	)
	tagPipelineSettings := mockPluginSettings()

	// pullRequestPipeline
	pullRequestPipelineWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(projectRootPath),
		wd_mock.FastPullRequest("1", "new pr", "feature-support", "main", "main"),
		wd_mock.WithCiForgeInfo(
			wd_mock.WithCiForgeType(valCiForgeType),
			wd_mock.WithCiForgeUrl(valCiForgeUrl),
		),
		wd_mock.WithCiSystemInfo(
			wd_mock.WithCiSystemHost(valCiSystemHost),
			wd_mock.WithCiSystemUrl(valCiSystemUrl),
		),
		wd_mock.WithRepositoryInfo(
			wd_mock.WithCIRepoName(valCiRepoName),
			wd_mock.WithCIRepoOwner(valCiRepoOwner),
			wd_mock.WithCIRepo(fmt.Sprintf("%s/%s", valCiRepoOwner, valCiRepoName)),
		),
	)
	pullRequestPipelineSettings := mockPluginSettings()

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       gitea_publish_golang.Settings
		workRoot       string

		ossTransferKey  string
		ossTransferData interface{}

		isDryRun bool
		wantErr  bool
	}{
		{
			name:           "tagPipeline",
			woodpeckerInfo: tagPipelineWoodpeckerInfo,
			settings:       tagPipelineSettings,
			isDryRun:       true,
		},
		{
			name:           "pullRequestPipeline",
			woodpeckerInfo: pullRequestPipelineWoodpeckerInfo,
			settings:       pullRequestPipelineSettings,
			isDryRun:       true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.Settings.DryRun = tc.isDryRun
			if tc.ossTransferKey != "" {
				errGenTransferData := generateTransferStepsOut(
					p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}
			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
