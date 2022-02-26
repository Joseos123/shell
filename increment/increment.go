package main

import (
	"log"
	"os/exec"
	"os"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
)

func main() {
	// control variable for testing
	// getlatesttag := "v10.9.9"
	// map to bash: TAG -> getlatesttag, TAGY -> getlastdigit, TAGX -> incrementlastdigit, TAGA -> getfirstdigits, TAGL -> getseconddigit, TAGO -> getfirstdigit, TAGF -> finaltag

	key := os.Getenv("key")

	uuuuuuuuu := sentry.Init(sentry.ClientOptions{
		Dsn: key,
	})
	if uuuuuuuuu != nil {
		log.Fatalf("sentry.Init: %s", uuuuuuuuu)
	}

	lmao, asdf := exec.Command("git", "describe", "--abbrev=0", "--tags").Output()
	if asdf != nil {
		sentry.CaptureException(asdf)
	}

	getlatesttag := string(lmao)
	log.Println("the initial (latest) tag is: ", getlatesttag)

	getlastdigit := getlatesttag[5:6]
	getfirstdigits := getlatesttag[0:5]
	getfirstdigit := getlatesttag[1:2]
	getseconddigit := getlatesttag[3:4]
	lastdigitconvertstringtonumber, sahdiahd := strconv.Atoi(getlastdigit)
	incrementlastdigit := lastdigitconvertstringtonumber + 1
	lastdigitconvertnumbertostring := strconv.Itoa(incrementlastdigit)
	finaltag := strings.Join([]string{getfirstdigits, lastdigitconvertnumbertostring}, "")

	if sahdiahd != nil {
		sentry.CaptureException(sahdiahd)
	}	

	// if the last digit is 9, eg. v0.0.9,
	if getlastdigit == "9" {
		// make the last digit 0 (and add one to the second digit later)
		lastdigitconvertnumbertostring = "0"
		if getseconddigit == "9" {
			// if the second digit is also 9, eg. v0.9.9
			// make the second digit 0 {
			getseconddigit = "0"
			// } add one to the first digit {
			firstdigitconvertstringtonumber, wkauhfsevuiejroefw := strconv.Atoi(getfirstdigit) // Convert string to int
			newfirstdigit := firstdigitconvertstringtonumber + 1 // add one
			getfirstdigit = strconv.Itoa(newfirstdigit) // Convert int to string as per variable type
			// } result: 1.0.0
			if wkauhfsevuiejroefw != nil {
				sentry.CaptureException(wkauhfsevuiejroefw)
			}
		} else {
			// else if it is not 9, eg. v0.8.9
			// add one to the second digit {
			seconddigitconvertstringtonumber, ueworiyiou4783788 := strconv.Atoi(getseconddigit) // Convert string to int
			newseconddigit := seconddigitconvertstringtonumber + 1 // add one
			getseconddigit = strconv.Itoa(newseconddigit) // Convert int to string as per variable type
			// } result: v0.9.0
			if ueworiyiou4783788 != nil {
				sentry.CaptureException(ueworiyiou4783788)
			}
		}
		almostfinaltag := strings.Join([]string{getfirstdigit, getseconddigit, lastdigitconvertnumbertostring}, ".") //"v$TAGO.$TAGL.$TAGX"
		finaltag = strings.Join([]string{"v", almostfinaltag}, "")
	}

	// get the length of the final tag
	TAGLENGTH := len(finaltag)
	// if the length of the final tag is over 6, trim it
	if TAGLENGTH > 6 {
		log.Println("the length of the version number is over 6, will trim")
		log.Println(finaltag)
		log.Println("tag length: ", TAGLENGTH)
		finaltag = finaltag[0:6]
	}

	log.Println("the new tag is: ", finaltag)

	gitadd := exec.Command("git", "add", ".").Run()
	if gitadd != nil {
		sentry.CaptureException(gitadd)
		log.Println("there was an error when performing git add .")
	}

	gitcommit := exec.Command("git", "commit", "-m", "🫣").Run()
	if gitcommit != nil {
		sentry.CaptureException(gitcommit)
		log.Println("there was an error when performing git commit")
	}

	gittag := exec.Command("git", "tag", "-a", finaltag, "-m", "its new release time!! ✨").Run()
	if gittag != nil {
		sentry.CaptureException(gittag)
		log.Println("there was an error when performing git tag")
	}

	gitpushtag := exec.Command("git", "push", "origin", finaltag).Run()
	if gitpushtag != nil {
		sentry.CaptureException(gitpushtag)
		log.Println("there was an error when performing git push origin tag")
	}

	gitpushmain := exec.Command("git", "push", "origin", "main").Run()
	if gitpushmain != nil {
		sentry.CaptureException(gitpushmain)
		log.Println("there was an error when performing git push origin main")
	}
}