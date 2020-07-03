package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string // alias 처럼 쓰는 것임.

// Type에도 Method를 추가할 수 있음.

var (
	errNotFound   = errors.New("Word Not Found")
	errWordExists = errors.New("That Word already exists!")
	errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant Delete non-existing word")
)

// Search Method
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound // string, error를 return해야 하므로 ""도 넣어줘야.
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// err 가 발생한다면,, Dictionary에 해당 단어가 없다는 말.
	// if 구문을 통해 구현
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil

	// switch를 통해 구현
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil

}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete

	}
	return nil

}
