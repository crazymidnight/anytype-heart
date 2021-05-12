package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
)

const (
	relPbPkg   = "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	typePrefix = "_ot"
)

type Relation struct {
	Format      string   `json:"format"`
	Hidden      bool     `json:"hidden"`
	Key         string   `json:"key"`
	MaxCount    int      `json:"maxCount"`
	Name        string   `json:"name"`
	ObjectTypes []string `json:"objectTypes"`
	Readonly    bool     `json:"readonly"`
	Source      string   `json:"source"`
	Description string   `json:"description"`
}

type ObjectType struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Types       []string `json:"types"`
	Emoji       string   `json:"emoji"`
	Hidden      bool     `json:"hidden"`
	Layout      string   `json:"layout"`
	Relations   []string `json:"relations"`
	Description string   `json:"description"`
}

type Layout struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	RequiredRelations []string `json:"requiredRelations"`
}

func main() {
	err := generateRelations()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = generateTypes()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = generateLayouts()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func relConst(key string) string {
	return "RelationKey" + strings.ToUpper(key[0:1]) + key[1:]
}

func typeConst(key string) string {
	return "TypeKey" + strings.ToUpper(key[0:1]) + key[1:]
}

func sbTypeConst(key string) string {
	return "SmartBlockType_" + strings.ToUpper(key[0:1]) + key[1:]
}

func generateRelations() error {
	b, err := ioutil.ReadFile("pkg/lib/bundle/relations.json")
	if err != nil {
		return err
	}
	checkSum := sha256.Sum256(b)

	var relations []Relation
	err = json.Unmarshal(b, &relations)
	if err != nil {
		return err
	}

	f := NewFile("bundle")
	f.PackageComment("Code generated by pkg/lib/bundle/generator. DO NOT EDIT.\nsource: pkg/lib/bundle/relations.json")

	f.ImportName(relPbPkg, "model")
	f.Const().Id("RelationChecksum").Op("=").Lit(fmt.Sprintf("%x", checkSum))
	relConst := func(key string) string {
		return "RelationKey" + strings.ToUpper(key[0:1]) + key[1:]
	}

	f.Type().Id("RelationKey").String()
	f.Func().Params(
		Id("rk").Id("RelationKey"),
	).Id("String").Params().String().Block(
		Return(String().Params(Id("rk"))),
	)

	f.Const().DefsFunc(func(g *Group) {
		for _, relation := range relations {
			g.Id(relConst(relation.Key)).Id("RelationKey").Op("=").Lit(relation.Key)
		}
	})

	f.Var().DefsFunc(func(g *Group) {
		var dict = make(map[Code]Code)
		for _, relation := range relations {
			dictS := Dict{
				Id("Key"):         Lit(relation.Key),
				Id("Name"):        Lit(relation.Name),
				Id("Format"):      Qual(relPbPkg, "RelationFormat_"+relation.Format),
				Id("DataSource"):  Qual(relPbPkg, "Relation_"+relation.Source),
				Id("ReadOnly"):    Lit(relation.Readonly),
				Id("Description"): Lit(relation.Description),
				Id("Scope"):       Qual(relPbPkg, "Relation_type"),
			}
			if relation.Hidden {
				dictS[Id("Hidden")] = Lit(relation.Hidden)
			}
			if relation.MaxCount != 0 {
				dictS[Id("MaxCount")] = Lit(relation.MaxCount)
			}
			if len(relation.ObjectTypes) > 0 {
				var t []Code
				for _, ot := range relation.ObjectTypes {
					t = append(t, Id("TypePrefix").Op("+").Lit(ot))
				}
				map[Code]Code(dictS)[Id("ObjectTypes")] = Index().String().Values(t...)
			}

			dict[Id(relConst(relation.Key))] = Block(dictS)
		}
		g.Id("relations").Op("=").Map(Id("RelationKey")).Op("*").Qual(relPbPkg, "Relation").Values(Dict(dict))
	})

	file, err := os.OpenFile("pkg/lib/bundle/relation.gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	fmt.Fprintf(file, "%#v", f)
	return nil
}

func generateTypes() error {
	b, err := ioutil.ReadFile("pkg/lib/bundle/types.json")
	if err != nil {
		return err
	}
	checkSum := sha256.Sum256(b)

	var types []ObjectType
	err = json.Unmarshal(b, &types)
	if err != nil {
		return err
	}

	f := NewFile("bundle")
	f.PackageComment("Code generated by pkg/lib/bundle/generator. DO NOT EDIT.\nsource: pkg/lib/bundle/types.json")
	f.ImportName(relPbPkg, "model")

	f.Const().Id("TypeChecksum").Op("=").Lit(fmt.Sprintf("%x", checkSum))

	f.Type().Id("TypeKey").String()
	f.Func().Params(
		Id("tk").Id("TypeKey"),
	).Id("String").Params().String().Block(
		Return(String().Params(Id("tk"))),
	)

	f.Func().Params(
		Id("tk").Id("TypeKey"),
	).Id("URL").Params().String().Block(
		Return(String().Params(Id("TypePrefix").Op("+").Id("tk"))),
	)

	f.Const().DefsFunc(func(g *Group) {
		g.Id("TypePrefix").Op("=").Lit(typePrefix)
	})

	f.Const().DefsFunc(func(g *Group) {
		for _, ot := range types {
			g.Id(typeConst(ot.ID)).Id("TypeKey").Op("=").Lit(ot.ID)
		}
	})

	f.Var().DefsFunc(func(g *Group) {
		var dict = make(map[Code]Code)
		for _, ot := range types {

			dictS := Dict{
				Id("Url"):         Id("TypePrefix").Op("+").Lit(ot.ID),
				Id("Name"):        Lit(ot.Name),
				Id("Layout"):      Qual(relPbPkg, "ObjectType_"+ot.Layout),
				Id("Description"): Lit(ot.Description),
			}
			if ot.Hidden {
				dictS[Id("Hidden")] = Lit(ot.Hidden)
			}
			if ot.Emoji != "" {
				dictS[Id("IconEmoji")] = Lit(ot.Emoji)
			}

			if len(ot.Relations) > 0 {
				var t []Code
				for _, rel := range ot.Relations {
					t = append(t, Id("relations").Index(Id(relConst(rel))))
				}
				map[Code]Code(dictS)[Id("Relations")] = Index().Op("*").Qual(relPbPkg, "Relation").Values(t...)
			}
			if len(ot.Types) > 0 {
				var t []Code
				for _, sbt := range ot.Types {
					t = append(t, Qual(relPbPkg, sbTypeConst(sbt)))
				}
				dictS[Id("Types")] = Index().Qual(relPbPkg, "SmartBlockType").Values(t...)
			}

			dict[Id(typeConst(ot.ID))] = Block(dictS)
		}
		g.Id("types").Op("=").Map(Id("TypeKey")).Op("*").Qual(relPbPkg, "ObjectType").Values(Dict(dict))
	})

	file, err := os.OpenFile("pkg/lib/bundle/types.gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	fmt.Fprintf(file, "%#v", f)
	return nil
}

func generateLayouts() error {
	b, err := ioutil.ReadFile("pkg/lib/bundle/layouts.json")
	if err != nil {
		return err
	}
	checkSum := sha256.Sum256(b)

	var layouts []Layout
	err = json.Unmarshal(b, &layouts)
	if err != nil {
		return err
	}

	f := NewFile("bundle")
	f.PackageComment("Code generated by pkg/lib/bundle/generator. DO NOT EDIT.\nsource: pkg/lib/bundle/layouts.json")

	f.ImportName(relPbPkg, "model")
	f.Const().Id("LayoutChecksum").Op("=").Lit(fmt.Sprintf("%x", checkSum))

	f.Var().DefsFunc(func(g *Group) {
		var dict = make(map[Code]Code)
		for _, lt := range layouts {
			dictS := Dict{
				Id("Id"):   Qual(relPbPkg, "ObjectType_"+lt.ID),
				Id("Name"): Lit(lt.Name),
			}
			if len(lt.RequiredRelations) > 0 {
				var t []Code
				for _, rel := range lt.RequiredRelations {
					t = append(t, Id("relations").Index(Id(relConst(rel))))
				}
				map[Code]Code(dictS)[Id("RequiredRelations")] = Index().Op("*").Qual(relPbPkg, "Relation").Values(t...)
			}

			dict[Qual(relPbPkg, "ObjectType_"+lt.ID)] = Block(dictS)
		}
		g.Id("Layouts").Op("=").Map(Qual(relPbPkg, "ObjectTypeLayout")).Qual(relPbPkg, "Layout").Values(Dict(dict))
	})
	file, err := os.OpenFile("pkg/lib/bundle/layout.gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	fmt.Fprintf(file, "%#v", f)
	return nil
}
