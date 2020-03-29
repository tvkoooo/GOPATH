@echo off
set c1="tvkoooo:dk2012@tcp(127.0.0.1:3306)/db_mm_cherry?charset=utf8"
set c2="select `id`,`name` from t_user_basic;"
set c3="select * from t_user_basic;"
set c4="select * from t_not;"
set c5="call p_userinfo_search(@_code, 3, '1');"
set c6="call p_userinfo_search(@_code, 3, '777');"
start cmd /k "sqltest.exe %c1% %c2% %c3% %c4% %c5% %c6%"
