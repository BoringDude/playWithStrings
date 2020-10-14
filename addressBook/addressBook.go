package addressBook

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var states = map[string]string{
	"AZ": "Arizona",
	"CA": "California",
	"ID": "Idaho",
	"IN": "Indiana",
	"MA": "Massachusetts",
	"OK": "Oklahoma",
	"PA": "Pennsylvania",
	"VA": "Virginia",
}

const prefix = "..... "

type Book struct {
}
type iBook interface {
	Format(list []string) (string, error)
}

func (b *Book) Format(list []string) (string, error) {
	fmt.Println(list)
	//get parsed book in map and states list in array
	book, st, err := prepareBook(list)
	if err != nil {
		return "", err
	}

	//sort states list
	sort.Strings(st)

	var res string
	//go through sorted state array and append addresses sorted by name
	for i, v := range st {
		if i > 0 {
			res += " "
		}
		res += v + "\n"
		addr := book[v]
		sort.Strings(addr)
		for _, e := range addr {
			e = strings.ReplaceAll(e, ",", "")
			res += prefix + e + v + "\n"
		}
	}
	res = res[:len(res)-1]
	return res, nil
}
func prepareBook(list []string) (book map[string][]string, stateList []string, err error) {
	//Init book template and states list
	book = map[string][]string{}
	stateList = []string{}

	for i, v := range list {
		//Check if empty entry was provided
		if v == "" {
			continue
		}
		//check if correct state was provided
		if _, ok := states[v[len(v)-2:]]; !ok {
			return book, stateList, errors.New("bad state provided in item " + strconv.Itoa(i))
		}
		//add new state to the book map
		//if given strings' state is not there
		if _, ok := book[states[v[len(v)-2:]]]; !ok {
			book[states[v[len(v)-2:]]] = []string{v[:len(v)-2]}
			stateList = append(stateList, states[v[len(v)-2:]])
		} else {
			//if state is presented in map, just append address string
			book[states[v[len(v)-2:]]] = append(book[states[v[len(v)-2:]]], v[:len(v)-2])
		}
	}

	return
}
