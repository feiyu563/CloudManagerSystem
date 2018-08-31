插件:
go get github.com/satori/go.uuid

问题解决；（which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by）
#SET GLOBAL sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
