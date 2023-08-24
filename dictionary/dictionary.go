package dictionary

type Dictionary map[string]string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("this word already exists")
	ErrWordDoesNotExists = DictionaryErr("this word doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] //map lookup has interesting feature -> can return twwo values, second being a boolean if a key was found successfully
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil //if word exists, then return the word and a nil error
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil: //if you get a nil error, you already have a word
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists //could have reused ErrNotFound but better to have a precise error for when an update fails, and can perform actions based on specific failure
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
