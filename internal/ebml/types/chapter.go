package types

import "github.com/google/uuid"

//Entry is one of the chapters associate with the document
type ChapterEntry struct {
	UID         uint
	flagHidden  uint
	flagDefault uint
	flagOrdered uint
	atoms       []Atom
}

//Atom is individual details about the contained chapters
type Atom struct {
	child             *Atom
	chapterUID        uint
	chapterStringUID  uint
	timeStart         uint
	timeEnd           uint
	flagHidden        uint
	flagEnabled       uint
	segmentUID        uuid.UUID
	segmentEditionUID uint
	physical          uint
	tracks            []uint
	displays          []Display
	processes         []Process
}

//Display is a collection of strings to display for each chapter
type Display struct {
	chapterString string
	languages     []string
	LanguageIETF  string
	countries     []string
}

//Process is a collection of commands belonging to an atom
type Process struct {
	codecID         uint
	processCommands []ProcessCommand
	data            []byte
}

//ProcessCommand is the definition of the individual command and when they should happen
type ProcessCommand struct {
	time uint
	data []byte
}
