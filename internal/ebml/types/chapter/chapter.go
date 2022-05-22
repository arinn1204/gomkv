package chapter

import "github.com/google/uuid"

type Entry struct {
	uid         uint
	flagHidden  uint
	flagDefault uint
	flagOrdered uint
	atoms       []Atom
}

type Atom struct {
	child             *Atom
	chapterUid        uint
	chapterStringUid  uint
	timeStart         uint
	timeEnd           uint
	flagHidden        uint
	flagEnabled       uint
	segmentUid        uuid.UUID
	segmentEditionUid uint
	physical          uint
	track             Track
	displays          []Display
	processes         []Process
}

type Track struct {
	numbers []uint
}

type Display struct {
	chapterString string
	languages     []string
	LanguageIETF  string
	countries     []string
}

type Process struct {
	codecId         uint
	processCommands []ProcessCommand
	data            []byte
}

type ProcessCommand struct {
	time uint
	data []byte
}
