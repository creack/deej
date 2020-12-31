package deej

import (
	"github.com/getlantern/systray"

	"github.com/omriharel/deej/icon"
	"github.com/omriharel/deej/util"
)

func (d *Deej) initializeTray(onDone func()) {
	logger := d.logger.Named("tray")

	onReady := func() {
		logger.Debug("Tray instance ready.")

		systray.SetTemplateIcon(icon.DeejLogo, icon.DeejLogo)
		systray.SetTitle("deej")
		systray.SetTooltip("deej")

		editConfig := systray.AddMenuItem("Edit configuration", "Open config file with notepad")
		editConfig.SetIcon(icon.EditConfig)

		refreshSessions := systray.AddMenuItem("Re-scan audio sessions", "Manually refresh audio sessions if something's stuck")
		refreshSessions.SetIcon(icon.RefreshSessions)

		if d.version != "" {
			systray.AddSeparator()
			versionInfo := systray.AddMenuItem(d.version, "")
			versionInfo.Disable()
		}

		systray.AddSeparator()
		quit := systray.AddMenuItem("Quit", "Stop deej and quit")

		// Wait on things to happen.
		go func() {
			for {
				select {
				// Quit.
				case <-quit.ClickedCh:
					logger.Info("Quit menu item clicked, stopping.")

					d.signalStop()

				// Edit config.
				case <-editConfig.ClickedCh:
					logger.Info("Edit config menu item clicked, opening config for editing.")

					editor := util.DefaultEditor
					if err := util.OpenExternal(logger, editor, userConfigFilepath); err != nil {
						logger.Warnw("Failed to open config file for editing.", "error", err)
					}

				// Refresh sessions.
				case <-refreshSessions.ClickedCh:
					logger.Info("Refresh sessions menu item clicked, triggering session map refresh.")

					// PERFORMANCE: The reason that forcing a refresh here is okay is that users can't spam the
					//              right-click -> select-this-option sequence at a rate that's meaningful to performance.
					d.sessions.refreshSessions(true)
				}
			}
		}()

		// Actually start the main runtime.
		onDone()
	}

	onExit := func() {
		logger.Debug("Tray exited.")
	}

	// Start the tray icon.
	logger.Debug("Running in tray.")
	systray.Run(onReady, onExit)
}

func (d *Deej) stopTray() {
	d.logger.Debug("Quitting tray.")
	systray.Quit()
}
