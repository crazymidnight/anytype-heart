package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/addr"
	"github.com/anyproto/anytype-heart/util/strutil"

	. "github.com/dave/jennifer/jen"
)

const (
	relPbPkg  = "github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	addrPkg   = "github.com/anyproto/anytype-heart/pkg/lib/localstore/addr"
	domainPkg = "github.com/anyproto/anytype-heart/core/domain"

	pkgPrefix = "pkg/lib/bundle/"
	jsonExt   = ".json"

	typePrefix            = "_ot"
	systemRelationsName   = "systemRelations"
	internalRelationsName = "internalRelations"
	systemTypesName       = "systemTypes"
	internalTypesName     = "internalTypes"
	relationsName         = "relations"
	typesName             = "types"

	relationAssertionError = "relations validation has failed"
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
	Revision    int      `json:"revision"`
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
	Revision    int      `json:"revision"`
}

type Layout struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	RequiredRelations []string `json:"requiredRelations"`
}

func main() {
	err := generateRelations()
	exitOnError(err)

	err = generateTypes()
	exitOnError(err)

	err = generateLayouts()
	exitOnError(err)

	err = generateRelationsLists(
		internalRelationsName,
		writeInternalRelations,
		addInternalRelationsComment,
		nil,
	)
	exitOnError(err)

	err = generateRelationsLists(
		systemRelationsName,
		appendInternalToSystemRelations,
		addSystemRelationsComment,
		excludeInternalRelations,
	)

	exitOnError(err)

	err = generateTypesLists(
		internalTypesName,
		writeInternalTypes,
		addInternalTypesComment,
		nil,
	)
	exitOnError(err)

	err = generateTypesLists(
		systemTypesName,
		appendInternalToSystemTypes,
		addSystemTypesComment,
		excludeInternalTypes,
	)

	exitOnError(err)
}

func excludeInternalTypes(allSystemKeys []domain.TypeKey) []domain.TypeKey {
	var sourceName = pkgPrefix + internalTypesName + jsonExt
	internalTypeKeys, _, err := readTypes(sourceName, nil)
	exitOnError(err)

	assertTypesIncluded(
		internalTypeKeys,
		allSystemKeys,
	)

	return lo.Without(allSystemKeys, internalTypeKeys...)
}

func excludeInternalRelations(allSystemKeys []domain.RelationKey) []domain.RelationKey {
	var sourceName = pkgPrefix + internalRelationsName + jsonExt
	internalRelationKeys, _, err := readRelations(sourceName, nil)
	exitOnError(err)

	assertRelationsIncluded(
		internalRelationKeys,
		allSystemKeys,
	)

	return lo.Without(allSystemKeys, internalRelationKeys...)
}

func exitOnError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func relConst(key string) string {
	return "RelationKey" + strutil.CapitalizeFirstLetter(key)
}

func typeConst(key string) string {
	return "TypeKey" + strutil.CapitalizeFirstLetter(key)
}

func sbTypeConst(key string) string {
	return "SmartBlockType_" + strutil.CapitalizeFirstLetter(key)
}

func generateRelations() error {
	b, err := os.ReadFile("pkg/lib/bundle/relations.json")
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
	f.ImportName(domainPkg, "domain")
	f.Const().Id("RelationChecksum").Op("=").Lit(fmt.Sprintf("%x", checkSum))
	relConst := func(key string) string {
		return "RelationKey" + strings.ToUpper(key[0:1]) + key[1:]
	}

	f.Const().DefsFunc(func(g *Group) {
		for _, relation := range relations {
			g.Id(relConst(relation.Key)).Qual(domainPkg, "RelationKey").Op("=").Lit(relation.Key)
		}
	})

	f.Var().DefsFunc(func(g *Group) {
		var dict = make(map[Code]Code)
		for _, relation := range relations {
			dictS := Dict{
				Id("Id"):               Lit(addr.BundledRelationURLPrefix + relation.Key),
				Id("Key"):              Lit(relation.Key),
				Id("Name"):             Lit(relation.Name),
				Id("Format"):           Qual(relPbPkg, "RelationFormat_"+relation.Format),
				Id("DataSource"):       Qual(relPbPkg, "Relation_"+relation.Source),
				Id("ReadOnly"):         Lit(relation.Readonly),
				Id("ReadOnlyRelation"): Lit(true),
				Id("Description"):      Lit(relation.Description),
				Id("Scope"):            Qual(relPbPkg, "Relation_type"),
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
			if relation.Revision != 0 {
				dictS[Id("Revision")] = Lit(relation.Revision)
			}

			dict[Id(relConst(relation.Key))] = Block(dictS)
		}
		g.Id(relationsName).Op("=").Map(Qual(domainPkg, "RelationKey")).Op("*").Qual(relPbPkg, "Relation").Values(Dict(dict))
	})

	file, err := os.OpenFile("pkg/lib/bundle/relation.gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(file, "%#v", f)
	return nil
}

func generateTypes() error {
	b, err := os.ReadFile("pkg/lib/bundle/types.json")
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
	f.ImportName(addrPkg, "addr")
	f.ImportName(domainPkg, "domain")

	f.Const().Id("TypeChecksum").Op("=").Lit(fmt.Sprintf("%x", checkSum))

	f.Const().DefsFunc(func(g *Group) {
		g.Id("TypePrefix").Op("=").Lit(typePrefix)
	})

	f.Const().DefsFunc(func(g *Group) {
		for _, ot := range types {
			g.Id(typeConst(ot.ID)).Qual(domainPkg, "TypeKey").Op("=").Lit(ot.ID)
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

			dictS[Id("Readonly")] = Lit(true)

			if len(ot.Relations) > 0 {
				var t []Code
				var m = make(map[string]struct{}, len(ot.Relations))
				for _, rel := range ot.Relations {
					if _, exists := m[rel]; exists {
						log.Fatalf("duplicate relation '%s' for object type '%s'", rel, ot.ID)
					}
					m[rel] = struct{}{}
					t = append(t, Id("MustGetRelationLink").Add(CallFunc(func(g *Group) { g.Id(relConst(rel)) })))
				}
				map[Code]Code(dictS)[Id("RelationLinks")] = Index().Op("*").Qual(relPbPkg, "RelationLink").Values(t...)
			}
			if len(ot.Types) > 0 {
				var t []Code
				for _, sbt := range ot.Types {
					t = append(t, Qual(relPbPkg, sbTypeConst(sbt)))
				}
				dictS[Id("Types")] = Index().Qual(relPbPkg, "SmartBlockType").Values(t...)
			}
			if ot.Revision != 0 {
				dictS[Id("Revision")] = Lit(ot.Revision)
			}

			dict[Id(typeConst(ot.ID))] = Block(dictS)
		}
		g.Id("types").Op("=").Map(Qual(domainPkg, "TypeKey")).Op("*").Qual(relPbPkg, "ObjectType").Values(Dict(dict))
	})

	file, err := os.OpenFile("pkg/lib/bundle/types.gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(file, "%#v", f)
	return nil
}

func generateLayouts() error {
	b, err := os.ReadFile("pkg/lib/bundle/layouts.json")
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
					t = append(t, Id(relationsName).Index(Id(relConst(rel))))
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
	_, _ = fmt.Fprintf(file, "%#v", f)
	return nil
}

func generateRelationsLists(
	name string,
	writeRelations func(genFile *File, list []Code),
	comment func(genFile *File),
	filter func([]domain.RelationKey) []domain.RelationKey,
) error {
	var sourceName = pkgPrefix + name + jsonExt

	relationKeys, checkSum, err := readRelations(sourceName, filter)
	if err != nil {
		return err
	}

	genFile := NewFile("bundle")
	addHeader(genFile, name, sourceName, checkSum, comment)

	relations := generateRelationsList(relationKeys)
	writeRelations(genFile, relations)

	return writeGeneratedCodeToFile(name, genFile)
}

func generateTypesLists(
	name string,
	writeTypes func(genFile *File, list []Code),
	comment func(genFile *File),
	filter func([]domain.TypeKey) []domain.TypeKey,
) error {
	var sourceName = pkgPrefix + name + jsonExt

	typeKeys, checkSum, err := readTypes(sourceName, filter)
	if err != nil {
		return err
	}

	genFile := NewFile("bundle")
	addHeader(genFile, name, sourceName, checkSum, comment)

	types := generateTypesList(typeKeys)
	writeTypes(genFile, types)

	return writeGeneratedCodeToFile(name, genFile)
}

func writeGeneratedCodeToFile(name string, genFile *File) error {
	outPutFile, err := os.OpenFile(pkgPrefix+name+".gen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(outPutFile, "%#v", genFile)
	return nil
}

func readTypes(
	sourceName string,
	filter func([]domain.TypeKey) []domain.TypeKey,
) ([]domain.TypeKey, [32]byte, error) {
	bytes, err := os.ReadFile(sourceName)
	if err != nil {
		return nil, [32]byte{}, err
	}

	checkSum := sha256.Sum256(bytes)

	typesKeys, err := parseTypes(bytes, filter)
	if err != nil {
		return nil, [32]byte{}, err
	}

	return typesKeys, checkSum, err
}

func readRelations(
	sourceName string,
	filter func([]domain.RelationKey) []domain.RelationKey,
) ([]domain.RelationKey, [32]byte, error) {
	bytes, err := os.ReadFile(sourceName)
	if err != nil {
		return nil, [32]byte{}, err
	}

	checkSum := sha256.Sum256(bytes)

	relationKeys, err := parseRelations(bytes, filter)
	if err != nil {
		return nil, [32]byte{}, err
	}

	return relationKeys, checkSum, err
}

func parseTypes(bytes []byte, filter func([]domain.TypeKey) []domain.TypeKey) ([]domain.TypeKey, error) {
	var typesKeys []domain.TypeKey
	err := json.Unmarshal(bytes, &typesKeys)
	if err != nil {
		return nil, err
	}

	allTypesKeys, err := readAllTypesKeys()
	if err != nil {
		return nil, err
	}
	assertTypesIncluded(typesKeys, allTypesKeys)
	if filter != nil {
		typesKeys = filter(typesKeys)
	}
	return typesKeys, nil
}

func parseRelations(bytes []byte, filter func([]domain.RelationKey) []domain.RelationKey) ([]domain.RelationKey, error) {
	var relationKeys []domain.RelationKey
	err := json.Unmarshal(bytes, &relationKeys)
	if err != nil {
		return nil, err
	}

	allRelationKeys, err := readAllRelationKeys()
	if err != nil {
		return nil, err
	}
	assertRelationsIncluded(relationKeys, allRelationKeys)
	if filter != nil {
		relationKeys = filter(relationKeys)
	}
	return relationKeys, nil
}

func assertTypesIncluded(
	whatIncluded []domain.TypeKey,
	whereIncluded []domain.TypeKey,
) {

	err := validateRelationsIncluded(
		lo.Map(whatIncluded, typeToString()),
		lo.Map(whereIncluded, typeToString()),
	)
	if err != nil {
		exitOnError(fmt.Errorf("%s: %w", relationAssertionError, err))
	}
}

func assertRelationsIncluded(
	whatIncluded []domain.RelationKey,
	whereIncluded []domain.RelationKey,
) {

	err := validateRelationsIncluded(
		lo.Map(whatIncluded, relationToString()),
		lo.Map(whereIncluded, relationToString()),
	)
	if err != nil {
		exitOnError(fmt.Errorf("%s: %w", relationAssertionError, err))
	}
}

func relationToString() func(item domain.RelationKey, index int) string {
	return func(item domain.RelationKey, index int) string { return item.String() }
}

func typeToString() func(item domain.TypeKey, index int) string {
	return func(item domain.TypeKey, index int) string { return item.String() }
}

func addHeader(genFile *File, name string, sourceName string, checkSum [32]byte, comment func(genFile *File)) {
	genFile.PackageComment(
		"Code generated by pkg/lib/bundle/generator. DO NOT EDIT.\n" +
			"source: " + sourceName,
	)
	genFile.ImportName(relPbPkg, "model")
	writeCheckSum(genFile, name, checkSum)
	comment(genFile)
}

func writeCheckSum(genFile *File, name string, checkSum [32]byte) *Statement {
	return genFile.Const().
		Id(strutil.CapitalizeFirstLetter(name) + "Checksum").
		Op("=").
		Lit(fmt.Sprintf("%x", checkSum))
}

func generateRelationsList(relationKeys []domain.RelationKey) []Code {
	var list = make([]Code, len(relationKeys))
	for _, relationKey := range relationKeys {
		list = append(list, Line().Id(relConst(relationKey.String())))
	}
	list = append(list, Line())
	return list
}

func generateTypesList(typesKeys []domain.TypeKey) []Code {
	var list = make([]Code, len(typesKeys))
	for _, typeKey := range typesKeys {
		list = append(list, Line().Id(typeConst(typeKey.String())))
	}
	list = append(list, Line())
	return list
}

func appendInternalToSystemRelations(genFile *File, list []Code) {
	genFile.
		Var().
		Id("SystemRelations").
		Op("=").
		Append(
			Id("RequiredInternalRelations").
				Op(",").
				Index().
				Qual(domainPkg, "RelationKey").
				Values(list...).
				Op("..."),
		)
}

func appendInternalToSystemTypes(genFile *File, list []Code) {
	genFile.
		Var().
		Id("SystemTypes").
		Op("=").
		Append(
			Id("InternalTypes").
				Op(",").
				Index().
				Qual(domainPkg, "TypeKey").
				Values(list...).
				Op("..."),
		)
}

func writeInternalRelations(genFile *File, list []Code) {
	genFile.
		Var().
		Id("RequiredInternalRelations").
		Op("=").
		Index().
		Qual(domainPkg, "RelationKey").
		Values(list...)
}

func writeInternalTypes(genFile *File, list []Code) {
	genFile.
		Var().
		Id("InternalTypes").
		Op("=").
		Index().
		Qual(domainPkg, "TypeKey").
		Values(list...)
}

func addInternalRelationsComment(genFile *File) {
	genFile.Comment("RequiredInternalRelations contains internal relations that will be added to EVERY new or existing object")
	genFile.Comment("if this relation only needs SPECIFIC objects(e.g. of some type) add it to the SystemRelations")
}

func addSystemRelationsComment(genFile *File) {
	genFile.Comment("SystemRelations contains relations that have some special biz logic depends on them in some objects")
	genFile.Comment("in case EVERY object depend on the relation please add it to RequiredInternalRelations")
}

func addInternalTypesComment(genFile *File) {
	genFile.Comment("InternalTypes contains the list of types that are not possible to create directly via ObjectCreate")
	genFile.Comment("to create as a general object because they have specific logic")
}

func addSystemTypesComment(genFile *File) {
	genFile.Comment("SystemTypes contains types that have some special biz logic depends on them in some objects")
	genFile.Comment("they shouldn't be removed or edited in any way")
}

func validateRelationsIncluded(
	relationsKeysSubSet []string,
	allRelationsKeys []string,
) error {
	for _, subKey := range relationsKeysSubSet {
		if !slices.Contains(allRelationsKeys, subKey) {
			return errors.New(subKey + " is absent in relations list!")
		}
	}
	return nil
}

func readAllRelationKeys() ([]domain.RelationKey, error) {
	bytes, err := os.ReadFile(pkgPrefix + relationsName + jsonExt)
	if err != nil {
		return []domain.RelationKey{}, err
	}

	var allRelations []Relation
	err = json.Unmarshal(bytes, &allRelations)
	if err != nil {
		return []domain.RelationKey{}, err
	}
	var allRelationsKeys = lo.Map(
		allRelations,
		func(item Relation, index int) domain.RelationKey { return domain.RelationKey(item.Key) },
	)
	return allRelationsKeys, nil
}

func readAllTypesKeys() ([]domain.TypeKey, error) {
	bytes, err := os.ReadFile(pkgPrefix + typesName + jsonExt)
	if err != nil {
		return []domain.TypeKey{}, err
	}

	var allTypes []ObjectType
	err = json.Unmarshal(bytes, &allTypes)
	if err != nil {
		return []domain.TypeKey{}, err
	}
	var allTypesKey = lo.Map(
		allTypes,
		func(item ObjectType, index int) domain.TypeKey { return domain.TypeKey(item.ID) },
	)
	return allTypesKey, nil
}
