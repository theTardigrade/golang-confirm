package confirm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ask(question string, maxRepeats int) (response bool, err error) {
	reader := bufio.NewReader(os.Stdin)
	var responseText string

	if maxRepeats < 0 {
		maxRepeats = 0
	}

	question = fmt.Sprintf("%s [y/N]\n", question)

	for i := maxRepeats; i >= 0; i-- {
		fmt.Print(question)

		responseText, err = reader.ReadString('\n')
		if err != nil {
			return
		}

		responseText = strings.ToLower(strings.TrimSpace(responseText))

		if responseText == "y" || responseText == "yes" {
			response = true
			break
		}
	}

	return
}
