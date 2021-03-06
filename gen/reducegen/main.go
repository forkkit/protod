package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/dave/protod/gen/behaviour"
)

const PATH = "github.com/dave/protod/delta"
const WRITE_DEFS = false

func needsSpecialCase(op1Type, op2Type behaviour.OpType) bool {
	// EDIT   - EDIT*, SET*, INSERT, MOVE, DELETE*, RENAME*
	// SET    - EDIT*, SET*, INSERT, MOVE, DELETE*, RENAME*
	// INSERT - EDIT, SET*, INSERT, MOVE*, DELETE*, RENAME
	// MOVE   - EDIT, SET, INSERT, MOVE*, DELETE*, RENAME
	// DELETE - EDIT, SET, INSERT, MOVE, DELETE, RENAME
	// RENAME - EDIT, SET, INSERT, MOVE, DELETE, RENAME
	data := map[behaviour.OpType]map[behaviour.OpType]bool{
		behaviour.EDIT: {
			behaviour.EDIT:   true, // e.g. EDIT A d1, EDIT A d2 => EDIT A d3 (use quill to merge d1 and d2)
			behaviour.SET:    true, // e.g. EDIT A, SET A => SET A
			behaviour.DELETE: true, // e.g. EDIT A, DELETE A => DELETE A
			behaviour.RENAME: true, // e.g. EDIT A, RENAME B to A => RENAME B to A
		},
		behaviour.SET: {
			behaviour.EDIT:   true, // e.g. SET A v1, EDIT A d1 => SET A v2 (use quill to calculate v2)
			behaviour.SET:    true, // e.g. SET A v1, SET A v2 => SET A v2
			behaviour.DELETE: true, // e.g. SET A, DELETE A => DELETE A
			behaviour.RENAME: true, // e.g. SET A, RENAME B to A => RENAME B to A
		},
		behaviour.INSERT: {
			behaviour.SET:    true, // e.g. INSERT A v1, SET A v2 => INSERT A v2
			behaviour.MOVE:   true, // e.g. INSERT A, MOVE A to B => INSERT B
			behaviour.DELETE: true, // e.g. INSERT A, DELETE A => null
		},
		behaviour.MOVE: {
			behaviour.MOVE:   true, // e.g. MOVE A to B, MOVE B to A => null | MOVE A to B, MOVE B to C => MOVE A to C
			behaviour.DELETE: true, // e.g. MOVE A to B, DELETE B => DELETE A
		},
	}
	if data[op1Type] != nil {
		return data[op1Type][op2Type]
	}
	return false
}

func main() {
	genGo()
	genDart()
}

func genDart() {
	var sb strings.Builder

	sb.WriteString("import 'package:protod/delta/delta.dart';\n")
	sb.WriteString("import 'package:protod/delta/delta.pb.dart' as pb;\n")
	sb.WriteString("import 'package:protod/delta/delta_reduce.dart';\n")
	sb.WriteString("\n")

	sb.WriteString("pb.Op reduceGenerated(pb.Op r, pb.Op op) {\n")
	sb.WriteString("  switch (r.type) {\n")
	for _, rOpType := range behaviour.OpTypes {
		rData := behaviour.Behaviours[rOpType]
		sb.WriteString(fmt.Sprintf("  case pb.Op_Type.%s:\n", rData.Name))
		sb.WriteString("    final rItem = item(r);\n")
		for i, rLocType := range rData.Locators {
			rLoc := behaviour.Locators[rLocType]
			var elseString = ""
			if i > 0 {
				elseString = "} else "
			}
			sb.WriteString(fmt.Sprintf("    %sif (rItem.%s()) {\n", elseString, rLoc.Dart))
			sb.WriteString("      switch (op.type) {\n")
			for _, opOpType := range behaviour.OpTypes {
				opData := behaviour.Behaviours[opOpType]
				sb.WriteString(fmt.Sprintf("      case pb.Op_Type.%s:\n", opData.Name))
				sb.WriteString("        final opItem = item(op);\n")
				for i, opLocType := range opData.Locators {
					opLoc := behaviour.Locators[opLocType]
					//opLoc.Type
					var elseString = ""
					if i > 0 {
						elseString = "} else "
					}
					sb.WriteString(fmt.Sprintf("        %sif (opItem.%s()) {\n", elseString, opLoc.Dart))
					if rLocType == opLocType && needsSpecialCase(rOpType, opOpType) {
						sb.WriteString(fmt.Sprintf("          return r%s%s%s%s(t, op);\n", rData.Name, rLoc.Name, opData.Name, opLoc.Name))
					} else {
						sb.WriteString("          return rIndependent(t, op);\n")
					}
				}
				sb.WriteString("        } else {\n")
				sb.WriteString("          throw Exception('invalid op');\n")
				sb.WriteString("        }\n")
				sb.WriteString("        break;\n")
			}
			sb.WriteString("      default:\n")
			sb.WriteString("        throw Exception('invalid op');\n")
			sb.WriteString("      }\n")
		}
		sb.WriteString("    } else {\n")
		sb.WriteString("      throw Exception('invalid op');\n")
		sb.WriteString("    }\n")
		sb.WriteString("    break;\n")
	}
	sb.WriteString("  default:\n")
	sb.WriteString("    throw Exception('invalid op');\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	if err := ioutil.WriteFile("../lib/delta/delta_reduce_generated.dart", []byte(sb.String()), 0666); err != nil {
		panic(err)
	}

}

func genGo() {
	f := NewFilePathName(PATH, "delta")
	fDefs := NewFilePathName(PATH, "delta")
	// func (r *Op) reduce(op *Op) []*Op {
	f.Comment("// reduce reduces two consecutive non-compound operations. If the operations are independent, the result will return\n// both inputs. If the operations can be reduced, the result will be a single operation. If the operations cancel each\n// other out, the result will be an empty slice.")
	f.Func().Params(
		Id("r").Op("*").Qual(PATH, "Op"),
	).Id("reduce").Params(
		Id("op").Op("*").Qual(PATH, "Op"),
	).Index().Op("*").Id("Op").Block(
		/*
			switch r.Type {
				case Op_Edit:
		*/
		Switch(Id("r").Dot("Type")).BlockFunc(func(g *Group) {
			for _, rOpType := range behaviour.OpTypes {
				rData := behaviour.Behaviours[rOpType]
				g.Case(Id(rData.Type)).Block(
					// _, tItem := r.Pop()
					List(Op("_"), Id("tItem")).Op(":=").Id("r").Dot("Pop").Call(),
					Switch(Id("tItem").Dot("V").Assert(Type())).BlockFunc(func(g *Group) {
						for _, rLocType := range rData.Locators {
							rLoc := behaviour.Locators[rLocType]
							g.Case(Op("*").Id(rLoc.Type)).Block(
								// switch op.Type {
								//	case Op_Edit:
								Switch(Id("op").Dot("Type")).BlockFunc(func(g *Group) {
									for _, opOpType := range behaviour.OpTypes {
										opData := behaviour.Behaviours[opOpType]
										g.Case(Id(opData.Type)).Block(
											// _, opItem := op.Pop()
											List(Op("_"), Id("opItem")).Op(":=").Id("op").Dot("Pop").Call(),
											Switch(Id("opItem").Dot("V").Assert(Type())).BlockFunc(func(g *Group) {
												for _, opLocType := range opData.Locators {
													opLoc := behaviour.Locators[opLocType]
													g.Case(Op("*").Id(opLoc.Type)).BlockFunc(func(g *Group) {
														if rLocType == opLocType && needsSpecialCase(rOpType, opOpType) {
															// return rEditFieldEditField(r, op, priority)
															g.Return(Id(fmt.Sprintf("r%s%s%s%s", rData.Name, rLoc.Name, opData.Name, opLoc.Name)).Call(
																Id("r"),
																Id("op"),
															))
															if WRITE_DEFS {
																fDefs.Func().Id(fmt.Sprintf("r%s%s%s%s", rData.Name, rLoc.Name, opData.Name, opLoc.Name)).Params(Id("op1"), Id("op2").Op("*").Qual(PATH, "Op")).Index().Op("*").Qual(PATH, "Op").Block(Panic(Lit("not implemented")))
															}
														} else {
															g.Return(Id("rIndependent").Call(Id("r"), Id("op")))
														}
													})
												}
												g.Default().Block(Panic(Lit("invalid op")))
											}),
										)
									}
									g.Default().Block(Panic(Lit("invalid op")))
								}),
							)
						}
						g.Default().Block(Panic(Lit("invalid op")))
					}),
				)
			}
			g.Default().Block(Panic(Lit("invalid op")))
		}),
	)

	if err := f.Save("../delta/reduce-generated.go"); err != nil {
		panic(err)
	}
	if WRITE_DEFS {
		if err := fDefs.Save("../delta/reduce-defs-generated.go"); err != nil {
			panic(err)
		}
	}
}
