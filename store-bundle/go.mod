module github.com/tigershi/singletongo/store-bundle
require github.com/tigershi/singletongo/service-core v0.0.0
replace github.com/tigershi/singletongo/service-core => ../service-core
require github.com/tigershi/singletongo/service-facade v0.0.0
replace github.com/tigershi/singletongo/service-facade => ../service-facade

go 1.12
