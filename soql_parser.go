package main

import "github.com/tzmfreedom/soql-parser/parser"

type SoqlParser struct {
	*parser.BaseSOQLVisitor
}

func (v *SoqlParser) VisitKeywords_alias_allowed(ctx *parser.Keywords_alias_allowedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitKeywords_name_allowed(ctx *parser.Keywords_name_allowedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitName(ctx *parser.NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitObject_name(ctx *parser.Object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitField_name(ctx *parser.Field_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFilter_scope_name(ctx *parser.Filter_scope_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitData_category_group_name(ctx *parser.Data_category_group_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitData_category_name(ctx *parser.Data_category_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitAlias_name(ctx *parser.Alias_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitAlias(ctx *parser.AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitDate_formula_literal(ctx *parser.Date_formula_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitDate_formula_n_literal_name(ctx *parser.Date_formula_n_literal_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitDate_formula_n_literal(ctx *parser.Date_formula_n_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitDatetime_literal(ctx *parser.Datetime_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitDate_literal(ctx *parser.Date_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitInteger_literal(ctx *parser.Integer_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitReal_literal(ctx *parser.Real_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitString_literal(ctx *parser.String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitBoolean_literal(ctx *parser.Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitNull_literal(ctx *parser.Null_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_name(ctx *parser.Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_date(ctx *parser.Function_dateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_aggregate(ctx *parser.Function_aggregateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_location(ctx *parser.Function_locationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_other(ctx *parser.Function_otherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSoql_query(ctx *parser.Soql_queryContext) interface{} {
	fields := ctx.Select_clause().Accept(v)
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSelect_clause(ctx *parser.Select_clauseContext) interface{} {
	fields := make([]string, len(ctx.AllSelect_spec()))
	for i, spec := range ctx.AllSelect_spec() {
		fields[i] = spec.Accept(v).(string)
	}
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFrom_clause(ctx *parser.From_clauseContext) interface{} {
	objects := make([]string, len(ctx.AllObject_spec()))
	for i, object := range ctx.AllObject_spec() {
		objects[i] = object.Accept(v).(string)
	}
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitUsing_clause(ctx *parser.Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitWhere_clause(ctx *parser.Where_clauseContext) interface{} {
	return ctx.Condition().Accept(v)
}

func (v *SoqlParser) VisitGroupby_clause(ctx *parser.Groupby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitHaving_clause(ctx *parser.Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOrderby_clause(ctx *parser.Orderby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitLimit_clause(ctx *parser.Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOffset_clause(ctx *parser.Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFor_clause(ctx *parser.For_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitUpdate_clause(ctx *parser.Update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSoql_subquery(ctx *parser.Soql_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSubquery_select_clause(ctx *parser.Subquery_select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSelect_spec(ctx *parser.Select_specContext) interface{} {
	if ctx.Field_spec() != nil {
		ctx.Field_spec().Accept(v)
	} else if ctx.Function_call_spec() != nil {
		ctx.Function_call_spec().Accept(v)
	} else if ctx.Soql_subquery() != nil {
		ctx.Soql_subquery().Accept(v)
	} else if ctx.Typeof_spec() != nil {
		ctx.Typeof_spec().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSubquery_select_spec(ctx *parser.Subquery_select_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitField_spec(ctx *parser.Field_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_call_spec(ctx *parser.Function_call_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitField(ctx *parser.FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_call(ctx *parser.Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_parameter_list(ctx *parser.Function_parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFunction_parameter(ctx *parser.Function_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitTypeof_spec(ctx *parser.Typeof_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitTypeof_when_then_clause_list(ctx *parser.Typeof_when_then_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitTypeof_when_then_clause(ctx *parser.Typeof_when_then_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitTypeof_then_clause(ctx *parser.Typeof_then_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitTypeof_else_clause(ctx *parser.Typeof_else_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitField_list(ctx *parser.Field_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitObject_spec(ctx *parser.Object_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitObject_prefix(ctx *parser.Object_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitComparison_operator(ctx *parser.Comparison_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSet_operator(ctx *parser.Set_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitCondition(ctx *parser.ConditionContext) interface{} {
	//for i, condition := range ctx.AllCondition1() {
	//	condition.Accept(v)
	//}
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitCondition1(ctx *parser.Condition1Context) interface{} {
	if ctx.Simple_condition() != nil {
		ctx.Simple_condition().Accept(v)
	} else if ctx.Parenthesis() != nil {
		ctx.Parenthesis().Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSimple_condition(ctx *parser.Simple_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitField_based_condition(ctx *parser.Field_based_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSet_based_condition(ctx *parser.Set_based_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitLike_based_condition(ctx *parser.Like_based_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitCondition_field(ctx *parser.Condition_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSet_values(ctx *parser.Set_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitSet_value_list(ctx *parser.Set_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitWith_clause(ctx *parser.With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitWith_plain_clause(ctx *parser.With_plain_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitWith_data_category_clause(ctx *parser.With_data_category_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitData_category_spec_list(ctx *parser.Data_category_spec_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitData_category_spec(ctx *parser.Data_category_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitData_category_parameter_list(ctx *parser.Data_category_parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitData_category_selector(ctx *parser.Data_category_selectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitGroup_by_plain_clause(ctx *parser.Group_by_plain_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitGroup_by_rollup_clause(ctx *parser.Group_by_rollup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitGroup_by_cube_clause(ctx *parser.Group_by_cube_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitGroup_by_list(ctx *parser.Group_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitGroup_by_spec(ctx *parser.Group_by_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOrder_by_list(ctx *parser.Order_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOrder_by_spec(ctx *parser.Order_by_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOrder_by_direction_clause(ctx *parser.Order_by_direction_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOrder_by_nulls_clause(ctx *parser.Order_by_nulls_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitOrder_by_field(ctx *parser.Order_by_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitFor_value(ctx *parser.For_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *SoqlParser) VisitUpdate_value(ctx *parser.Update_valueContext) interface{} {
	return v.VisitChildren(ctx)
}
