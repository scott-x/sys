# sys

### API

- `func RunCommand(cmd string) (string, error)`: run your linux command, eg: `ls -al`
- `func IsProcessRunning(serverName string) (bool, error)`:check if the process is running by given name
- `func GetPid(serverName string) (string, error)`: get the process id by given name

