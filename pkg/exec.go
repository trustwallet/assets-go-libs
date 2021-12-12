package pkg

import (
	"os/exec"
)

func RunCmd(cmd string, shell bool) ([]byte, error) {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return nil, err
		}

		return out, nil
	}

	out, err := exec.Command(cmd).Output()
	if err != nil {
		return nil, err
	}

	return out, nil
}
