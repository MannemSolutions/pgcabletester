package internal

import (
	"io/ioutil"
	"log"
	"os"
)

func string2File(data string) (fileName string) {
	if data == "" {
	} else if tmpFile, err := ioutil.TempFile("", "pgQuartsInlineCommand"); err != nil {
		log.Panicf("error creating tempfile: %e", err)
	} else if _, err = tmpFile.WriteString(data); err != nil {
		log.Panicf("error writing contents to tempfile: %e", err)
	} else if err = tmpFile.Close(); err != nil {
		log.Panicf("error closing the tmpfile: %e", err)
	} else if err = os.Chmod(tmpFile.Name(), 0600); err != nil {
		log.Panicf("error making inline tempfile script executable: %e", err)
	} else {
		return tmpFile.Name()
	}
	// Data was "" (log.Panicf would not return to end up here)
	return ""
}
