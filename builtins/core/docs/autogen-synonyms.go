package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!set`:            `set`,
	`!event`:          `event`,
	`echo`:            `out`,
	`unset`:           `unset`,
	`prepend`:         `prepend`,
	`try`:             `try`,
	`trypipe`:         `trypipe`,
	`set`:             `set`,
	`alter`:           `alter`,
	`tout`:            `tout`,
	`pt`:              `pt`,
	`swivel-table`:    `swivel-table`,
	`>>`:              `>>`,
	`g`:               `g`,
	`out`:             `out`,
	`get`:             `get`,
	`err`:             `err`,
	`f`:               `f`,
	`catch`:           `catch`,
	`append`:          `append`,
	`swivel-datatype`: `swivel-datatype`,
	`murex-docs`:      `murex-docs`,
	`getfile`:         `getfile`,
	`post`:            `post`,
	`>`:               `>`,
	`ttyfd`:           `ttyfd`,
	`rx`:              `rx`,
	`if`:              `if`,
	`event`:           `event`,
	`print`:           `print`,
}
