// Code generated from SOQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SOQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SOQLListener is a complete listener for a parse tree produced by SOQLParser.
type SOQLListener interface {
	antlr.ParseTreeListener

	// EnterKeywords_alias_allowed is called when entering the keywords_alias_allowed production.
	EnterKeywords_alias_allowed(c *Keywords_alias_allowedContext)

	// EnterKeywords_name_allowed is called when entering the keywords_name_allowed production.
	EnterKeywords_name_allowed(c *Keywords_name_allowedContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterObject_name is called when entering the object_name production.
	EnterObject_name(c *Object_nameContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// EnterFilter_scope_name is called when entering the filter_scope_name production.
	EnterFilter_scope_name(c *Filter_scope_nameContext)

	// EnterData_category_group_name is called when entering the data_category_group_name production.
	EnterData_category_group_name(c *Data_category_group_nameContext)

	// EnterData_category_name is called when entering the data_category_name production.
	EnterData_category_name(c *Data_category_nameContext)

	// EnterAlias_name is called when entering the alias_name production.
	EnterAlias_name(c *Alias_nameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterDate_formula_literal is called when entering the date_formula_literal production.
	EnterDate_formula_literal(c *Date_formula_literalContext)

	// EnterDate_formula_n_literal_name is called when entering the date_formula_n_literal_name production.
	EnterDate_formula_n_literal_name(c *Date_formula_n_literal_nameContext)

	// EnterDate_formula_n_literal is called when entering the date_formula_n_literal production.
	EnterDate_formula_n_literal(c *Date_formula_n_literalContext)

	// EnterDatetime_literal is called when entering the datetime_literal production.
	EnterDatetime_literal(c *Datetime_literalContext)

	// EnterDate_literal is called when entering the date_literal production.
	EnterDate_literal(c *Date_literalContext)

	// EnterInteger_literal is called when entering the integer_literal production.
	EnterInteger_literal(c *Integer_literalContext)

	// EnterReal_literal is called when entering the real_literal production.
	EnterReal_literal(c *Real_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterBoolean_literal is called when entering the boolean_literal production.
	EnterBoolean_literal(c *Boolean_literalContext)

	// EnterNull_literal is called when entering the null_literal production.
	EnterNull_literal(c *Null_literalContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterFunction_date is called when entering the function_date production.
	EnterFunction_date(c *Function_dateContext)

	// EnterFunction_aggregate is called when entering the function_aggregate production.
	EnterFunction_aggregate(c *Function_aggregateContext)

	// EnterFunction_location is called when entering the function_location production.
	EnterFunction_location(c *Function_locationContext)

	// EnterFunction_other is called when entering the function_other production.
	EnterFunction_other(c *Function_otherContext)

	// EnterSoql_query is called when entering the soql_query production.
	EnterSoql_query(c *Soql_queryContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterUsing_clause is called when entering the using_clause production.
	EnterUsing_clause(c *Using_clauseContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterGroupby_clause is called when entering the groupby_clause production.
	EnterGroupby_clause(c *Groupby_clauseContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterOrderby_clause is called when entering the orderby_clause production.
	EnterOrderby_clause(c *Orderby_clauseContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterOffset_clause is called when entering the offset_clause production.
	EnterOffset_clause(c *Offset_clauseContext)

	// EnterFor_clause is called when entering the for_clause production.
	EnterFor_clause(c *For_clauseContext)

	// EnterUpdate_clause is called when entering the update_clause production.
	EnterUpdate_clause(c *Update_clauseContext)

	// EnterSoql_subquery is called when entering the soql_subquery production.
	EnterSoql_subquery(c *Soql_subqueryContext)

	// EnterSubquery_select_clause is called when entering the subquery_select_clause production.
	EnterSubquery_select_clause(c *Subquery_select_clauseContext)

	// EnterSelect_spec is called when entering the select_spec production.
	EnterSelect_spec(c *Select_specContext)

	// EnterSubquery_select_spec is called when entering the subquery_select_spec production.
	EnterSubquery_select_spec(c *Subquery_select_specContext)

	// EnterField_spec is called when entering the field_spec production.
	EnterField_spec(c *Field_specContext)

	// EnterFunction_call_spec is called when entering the function_call_spec production.
	EnterFunction_call_spec(c *Function_call_specContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterFunction_parameter_list is called when entering the function_parameter_list production.
	EnterFunction_parameter_list(c *Function_parameter_listContext)

	// EnterFunction_parameter is called when entering the function_parameter production.
	EnterFunction_parameter(c *Function_parameterContext)

	// EnterTypeof_spec is called when entering the typeof_spec production.
	EnterTypeof_spec(c *Typeof_specContext)

	// EnterTypeof_when_then_clause_list is called when entering the typeof_when_then_clause_list production.
	EnterTypeof_when_then_clause_list(c *Typeof_when_then_clause_listContext)

	// EnterTypeof_when_then_clause is called when entering the typeof_when_then_clause production.
	EnterTypeof_when_then_clause(c *Typeof_when_then_clauseContext)

	// EnterTypeof_then_clause is called when entering the typeof_then_clause production.
	EnterTypeof_then_clause(c *Typeof_then_clauseContext)

	// EnterTypeof_else_clause is called when entering the typeof_else_clause production.
	EnterTypeof_else_clause(c *Typeof_else_clauseContext)

	// EnterField_list is called when entering the field_list production.
	EnterField_list(c *Field_listContext)

	// EnterObject_spec is called when entering the object_spec production.
	EnterObject_spec(c *Object_specContext)

	// EnterObject_prefix is called when entering the object_prefix production.
	EnterObject_prefix(c *Object_prefixContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterSet_operator is called when entering the set_operator production.
	EnterSet_operator(c *Set_operatorContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterCondition1 is called when entering the condition1 production.
	EnterCondition1(c *Condition1Context)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterSimple_condition is called when entering the simple_condition production.
	EnterSimple_condition(c *Simple_conditionContext)

	// EnterField_based_condition is called when entering the field_based_condition production.
	EnterField_based_condition(c *Field_based_conditionContext)

	// EnterSet_based_condition is called when entering the set_based_condition production.
	EnterSet_based_condition(c *Set_based_conditionContext)

	// EnterLike_based_condition is called when entering the like_based_condition production.
	EnterLike_based_condition(c *Like_based_conditionContext)

	// EnterCondition_field is called when entering the condition_field production.
	EnterCondition_field(c *Condition_fieldContext)

	// EnterSet_values is called when entering the set_values production.
	EnterSet_values(c *Set_valuesContext)

	// EnterSet_value_list is called when entering the set_value_list production.
	EnterSet_value_list(c *Set_value_listContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterWith_plain_clause is called when entering the with_plain_clause production.
	EnterWith_plain_clause(c *With_plain_clauseContext)

	// EnterWith_data_category_clause is called when entering the with_data_category_clause production.
	EnterWith_data_category_clause(c *With_data_category_clauseContext)

	// EnterData_category_spec_list is called when entering the data_category_spec_list production.
	EnterData_category_spec_list(c *Data_category_spec_listContext)

	// EnterData_category_spec is called when entering the data_category_spec production.
	EnterData_category_spec(c *Data_category_specContext)

	// EnterData_category_parameter_list is called when entering the data_category_parameter_list production.
	EnterData_category_parameter_list(c *Data_category_parameter_listContext)

	// EnterData_category_selector is called when entering the data_category_selector production.
	EnterData_category_selector(c *Data_category_selectorContext)

	// EnterGroup_by_plain_clause is called when entering the group_by_plain_clause production.
	EnterGroup_by_plain_clause(c *Group_by_plain_clauseContext)

	// EnterGroup_by_rollup_clause is called when entering the group_by_rollup_clause production.
	EnterGroup_by_rollup_clause(c *Group_by_rollup_clauseContext)

	// EnterGroup_by_cube_clause is called when entering the group_by_cube_clause production.
	EnterGroup_by_cube_clause(c *Group_by_cube_clauseContext)

	// EnterGroup_by_list is called when entering the group_by_list production.
	EnterGroup_by_list(c *Group_by_listContext)

	// EnterGroup_by_spec is called when entering the group_by_spec production.
	EnterGroup_by_spec(c *Group_by_specContext)

	// EnterOrder_by_list is called when entering the order_by_list production.
	EnterOrder_by_list(c *Order_by_listContext)

	// EnterOrder_by_spec is called when entering the order_by_spec production.
	EnterOrder_by_spec(c *Order_by_specContext)

	// EnterOrder_by_direction_clause is called when entering the order_by_direction_clause production.
	EnterOrder_by_direction_clause(c *Order_by_direction_clauseContext)

	// EnterOrder_by_nulls_clause is called when entering the order_by_nulls_clause production.
	EnterOrder_by_nulls_clause(c *Order_by_nulls_clauseContext)

	// EnterOrder_by_field is called when entering the order_by_field production.
	EnterOrder_by_field(c *Order_by_fieldContext)

	// EnterFor_value is called when entering the for_value production.
	EnterFor_value(c *For_valueContext)

	// EnterUpdate_value is called when entering the update_value production.
	EnterUpdate_value(c *Update_valueContext)

	// ExitKeywords_alias_allowed is called when exiting the keywords_alias_allowed production.
	ExitKeywords_alias_allowed(c *Keywords_alias_allowedContext)

	// ExitKeywords_name_allowed is called when exiting the keywords_name_allowed production.
	ExitKeywords_name_allowed(c *Keywords_name_allowedContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitObject_name is called when exiting the object_name production.
	ExitObject_name(c *Object_nameContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)

	// ExitFilter_scope_name is called when exiting the filter_scope_name production.
	ExitFilter_scope_name(c *Filter_scope_nameContext)

	// ExitData_category_group_name is called when exiting the data_category_group_name production.
	ExitData_category_group_name(c *Data_category_group_nameContext)

	// ExitData_category_name is called when exiting the data_category_name production.
	ExitData_category_name(c *Data_category_nameContext)

	// ExitAlias_name is called when exiting the alias_name production.
	ExitAlias_name(c *Alias_nameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitDate_formula_literal is called when exiting the date_formula_literal production.
	ExitDate_formula_literal(c *Date_formula_literalContext)

	// ExitDate_formula_n_literal_name is called when exiting the date_formula_n_literal_name production.
	ExitDate_formula_n_literal_name(c *Date_formula_n_literal_nameContext)

	// ExitDate_formula_n_literal is called when exiting the date_formula_n_literal production.
	ExitDate_formula_n_literal(c *Date_formula_n_literalContext)

	// ExitDatetime_literal is called when exiting the datetime_literal production.
	ExitDatetime_literal(c *Datetime_literalContext)

	// ExitDate_literal is called when exiting the date_literal production.
	ExitDate_literal(c *Date_literalContext)

	// ExitInteger_literal is called when exiting the integer_literal production.
	ExitInteger_literal(c *Integer_literalContext)

	// ExitReal_literal is called when exiting the real_literal production.
	ExitReal_literal(c *Real_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitBoolean_literal is called when exiting the boolean_literal production.
	ExitBoolean_literal(c *Boolean_literalContext)

	// ExitNull_literal is called when exiting the null_literal production.
	ExitNull_literal(c *Null_literalContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitFunction_date is called when exiting the function_date production.
	ExitFunction_date(c *Function_dateContext)

	// ExitFunction_aggregate is called when exiting the function_aggregate production.
	ExitFunction_aggregate(c *Function_aggregateContext)

	// ExitFunction_location is called when exiting the function_location production.
	ExitFunction_location(c *Function_locationContext)

	// ExitFunction_other is called when exiting the function_other production.
	ExitFunction_other(c *Function_otherContext)

	// ExitSoql_query is called when exiting the soql_query production.
	ExitSoql_query(c *Soql_queryContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitUsing_clause is called when exiting the using_clause production.
	ExitUsing_clause(c *Using_clauseContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitGroupby_clause is called when exiting the groupby_clause production.
	ExitGroupby_clause(c *Groupby_clauseContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitOrderby_clause is called when exiting the orderby_clause production.
	ExitOrderby_clause(c *Orderby_clauseContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitOffset_clause is called when exiting the offset_clause production.
	ExitOffset_clause(c *Offset_clauseContext)

	// ExitFor_clause is called when exiting the for_clause production.
	ExitFor_clause(c *For_clauseContext)

	// ExitUpdate_clause is called when exiting the update_clause production.
	ExitUpdate_clause(c *Update_clauseContext)

	// ExitSoql_subquery is called when exiting the soql_subquery production.
	ExitSoql_subquery(c *Soql_subqueryContext)

	// ExitSubquery_select_clause is called when exiting the subquery_select_clause production.
	ExitSubquery_select_clause(c *Subquery_select_clauseContext)

	// ExitSelect_spec is called when exiting the select_spec production.
	ExitSelect_spec(c *Select_specContext)

	// ExitSubquery_select_spec is called when exiting the subquery_select_spec production.
	ExitSubquery_select_spec(c *Subquery_select_specContext)

	// ExitField_spec is called when exiting the field_spec production.
	ExitField_spec(c *Field_specContext)

	// ExitFunction_call_spec is called when exiting the function_call_spec production.
	ExitFunction_call_spec(c *Function_call_specContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitFunction_parameter_list is called when exiting the function_parameter_list production.
	ExitFunction_parameter_list(c *Function_parameter_listContext)

	// ExitFunction_parameter is called when exiting the function_parameter production.
	ExitFunction_parameter(c *Function_parameterContext)

	// ExitTypeof_spec is called when exiting the typeof_spec production.
	ExitTypeof_spec(c *Typeof_specContext)

	// ExitTypeof_when_then_clause_list is called when exiting the typeof_when_then_clause_list production.
	ExitTypeof_when_then_clause_list(c *Typeof_when_then_clause_listContext)

	// ExitTypeof_when_then_clause is called when exiting the typeof_when_then_clause production.
	ExitTypeof_when_then_clause(c *Typeof_when_then_clauseContext)

	// ExitTypeof_then_clause is called when exiting the typeof_then_clause production.
	ExitTypeof_then_clause(c *Typeof_then_clauseContext)

	// ExitTypeof_else_clause is called when exiting the typeof_else_clause production.
	ExitTypeof_else_clause(c *Typeof_else_clauseContext)

	// ExitField_list is called when exiting the field_list production.
	ExitField_list(c *Field_listContext)

	// ExitObject_spec is called when exiting the object_spec production.
	ExitObject_spec(c *Object_specContext)

	// ExitObject_prefix is called when exiting the object_prefix production.
	ExitObject_prefix(c *Object_prefixContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitSet_operator is called when exiting the set_operator production.
	ExitSet_operator(c *Set_operatorContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitCondition1 is called when exiting the condition1 production.
	ExitCondition1(c *Condition1Context)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitSimple_condition is called when exiting the simple_condition production.
	ExitSimple_condition(c *Simple_conditionContext)

	// ExitField_based_condition is called when exiting the field_based_condition production.
	ExitField_based_condition(c *Field_based_conditionContext)

	// ExitSet_based_condition is called when exiting the set_based_condition production.
	ExitSet_based_condition(c *Set_based_conditionContext)

	// ExitLike_based_condition is called when exiting the like_based_condition production.
	ExitLike_based_condition(c *Like_based_conditionContext)

	// ExitCondition_field is called when exiting the condition_field production.
	ExitCondition_field(c *Condition_fieldContext)

	// ExitSet_values is called when exiting the set_values production.
	ExitSet_values(c *Set_valuesContext)

	// ExitSet_value_list is called when exiting the set_value_list production.
	ExitSet_value_list(c *Set_value_listContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitWith_plain_clause is called when exiting the with_plain_clause production.
	ExitWith_plain_clause(c *With_plain_clauseContext)

	// ExitWith_data_category_clause is called when exiting the with_data_category_clause production.
	ExitWith_data_category_clause(c *With_data_category_clauseContext)

	// ExitData_category_spec_list is called when exiting the data_category_spec_list production.
	ExitData_category_spec_list(c *Data_category_spec_listContext)

	// ExitData_category_spec is called when exiting the data_category_spec production.
	ExitData_category_spec(c *Data_category_specContext)

	// ExitData_category_parameter_list is called when exiting the data_category_parameter_list production.
	ExitData_category_parameter_list(c *Data_category_parameter_listContext)

	// ExitData_category_selector is called when exiting the data_category_selector production.
	ExitData_category_selector(c *Data_category_selectorContext)

	// ExitGroup_by_plain_clause is called when exiting the group_by_plain_clause production.
	ExitGroup_by_plain_clause(c *Group_by_plain_clauseContext)

	// ExitGroup_by_rollup_clause is called when exiting the group_by_rollup_clause production.
	ExitGroup_by_rollup_clause(c *Group_by_rollup_clauseContext)

	// ExitGroup_by_cube_clause is called when exiting the group_by_cube_clause production.
	ExitGroup_by_cube_clause(c *Group_by_cube_clauseContext)

	// ExitGroup_by_list is called when exiting the group_by_list production.
	ExitGroup_by_list(c *Group_by_listContext)

	// ExitGroup_by_spec is called when exiting the group_by_spec production.
	ExitGroup_by_spec(c *Group_by_specContext)

	// ExitOrder_by_list is called when exiting the order_by_list production.
	ExitOrder_by_list(c *Order_by_listContext)

	// ExitOrder_by_spec is called when exiting the order_by_spec production.
	ExitOrder_by_spec(c *Order_by_specContext)

	// ExitOrder_by_direction_clause is called when exiting the order_by_direction_clause production.
	ExitOrder_by_direction_clause(c *Order_by_direction_clauseContext)

	// ExitOrder_by_nulls_clause is called when exiting the order_by_nulls_clause production.
	ExitOrder_by_nulls_clause(c *Order_by_nulls_clauseContext)

	// ExitOrder_by_field is called when exiting the order_by_field production.
	ExitOrder_by_field(c *Order_by_fieldContext)

	// ExitFor_value is called when exiting the for_value production.
	ExitFor_value(c *For_valueContext)

	// ExitUpdate_value is called when exiting the update_value production.
	ExitUpdate_value(c *Update_valueContext)
}
