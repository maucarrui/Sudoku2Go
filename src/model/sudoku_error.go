package model

type SudokuErrorType int64

const (
	DEFAULT_ERROR  SudokuErrorType = -1
	INVALID_ROW    SudokuErrorType = 0
	INVALID_COLUMN SudokuErrorType = 1
	INVALID_VALUE  SudokuErrorType = 2
	REPEATED_VALUE SudokuErrorType = 3
	INVALID_SUDOKU SudokuErrorType = 4
)

type SudokuError struct {
	errorType      SudokuErrorType
	conflictRow    int
	conflictColumn int
	conflictValue  int
	errorMessage   string
	stackError     *SudokuError
}

func NewSudokuError() *SudokuError {
	sudokuError := &SudokuError{
		errorType:      -1,
		conflictRow:    0,
		conflictColumn: 0,
		conflictValue:  0,
		errorMessage:   "",
		stackError:     nil,
	}

	return sudokuError
}

func (sudokuError *SudokuError) SetErrorType(errorType SudokuErrorType) *SudokuError {
	sudokuError.errorType = errorType
	return sudokuError
}

func (sudokuError *SudokuError) SetConflictRow(conflictRow int) *SudokuError {
	sudokuError.conflictRow = conflictRow
	return sudokuError
}

func (sudokuError *SudokuError) GetConflictRow() int {
	return sudokuError.conflictRow
}

func (sudokuError *SudokuError) SetConflictColumn(conflictColumn int) *SudokuError {
	sudokuError.conflictColumn = conflictColumn
	return sudokuError
}

func (sudokuError *SudokuError) GetConflictColumn() int {
	return sudokuError.conflictColumn
}

func (sudokuError *SudokuError) SetConflictValue(conflictValue int) *SudokuError {
	sudokuError.conflictValue = conflictValue
	return sudokuError
}

func (sudokuError *SudokuError) SetErrorMessage(errorMessage string) *SudokuError {
	sudokuError.errorMessage = errorMessage
	return sudokuError
}

func (sudokuError *SudokuError) SetStackError(stackError *SudokuError) *SudokuError {
	sudokuError.stackError = stackError
	return sudokuError
}
