package isikukood

import (
	"fmt"
	"strconv"
	"time"
)

type Isikukood struct {
	code string
}

func (i Isikukood) Code() string {
	return i.code
}

func (i Isikukood) Validate() bool {
	lastdigit, _ := strconv.Atoi(string(i.code[len(i.code)-1]))
	return ((len(i.code) == 11) && (checksum(i) == lastdigit))
}

func (i Isikukood) BirthPlace() string {

	bi, _ := strconv.Atoi(string(i.code[7:10]))

	if i.Birthday().Year() >= 2013 {
		return "Don't know"
	}

	if bi >= 1 && bi <= 10 {
		return "Kuressaare Haigla"
	}
	if bi >= 11 && bi <= 19 {
		return "Tartu Ülikooli Naistekliinik, Tartumaa, Tartu"
	}
	if bi >= 21 && bi <= 220 {
		return "Ida-Tallinna Keskhaigla, Pelgulinna sünnitusmaja, Hiiumaa, Keila, Rapla haigla"
	}
	if bi >= 221 && bi <= 270 {
		return "Ida-Viru Keskhaigla (Kohtla-Järve, endine Jõhvi)"
	}
	if bi >= 271 && bi <= 370 {
		return "Maarjamõisa Kliinikum (Tartu), Jõgeva Haigla"
	}
	if bi >= 371 && bi <= 420 {
		return "Narva Haigla"
	}
	if bi >= 421 && bi <= 470 {
		return "Pärnu Haigla"
	}
	if bi >= 471 && bi <= 490 {
		return "Pelgulinna Sünnitusmaja (Tallinn), Haapsalu haigla"
	}
	if bi >= 491 && bi <= 520 {
		return "Järvamaa Haigla (Paide)"
	}
	if bi >= 521 && bi <= 570 {
		return "Rakvere, Tapa haigla"
	}
	if bi >= 571 && bi <= 600 {
		return "Valga Haigla"
	}
	if bi >= 601 && bi <= 650 {
		return "Viljandi Haigla"
	}
	if bi >= 651 && bi <= 710 {
		return "Lõuna-Eesti Haigla (Võru), Põlva Haigla"
	}

	return "_"
}

func (i Isikukood) Sex() string {
	sex_identifier, err := strconv.Atoi(string(i.code[:1]))
	if err != nil {

	}
	switch sex_identifier {
	case 1, 3, 5, 7:
		return "M"
	case 2, 4, 6, 8:
		return "F"
	default:
		return "Tulnukas"
	}
}

func (i Isikukood) Birthday() time.Time {

	y, _ := strconv.Atoi(i.code[1:3])
	y += century(i)

	year := strconv.Itoa(y)
	month := i.code[3:5]
	day := i.code[5:7]

	date := fmt.Sprint(month + "/" + day + "/" + year)

	t, _ := time.Parse("01/02/2006", date)

	return t
}

/* Based on: https://et.wikipedia.org/wiki/Isikukood */
func checksum(ik Isikukood) int {
	first_degree := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	second_degree := []int{3, 4, 5, 6, 7, 8, 9, 1, 2, 3}

	checksum := 0
	checksum_numbers := [10]int{}
	for i := 0; i < len(ik.code)-1; i++ {
		n, _ := strconv.Atoi(string(ik.code[i]))
		checksum += first_degree[i] * n
		checksum_numbers[i] = checksum
	}

	checksum = checksum % 11

	if checksum%11 != 10 {
		return checksum
	}

	if checksum%11 == 10 {
		c := 0
		for i := 0; i < len(checksum_numbers); i++ {
			c += checksum_numbers[i] * second_degree[i]
		}
		checksum := c % 11
		if checksum != 10 {
			checksum = 0
			return checksum
		}
	}
	return -1
}

func century(i Isikukood) int {
	sex_identifier, _ := strconv.Atoi(string(i.code[:1]))

	switch sex_identifier {
	case 1, 2:
		return 1800
	case 3, 4:
		return 1900
	case 5, 6:
		return 2000
	case 7, 8:
		return 2100
	default:
		return 0
	}
}
