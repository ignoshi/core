# Ignoshi Core

Two minds are better than one. Ignoshi is your backup memory which remembers alot of stuff that you might forget.
Consider Ignoshi as your friend and your manager at the same time, it will help you organize your daily life, answers
questions you have asked before, also reminds you with stuff you forgot.

#### Running Locally
using docker compose we can run the project for local development

```bash
docker-compose up
```

#### Brain Storming
 - CRUD Global Tags
 - CRUD Code Snipptes with tags
 - CRUD Notes with Tags
 - URL Bookmarks, Reading List
 - Knowledge Base (answers and questions)
 - Global search, or search per feature


#### Project Structure

The project will be spearated into multiple microservices, as following:

 - Ignosi App (a ui microservice written in Vuejs)
 - Ignoshi Tags (a microservice responsible for tags operations)
 - Ignoshi Snippets
 - Ignoshi Notes
 - Ignoshi Notifications
 - Ignoshi Bookmarks
 - Ignoshi Reader
 - Ignoshi Calendar 
 - Integration with gitlab/github/bitbucket issues


#### Questions ?

 - Do we need a microservice for search functionality ?
 - How will microservices communicate with each others ? just http calls for now ?
