# online-store

Design a RESTful API for an online store which can be used to manage different products and track home delivery agents.

Details: You need to send us a design document on how you would implement a RESTful API for an online store. It should support basic CRUD operations. Once a user places a buy order, an agent is assigned according to availability. Also, track the agent's live location.  You are free to assume everything else but make sure you document them. 

You are creating the API for a mobile developer who will use it to create a mobile app. It would be great if you can also include some example scenarios along with the expected request/response objects.

## Note:
A sample html page named client.html has been created in the root directory of the project. Edit the value of agent in the query parameter of the method to `ws://34.220.237.190/ws?agent=<agent_id>` and you can fetch lng/lat in real time. 