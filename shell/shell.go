package shell

// import "io/ioutil"
import "os"

func GetCurrentWorkingDirectory() (dir string, err error) {
	return os.Getwd()
}

func GetHostname() (hostname string, err error) {
	return "hourglass", nil
}

func GetUsername() (username string, err error) {
	return "bloodred", nil
}

func GetBasicShellInfo() (shellInfo map[string]string, errs []error) {
	pwd, err_pwd := GetCurrentWorkingDirectory()
	host, err_host := GetHostname()
	user, err_user := GetUsername()
	info := map[string]string {
		"pwd": pwd,
		"host": host,
		"user": user,
	}
	return info, []error{err_pwd, err_host, err_user}
}