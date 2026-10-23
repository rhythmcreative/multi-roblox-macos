//go:build darwin
// +build darwin

package ps

/*
#include <libproc.h>
#include <stdlib.h>
#include <string.h>

static int get_proc_list(pid_t **pids_out) {
    int bytes = proc_listpids(PROC_ALL_PIDS, 0, NULL, 0);
    if (bytes <= 0) return 0;

    pid_t *pids = (pid_t *)malloc(bytes);
    if (!pids) return 0;

    int actual_bytes = proc_listpids(PROC_ALL_PIDS, 0, pids, bytes);
    if (actual_bytes <= 0) {
        free(pids);
        return 0;
    }

    *pids_out = pids;
    return actual_bytes / sizeof(pid_t);
}

static int get_proc_info(pid_t pid, int *ppid_out, char *name_out, int name_len) {
    struct proc_bsdinfo info;
    int ret = proc_pidinfo(pid, PROC_PIDTBSDINFO, 0, &info, sizeof(info));
    if (ret != sizeof(info)) return -1;

    *ppid_out = info.pbi_ppid;
    if (info.pbi_name[0] != '\0') {
        strncpy(name_out, info.pbi_name, name_len - 1);
        name_out[name_len - 1] = '\0';
    } else {
        name_out[0] = '\0';
    }
    return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type DarwinProcess struct {
	pid    int
	ppid   int
	binary string
}

func (p *DarwinProcess) Pid() int {
	return p.pid
}

func (p *DarwinProcess) PPid() int {
	return p.ppid
}

func (p *DarwinProcess) Executable() string {
	return p.binary
}

func findProcess(pid int) (Process, error) {
	var ppid C.int
	var nameBuf [256]C.char

	ret := C.get_proc_info(C.pid_t(pid), &ppid, &nameBuf[0], 256)
	if ret != 0 {
		return nil, nil
	}

	return &DarwinProcess{
		pid:    pid,
		ppid:   int(ppid),
		binary: C.GoString(&nameBuf[0]),
	}, nil
}

func processes() ([]Process, error) {
	var pids *C.pid_t
	count := int(C.get_proc_list(&pids))
	if count <= 0 {
		return nil, fmt.Errorf("failed to retrieve process list")
	}
	defer C.free(unsafe.Pointer(pids))

	slice := (*[1 << 28]C.pid_t)(unsafe.Pointer(pids))[:count:count]
	procs := make([]Process, 0, count)

	var ppid C.int
	var nameBuf [256]C.char

	for _, pid := range slice {
		if pid <= 0 {
			continue
		}
		if C.get_proc_info(pid, &ppid, &nameBuf[0], 256) == 0 {
			procs = append(procs, &DarwinProcess{
				pid:    int(pid),
				ppid:   int(ppid),
				binary: C.GoString(&nameBuf[0]),
			})
		}
	}

	return procs, nil
}
