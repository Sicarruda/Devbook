package router

import "api/src/controllers"

var userRoutes = []Router{
	{
		URI:             "/users",
		Method:          "POST",
		Function:        controllers.CreateUser,
		IsAutentication: false,
	},
	{
		URI:             "/users",
		Method:          "GET",
		Function:        controllers.GetUsers,
		IsAutentication: false,
	},
	{
		URI:             "/users/{userId}",
		Method:          "GET",
		Function:        controllers.GetUser,
		IsAutentication: false,
	},
	{
		URI:             "/users/{userId}",
		Method:          "PUT",
		Function:        controllers.UpdateUser,
		IsAutentication: false,
	},
	{
		URI:             "/users/{userId}",
		Method:          "DELETE",
		Function:        controllers.DeleteUser,
		IsAutentication: false,
	},
}
