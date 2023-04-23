// name    :	main_test.go
// version :	0.0.1
// date    :	20230423
// author  :	Leam Hall
// desc    :    Test wc program

package main

import (
    "bytes"
    "testing"
)

func TestCountWords(t *testing.T) {
    b := bytes.NewBufferString("Al and Wilbur,\nAlone at last,\nsort of\n")
    exp := 8
    _, res, _ := count(b)
    if res != exp {
        t.Errorf("Expected %d, got %d\n", exp, res)
    }
}

func TestCountLines(t *testing.T) {
    b := bytes.NewBufferString("Al and Wilbur,\nAlone at last,\nsort of\n")
    exp := 3
    res, _, _ := count(b)
    if res != exp {
        t.Errorf("Expected %d, got %d\n", exp, res)
    }
}

func TestCountCharacters(t *testing.T) {
    b := bytes.NewBufferString("Al and Wilbur,\nAlone at last,\nsort of\n")
    exp :=30 
    _, _, res := count(b)
    if res != exp {
        t.Errorf("Expected %d, got %d\n", exp, res)
    }
}
