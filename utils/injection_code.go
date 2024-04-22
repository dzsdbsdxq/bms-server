package utils

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

//@author: [LeonardWang](https://github.com/WangLeonard)
//@function: AutoInjectionCode
//@description: 向文件中固定注释位置写入代码
//@param: filepath string, funcName string, codeData string
//@return: error

const (
	startComment = "Code generated by bms-server Begin; DO NOT EDIT."
	endComment   = "Code generated by bms-server End; DO NOT EDIT."
)

//@author: [LeonardWang](https://github.com/WangLeonard)
//@function: AutoInjectionCode
//@description: 向文件中固定注释位置写入代码
//@param: filepath string, funcName string, codeData string
//@return: error

func AutoInjectionCode(filepath string, funcName string, codeData string) error {
	srcData, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	srcDataLen := len(srcData)
	fset := token.NewFileSet()
	fparser, err := parser.ParseFile(fset, filepath, srcData, parser.ParseComments)
	if err != nil {
		return err
	}
	codeData = strings.TrimSpace(codeData)
	codeStartPos := -1
	codeEndPos := srcDataLen
	var expectedFunction *ast.FuncDecl

	startCommentPos := -1
	endCommentPos := srcDataLen

	// 如果指定了函数名，先寻找对应函数
	if funcName != "" {
		for _, decl := range fparser.Decls {
			if funDecl, ok := decl.(*ast.FuncDecl); ok && funDecl.Name.Name == funcName {
				expectedFunction = funDecl
				codeStartPos = int(funDecl.Body.Lbrace)
				codeEndPos = int(funDecl.Body.Rbrace)
				break
			}
		}
	}

	// 遍历所有注释
	for _, comment := range fparser.Comments {
		if int(comment.Pos()) > codeStartPos && int(comment.End()) <= codeEndPos {
			if startComment != "" && strings.Contains(comment.Text(), startComment) {
				startCommentPos = int(comment.Pos()) // Note: Pos is the second '/'
			}
			if endComment != "" && strings.Contains(comment.Text(), endComment) {
				endCommentPos = int(comment.Pos()) // Note: Pos is the second '/'
			}
		}
	}

	if endCommentPos == srcDataLen {
		return fmt.Errorf("comment:%s not found", endComment)
	}

	// 在指定函数名，且函数中startComment和endComment都存在时，进行区间查重
	if (codeStartPos != -1 && codeEndPos <= srcDataLen) && (startCommentPos != -1 && endCommentPos != srcDataLen) && expectedFunction != nil {
		if exist := checkExist(&srcData, startCommentPos, endCommentPos, expectedFunction.Body, codeData); exist {
			fmt.Printf("文件 %s 待插入数据 %s 已存在\n", filepath, codeData)
			return nil // 这里不需要返回错误？
		}
	}

	// 两行注释中间没有换行时，会被认为是一条Comment
	if startCommentPos == endCommentPos {
		endCommentPos = startCommentPos + strings.Index(string(srcData[startCommentPos:]), endComment)
		for srcData[endCommentPos] != '/' {
			endCommentPos--
		}
	}

	// 记录"//"之前的空字符，保持写入后的格式一致
	tmpSpace := make([]byte, 0, 10)
	for tmp := endCommentPos - 2; tmp >= 0; tmp-- {
		if srcData[tmp] != '\n' {
			tmpSpace = append(tmpSpace, srcData[tmp])
		} else {
			break
		}
	}

	reverseSpace := make([]byte, 0, len(tmpSpace))
	for index := len(tmpSpace) - 1; index >= 0; index-- {
		reverseSpace = append(reverseSpace, tmpSpace[index])
	}

	// 插入数据
	indexPos := endCommentPos - 1
	insertData := []byte(append([]byte(codeData+"\n"), reverseSpace...))

	remainData := append([]byte{}, srcData[indexPos:]...)
	srcData = append(append(srcData[:indexPos], insertData...), remainData...)

	// 写回数据
	return os.WriteFile(filepath, srcData, 0o600)
}

func checkExist(srcData *[]byte, startPos int, endPos int, blockStmt *ast.BlockStmt, target string) bool {
	for _, list := range blockStmt.List {
		switch stmt := list.(type) {
		case *ast.ExprStmt:
			if callExpr, ok := stmt.X.(*ast.CallExpr); ok &&
				int(callExpr.Pos()) > startPos && int(callExpr.End()) < endPos {
				text := string((*srcData)[int(callExpr.Pos()-1):int(callExpr.End())])
				key := strings.TrimSpace(text)
				if key == target {
					return true
				}
			}
		case *ast.BlockStmt:
			if checkExist(srcData, startPos, endPos, stmt, target) {
				return true
			}
		case *ast.AssignStmt:
			// 为 model 中的代码进行检查
			if len(stmt.Rhs) > 0 {
				if callExpr, ok := stmt.Rhs[0].(*ast.CallExpr); ok {
					for _, arg := range callExpr.Args {
						if int(arg.Pos()) > startPos && int(arg.End()) < endPos {
							text := string((*srcData)[int(arg.Pos()-1):int(arg.End())])
							key := strings.TrimSpace(text)
							if key == target {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

func AutoClearCode(filepath string, codeData string) error {
	srcData, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	srcData, err = cleanCode(codeData, string(srcData))
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, srcData, 0o600)
}

func cleanCode(clearCode string, srcData string) ([]byte, error) {
	bf := make([]rune, 0, 1024)
	for i, v := range srcData {
		if v == '\n' {
			if strings.TrimSpace(string(bf)) == clearCode {
				return append([]byte(srcData[:i-len(bf)]), []byte(srcData[i+1:])...), nil
			}
			bf = (bf)[:0]
			continue
		}
		bf = append(bf, v)
	}
	return []byte(srcData), errors.New("未找到内容")
}
