package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!global`:         `global`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!if`:             `if`,
	`echo`:            `out`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`err`:             `err`,
	`event`:           `event`,
	`pt`:              `pt`,
	`swivel-table`:    `swivel-table`,
	`getfile`:         `getfile`,
	`murex-docs`:      `murex-docs`,
	`brace-quote`:     `brace-quote`,
	`tout`:            `tout`,
	`>>`:              `>>`,
	`tread`:           `tread`,
	`if`:              `if`,
	`append`:          `append`,
	`prepend`:         `prepend`,
	`catch`:           `catch`,
	`rx`:              `rx`,
	`or`:              `or`,
	`alter`:           `alter`,
	`>`:               `>`,
	`g`:               `g`,
	`export`:          `export`,
	`set`:             `set`,
	`swivel-datatype`: `swivel-datatype`,
	`out`:             `out`,
	`ttyfd`:           `ttyfd`,
	`read`:            `read`,
	`trypipe`:         `trypipe`,
	`get`:             `get`,
	`post`:            `post`,
	`try`:             `try`,
	`global`:          `global`,
	`f`:               `f`,
	`and`:             `and`,
}
