// Code generated from SOQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SOQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SOQLParser.
type SOQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SOQLParser#keywords_alias_allowed.
	VisitKeywords_alias_allowed(ctx *Keywords_alias_allowedContext) interface{}

	// Visit a parse tree produced by SOQLParser#keywords_name_allowed.
	VisitKeywords_name_allowed(ctx *Keywords_name_allowedContext) interface{}

	// Visit a parse tree produced by SOQLParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by SOQLParser#object_name.
	VisitObject_name(ctx *Object_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#field_name.
	VisitField_name(ctx *Field_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#filter_scope_name.
	VisitFilter_scope_name(ctx *Filter_scope_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#data_category_group_name.
	VisitData_category_group_name(ctx *Data_category_group_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#data_category_name.
	VisitData_category_name(ctx *Data_category_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#alias_name.
	VisitAlias_name(ctx *Alias_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by SOQLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by SOQLParser#date_formula_literal.
	VisitDate_formula_literal(ctx *Date_formula_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#date_formula_n_literal_name.
	VisitDate_formula_n_literal_name(ctx *Date_formula_n_literal_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#date_formula_n_literal.
	VisitDate_formula_n_literal(ctx *Date_formula_n_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#datetime_literal.
	VisitDatetime_literal(ctx *Datetime_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#date_literal.
	VisitDate_literal(ctx *Date_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#integer_literal.
	VisitInteger_literal(ctx *Integer_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#real_literal.
	VisitReal_literal(ctx *Real_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#boolean_literal.
	VisitBoolean_literal(ctx *Boolean_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#null_literal.
	VisitNull_literal(ctx *Null_literalContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_date.
	VisitFunction_date(ctx *Function_dateContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_aggregate.
	VisitFunction_aggregate(ctx *Function_aggregateContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_location.
	VisitFunction_location(ctx *Function_locationContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_other.
	VisitFunction_other(ctx *Function_otherContext) interface{}

	// Visit a parse tree produced by SOQLParser#soql_query.
	VisitSoql_query(ctx *Soql_queryContext) interface{}

	// Visit a parse tree produced by SOQLParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#groupby_clause.
	VisitGroupby_clause(ctx *Groupby_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#orderby_clause.
	VisitOrderby_clause(ctx *Orderby_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#offset_clause.
	VisitOffset_clause(ctx *Offset_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#for_clause.
	VisitFor_clause(ctx *For_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#update_clause.
	VisitUpdate_clause(ctx *Update_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#soql_subquery.
	VisitSoql_subquery(ctx *Soql_subqueryContext) interface{}

	// Visit a parse tree produced by SOQLParser#subquery_select_clause.
	VisitSubquery_select_clause(ctx *Subquery_select_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#select_spec.
	VisitSelect_spec(ctx *Select_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#subquery_select_spec.
	VisitSubquery_select_spec(ctx *Subquery_select_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#field_spec.
	VisitField_spec(ctx *Field_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_call_spec.
	VisitFunction_call_spec(ctx *Function_call_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_call.
	VisitFunction_call(ctx *Function_callContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_parameter_list.
	VisitFunction_parameter_list(ctx *Function_parameter_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#function_parameter.
	VisitFunction_parameter(ctx *Function_parameterContext) interface{}

	// Visit a parse tree produced by SOQLParser#typeof_spec.
	VisitTypeof_spec(ctx *Typeof_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#typeof_when_then_clause_list.
	VisitTypeof_when_then_clause_list(ctx *Typeof_when_then_clause_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#typeof_when_then_clause.
	VisitTypeof_when_then_clause(ctx *Typeof_when_then_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#typeof_then_clause.
	VisitTypeof_then_clause(ctx *Typeof_then_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#typeof_else_clause.
	VisitTypeof_else_clause(ctx *Typeof_else_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#field_list.
	VisitField_list(ctx *Field_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#object_spec.
	VisitObject_spec(ctx *Object_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#object_prefix.
	VisitObject_prefix(ctx *Object_prefixContext) interface{}

	// Visit a parse tree produced by SOQLParser#comparison_operator.
	VisitComparison_operator(ctx *Comparison_operatorContext) interface{}

	// Visit a parse tree produced by SOQLParser#set_operator.
	VisitSet_operator(ctx *Set_operatorContext) interface{}

	// Visit a parse tree produced by SOQLParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by SOQLParser#condition1.
	VisitCondition1(ctx *Condition1Context) interface{}

	// Visit a parse tree produced by SOQLParser#parenthesis.
	VisitParenthesis(ctx *ParenthesisContext) interface{}

	// Visit a parse tree produced by SOQLParser#simple_condition.
	VisitSimple_condition(ctx *Simple_conditionContext) interface{}

	// Visit a parse tree produced by SOQLParser#field_based_condition.
	VisitField_based_condition(ctx *Field_based_conditionContext) interface{}

	// Visit a parse tree produced by SOQLParser#set_based_condition.
	VisitSet_based_condition(ctx *Set_based_conditionContext) interface{}

	// Visit a parse tree produced by SOQLParser#like_based_condition.
	VisitLike_based_condition(ctx *Like_based_conditionContext) interface{}

	// Visit a parse tree produced by SOQLParser#condition_field.
	VisitCondition_field(ctx *Condition_fieldContext) interface{}

	// Visit a parse tree produced by SOQLParser#set_values.
	VisitSet_values(ctx *Set_valuesContext) interface{}

	// Visit a parse tree produced by SOQLParser#set_value_list.
	VisitSet_value_list(ctx *Set_value_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#with_clause.
	VisitWith_clause(ctx *With_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#with_plain_clause.
	VisitWith_plain_clause(ctx *With_plain_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#with_data_category_clause.
	VisitWith_data_category_clause(ctx *With_data_category_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#data_category_spec_list.
	VisitData_category_spec_list(ctx *Data_category_spec_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#data_category_spec.
	VisitData_category_spec(ctx *Data_category_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#data_category_parameter_list.
	VisitData_category_parameter_list(ctx *Data_category_parameter_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#data_category_selector.
	VisitData_category_selector(ctx *Data_category_selectorContext) interface{}

	// Visit a parse tree produced by SOQLParser#group_by_plain_clause.
	VisitGroup_by_plain_clause(ctx *Group_by_plain_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#group_by_rollup_clause.
	VisitGroup_by_rollup_clause(ctx *Group_by_rollup_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#group_by_cube_clause.
	VisitGroup_by_cube_clause(ctx *Group_by_cube_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#group_by_list.
	VisitGroup_by_list(ctx *Group_by_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#group_by_spec.
	VisitGroup_by_spec(ctx *Group_by_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#order_by_list.
	VisitOrder_by_list(ctx *Order_by_listContext) interface{}

	// Visit a parse tree produced by SOQLParser#order_by_spec.
	VisitOrder_by_spec(ctx *Order_by_specContext) interface{}

	// Visit a parse tree produced by SOQLParser#order_by_direction_clause.
	VisitOrder_by_direction_clause(ctx *Order_by_direction_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#order_by_nulls_clause.
	VisitOrder_by_nulls_clause(ctx *Order_by_nulls_clauseContext) interface{}

	// Visit a parse tree produced by SOQLParser#order_by_field.
	VisitOrder_by_field(ctx *Order_by_fieldContext) interface{}

	// Visit a parse tree produced by SOQLParser#for_value.
	VisitFor_value(ctx *For_valueContext) interface{}

	// Visit a parse tree produced by SOQLParser#update_value.
	VisitUpdate_value(ctx *Update_valueContext) interface{}
}
