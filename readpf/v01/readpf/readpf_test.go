package readpf

import (
  "testing"
)

func TestParser(t *testing.T) {

  script := `
    CCCC 0001
    DDDD 0004
    AAAA 0001 0002 0003 //comment 
    AAAA 0004 0005 0006

    // comment comment
    BBBB 0006 9999
  `

  sc, rerr := ReadFromString(script)
  if rerr != nil {
    t.Fatalf(rerr.Message)
  }

  stmts, perr := Parse(sc)
  if perr != nil {
    t.Fatalf(perr.Message)
  }

  if len(stmts) != 5 {
    t.Fatalf("statements is not expected length")
  }


  if s, ok := stmts[0].(CCCCStatement) ; ok {
    if s.InputX != "0001" {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 2 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[1].(DDDDStatement) ; ok {
    if s.InputX != "0004" {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 3 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[2].(AAAAStatement) ; ok {
    if (s.InputX != "0001") || (s.InputY != "0002") || (s.InputZ != "0003") {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 4 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[3].(AAAAStatement) ; ok {
    if (s.InputX != "0004") || (s.InputY != "0005") || (s.InputZ != "0006") {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 5 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

  if s, ok := stmts[4].(BBBBStatement) ; ok {
    if (s.InputX != "0006") || (s.InputY != "9999") {
      t.Fatalf("statements is not expected parameters")
    }
    if s.LineNo != 8 {
      t.Fatalf("statements is not expected lineno")
    }
  } else {
    t.Fatalf("statements is not expected type")
  }

}

