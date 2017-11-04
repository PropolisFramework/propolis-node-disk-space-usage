import "syscall"
import "os"

var stat syscall.Statfs_t

wd, err := os.Getwd()

syscall.Statfs(wd, &stat)

// Available blocks * size per block = available space in bytes
fmt.Println(stat.Bavail * uint64(stat.Bsize))
