package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`echo`:            `out`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!set`:            `set`,
	`!event`:          `event`,
	`prepend`:         `prepend`,
	`out`:             `out`,
	`print`:           `print`,
	`ttyfd`:           `ttyfd`,
	`event`:           `event`,
	`get`:             `get`,
	`>>`:              `>>`,
	`pt`:              `pt`,
	`set`:             `set`,
	`unset`:           `unset`,
	`swivel-datatype`: `swivel-datatype`,
	`post`:            `post`,
	`catch`:           `catch`,
	`trypipe`:         `trypipe`,
	`err`:             `err`,
	`>`:               `>`,
	`f`:               `f`,
	`try`:             `try`,
	`rx`:              `rx`,
	`alter`:           `alter`,
	`append`:          `append`,
	`if`:              `if`,
	`swivel-table`:    `swivel-table`,
	`murex-docs`:      `murex-docs`,
	`getfile`:         `getfile`,
	`tout`:            `tout`,
	`g`:               `g`,
}
