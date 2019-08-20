// Code generated from SOQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SOQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSOQLListener is a complete listener for a parse tree produced by SOQLParser.
type BaseSOQLListener struct{}

var _ SOQLListener = &BaseSOQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSOQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSOQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSOQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSOQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKeywords_alias_allowed is called when production keywords_alias_allowed is entered.
func (s *BaseSOQLListener) EnterKeywords_alias_allowed(ctx *Keywords_alias_allowedContext) {}

// ExitKeywords_alias_allowed is called when production keywords_alias_allowed is exited.
func (s *BaseSOQLListener) ExitKeywords_alias_allowed(ctx *Keywords_alias_allowedContext) {}

// EnterKeywords_name_allowed is called when production keywords_name_allowed is entered.
func (s *BaseSOQLListener) EnterKeywords_name_allowed(ctx *Keywords_name_allowedContext) {}

// ExitKeywords_name_allowed is called when production keywords_name_allowed is exited.
func (s *BaseSOQLListener) ExitKeywords_name_allowed(ctx *Keywords_name_allowedContext) {}

// EnterName is called when production name is entered.
func (s *BaseSOQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSOQLListener) ExitName(ctx *NameContext) {}

// EnterObject_name is called when production object_name is entered.
func (s *BaseSOQLListener) EnterObject_name(ctx *Object_nameContext) {}

// ExitObject_name is called when production object_name is exited.
func (s *BaseSOQLListener) ExitObject_name(ctx *Object_nameContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BaseSOQLListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BaseSOQLListener) ExitField_name(ctx *Field_nameContext) {}

// EnterFilter_scope_name is called when production filter_scope_name is entered.
func (s *BaseSOQLListener) EnterFilter_scope_name(ctx *Filter_scope_nameContext) {}

// ExitFilter_scope_name is called when production filter_scope_name is exited.
func (s *BaseSOQLListener) ExitFilter_scope_name(ctx *Filter_scope_nameContext) {}

// EnterData_category_group_name is called when production data_category_group_name is entered.
func (s *BaseSOQLListener) EnterData_category_group_name(ctx *Data_category_group_nameContext) {}

// ExitData_category_group_name is called when production data_category_group_name is exited.
func (s *BaseSOQLListener) ExitData_category_group_name(ctx *Data_category_group_nameContext) {}

// EnterData_category_name is called when production data_category_name is entered.
func (s *BaseSOQLListener) EnterData_category_name(ctx *Data_category_nameContext) {}

// ExitData_category_name is called when production data_category_name is exited.
func (s *BaseSOQLListener) ExitData_category_name(ctx *Data_category_nameContext) {}

// EnterAlias_name is called when production alias_name is entered.
func (s *BaseSOQLListener) EnterAlias_name(ctx *Alias_nameContext) {}

// ExitAlias_name is called when production alias_name is exited.
func (s *BaseSOQLListener) ExitAlias_name(ctx *Alias_nameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseSOQLListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseSOQLListener) ExitAlias(ctx *AliasContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSOQLListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSOQLListener) ExitLiteral(ctx *LiteralContext) {}

// EnterDate_formula_literal is called when production date_formula_literal is entered.
func (s *BaseSOQLListener) EnterDate_formula_literal(ctx *Date_formula_literalContext) {}

// ExitDate_formula_literal is called when production date_formula_literal is exited.
func (s *BaseSOQLListener) ExitDate_formula_literal(ctx *Date_formula_literalContext) {}

// EnterDate_formula_n_literal_name is called when production date_formula_n_literal_name is entered.
func (s *BaseSOQLListener) EnterDate_formula_n_literal_name(ctx *Date_formula_n_literal_nameContext) {}

// ExitDate_formula_n_literal_name is called when production date_formula_n_literal_name is exited.
func (s *BaseSOQLListener) ExitDate_formula_n_literal_name(ctx *Date_formula_n_literal_nameContext) {}

// EnterDate_formula_n_literal is called when production date_formula_n_literal is entered.
func (s *BaseSOQLListener) EnterDate_formula_n_literal(ctx *Date_formula_n_literalContext) {}

// ExitDate_formula_n_literal is called when production date_formula_n_literal is exited.
func (s *BaseSOQLListener) ExitDate_formula_n_literal(ctx *Date_formula_n_literalContext) {}

// EnterDatetime_literal is called when production datetime_literal is entered.
func (s *BaseSOQLListener) EnterDatetime_literal(ctx *Datetime_literalContext) {}

// ExitDatetime_literal is called when production datetime_literal is exited.
func (s *BaseSOQLListener) ExitDatetime_literal(ctx *Datetime_literalContext) {}

// EnterDate_literal is called when production date_literal is entered.
func (s *BaseSOQLListener) EnterDate_literal(ctx *Date_literalContext) {}

// ExitDate_literal is called when production date_literal is exited.
func (s *BaseSOQLListener) ExitDate_literal(ctx *Date_literalContext) {}

// EnterInteger_literal is called when production integer_literal is entered.
func (s *BaseSOQLListener) EnterInteger_literal(ctx *Integer_literalContext) {}

// ExitInteger_literal is called when production integer_literal is exited.
func (s *BaseSOQLListener) ExitInteger_literal(ctx *Integer_literalContext) {}

// EnterReal_literal is called when production real_literal is entered.
func (s *BaseSOQLListener) EnterReal_literal(ctx *Real_literalContext) {}

// ExitReal_literal is called when production real_literal is exited.
func (s *BaseSOQLListener) ExitReal_literal(ctx *Real_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseSOQLListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseSOQLListener) ExitString_literal(ctx *String_literalContext) {}

// EnterBoolean_literal is called when production boolean_literal is entered.
func (s *BaseSOQLListener) EnterBoolean_literal(ctx *Boolean_literalContext) {}

// ExitBoolean_literal is called when production boolean_literal is exited.
func (s *BaseSOQLListener) ExitBoolean_literal(ctx *Boolean_literalContext) {}

// EnterNull_literal is called when production null_literal is entered.
func (s *BaseSOQLListener) EnterNull_literal(ctx *Null_literalContext) {}

// ExitNull_literal is called when production null_literal is exited.
func (s *BaseSOQLListener) ExitNull_literal(ctx *Null_literalContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseSOQLListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseSOQLListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_date is called when production function_date is entered.
func (s *BaseSOQLListener) EnterFunction_date(ctx *Function_dateContext) {}

// ExitFunction_date is called when production function_date is exited.
func (s *BaseSOQLListener) ExitFunction_date(ctx *Function_dateContext) {}

// EnterFunction_aggregate is called when production function_aggregate is entered.
func (s *BaseSOQLListener) EnterFunction_aggregate(ctx *Function_aggregateContext) {}

// ExitFunction_aggregate is called when production function_aggregate is exited.
func (s *BaseSOQLListener) ExitFunction_aggregate(ctx *Function_aggregateContext) {}

// EnterFunction_location is called when production function_location is entered.
func (s *BaseSOQLListener) EnterFunction_location(ctx *Function_locationContext) {}

// ExitFunction_location is called when production function_location is exited.
func (s *BaseSOQLListener) ExitFunction_location(ctx *Function_locationContext) {}

// EnterFunction_other is called when production function_other is entered.
func (s *BaseSOQLListener) EnterFunction_other(ctx *Function_otherContext) {}

// ExitFunction_other is called when production function_other is exited.
func (s *BaseSOQLListener) ExitFunction_other(ctx *Function_otherContext) {}

// EnterSoql_query is called when production soql_query is entered.
func (s *BaseSOQLListener) EnterSoql_query(ctx *Soql_queryContext) {}

// ExitSoql_query is called when production soql_query is exited.
func (s *BaseSOQLListener) ExitSoql_query(ctx *Soql_queryContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseSOQLListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseSOQLListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseSOQLListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseSOQLListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseSOQLListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseSOQLListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseSOQLListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseSOQLListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterGroupby_clause is called when production groupby_clause is entered.
func (s *BaseSOQLListener) EnterGroupby_clause(ctx *Groupby_clauseContext) {}

// ExitGroupby_clause is called when production groupby_clause is exited.
func (s *BaseSOQLListener) ExitGroupby_clause(ctx *Groupby_clauseContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseSOQLListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseSOQLListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterOrderby_clause is called when production orderby_clause is entered.
func (s *BaseSOQLListener) EnterOrderby_clause(ctx *Orderby_clauseContext) {}

// ExitOrderby_clause is called when production orderby_clause is exited.
func (s *BaseSOQLListener) ExitOrderby_clause(ctx *Orderby_clauseContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseSOQLListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseSOQLListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterOffset_clause is called when production offset_clause is entered.
func (s *BaseSOQLListener) EnterOffset_clause(ctx *Offset_clauseContext) {}

// ExitOffset_clause is called when production offset_clause is exited.
func (s *BaseSOQLListener) ExitOffset_clause(ctx *Offset_clauseContext) {}

// EnterFor_clause is called when production for_clause is entered.
func (s *BaseSOQLListener) EnterFor_clause(ctx *For_clauseContext) {}

// ExitFor_clause is called when production for_clause is exited.
func (s *BaseSOQLListener) ExitFor_clause(ctx *For_clauseContext) {}

// EnterUpdate_clause is called when production update_clause is entered.
func (s *BaseSOQLListener) EnterUpdate_clause(ctx *Update_clauseContext) {}

// ExitUpdate_clause is called when production update_clause is exited.
func (s *BaseSOQLListener) ExitUpdate_clause(ctx *Update_clauseContext) {}

// EnterSoql_subquery is called when production soql_subquery is entered.
func (s *BaseSOQLListener) EnterSoql_subquery(ctx *Soql_subqueryContext) {}

// ExitSoql_subquery is called when production soql_subquery is exited.
func (s *BaseSOQLListener) ExitSoql_subquery(ctx *Soql_subqueryContext) {}

// EnterSubquery_select_clause is called when production subquery_select_clause is entered.
func (s *BaseSOQLListener) EnterSubquery_select_clause(ctx *Subquery_select_clauseContext) {}

// ExitSubquery_select_clause is called when production subquery_select_clause is exited.
func (s *BaseSOQLListener) ExitSubquery_select_clause(ctx *Subquery_select_clauseContext) {}

// EnterSelect_spec is called when production select_spec is entered.
func (s *BaseSOQLListener) EnterSelect_spec(ctx *Select_specContext) {}

// ExitSelect_spec is called when production select_spec is exited.
func (s *BaseSOQLListener) ExitSelect_spec(ctx *Select_specContext) {}

// EnterSubquery_select_spec is called when production subquery_select_spec is entered.
func (s *BaseSOQLListener) EnterSubquery_select_spec(ctx *Subquery_select_specContext) {}

// ExitSubquery_select_spec is called when production subquery_select_spec is exited.
func (s *BaseSOQLListener) ExitSubquery_select_spec(ctx *Subquery_select_specContext) {}

// EnterField_spec is called when production field_spec is entered.
func (s *BaseSOQLListener) EnterField_spec(ctx *Field_specContext) {}

// ExitField_spec is called when production field_spec is exited.
func (s *BaseSOQLListener) ExitField_spec(ctx *Field_specContext) {}

// EnterFunction_call_spec is called when production function_call_spec is entered.
func (s *BaseSOQLListener) EnterFunction_call_spec(ctx *Function_call_specContext) {}

// ExitFunction_call_spec is called when production function_call_spec is exited.
func (s *BaseSOQLListener) ExitFunction_call_spec(ctx *Function_call_specContext) {}

// EnterField is called when production field is entered.
func (s *BaseSOQLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseSOQLListener) ExitField(ctx *FieldContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseSOQLListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseSOQLListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterFunction_parameter_list is called when production function_parameter_list is entered.
func (s *BaseSOQLListener) EnterFunction_parameter_list(ctx *Function_parameter_listContext) {}

// ExitFunction_parameter_list is called when production function_parameter_list is exited.
func (s *BaseSOQLListener) ExitFunction_parameter_list(ctx *Function_parameter_listContext) {}

// EnterFunction_parameter is called when production function_parameter is entered.
func (s *BaseSOQLListener) EnterFunction_parameter(ctx *Function_parameterContext) {}

// ExitFunction_parameter is called when production function_parameter is exited.
func (s *BaseSOQLListener) ExitFunction_parameter(ctx *Function_parameterContext) {}

// EnterTypeof_spec is called when production typeof_spec is entered.
func (s *BaseSOQLListener) EnterTypeof_spec(ctx *Typeof_specContext) {}

// ExitTypeof_spec is called when production typeof_spec is exited.
func (s *BaseSOQLListener) ExitTypeof_spec(ctx *Typeof_specContext) {}

// EnterTypeof_when_then_clause_list is called when production typeof_when_then_clause_list is entered.
func (s *BaseSOQLListener) EnterTypeof_when_then_clause_list(ctx *Typeof_when_then_clause_listContext) {
}

// ExitTypeof_when_then_clause_list is called when production typeof_when_then_clause_list is exited.
func (s *BaseSOQLListener) ExitTypeof_when_then_clause_list(ctx *Typeof_when_then_clause_listContext) {
}

// EnterTypeof_when_then_clause is called when production typeof_when_then_clause is entered.
func (s *BaseSOQLListener) EnterTypeof_when_then_clause(ctx *Typeof_when_then_clauseContext) {}

// ExitTypeof_when_then_clause is called when production typeof_when_then_clause is exited.
func (s *BaseSOQLListener) ExitTypeof_when_then_clause(ctx *Typeof_when_then_clauseContext) {}

// EnterTypeof_then_clause is called when production typeof_then_clause is entered.
func (s *BaseSOQLListener) EnterTypeof_then_clause(ctx *Typeof_then_clauseContext) {}

// ExitTypeof_then_clause is called when production typeof_then_clause is exited.
func (s *BaseSOQLListener) ExitTypeof_then_clause(ctx *Typeof_then_clauseContext) {}

// EnterTypeof_else_clause is called when production typeof_else_clause is entered.
func (s *BaseSOQLListener) EnterTypeof_else_clause(ctx *Typeof_else_clauseContext) {}

// ExitTypeof_else_clause is called when production typeof_else_clause is exited.
func (s *BaseSOQLListener) ExitTypeof_else_clause(ctx *Typeof_else_clauseContext) {}

// EnterField_list is called when production field_list is entered.
func (s *BaseSOQLListener) EnterField_list(ctx *Field_listContext) {}

// ExitField_list is called when production field_list is exited.
func (s *BaseSOQLListener) ExitField_list(ctx *Field_listContext) {}

// EnterObject_spec is called when production object_spec is entered.
func (s *BaseSOQLListener) EnterObject_spec(ctx *Object_specContext) {}

// ExitObject_spec is called when production object_spec is exited.
func (s *BaseSOQLListener) ExitObject_spec(ctx *Object_specContext) {}

// EnterObject_prefix is called when production object_prefix is entered.
func (s *BaseSOQLListener) EnterObject_prefix(ctx *Object_prefixContext) {}

// ExitObject_prefix is called when production object_prefix is exited.
func (s *BaseSOQLListener) ExitObject_prefix(ctx *Object_prefixContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BaseSOQLListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BaseSOQLListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterSet_operator is called when production set_operator is entered.
func (s *BaseSOQLListener) EnterSet_operator(ctx *Set_operatorContext) {}

// ExitSet_operator is called when production set_operator is exited.
func (s *BaseSOQLListener) ExitSet_operator(ctx *Set_operatorContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseSOQLListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseSOQLListener) ExitCondition(ctx *ConditionContext) {}

// EnterCondition1 is called when production condition1 is entered.
func (s *BaseSOQLListener) EnterCondition1(ctx *Condition1Context) {}

// ExitCondition1 is called when production condition1 is exited.
func (s *BaseSOQLListener) ExitCondition1(ctx *Condition1Context) {}

// EnterParenthesis is called when production parenthesis is entered.
func (s *BaseSOQLListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production parenthesis is exited.
func (s *BaseSOQLListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterSimple_condition is called when production simple_condition is entered.
func (s *BaseSOQLListener) EnterSimple_condition(ctx *Simple_conditionContext) {}

// ExitSimple_condition is called when production simple_condition is exited.
func (s *BaseSOQLListener) ExitSimple_condition(ctx *Simple_conditionContext) {}

// EnterField_based_condition is called when production field_based_condition is entered.
func (s *BaseSOQLListener) EnterField_based_condition(ctx *Field_based_conditionContext) {}

// ExitField_based_condition is called when production field_based_condition is exited.
func (s *BaseSOQLListener) ExitField_based_condition(ctx *Field_based_conditionContext) {}

// EnterSet_based_condition is called when production set_based_condition is entered.
func (s *BaseSOQLListener) EnterSet_based_condition(ctx *Set_based_conditionContext) {}

// ExitSet_based_condition is called when production set_based_condition is exited.
func (s *BaseSOQLListener) ExitSet_based_condition(ctx *Set_based_conditionContext) {}

// EnterLike_based_condition is called when production like_based_condition is entered.
func (s *BaseSOQLListener) EnterLike_based_condition(ctx *Like_based_conditionContext) {}

// ExitLike_based_condition is called when production like_based_condition is exited.
func (s *BaseSOQLListener) ExitLike_based_condition(ctx *Like_based_conditionContext) {}

// EnterCondition_field is called when production condition_field is entered.
func (s *BaseSOQLListener) EnterCondition_field(ctx *Condition_fieldContext) {}

// ExitCondition_field is called when production condition_field is exited.
func (s *BaseSOQLListener) ExitCondition_field(ctx *Condition_fieldContext) {}

// EnterSet_values is called when production set_values is entered.
func (s *BaseSOQLListener) EnterSet_values(ctx *Set_valuesContext) {}

// ExitSet_values is called when production set_values is exited.
func (s *BaseSOQLListener) ExitSet_values(ctx *Set_valuesContext) {}

// EnterSet_value_list is called when production set_value_list is entered.
func (s *BaseSOQLListener) EnterSet_value_list(ctx *Set_value_listContext) {}

// ExitSet_value_list is called when production set_value_list is exited.
func (s *BaseSOQLListener) ExitSet_value_list(ctx *Set_value_listContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseSOQLListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseSOQLListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterWith_plain_clause is called when production with_plain_clause is entered.
func (s *BaseSOQLListener) EnterWith_plain_clause(ctx *With_plain_clauseContext) {}

// ExitWith_plain_clause is called when production with_plain_clause is exited.
func (s *BaseSOQLListener) ExitWith_plain_clause(ctx *With_plain_clauseContext) {}

// EnterWith_data_category_clause is called when production with_data_category_clause is entered.
func (s *BaseSOQLListener) EnterWith_data_category_clause(ctx *With_data_category_clauseContext) {}

// ExitWith_data_category_clause is called when production with_data_category_clause is exited.
func (s *BaseSOQLListener) ExitWith_data_category_clause(ctx *With_data_category_clauseContext) {}

// EnterData_category_spec_list is called when production data_category_spec_list is entered.
func (s *BaseSOQLListener) EnterData_category_spec_list(ctx *Data_category_spec_listContext) {}

// ExitData_category_spec_list is called when production data_category_spec_list is exited.
func (s *BaseSOQLListener) ExitData_category_spec_list(ctx *Data_category_spec_listContext) {}

// EnterData_category_spec is called when production data_category_spec is entered.
func (s *BaseSOQLListener) EnterData_category_spec(ctx *Data_category_specContext) {}

// ExitData_category_spec is called when production data_category_spec is exited.
func (s *BaseSOQLListener) ExitData_category_spec(ctx *Data_category_specContext) {}

// EnterData_category_parameter_list is called when production data_category_parameter_list is entered.
func (s *BaseSOQLListener) EnterData_category_parameter_list(ctx *Data_category_parameter_listContext) {
}

// ExitData_category_parameter_list is called when production data_category_parameter_list is exited.
func (s *BaseSOQLListener) ExitData_category_parameter_list(ctx *Data_category_parameter_listContext) {
}

// EnterData_category_selector is called when production data_category_selector is entered.
func (s *BaseSOQLListener) EnterData_category_selector(ctx *Data_category_selectorContext) {}

// ExitData_category_selector is called when production data_category_selector is exited.
func (s *BaseSOQLListener) ExitData_category_selector(ctx *Data_category_selectorContext) {}

// EnterGroup_by_plain_clause is called when production group_by_plain_clause is entered.
func (s *BaseSOQLListener) EnterGroup_by_plain_clause(ctx *Group_by_plain_clauseContext) {}

// ExitGroup_by_plain_clause is called when production group_by_plain_clause is exited.
func (s *BaseSOQLListener) ExitGroup_by_plain_clause(ctx *Group_by_plain_clauseContext) {}

// EnterGroup_by_rollup_clause is called when production group_by_rollup_clause is entered.
func (s *BaseSOQLListener) EnterGroup_by_rollup_clause(ctx *Group_by_rollup_clauseContext) {}

// ExitGroup_by_rollup_clause is called when production group_by_rollup_clause is exited.
func (s *BaseSOQLListener) ExitGroup_by_rollup_clause(ctx *Group_by_rollup_clauseContext) {}

// EnterGroup_by_cube_clause is called when production group_by_cube_clause is entered.
func (s *BaseSOQLListener) EnterGroup_by_cube_clause(ctx *Group_by_cube_clauseContext) {}

// ExitGroup_by_cube_clause is called when production group_by_cube_clause is exited.
func (s *BaseSOQLListener) ExitGroup_by_cube_clause(ctx *Group_by_cube_clauseContext) {}

// EnterGroup_by_list is called when production group_by_list is entered.
func (s *BaseSOQLListener) EnterGroup_by_list(ctx *Group_by_listContext) {}

// ExitGroup_by_list is called when production group_by_list is exited.
func (s *BaseSOQLListener) ExitGroup_by_list(ctx *Group_by_listContext) {}

// EnterGroup_by_spec is called when production group_by_spec is entered.
func (s *BaseSOQLListener) EnterGroup_by_spec(ctx *Group_by_specContext) {}

// ExitGroup_by_spec is called when production group_by_spec is exited.
func (s *BaseSOQLListener) ExitGroup_by_spec(ctx *Group_by_specContext) {}

// EnterOrder_by_list is called when production order_by_list is entered.
func (s *BaseSOQLListener) EnterOrder_by_list(ctx *Order_by_listContext) {}

// ExitOrder_by_list is called when production order_by_list is exited.
func (s *BaseSOQLListener) ExitOrder_by_list(ctx *Order_by_listContext) {}

// EnterOrder_by_spec is called when production order_by_spec is entered.
func (s *BaseSOQLListener) EnterOrder_by_spec(ctx *Order_by_specContext) {}

// ExitOrder_by_spec is called when production order_by_spec is exited.
func (s *BaseSOQLListener) ExitOrder_by_spec(ctx *Order_by_specContext) {}

// EnterOrder_by_direction_clause is called when production order_by_direction_clause is entered.
func (s *BaseSOQLListener) EnterOrder_by_direction_clause(ctx *Order_by_direction_clauseContext) {}

// ExitOrder_by_direction_clause is called when production order_by_direction_clause is exited.
func (s *BaseSOQLListener) ExitOrder_by_direction_clause(ctx *Order_by_direction_clauseContext) {}

// EnterOrder_by_nulls_clause is called when production order_by_nulls_clause is entered.
func (s *BaseSOQLListener) EnterOrder_by_nulls_clause(ctx *Order_by_nulls_clauseContext) {}

// ExitOrder_by_nulls_clause is called when production order_by_nulls_clause is exited.
func (s *BaseSOQLListener) ExitOrder_by_nulls_clause(ctx *Order_by_nulls_clauseContext) {}

// EnterOrder_by_field is called when production order_by_field is entered.
func (s *BaseSOQLListener) EnterOrder_by_field(ctx *Order_by_fieldContext) {}

// ExitOrder_by_field is called when production order_by_field is exited.
func (s *BaseSOQLListener) ExitOrder_by_field(ctx *Order_by_fieldContext) {}

// EnterFor_value is called when production for_value is entered.
func (s *BaseSOQLListener) EnterFor_value(ctx *For_valueContext) {}

// ExitFor_value is called when production for_value is exited.
func (s *BaseSOQLListener) ExitFor_value(ctx *For_valueContext) {}

// EnterUpdate_value is called when production update_value is entered.
func (s *BaseSOQLListener) EnterUpdate_value(ctx *Update_valueContext) {}

// ExitUpdate_value is called when production update_value is exited.
func (s *BaseSOQLListener) ExitUpdate_value(ctx *Update_valueContext) {}
