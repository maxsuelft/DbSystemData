package disk

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/shirou/gopsutil/v4/disk"

	"dbsystemdata-backend/internal/config"
)

type DiskService struct{}

func (s *DiskService) GetDiskUsage() (*DiskUsage, error) {
	if config.GetEnv().IsCloud {
		return &DiskUsage{
			Platform:        PlatformLinux,
			TotalSpaceBytes: 100,
			UsedSpaceBytes:  0,
			FreeSpaceBytes:  100,
		}, nil
	}

	platform := s.detectPlatform()

	var path string

	if platform == PlatformWindows {
		path = "C:\\"
	} else {
		// Use dbsystemdata-data folder location for Linux (Docker)
		cfg := config.GetEnv()
		path = filepath.Dir(cfg.DataFolder) // Gets /dbsystemdata-data from /dbsystemdata-data/backups
	}

	diskUsage, err := disk.Usage(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get disk usage for path %s: %w", path, err)
	}

	return &DiskUsage{
		Platform:        platform,
		TotalSpaceBytes: int64(diskUsage.Total),
		UsedSpaceBytes:  int64(diskUsage.Used),
		FreeSpaceBytes:  int64(diskUsage.Free),
	}, nil
}

func (s *DiskService) detectPlatform() Platform {
	switch runtime.GOOS {
	case "windows":
		return PlatformWindows
	default:
		return PlatformLinux
	}
}
