// Code generated from SOQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SOQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSOQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSOQLVisitor) VisitKeywords_alias_allowed(ctx *Keywords_alias_allowedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitKeywords_name_allowed(ctx *Keywords_name_allowedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitObject_name(ctx *Object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitField_name(ctx *Field_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFilter_scope_name(ctx *Filter_scope_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitData_category_group_name(ctx *Data_category_group_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitData_category_name(ctx *Data_category_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitAlias_name(ctx *Alias_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitDate_formula_literal(ctx *Date_formula_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitDate_formula_n_literal_name(ctx *Date_formula_n_literal_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitDate_formula_n_literal(ctx *Date_formula_n_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitDatetime_literal(ctx *Datetime_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitDate_literal(ctx *Date_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitInteger_literal(ctx *Integer_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitReal_literal(ctx *Real_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitBoolean_literal(ctx *Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitNull_literal(ctx *Null_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_date(ctx *Function_dateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_aggregate(ctx *Function_aggregateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_location(ctx *Function_locationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_other(ctx *Function_otherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSoql_query(ctx *Soql_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitGroupby_clause(ctx *Groupby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOrderby_clause(ctx *Orderby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOffset_clause(ctx *Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFor_clause(ctx *For_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitUpdate_clause(ctx *Update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSoql_subquery(ctx *Soql_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSubquery_select_clause(ctx *Subquery_select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSelect_spec(ctx *Select_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSubquery_select_spec(ctx *Subquery_select_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitField_spec(ctx *Field_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_call_spec(ctx *Function_call_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_parameter_list(ctx *Function_parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFunction_parameter(ctx *Function_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitTypeof_spec(ctx *Typeof_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitTypeof_when_then_clause_list(ctx *Typeof_when_then_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitTypeof_when_then_clause(ctx *Typeof_when_then_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitTypeof_then_clause(ctx *Typeof_then_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitTypeof_else_clause(ctx *Typeof_else_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitField_list(ctx *Field_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitObject_spec(ctx *Object_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitObject_prefix(ctx *Object_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitComparison_operator(ctx *Comparison_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSet_operator(ctx *Set_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitCondition1(ctx *Condition1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitParenthesis(ctx *ParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSimple_condition(ctx *Simple_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitField_based_condition(ctx *Field_based_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSet_based_condition(ctx *Set_based_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitLike_based_condition(ctx *Like_based_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitCondition_field(ctx *Condition_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSet_values(ctx *Set_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitSet_value_list(ctx *Set_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitWith_plain_clause(ctx *With_plain_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitWith_data_category_clause(ctx *With_data_category_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitData_category_spec_list(ctx *Data_category_spec_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitData_category_spec(ctx *Data_category_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitData_category_parameter_list(ctx *Data_category_parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitData_category_selector(ctx *Data_category_selectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitGroup_by_plain_clause(ctx *Group_by_plain_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitGroup_by_rollup_clause(ctx *Group_by_rollup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitGroup_by_cube_clause(ctx *Group_by_cube_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitGroup_by_list(ctx *Group_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitGroup_by_spec(ctx *Group_by_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOrder_by_list(ctx *Order_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOrder_by_spec(ctx *Order_by_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOrder_by_direction_clause(ctx *Order_by_direction_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOrder_by_nulls_clause(ctx *Order_by_nulls_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitOrder_by_field(ctx *Order_by_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitFor_value(ctx *For_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOQLVisitor) VisitUpdate_value(ctx *Update_valueContext) interface{} {
	return v.VisitChildren(ctx)
}
