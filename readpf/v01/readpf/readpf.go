package readpf

import (
  "fmt"
  "os"
  "io/ioutil"
  "regexp"
)

type ParameterFile struct {
  Line []string
}

type ReaderError struct {
  Message string
}

func ReadFromString(s string) (pf *ParameterFile, err *ReaderError) {

  pf = &ParameterFile {
    Line: []string{},
  }
  pf.Line = regexp.MustCompile(`\r\n|\n`).Split(s, -1)

  return pf, nil
}

func ReadFromFile(filepath string) (pf *ParameterFile, err *ReaderError) {

  f, e := os.OpenFile(filepath, os.O_RDONLY, 0666)

  if e != nil {
    err = &ReaderError{ Message: fmt.Sprintf("%s", e.Error()), }
    return nil, err
  }
  defer f.Close()

  b, e := ioutil.ReadAll(f)
  if e != nil {
    err = &ReaderError{ Message: fmt.Sprintf("%s", e.Error()), }
    return nil, err
  }

  return ReadFromString(string(b))
}

type ParserError struct {
  Message string
}

type Statement interface {
  GetString()  string
  GetLineNo() int
}

type AAAAStatement struct {
  InputX string
  InputY string
  InputZ string
  LineNo int
}

type BBBBStatement struct {
  InputX string
  InputY string
  LineNo int
}

type CCCCStatement struct {
  InputX string
  LineNo int
}

type DDDDStatement struct {
  InputX string
  LineNo int
}

func (a AAAAStatement) GetString() string {
  return fmt.Sprintf("AAAA %s %s %s", a.InputX, a.InputY, a.InputZ)
}

func (a AAAAStatement) GetLineNo() int {
  return a.LineNo
}

func (b BBBBStatement) GetString() string {
  return fmt.Sprintf("BBBB %s %s", b.InputX, b.InputY)
}

func (b BBBBStatement) GetLineNo() int {
  return b.LineNo
}

func (c CCCCStatement) GetString() string {
  return fmt.Sprintf("CCCC %s", c.InputX)
}

func (c CCCCStatement) GetLineNo() int {
  return c.LineNo
}

func (d DDDDStatement) GetString() string {
  return fmt.Sprintf("DDDD %s", d.InputX)
}

func (d DDDDStatement) GetLineNo() int {
  return d.LineNo
}



func Parse(pf *ParameterFile) (stmts []Statement, err *ParserError) {

  RE_ARG3  := regexp.MustCompile(`^\s*([A-Z]+)\s+(\d+)\s+(\d+)\s+(\d+)\s*(//.*)?$`)
  RE_ARG2  := regexp.MustCompile(`^\s*([A-Z]+)\s+(\d+)\s+(\d+)\s*(//.*)?$`)
  RE_ARG1  := regexp.MustCompile(`^\s*([A-Z]+)\s+(\d+)\s*(//.*)?$`)
  RE_BLANK := regexp.MustCompile(`^\s*(//.*)?$`)

  for i, s := range pf.Line {
    lineno := i + 1

    var stmt Statement
    switch {
    case RE_ARG3.MatchString(s):
      ma := RE_ARG3.FindStringSubmatch(s)
      switch ma[1] {
      case "AAAA":
        stmt = AAAAStatement{
          InputX: ma[2],
          InputY: ma[3],
          InputZ: ma[4],
          LineNo: lineno,
        }
      default:
        err = &ParserError{
          Message: fmt.Sprintf("line %d, unknown keyword `%s`", lineno, ma[1]),
        }
        return stmts, err
      }

    case RE_ARG2.MatchString(s):
      ma := RE_ARG2.FindStringSubmatch(s)
      switch ma[1] {
      case "BBBB":
        stmt = BBBBStatement{
          InputX: ma[2],
          InputY: ma[3],
          LineNo: lineno,
        }
      default:
        err = &ParserError{
          Message: fmt.Sprintf("line %d, unknown keyword `%s`", lineno, ma[1]),
        }
        return stmts, err
      }

    case RE_ARG1.MatchString(s):
      ma := RE_ARG1.FindStringSubmatch(s)
      switch ma[1] {
      case "CCCC":
        stmt = CCCCStatement{
          InputX: ma[2],
          LineNo: lineno,
        }
      case "DDDD":
        stmt = DDDDStatement{
          InputX: ma[2],
          LineNo: lineno,
        }
      default:
        err = &ParserError{
          Message: fmt.Sprintf("line %d, unknown define type `%s`", lineno, ma[1]),
        }
        return stmts, err
      }

    case RE_BLANK.MatchString(s):
      continue

    default:
      err = &ParserError{
        Message: fmt.Sprintf("line %d, unknown statement `%s`", lineno, s),
      }
      return stmts, err
    }

    stmts = append(stmts, stmt)
  }

  return stmts, nil
}


