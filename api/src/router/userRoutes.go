package router

var userRoutes = []Router{
	{
		URI:             "/users",
		Method:          "POST",
		Function:        CreateUser,
		IsAutentication: false,
	},
	{
		URI:             "/users",
		Method:          "GET",
		Function:        GetUsers,
		IsAutentication: false,
	},
	{
		URI:             "/users/{userId}",
		Method:          "GET",
		Function:        GetUser,
		IsAutentication: false,
	},
	{
		URI:             "/users/{userId}",
		Method:          "PUT",
		Function:        UpdateUser,
		IsAutentication: false,
	},
	{
		URI:             "/users/{userId}",
		Method:          "DELETE",
		Function:        DeleteUser,
		IsAutentication: false,
	},
}
