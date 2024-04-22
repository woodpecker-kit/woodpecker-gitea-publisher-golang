package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestPanicGlobal_init_without_package_json(t *testing.T) {
	// mock TestPanicGlobal_init_without_package_json

	errString := "pkg_kit must use InitPkgJsonContent(content), then use"

	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicGlobal_init_without_package_json
		app := cli.NewApp()
		app.Before = GlobalBeforeAction
		app.Action = GlobalAction
		app.After = GlobalAfterAction
		err := app.Run([]string{"-h"})
		assert.Nil(t, err)
	}) {
		// verify TestPanicGlobal_init_without_package_json
		t.Fatalf("TestPanicGlobal_init_without_package_json should panic")
	}
}

func TestPanicGlobalAction(t *testing.T) {
	// mock TestPanicGlobalAction

	errString := "must success run GlobalBeforeAction then run GlobalAction"

	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicGlobalAction
		app := cli.NewApp()
		app.Action = GlobalAction
		app.After = GlobalAfterAction
		err := app.Run([]string{"-h"})
		assert.Nil(t, err)
	}) {
		// verify TestPanicGlobalAction
		t.Fatalf("TestPanicGlobalAction should panic")
	}
}

func TestGlobalAfterAction(t *testing.T) {
	t.Logf("~> mock GlobalAfterAction")
	// mock GlobalAfterAction
	app := cli.NewApp()
	app.After = GlobalAfterAction

	t.Logf("~> do GlobalAfterAction")
	// do GlobalAfterAction
	err := app.Run([]string{"-h"})

	// verify GlobalAfterAction
	assert.Nil(t, err)
}
