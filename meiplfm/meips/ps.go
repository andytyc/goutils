package meips

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// GetCPUPercent CPU使用率
func GetCPUPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

// GetMemPercent 内存使用率
func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

// GetDiskPercent 磁盘使用率
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

//HardDiskUsege 磁盘使用率
func HardDiskUsege() float64 {
	hardDisk, _ := disk.Usage("/")
	return hardDisk.UsedPercent
}

// GetHostInfo 主机信息
func GetHostInfo() *host.InfoStat {
	info, _ := host.Info()
	return info
}

// GetCPUInfo CPU信息
func GetCPUInfo() []cpu.InfoStat {
	info, _ := cpu.Info()
	return info
}

// GetCPULogicalCount CPU逻辑数量
func GetCPULogicalCount() int {
	count, _ := cpu.Counts(true)
	return count
}

// GetCPUCoreCount CPU物理核心数量, 如果是2说明是双核超线程, 如果是4则是4核非超线程
func GetCPUCoreCount() int {
	count, _ := cpu.Counts(false)
	return count
}

// GetCPUTime 用户CPU时间/系统CPU时间/空闲时间。。。等等
// 用户CPU时间：就是用户的进程获得了CPU资源以后，在用户态执行的时间。
// 系统CPU时间：用户进程获得了CPU资源以后，在内核态的执行时间。
func GetCPUTime() []cpu.TimesStat {
	info, _ := cpu.Times(false)
	return info
}

// GetMemVirtualInfo 获取物理内存信息{总内存大小，已使用大小，使用率...}
func GetMemVirtualInfo() *mem.VirtualMemoryStat {
	info, _ := mem.VirtualMemory()
	return info
}

// GetMemSwapInfo 获取交换区内存信息
func GetMemSwapInfo() *mem.SwapMemoryStat {
	info, _ := mem.SwapMemory()
	return info
}

// GetDiskBaseAllInfo 获取磁盘分区{所有分区}、磁盘使用率和磁盘IO信息：
func GetDiskBaseAllInfo() []disk.PartitionStat {
	info, _ := disk.Partitions(true)
	return info
}

// GetDiskPathUsageInfo 获取某路径的磁盘使用情况,{disk.Usage("E:")}
func GetDiskPathUsageInfo(path string) *disk.UsageStat {
	info, _ := disk.Usage(path)
	return info
}

// GetDiskIOInfo 获取磁盘IO信息 {names为空，就是获取所有磁盘io信息}
func GetDiskIOInfo(names ...string) map[string]disk.IOCountersStat {
	info, _ := disk.IOCounters(names...)
	return info
}

// GetNetInfo 当前网络连接信息 {可填入all、tcp、udp、tcp4、udp4等等}
func GetNetInfo(kind string) []net.ConnectionStat {
	info, _ := net.Connections(kind)
	return info
}

// GetNetIOCount 获取网络读写字节／包的个数
func GetNetIOCount(kind string) []net.IOCountersStat {
	info, _ := net.IOCounters(false)
	return info
}

// ProcessGetAllInfo 获取当前所有进程的信息
func ProcessGetAllInfo() []*process.Process {
	info, _ := process.Processes()
	return info
}

// ProcessGetAllPid 获取当前所有进程的pid
func ProcessGetAllPid() []int32 {
	pids, _ := process.Pids()
	return pids
}

// ProcessPidExists 进程是否存在
func ProcessPidExists(pid int32) bool {
	exists, _ := process.PidExists(pid)
	return exists
}

// ProcessIsRunning 进程是否运行{running}
func ProcessIsRunning(pid int32) bool {
	pidInstance, _ := process.NewProcess(pid) // NewProcess仅仅为一个进程实例，用于获取该进程信息及状态
	runing, _ := pidInstance.IsRunning()
	return runing
}

// ProcessCreateTime 进程的创建时间
func ProcessCreateTime(pid int32) int64 {
	pidInstance, _ := process.NewProcess(pid) // NewProcess仅仅为一个进程实例，用于获取该进程信息及状态
	createTime, _ := pidInstance.CreateTime()
	return createTime
}

// ProcessMemoryPercent 进程的内存使用率
func ProcessMemoryPercent(pid int32) float32 {
	pidInstance, _ := process.NewProcess(pid) // NewProcess仅仅为一个进程实例，用于获取该进程信息及状态
	mempercent, _ := pidInstance.MemoryPercent()
	return mempercent
}

// ProcessCPUPercent 进程的CPU使用率
func ProcessCPUPercent(pid int32) float64 {
	pidInstance, _ := process.NewProcess(pid) // NewProcess仅仅为一个进程实例，用于获取该进程信息及状态
	cpupercent, _ := pidInstance.CPUPercent()
	return cpupercent
}

// ----------- 实战有效{下边} -----------

//DeviceHelper DeviceHelper
type DeviceHelper struct {
}

//HardDiskTotal HardDiskTotal
func (d *DeviceHelper) HardDiskTotal() int {
	hardDisk, _ := disk.Usage("/")
	return int(hardDisk.Total / 1024 / 1024 / 1024)
}

//HardDiskUsege HardDiskUsege
func (d *DeviceHelper) HardDiskUsege() float32 {
	hardDisk, _ := disk.Usage("/")
	return float32(hardDisk.UsedPercent)
}

//MemoryTotal MemoryTotal
func (d *DeviceHelper) MemoryTotal() int {
	memory, _ := mem.VirtualMemory()
	return int(memory.Total / 1024 / 1024 / 1024)
}

//MemoryUsage MemoryUsage
func (d *DeviceHelper) MemoryUsage() float32 {
	memory, _ := mem.VirtualMemory()
	return float32(memory.UsedPercent)
}

//CPUUsage CPUUsage
func (d *DeviceHelper) CPUUsage() float32 {
	cpus, _ := cpu.Percent(time.Second, false)
	var total float32
	for _, cpu := range cpus {
		total += float32(cpu)
	}
	if len(cpus) > 0 {
		return float32(total / float32(len(cpus)))
	} else {
		return 0.0
	}
}

//CPUModelName CPUModelName
func (d *DeviceHelper) CPUModelName() string {
	cpuInfo, _ := cpu.Info()
	cpuModelName := ""

	for _, subCpu := range cpuInfo {
		modelname := subCpu.ModelName
		cpuModelName = modelname
		break
	}

	return cpuModelName
}
