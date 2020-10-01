package grules

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"

	"../data"
)

var Engine = engine.NewGruleEngine()

func GetKnowledgeBase() *ast.KnowledgeBase {
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	// Add the rule definition above into the library and name it 'TutorialRules'  version '0.0.1'
	byteArr := pkg.NewBytesResource([]byte(data.Drls))
	err := ruleBuilder.BuildRuleFromResource("HealthRules", "0.0.1", byteArr)
	if err != nil {
		panic(err)
	}

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("HealthRules", "0.0.1")
	return knowledgeBase
}

func GetDataContext() ast.IDataContext {
	dataCtx := ast.NewDataContext()
	for i, _:=range data.People {
		err := dataCtx.Add("PERSON", &data.People[i])
		if err != nil {
			panic(err)
		}
	}
	return dataCtx
}
