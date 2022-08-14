# learning-clean-code-golang-one
```go
    package main

    func main(){
		//ar := repository.NewMysqlArticleRepository(dbConn)
		//au := usecase.NewArticleUsecase(ar)
		//httpDeliver.NewArticleHttpHandler(e, au)

		/*
		   repository depends on db connection and serves as a concrete implementation of
		   business logic

		   Usecase needs a repository to be injected. more like a cleaner controller in MVC

		   Handler is an interface adapter for usecases

		   Handler -> Usecase -> Repository

		   Domain (models)
		   Infrastructure (db, persistence login, external lib)
		   Interfaces (httpHandlers)
		   Application (usecase)

		*/
    }
   
```