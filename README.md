# QUICK DOCUMENTATON OF API

+ ### POST /api/register
  >This endpoint allows to create account
+ ### POST /api/login
  >This endpoint allows to log in and returns jwt-token
+ ### GET /api/users
  >This endpoint allows to get a list of users 
+ ### GET /api/users/{id}
  >This endpoint allows to get a specific user by his id
+ ### GET /api/users/{id}/articles
  >This endpoint allows to get a list of user's articles by user id
+ ### GET /api/articles
  >This endpoint allows to get a list of articles (There is a pagination)
+ ### POST /api/articles
  >This endpoint allows to create an article, if you have a jwt-token
+ ### GET /api/articles/{id}
  >This endpoint allows to get an article of specific user by his id
+ ### DELETE /api/articles/{id}
  >This endpoint allows to delete an article by its id, if this article is yours and you have a jwt-token
+ ### GET /api/categories
  >This endpoint allows to get a list of categories
+ ### GET /api/categories/{id}
  >This endpoint allows to get a specific category by its id
+ ### GET /api/categories/{id}/articles
  >This endpoint allows to get a list of articles by category id (There is a pagination)
