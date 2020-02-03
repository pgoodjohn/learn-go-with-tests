package dictionary

type Dictionary map[string]string

const (
	ErrNotFound           = DictionaryErr("could not find the word you were looking for")
	ErrWordExists         = DictionaryErr("this word already exists")
	ErrWordDoesNotExist   = DictionaryErr("could not find the word you wanted to update")
	ErrWordAlreadyMissing = DictionaryErr("could not find the word you wanted to delete")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	if d.hasKey(word) {
		return ErrWordExists
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Update(word, definition string) error {

	if d.hasKey(word) == false {
		return ErrWordDoesNotExist
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Delete(word string) error {
	delete(d, word)

	if d.hasKey(word) == false {
		return ErrWordAlreadyMissing
	}

	return nil
}

func (d Dictionary) hasKey(word string) bool {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return false
	}
	return true
}
