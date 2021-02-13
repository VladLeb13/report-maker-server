package normalizer

import (
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/VladLeb13/report-maker-lib/datalib"
)

func Actions(data *datalib.Report) {
	normalizeEvent(&data.Events)
	normalizeHardware(&data.Hardware)
	normalizePerfomance(&data.Perfomance)
	normalizeSoftware(&data.Software)
}

func normalizeHardware(hardware *datalib.Hardware) {
	var hddsList []datalib.HDD
	for _, v := range hardware.HDDs {
		v.Size = func(sz uint64) uint64 {
			if sz > 1000000 {
				return sz / 1000000
			}
			return sz
		}(v.Size)
		hddsList = append(hddsList, v)
	}
	hardware.HDDs = hddsList

	var volumesList []datalib.Volume
	for _, v := range hardware.Volumes {
		v.FreeSpace, v.Size = func(fs uint64, s string) (uint64, string) {
			var (
				free_space  uint64
				size_string string
				norm_fs     bool
				norm_size   bool
			)
			if fs > 1000000 {
				free_space = fs / 1000000
				norm_fs = true
			}

			size, _ := strconv.ParseUint(s, 10, 64)

			if size > 1000000 {
				size_string = strconv.Itoa(int(size / 1000000))
				norm_size = true
			}

			if norm_fs && norm_size {
				return free_space, size_string
			}
			return fs, s
		}(v.FreeSpace, v.Size)
		volumesList = append(volumesList, v)
	}
	hardware.Volumes = volumesList

	var nicList []datalib.NIC
	for _, v := range hardware.NICs {
		if v.IPAddress != nil {
			nicList = append(nicList, v)
		}
	}
	hardware.NICs = nicList

	var ramList []datalib.RAM
	for _, v := range hardware.RAMs {
		if v.Capacity > 512 {
			v.Capacity = func(sz uint64) uint64 {
				return sz / 1000000000
			}(v.Capacity)
		}
		ramList = append(ramList, v)
	}
	hardware.RAMs = ramList
}

func normalizePerfomance(perfomance *datalib.Perfomance) {
	//TODO: что ниб придумать оригинальное
}

type EventList []datalib.Event

func (list EventList) Len() int {
	return len(list)
}
func (list EventList) Less(i, j int) bool {
	return list[i].Message == list[j].Message
}
func (list EventList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func normalizeEvent(events *datalib.Events) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	sl := events.List

	for i := 0; i < len(sl); i++ {
		sort.Sort(EventList(sl))
	}

	var (
		normlist   []datalib.Event
		oldmessage string
	)
	for _, v := range sl {
		if v.Message != oldmessage {
			if _, err := time.Parse(time.RFC3339, v.TimeWritten); err != nil {
				v.TimeWritten = normDate(v.TimeWritten, "20060102150405")
			}
			normlist = append(normlist, v)
		}
		oldmessage = v.Message
	}
	events.List = normlist
}

type ProgramList []datalib.Program

func (list ProgramList) Len() int {
	return len(list)
}
func (list ProgramList) Less(i, j int) bool {
	return list[i].Name == list[j].Name
}
func (list ProgramList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func normalizeSoftware(software *datalib.Software) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	os := software.OS
	if _, err := time.Parse(time.RFC3339, os.InstallDate); err != nil {
		os.InstallDate = normDate(os.InstallDate, "20060102150405")
	}
	software.OS = os

	var updateslist []datalib.Update
	for _, v := range software.Updates {
		if _, err := time.Parse(time.RFC3339, v.InstalledOn); err != nil {
			v.InstalledOn = normDate(v.InstalledOn, "01/02/2006")
		}
		updateslist = append(updateslist, v)
	}
	software.Updates = updateslist

	sl := software.Programs
	for i := 0; i < len(sl); i++ {
		sort.Sort(ProgramList(sl))
	}

	var (
		normlist []datalib.Program
		oldname  string
		oldv     string
	)
	for _, v := range sl {
		if v.Name != oldname && v.Version != oldv {
			if _, err := time.Parse(time.RFC3339, v.InstallDate); err != nil {
				v.InstallDate = normDate(v.InstallDate, "20060102")
			}
			normlist = append(normlist, v)
		}
		oldname = v.Name
		oldv = v.Version
	}
	software.Programs = normlist
}
