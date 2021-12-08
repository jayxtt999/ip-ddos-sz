package logging

import (
	"fmt"
	"time"
	"github.com/jayxtt999/ip-ddos-sz/pkg/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.ServerSetting.RuntimeRootPath, setting.LogSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LogSetting.LogSaveName,
		time.Now().Format(setting.LogSetting.TimeFormat),
		setting.LogSetting.LogFileExt,
	)
}
