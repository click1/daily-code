/**
 * 实现php 释放函数功能(function_unset)  类似变量 unset  
 *   核心函数：zend_hash_del_key_or_index   删除function_table hashtable中对应函数名
 */
PHP_FUNCTION(function_unset)
{
    char *name;
    int name_len;
    zend_function *func;
    char *lcname;
    zend_bool retval;

    if (zend_parse_parameters(ZEND_NUM_ARGS() TSRMLS_CC, "s", &name, &name_len) == FAILURE) {
        return;
    }   

    lcname = zend_str_tolower_dup(name, name_len);


    /* Ignore leading "\" */
    name = lcname;
    if (lcname[0] == '\\') {
        name = &lcname[1];
        name_len--;
    }   

    retval = (zend_hash_del_key_or_index(EG(function_table), name, name_len+1, 0, 0) == SUCCESS);

    efree(lcname);

    RETURN_BOOL(retval);
}

