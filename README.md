# List of English Words

![tests](https://github.com/gouniverse/enwords/workflows/tests/badge.svg)

This is a small package which returns a simple list of the English words as a ready to use array

Also contains a separate list of profane words, which can be useful to blacklist posts for review

## Installation

```golang
go get github.com/gouniverse/enwords
```

## Usage

1. To get a list of the regular English words

```golang
englishWordsList := enwords.EnglishWords()
```


2. To get a list of the profane English words

```golang
englishWordsProfaneList := enwords.EnglishWordsProfane()
```

## Origin

https://github.com/dwyl/english-words - regular English words

https://github.com/zacanger/profane-words - profane English words
