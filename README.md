# High School Programming Contest (HSPC)

The goal of this project is to build an web based application for the High School 
Programming Contest here at Kansas State university. Please visit 
[HSPC page](https://www.cs.ksu.edu/hspc) for more details.

Although the application is being developed specifically for HSPC, however any programming contest event which 
 follows a pattern similar to HSPC, should be able to use the application.

The architecture of the application is very simple
1. The data will be stored in PostgreSQL and MongoDB
2. Data will be accessible through REST API written in Go-lang
3. Front-end application (web, mobile) will use the REST API

We are using 
[OpenAPI](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#pathItemObject) v2, 
 which uses YAML format, to define the REST API . From the OpenAPI specification 
 file [go-swagger](https://github.com/go-swagger/go-swagger) can autogenerate 
 both the server and client side code in supported language.
 


__This is still a work in progress and hence expect major changes.__

Please email me for any suggestion, feedback, vulnerability report. 