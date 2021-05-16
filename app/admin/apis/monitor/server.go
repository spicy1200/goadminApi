package monitor

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"go-admin/common/apis"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Monitor struct {
	apis.Api
}

// @Summary 系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/settings/serverInfo [get]
func (e Monitor) ServerInfo(c *gin.Context) {
	e.Context = c

	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()
	osDic["ip"] = pkg.GetLocaHonst()
	osDic["projectDir"] = pkg.GetCurrentPath()

	dis, _ := disk.Usage("/")
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotalGB
	diskDic["free"] = diskFreeGB

	mem, _ := mem.VirtualMemory()
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent

	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuInfo"], _ = cpu.Info()
	percent, _ := cpu.Percent(0, false)
	cpuDic["Percent"] = pkg.Round(percent[0], 2)
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	e.Custom(gin.H{
		"code": 200,
		"os":   osDic,
		"mem":  memDic,
		"cpu":  cpuDic,
		"disk": diskDic,
	})
}
